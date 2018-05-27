PACKAGE=github.com/whiter4bbit/imdb-datasets
PACKAGES=`go list ./... | grep -v vendor`
BINARY=imdb-datasets
VENDOR=vendor

clean:
	rm -rf $(VENDOR) $(BINARY)

all: test install

test: $(VENDOR)
	for pkg in $(PACKAGES); do go test -v $$pkg; done

install: $(BINARY)	
	go install $(PACKAGE)

$(BINARY): $(VENDOR)
	go build

$(VENDOR):
	dep ensure

generate:
	go get github.com/tinylib/msgp
	for pkg in $(PACKAGES); do go generate $$pkg; done

.PHONY: generate $(VENDOR) test all