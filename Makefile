binary=setenv_linux
uname := $(shell uname -s)

ifeq ($(OS),Windows_NT)
    binary=setenv_windows
else
    ifeq ($(uname),Darwin)
		binary=setenv_mac
	endif
		apps=$(APPLICATION)
	ifeq ($(apps),)
    	apps=.
	endif
endif

build:
	go build -o $(binary) setenv.go

configure: $(binary)
	./$(binary) --service=$(apps)
