MYSQL_VERSION=latest
TERGUM_VERSION=v0.45.0
IMAGE=ghcr.io/tergum-extra/mysql-with-tergum:${MYSQL_VERSION}-tergum-${TERGUM_VERSION}
IMAGE_NEW=ghcr.io/tergumregistry/mysql-with-tergum:${MYSQL_VERSION}-tergum-${TERGUM_VERSION}

all:
	make build-push
	make build-push MYSQL_VERSION=8

build-push:
	@make build
	@make push

build:
	docker build \
		--pull \
		--platform linux/amd64 \
		--build-arg MYSQL_VERSION=${MYSQL_VERSION} \
		--build-arg TERGUM_VERSION=${TERGUM_VERSION} \
		-t ${IMAGE} \
		-t ${IMAGE_NEW} \
		.

push:
	docker push ${IMAGE}
	docker push ${IMAGE_NEW}
