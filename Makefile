SHELL := /bin/bash
.SHELLFLAGS = -e -c
.DEFAULT_GOAL := help
.ONESHELL:
.SILENT:

.PHONY: $(MAKECMDGOALS)

YAML_FILES := $(shell find . -not -path "*/node_modules/*" -not -path "**/helm-templates/**" -not -path "**/cluster-nodes/**/templates/**" -type f -name '*.yml')

UNAME_S := $(shell uname -s)

.PHONY: $(MAKECMDGOALS)

fix: fix/yaml ## [all] Fix format and lint errors

format: format/go format/python format/yaml ## [all] Format

format/go: ## [golang] Format
	go fmt ./...

fix/python: format/python ## [python] Fix
format/python:
	SHOULD_INSTALL=0
	if ! [[ -d ./venv ]]; then \
		python3 -m venv ./venv; \
		SHOULD_INSTALL=1; \
	fi
	. ./venv/bin/activate
	if [[ $${SHOULD_INSTALL} -eq 1 ]]; then \
		pip install --upgrade pip; \
		pip install -r ./requirements.dev.txt; \
	fi
	black .
	isort --profile black .

fix/yaml: format/yaml ## [yaml] Format
format/yaml:
	echo "[fix/format/yaml] begin"
	for file in $(YAML_FILES); do \
		yq -i -P $${file}; \
	done
	echo "[fix/format/yaml] end"

lint: lint/go lint/python lint/yaml ## [all] Lint

lint/go: lint/go/format lint/go/vet ## [all] Lint Golang

lint/go/format:
	gofmt -s -w .

lint/go/vet:
	# iterate over al folders with go.mod
	echo "[lint/vet/go] begin"
	for dir in $(shell find . -type f -name 'go.mod' -exec dirname {} \;); do \
		pushd $${dir}; \
		echo "[lint/vet/go] running go vet in $${dir}"; \
		go vet ./...; \
		popd; \
	done
	echo "[lint/vet/go] end"

lint/python: lint/python/format ## [all] Lint Python
lint/python/format:
	SHOULD_INSTALL=0
	if ! [[ -d ./venv ]]; then \
		python3 -m venv ./venv; \
		SHOULD_INSTALL=1; \
	fi
	. ./venv/bin/activate
	if [[ $${SHOULD_INSTALL} -eq 1 ]]; then \
		pip install --upgrade pip; \
		pip install -r ./requirements.dev.txt; \
	fi
	black --check .
	isort --profile black --check-only .

lint/yaml: lint/yaml/format ## [all] Lint YAML
lint/yaml/format:
	echo "[lint/format/yaml] begin"
	for file in $(YAML_FILES); do \
		yq -P $${file} > /dev/null; \
	done
	echo "[lint/format/yaml] end"

setup: setup/db

setup/db: ## [db] Setup DB
ifeq ($(UNAME_S),Linux)
	if ! command -v sqlite3 &> /dev/null; then \
		sudo apt-get install -y sqlite3; \
	fi
else ifeq ($(UNAME_S),Darwin)
	if ! command -v sqlite3 &> /dev/null; then \
		brew install sqlite; \
	fi
endif

env-%: ## Check for env var
	if [ -z "$($*)" ]; then \
		echo "Error: Environment variable '$*' is not set."; \
		exit 1; \
	fi

.PHONY: help
help: ## Displays help info
	awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
