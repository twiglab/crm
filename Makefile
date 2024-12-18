SHELL := /bin/bash

# Replace this with your registry address
REGISTRY := hub.pipi.dev:10443/erp

# Get all subdirectories under build/docker
SUBDIRS := $(shell find build/docker -mindepth 1 -type d)

# Build and push targets for each subdirectory
$(foreach dir, $(SUBDIRS), \
	$(eval SHORT_DIR := $(shell basename $(dir))) \
	$(eval IMAGE := $(REGISTRY)/$(SHORT_DIR)) \
	$(eval TAG := $(shell git describe --tags --always --dirty)) \
	$(info $(SHORT_DIR)_IMAGE: $(IMAGE)) \
	$(info $(SHORT_DIR)_TAG: $(TAG)) \
	$(eval $(SHORT_DIR).build: ; @echo "Building image: $(IMAGE):$(TAG)" && \
		docker build --build-arg  REGISTRY=$(REGISTRY) -t $(IMAGE):$(TAG) -f $(dir)/Dockerfile .) \
	$(eval $(SHORT_DIR).push: $(SHORT_DIR).build ; @echo "Pushing image: $(IMAGE):$(TAG)" && \
		docker push $(IMAGE):$(TAG)) \
	$(eval .PHONY: $(SHORT_DIR).build $(SHORT_DIR).push) \
)

# Add build-all, push-all, and help targets to .PHONY
.PHONY: build-all push-all help

# Build all images
build-all: $(foreach dir, $(SUBDIRS), $(shell basename $(dir)).build)

# Push all images
push-all: $(foreach dir, $(SUBDIRS), $(shell basename $(dir)).push)

# Help target
help:
	@echo "Usage:"
	@echo "  make <short_dir>.build   Build the Docker image for the specified subdirectory."
	@echo "  make <short_dir>.push    Push the Docker image for the specified subdirectory."
	@echo "  make build-all	   Build Docker images for all subdirectories."
	@echo "  make push-all	    Push Docker images for all subdirectories."
	@echo "  make help		Display this help message."
