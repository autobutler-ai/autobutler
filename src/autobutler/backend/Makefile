SHELL := /bin/bash
.SHELLFLAGS = -e -c
.DEFAULT_GOAL := help
.ONESHELL:
.SILENT:

.PHONY: $(MAKECMDGOALS)

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif
ifeq (run/build,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

build: ## Build backend
	go build -o ./build ./...

run/build: build ## Run backend after building
	$(MAKE) run $(RUN_ARGS)

run: ## Run backend
	go run ./... $(RUN_ARGS)

serve: ## Serve backend
	$(MAKE) run serve

version: ## Print version
	$(MAKE) run version

