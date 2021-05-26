AWS_ACCESS_KEY_ID ?=
AWS_SECRET_ACCESS_KEY ?=

dev-test-all: dev-test-local dev-test-aws

dev-test-local:
	go run tergum.go -config misc/example/config/tergum-local.json

dev-test-aws:
	go run tergum.go -config misc/example/config/tergum-aws.local.json

commit-go-mod-tidy:
	git add go.sum
	git commit -m "[auto] refactor: go mod tidy"
