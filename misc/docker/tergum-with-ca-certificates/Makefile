TERGUM_VERSION=v0.46.0
IMAGE=ghcr.io/tergum-extra/tergum-with-ca-certificates:${TERGUM_VERSION}
IMAGE_NEW=ghcr.io/tergumregistry/tergum-with-ca-certificates:${TERGUM_VERSION}
IMAGE_BACKPORT=ghcr.io/tergum-extra/tergum-ca-certificates:${TERGUM_VERSION}

all: build push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		-t ${IMAGE_NEW} \
		-t ${IMAGE_BACKPORT} \
		.

push:
	docker push ${IMAGE}
	docker push ${IMAGE_NEW}
	docker push ${IMAGE_BACKPORT}
