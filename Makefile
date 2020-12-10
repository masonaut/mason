.DEFAULT_GOAL := help

.SILENT : help
.PHONY  : apply clean deploy init install help plan validate

company  := Mason Systems
project  := Mason Workflow Orchestrator
version  := 0.1.0
# planfile := security.terraform.planfile

## build mason binary
build:
	@go build -o bin/mason

## clean up project
clean:
	@rm -rf ./bin

## install dependencies
install:
	@go get

## show help screen
help:
	printf "$(company) - $(project) v$(version)\n"
	printf "\nUsage:\n  make <target>\n\nTargets:\n"
	awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  \033[32m%-15s\033[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
