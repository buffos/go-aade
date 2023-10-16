ifeq ($(OS),Windows_NT)
	SHELL := pwsh.exe
	.SHELLFLAGS := -Command
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		SHELL := pwsh
		.SHELLFLAGS := -Command
	endif
	ifeq ($(UNAME_S),Darwin)
		SHELL := pwsh
		.SHELLFLAGS := -Command
	endif
endif

test_pdf:
	echo "Testing PDF generation" ; $$env:FONT_PATH = "../assets/fonts" ; $$env:LOGO_FILE_PATH = "../assets/logos/logoLight.png" ; & go test -v -run ./... -covermode=set -coverprofile cover/coverage.out  ; & go tool cover -html cover/coverage.out -o cover/coverage.html



.PHONY: test_pdf