include .env

upd:
	docker compose up -d

run api:
	go run cmd/api/main.go

mysql login:
	mysql -u app-user -p -h $(DB_HOST) --port $(DB_PORT) -D $(DB_NAME)
