run:
	go run cmd/kanban-api/main.go

lint:
	golangci-lint run

test:
	# test_db_up
	go test -v ./...

migrate:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" up

migrate_test:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user_test dbname=kanban_go_test_db password=12345" up

status:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" status

test_db_up:
	docker compose -f docker-compose.test.yml  up

psql_test_inspect:
	psql -d kanban_go_test_db -U db_user_test -W
