TERGUM_VERSION=v0.44.0
IMAGE=ghcr.io/tergum-extra/mssql-with-tergum:tergum-${TERGUM_VERSION}

all:
	make build-push

build-push:
	@make build
	@make push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
