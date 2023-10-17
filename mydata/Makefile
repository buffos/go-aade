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

test_mydata_service:
	# fill the environment variables with your credentials
	echo "Testing mydata Service" ; $$env:GSIS_VAT_USERNAME = "" ; $$env:GSIS_VAT_PASSWORD = "" ; & go test -v -run ./... -covermode=set -coverprofile cover/coverage.out  ; & go tool cover -html cover/coverage.out -o cover/coverage.html

.PHONY: test_mydata_service