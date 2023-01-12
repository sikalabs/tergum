MYSQL_VERSION=latest
TERGUM_VERSION=v0.32.0
IMAGE=ghcr.io/tergum-extra/mysql-with-tergum:${MYSQL_VERSION}-tergum-${TERGUM_VERSION}

all: build push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg MYSQL_VERSION=${MYSQL_VERSION} \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		.

push:
	docker push ${IMAGE}
