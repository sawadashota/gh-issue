.DEFAULT_GOAL := help

# https://gist.github.com/tadashi-aikawa/da73d277a3c1ec6767ed48d1335900f3
.PHONY: $(shell grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/://')

test: ## run all test
	go test -v ./...

goimports: ## fix format by goimports
	goimports -w .

download: ## install dependencies
	go mod download

vendor:
	go mod vendor

install: ## install gh-issue command
	go install github.com/sawadashota/gh-issue/cmd/gh-issue

# https://postd.cc/auto-documented-makefile/
help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
