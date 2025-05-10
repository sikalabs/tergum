up:
	docker compose up -d --build

down:
	docker compose down --volumes

test:
	echo "select * from example_users;" | psql "host=127.0.0.1 port=15432 user=postgres password=pg dbname=postgres"

dump:
	pg_dump "host=127.0.0.1 port=15432 user=postgres password=pg dbname=postgres"
