PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

MINIKUBE = $(shell minikube ip)

# DOCKER TASKS
# Build the container
all: build push 

build: build-ruby build-golang build-python

build-ruby: ## build ruby example
	docker build -t $(MINIKUBE):5000/prom_client_ruby -f ./ruby/Dockerfile ./ruby 

build-golang: ## build golang example
	docker build -t $(MINIKUBE):5000/prom_client_golang -f ./golang/Dockerfile ./golang

build-python: ## build ruby example
	docker build -t $(MINIKUBE):5000/prom_client_python -f ./python/Dockerfile ./python

push: push-ruby push-golang push-python

push-ruby:
	docker push $(MINIKUBE):5000/prom_client_ruby

push-golang:
	docker push $(MINIKUBE):5000/prom_client_golang

push-python:
	docker push $(MINIKUBE):5000/prom_client_python