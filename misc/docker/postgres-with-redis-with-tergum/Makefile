POSTGRES_VERSION=13
TERGUM_VERSION=v0.32.0
IMAGE=ghcr.io/tergum-extra/postgres-with-tergum-with-redis:${POSTGRES_VERSION}-tergum-${TERGUM_VERSION}

all: build push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg POSTGRES_VERSION=${POSTGRES_VERSION} \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
