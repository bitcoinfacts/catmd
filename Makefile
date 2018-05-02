COMMIT=`git rev-parse --short HEAD`
APP=catmd
REPO?=joaofnfernandes/$(APP)
TAG?=latest
BUILD?=-dev

.PHONY: image
image:
	@docker build --build-arg REPO=$(REPO) -t $(REPO):$(TAG) .
	@echo "Docker image created: $(REPO):$(TAG)"

.PHONY: push
push: image
	@docker push $(REPO):$(TAG)
	