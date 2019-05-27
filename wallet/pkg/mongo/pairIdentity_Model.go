package mongo

import (
	"golang.org/x/crypto/bcrypt"
)

// pairwiseIdentityModel is used to store the pairwise dids for the user
type pairIdentityModel struct {
	Identifier string

	UserName       string
	OtherPartyName string
	PublicKey      string
	PrivateKey     string
}

func newPairIdentityModel(username string, otherPartyName string) (*pairIdentityModel, error) {
	// add code to generate public private key pair
	pairId := pairIdentityModel{UserName: username, OtherPartyName: otherPartyName, PublicKey: "piD1", PrivateKey: "piD2"}
	err := pairId.generateUniqueIdentifier()
	if err != nil {
		return &pairId, err
	}

	return &pairId, nil
}

func (pId *pairIdentityModel) generateUniqueIdentifier() error {
	randomIdentifier := []byte(pId.UserName)

	hash, err := bcrypt.GenerateFromPassword(randomIdentifier, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pId.Identifier = string(hash[:])
	return nil
}
