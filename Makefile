
   

GOCMD=go
BINARY_NAME=whatip
VERSION?=0.1.0
OUTDIR=./out/bin

LDFLAGS=-ldflags '-X main.VERSION=${VERSION}'

.PHONY: vendor clean
.SILENT: build version
.DEFAULT_GOAL: build

## Build:
build: ## Build your project and put the output binary in out/bin/ | TODO: Iterate & Compress using tar 
	mkdir -p $(OUTDIR)/$(VERSION);
	rm -rf $(OUTDIR)/$(VERSION)/*; \
	for OS in darwin linux; do \
		for ARCH in amd64 arm64 arm; do \
			GOOS=$$OS GOARCH=$$ARCH GO111MODULE=on $(GOCMD) build -mod vendor ${LDFLAGS} -o $(OUTDIR)/$(VERSION)/$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH .; \
		done; \
	done; \
	for file in $(OUTDIR)/$(VERSION)/*; do \
		tar -cf "$$file.tar.gz" "$$file"; \
		rm $$file; \
	done;
clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

version: ## Output version
	echo $(VERSION)