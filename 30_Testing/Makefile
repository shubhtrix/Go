# Go parameters
GOCMD		=go
GOBUILD		=$(GOCMD) build
GOCLEAN		=$(GOCMD) clean
GOTEST		=$(GOCMD) test
GOGET		=$(GOCMD) get
BINARY_NAME	=binary
BINARY_UNIX	=$(BINARY_NAME)_unix

SRC_DIR		=${PWD}/src

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(SRC_DIR)/*.go

test: 
	$(GOTEST) -v ./... 

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop


# Cross compilation
#build-linux:
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

#docker-build:
#	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
