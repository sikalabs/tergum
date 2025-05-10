up:
	docker compose up -d --build

down:
	docker compose down --volumes

test:
	echo "select * from example_users;" | mysql -h 127.0.0.1 -P 13306 -u root -proot default

dump:
	mysqldump -h 127.0.0.1 -P 13306 -u root -proot default
