.PHONY: all build test clean docker-build docker-push deploy

BINARY_NAME=logistics-marketplace
DOCKER_IMAGE=your-docker-repo/logistics-marketplace
DOCKER_TAG=latest

all: build

build:
	go build -o $(BINARY_NAME) ./main.go

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

docker-build:
	docker build -f Dockerfile.prod -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-push:
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

deploy: docker-build docker-push
	bash scripts/deploy_to_aws.sh

# Add more targets as needed for linting, formatting, etc.
