AWS_ACCESS_KEY_ID ?=
AWS_SECRET_ACCESS_KEY ?=

release:
	goreleaser
	rm -rf ./dist

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o tergum-linux-amd64

dev-test-all-implementation1: dev-test-local-implementation1 dev-test-aws-implementation1

dev-test-local-implementation1:
	go run tergum.go backup --implementation1 --config misc/example/config/tergum-local.json

dev-test-aws-implementation1:
	go run tergum.go backup --implementation1 --config misc/example/config/tergum-aws.local.json

dev-test-all: dev-test-local dev-test-aws dev-test-local-yaml dev-test-aws-yaml

dev-test-local:
	go run tergum.go backup --config misc/example/config/tergum-local-v3.json

dev-test-local-yaml:
	go run tergum.go backup --config misc/example/config/tergum-local-v3.yaml

dev-test-aws:
	go run tergum.go backup --config misc/example/config/tergum-aws-v3.local.json

dev-test-aws-yaml:
	go run tergum.go backup --config misc/example/config/tergum-aws-v3.local.yaml

commit-go-mod-tidy:
	git add go.sum
	git commit -m "[auto] refactor: go mod tidy"

docker-build-and-push-all-images:
	(cd misc/docker/redis-with-tergum && make all)
	(cd misc/docker/mysql-with-tergum && make all)
	(cd misc/docker/postgres-with-tergum && make all)
	(cd misc/docker/postgres-with-mysqldump-mongodump-with-tergum && make all)
	(cd misc/docker/postgres-with-mysqldump-with-tergum && make all)
	(cd misc/docker/postgres-with-redis-with-tergum && make all)
	(cd misc/docker/tergum-with-ca-certificates && make all)
