MONGO_VERSION=latest
TERGUM_VERSION=v0.44.0
IMAGE=ghcr.io/tergum-extra/mongo-with-tergum:${MONGO_VERSION}-tergum-${TERGUM_VERSION}

all:
	make build-push MONGO_VERSION=8

build-push:
	@make build
	@make push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg MONGO_VERSION=${MONGO_VERSION} \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
