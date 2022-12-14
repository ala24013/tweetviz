# Cross platform support
ifeq ($(OS), Windows_NT)
  MV=move
  RM=del -Recurse -Force
  ENV= #$$env:TWITTER_BEARER_TOKEN='$(TWITTER_BEARER_TOKEN)';
else
  MV=mv -f
  RM=rm -rf
  ENV=TWITTER_BEARER_TOKEN=$(TWITTER_BEARER_TOKEN)
endif

.PHONY: install
install:
	cd client && $(MAKE) install
	cd server && $(MAKE) install

run: install
	echo "$(TWITTER_BEARER_TOKEN)"
	cd client && $(MAKE) build
	cd server && $(ENV) $(MAKE) run

build: install
	cd client && $(MAKE) build
	cd server && $(MAKE) build

test: install
	cd client && $(MAKE) test
	cd server && $(MAKE) test

.PHONY: precommit
precommit: pretty test

.PHONY: tidy
tidy:
	cd server && $(MAKE) tidy

.PHONY: pretty
pretty:
	cd client && $(MAKE) pretty
	cd server && $(MAKE) fmt

.PHONY: lint
lint:
	cd client && $(MAKE) lint
	cd server && $(MAKE) lint

.PHONY: clean
clean:
	cd client && $(MAKE) clean
	cd server && $(MAKE) clean

.PHONY: clean-all
clean-all:
	cd client && $(MAKE) clean-all
	cd server && $(MAKE) clean-all