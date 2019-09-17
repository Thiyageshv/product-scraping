# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u
GOINSTALL=$(GOCMD) install
EXTPKGPATH=github.com/

#package name
APP=app

#external packages
gocql = github.com/gocql/gocql
xxhash = github.com/OneOfOne/xxhash



all : clean deps build install
	@echo "Building server executable..."

deps:
	$(GOGET) $(gocql)
	$(GOGET) $(xxhash)

install :
	$(GOINSTALL)

build:
	$(GOBUILD)

.PHONY: clean
clean :
	rm -rf $(EXTPKGPATH)

.PHONY: gopath
gopath :
	@echo $(GOPATH)
