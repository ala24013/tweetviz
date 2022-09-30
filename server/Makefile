# Fix environment
include ./.env

# Cross platform support
ifeq ($(OS), Windows_NT)
  MV=move
  RM=del -Recurse -Force
  EXT=.exe
  GOPATH=
else
  MV=mv -f
  RM=rm -rf
  EXT=""
  ENV=TWITTER_CONSUMER_KEY=$(TWITTER_CONSUMER_KEY) TWITTER_CONSUMER_SECRET=$(TWITTER_CONSUMER_SECRET) TWITTER_BEARER_TOKEN=$(TWITTER_BEARER_TOKEN) TWITTER_ACCESS_KEY=$(TWITTER_ACCESS_KEY) TWITTER_ACCESS_SECRET=$(TWITTER_ACCESS_SECRET) 
  GOPATH=~/go/bin
endif

.PHONY: install
install:
	cd pkg && go get .

build: install
	go build -o ./dist/server$(EXT)

.PHONY: run
run: install
	$(ENV) go run ./cmd/main/main.go

.PHONY: test
test:
	go test -cover ./pkg/...

.PHONY: fmt
fmt:
	gofmt -w ./pkg/
	gofmt -w ./cmd/

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint: security critic

.PHONY: critic
critic: 
	$(GOPATH)/gocritic check -enableAll ./pkg/...

.PHONY: security
security:
	$(GOPATH)/gosec ./pkg/...

.PHONY: cover
cover:
	go test -coverprofile=c.out
	go tool cover -html=c.out

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	$(RM) dist
	$(RM) c.out

.PHONY: clean-all
clean-all: clean