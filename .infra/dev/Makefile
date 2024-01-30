export PROJECT := socnet
go-run:
	docker exec -it $(PROJECT)_app_1 bash -c "go run ./cmd/app/main.go"
go-build:
	docker exec -it $(PROJECT) bash -c "go build -o /go/bin/app ./cmd/app/main.go" ;docker-compose -f docker-compose.yml -f docker-compose-add.yml -p $(PROJECT) up
migr:
	docker exec -it $(PROJECT)_app_1 bash -c "cd internal/migrations/ ; goose up "	
up:
	docker-compose -p $(PROJECT) up
down:
	docker-compose -p $(PROJECT) down
docker_stats:
	docker stats $(PROJECT)
srv-psql:
	cd ../srv/psql;docker-compose -p $(PROJECT)-db up
