BINARY_NAME=my-singo

build:
	cd ./; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) 

dev:
	cd ./; go build -o $(BINARY_NAME) -v