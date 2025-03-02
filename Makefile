clean:
	docker compose down
	# docker kill tulsi/seed
	docker rmi tulsi/seed
	docker volume rm -f tulsi_db

# Live-reload
air:
	air

build_image:
	docker build . -t tulsi/seed

compose_up:
	docker compose up

run_test_image:
	docker run --name seed -e POSTGRES_PASSWORD=mysecretpassword -d tulsi/seed

exec_seed_image:
	docker exec -it seed sh

run:
	go run cmd/api/main.go

seed:
	go run cmd/seedDb/main.go

lint:
	golangci-lint run

test:
	# test_db_up
	go test -v ./...

migrate:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=tulsi password=12345" up

migrate_down:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=tulsi password=12345" down

migrate_reset:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=tulsi password=12345" reset

migrate_test:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user_test dbname=tulsi_test_db password=12345" up

status:
	goose -dir "./migrations" postgres "host=localhost port=5432 user=db_user dbname=tulsi password=12345" status


psql_test_inspect:
	psql -d tulsi -U db_user -W

build:
	go build ./...

setup: build_docker compose_up

run_all: compose_up run
