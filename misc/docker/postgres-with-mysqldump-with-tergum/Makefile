POSTGRES_VERSION=16
TERGUM_VERSION=v0.40.0
IMAGE=ghcr.io/tergum-extra/postgres-with-mysqldump-tergum:${POSTGRES_VERSION}-tergum-${TERGUM_VERSION}

all:
	make build-and-push POSTGRES_VERSION=16
	make build-and-push POSTGRES_VERSION=15
	make build-and-push POSTGRES_VERSION=13

build-and-push: build push

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
