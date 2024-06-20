lint:
	golangci-lint run

test:
	go test -v ./...

migrate:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" up

status:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" status
