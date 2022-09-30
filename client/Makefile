# Cross platform support
ifeq ($(OS), Windows_NT)
  MV ?= move
  RM ?= del -Recurse -Force
else
  MV ?= mv -f
  RM ?= rm -rf
endif

clean:
	$(RM) ./build