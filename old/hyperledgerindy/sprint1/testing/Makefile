# makefile to facilitate working with Indy

# starts the local indy pool
start:
	docker run -itd -p 9701-9708:9701-9708 indy_pool

# stops and cleans the docker containers of local indy pool
stop:
	bash clean.sh

# removes the folder storing the client credentials
clean:
	rm -rf ~/.indy_client

# runs the main.py file
main:
	python main.py

# checks whether the credentials files exists or not
check:
	ls -al ~ | grep "indy"

# for random testing of makefile commands
test: main clean
