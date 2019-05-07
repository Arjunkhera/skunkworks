import time
import asyncio

from indy import anoncreds, crypto, did, ledger, pool, wallet

import json
import logging
from typing import Optional

from indy.error import ErrorCode, IndyError
from utils import get_pool_genesis_txn_path, PROTOCOL_VERSION

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.INFO)

async def connectPool():

    logger.info("Getting started -> started")

    pool_name = 'pool1'
    logger.info("Open Pool Ledger: {}".format(pool_name))
    pool_genesis_txn_path = get_pool_genesis_txn_path(pool_name)
    pool_config = json.dumps({"genesis_txn": str(pool_genesis_txn_path)})

    # Set protocol version 2 to work with Indy Node 1.4
    await pool.set_protocol_version(PROTOCOL_VERSION)

    try:
        await pool.create_pool_ledger_config(pool_name, pool_config)
    except IndyError as ex:
        if ex.error_code == ErrorCode.PoolLedgerConfigAlreadyExistsError:
            pass
    pool_handle = await pool.open_pool_ledger(pool_name, None)

class wt:
    def __init__(self):
        print ("Wallet class called")

    async def openWallet(self):
        logger.info("\"Sovrin Steward\" -> Create wallet")
        self.steward_wallet_config = json.dumps({"id": "sovrin_steward_wallet"})
        self.steward_wallet_credentials = json.dumps({"key": "steward_wallet_key"})
        try:
            await wallet.create_wallet(self.steward_wallet_config, self.steward_wallet_credentials)
        except IndyError as ex:
            if ex.error_code == ErrorCode.WalletAlreadyExistsError:
                pass
        self.steward_wallet = await wallet.open_wallet(self.steward_wallet_config, self.steward_wallet_credentials)

    async def closeWallet(self):
        logger.info("==============================")

        logger.info(" \"Sovrin Steward\" -> Close and Delete wallet")
        await wallet.close_wallet(self.steward_wallet)
        await wallet.delete_wallet(self.steward_wallet_config, self.steward_wallet_credentials)

class stewDID:
    def __init__(self):
        print("DID class called")

    async def createAndStoreDid(self,steward_wallet):

        logger.info("\"Sovrin Steward\" -> Create and store in Wallet DID from seed")
        self.steward_did_info = {'seed': '000000000000000000000000Steward1'}
        (self.steward_did, self.steward_key) = await did.create_and_store_my_did(steward_wallet, json.dumps(self.steward_did_info))

    async def getDidVerKey(self,steward_wallet):
        logger.info("Obtaining Updated Verification Key from the Pool")
        print("The Old one : {}".format(self.steward_key))
        temp = await did.key_for_local_did(steward_wallet,self.steward_did)
        print("The Update one : {}".format(temp))

async def run():
    await connectPool()

    # Instantiate Class wallet
    stewardWallet = wt()
    # Open the wallet
    await stewardWallet.openWallet()

    # Instantiate DID class
    stdid = stewDID()
    # Create and store did in wallet
    await stdid.createAndStoreDid(stewardWallet.steward_wallet)
    # Query the verification key from the ledger
    await stdid.getDidVerKey(stewardWallet.steward_wallet)

    # Close the wallet
    await stewardWallet.closeWallet()

if __name__ == '__main__':

    print("Starting the testing script")

    loop = asyncio.get_event_loop()
    loop.run_until_complete(run())

    print("Ending the testing script")
