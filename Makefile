run:
	go run cmd/kanban-api/main.go

lint:
	golangci-lint run

test:
	# docker compose -f docker-compose.test.yml  up
	go test -v ./...

migrate:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" up

status:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" status


test_db_up:
	docker compose -f docker-compose.test.yml  up
