# Fix environment
include ./.env

# Cross platform support
ifeq ($(OS), Windows_NT)
  MV ?= move
  RM ?= del -Recurse -Force
  EXT ?= .exe
else
  MV ?= mv -f
  RM ?= rm -rf
  EXT ?= ""
  ENV ?= TWITTER_CONSUMER_KEY=$(TWITTER_CONSUMER_KEY) TWITTER_CONSUMER_SECRET=$(TWITTER_CONSUMER_SECRET) TWITTER_BEARER_TOKEN=$(TWITTER_BEARER_TOKEN) TWITTER_ACCESS_KEY=$(TWITTER_ACCESS_KEY) TWITTER_ACCESS_SECRET=$(TWITTER_ACCESS_SECRET) 
endif

.PHONY: deps
deps:
	cd pkg && go get .

build: deps
	go build -o ./dist/server$(EXT)

.PHONY: run
run: deps
	go run ./...

.PHONY: stream
stream: deps
	$(ENV) go run ./cmd/streaming/streaming.go

.PHONY: test
test:
	go test -cover ./...

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