POSTGRES_VERSION=13
TERGUM_VERSION=v0.12.0
IMAGE=ghcr.io/tergum-extra/postgres-with-tergum:${POSTGRES_VERSION}-tergum-${TERGUM_VERSION}

all: build push

build:
	docker build \
		--build-arg POSTGRES_VERSION=13 \
		--build-arg TERGUM_VERSION=v0.12.0 \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
