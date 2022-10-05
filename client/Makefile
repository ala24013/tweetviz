# Cross platform support
ifeq ($(OS), Windows_NT)
  MV=move
  RM=del -Recurse -Force
else
  MV=mv -f
  RM=rm -rf
endif

install:
	npm install

.PHONY: build	
build:
	npm run-script build

.PHONY: test
test:

.PHONY: pretty
pretty:
	npx prettier --write ./public
	npx prettier --write ./src

.PHONY: clean
clean:
	$(RM) ./build

.PHONY: clean-all
clean-all: clean
	$(RM) ./node_modules