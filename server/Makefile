# Fix environment
include .$(PWD)/.env

# Cross platform support
ifeq ($(OS), Windows_NT)
  MV ?= move
  RM ?= del -Recurse -Force
  EXT ?= .exe
else
  MV ?= mv -f
  RM ?= rm -rf
  EXT ?= ""
endif

.PHONY: deps
deps:
	cd pkg && go get .

build: deps
	go build -o ./dist/server$(EXT)

.PHONY: run
run: deps
	go run ./...

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