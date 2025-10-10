REDIS_VERSION=8
TERGUM_VERSION=v0.42.0
IMAGE=ghcr.io/tergum-extra/redis-with-tergum:${REDIS_VERSION}-tergum-${TERGUM_VERSION}

all:
	make build-and-push

build-and-push: build push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg REDIS_VERSION=${REDIS_VERSION} \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
