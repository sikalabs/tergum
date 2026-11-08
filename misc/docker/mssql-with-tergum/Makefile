TERGUM_VERSION=v0.46.0
IMAGE=ghcr.io/tergum-extra/mssql-with-tergum:tergum-${TERGUM_VERSION}
IMAGE_NEW=ghcr.io/tergumregistry/mssql-with-tergum:tergum-${TERGUM_VERSION}

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
		-t ${IMAGE_NEW} \
		.

push:
	docker push ${IMAGE}
	docker push ${IMAGE_NEW}
