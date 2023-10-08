docker_all:
	docker compose -f docker-compose.local.yml up -d --build

docker_wire_gen:
	docker compose -f docker-compose.local.yml exec gen wire api/di/wire.go
	docker compose -f docker-compose.local.yml exec gen wire batch/di/wire.go

docker_swag_gen:
	docker compose -f docker-compose.local.yml exec api swag init --dir=api --output=docs/swagger/api

docker_swag_mock:
	docker compose -f docker-compose.local.yml exec swagger prism mock ./docs/swagger/api/swagger.yaml --port=8000 --host=0.0.0.0

docker_migrate:
	docker compose -f docker-compose.local.yml exec db mysql --host=localhost --user=mysql_user --password=mysql_password echo_ddd_local

docker_test:
	docker compose -f docker-compose.test.yml up -d --build
