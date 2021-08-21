AWS_ACCESS_KEY_ID ?=
AWS_SECRET_ACCESS_KEY ?=

dev-test-all-implementation1: dev-test-local-implementation1 dev-test-aws-implementation1

dev-test-local-implementation1:
	go run tergum.go backup --implementation1 --config misc/example/config/tergum-local.json

dev-test-aws-implementation1:
	go run tergum.go backup --implementation1 --config misc/example/config/tergum-aws.local.json

dev-test-all: dev-test-local dev-test-aws

dev-test-local:
	go run tergum.go backup --config misc/example/config/tergum-local-v3.json

dev-test-aws:
	go run tergum.go backup --config misc/example/config/tergum-aws-v3.local.json

commit-go-mod-tidy:
	git add go.sum
	git commit -m "[auto] refactor: go mod tidy"
