import asyncio
import time

async def func1():
    print("Hello I am the first function\n")
    await asyncio.sleep(1)
    print("Man I am done\n")

async def func2():
    print("Hello I am the second function\n")
    await asyncio.sleep(1)
    print("Man I am also done\n")

async def run():
    await func1()
    await func2()

if __name__ == '__main__':
    print("Starting the functions")

    loop = asyncio.get_event_loop()
    loop.run_until_complete(run())

    print("Ending the program")
