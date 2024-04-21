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
tools-hl_to_table_postgres:
	docker exec -it $(PROJECT)_app_1 bash -c "go run ./tools/hl_to_table_postgres/main.go"		
docker_stats:
	docker stats $(PROJECT)
srv-psql-master:
	cd ../srv/psql/master;docker-compose -p $(PROJECT)-db-master up
srv-psql-s1:
	cd ../srv/psql/s1;docker-compose -p $(PROJECT)-db-s1 up	
srv-psql-s2:
	cd ../srv/psql/s2;docker-compose -p $(PROJECT)-db-s2 up		
srv-psql-pgadmin:
	cd ../srv/psql/pgadmin;docker-compose -p $(PROJECT)-db-pgadmin up	
srv-k6:
	cd ../srv/k6;docker-compose -p $(PROJECT)-hl-test up
srv-monitoring:
	cd ../srv/monitoring;docker-compose -p $(PROJECT)-hl-mon up
srv-grafana:
	cd ../srv/grafana;docker-compose -p $(PROJECT)-hl-grafana up	