build_docker:
	docker build . -t kanban-go/seed

compose_up:
	docker compose up

run_test_image:
	docker run --name seed -e POSTGRES_PASSWORD=mysecretpassword -d kanban-go/seed

exec_seed_image:
	docker exec -it seed sh

run:
	go run cmd/kanban-api/main.go

seed:
	go run cmd/seedDb/main.go

lint:
	golangci-lint run

test:
	# test_db_up
	go test -v ./...

migrate:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" up

migrate_reset:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" reset

migrate_test:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user_test dbname=kanban_go_test_db password=12345" up

status:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=kanban-go password=12345" status


psql_test_inspect:
	psql -d kanban-go -U db_user -W

build:
	go build ./...

setup: build_docker compose_up

run_all: compose_up run
