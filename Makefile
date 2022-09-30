# Cross platform support
ifeq ($(OS), Windows_NT)
  MV=move
  RM=del -Recurse -Force
else
  MV=mv -f
  RM=rm -rf
endif

.PHONY: install
install:
	cd client && $(MAKE) install
	cd server && $(MAKE) install

run: install
	cd client && $(MAKE) build
	cd server && $(MAKE) run

test:
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