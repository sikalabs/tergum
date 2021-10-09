TERGUM_VERSION=v0.19.0
IMAGE=ghcr.io/tergum-extra/tergum-ca-certificates:${TERGUM_VERSION}

all: build push

build:
	docker build \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
