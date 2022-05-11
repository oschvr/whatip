
   

GOCMD=go
BINARY_NAME=whatip
VERSION?=0.0.4

.PHONY: build vendor clean

## Build:
build: ## Build your project and put the output binary in out/bin/ | TODO: Iterate & Compress using tar 
	mkdir -p out/bin/$(VERSION)
	for OS in darwin linux; do for ARCH in amd64 arm64 arm; do GOOS=$$OS GOARCH=$$ARCH GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(VERSION)/$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH .; done; done
	
clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor