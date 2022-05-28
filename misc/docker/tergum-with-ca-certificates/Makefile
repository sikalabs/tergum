TERGUM_VERSION=v0.29.0
IMAGE=ghcr.io/tergum-extra/tergum-with-ca-certificates:${TERGUM_VERSION}
IMAGE_BACKPORT=ghcr.io/tergum-extra/tergum-ca-certificates:${TERGUM_VERSION}

all: build push

build:
	docker build \
		--platform linux/amd64 \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		-t ${IMAGE_BACKPORT} \
		.

push:
	docker push ${IMAGE}
	docker push ${IMAGE_BACKPORT}
