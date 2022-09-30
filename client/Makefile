# Cross platform support
ifeq ($(OS), Windows_NT)
  MV=move
  RM=del -Recurse -Force
else
  MV=mv -f
  RM=rm -rf
endif

install:
	npm run-script install

build:
	npm run-script build

.PHONY: clean
clean:
	$(RM) ./build

.PHONY: clean-all
clean-all: clean
	$(RM) ./node_modules