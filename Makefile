dev-test-all: dev-test-mysql-to-stdout dev-test-mysql-to-file-full-path dev-test-mysql-to-file-with-time

dev-test-mysql-to-stdout:
	go run tergum.go -src mysql -src-mysql-host 127.0.0.1 -src-mysql-port 13306 -src-mysql-user root -src-mysql-password root -src-mysql-database default -dst stdout > ./tmp/backup.stdout.sql

dev-test-mysql-to-file-full-path:
	go run tergum.go -src mysql -src-mysql-host 127.0.0.1 -src-mysql-port 13306 -src-mysql-user root -src-mysql-password root -src-mysql-database default -dst file -dst-file-path tmp/backup.sql

dev-test-mysql-to-file-with-time:
	go run tergum.go -src mysql -src-mysql-host 127.0.0.1 -src-mysql-port 13306 -src-mysql-user root -src-mysql-password root -src-mysql-database default -dst file -dst-file-dir tmp -dst-file-prefix default -dst-file-suffix sql
