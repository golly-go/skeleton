GO=go
GOCOVER=$(GO) tool cover
GOTEST=TZ=UTC $(GO) test

deps: 
	go get -u ./...
	
vet: 
	go vet ./...

test:
	$(GOTEST) ./... -cover 

.PHONY: test/cover
test/cover: vet
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out -o coverage.html
	@unlink coverage.out

development:
	air -c .air.toml
