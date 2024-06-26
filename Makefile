# 全てのコンテナを起動
docker_all:
	docker compose -f docker-compose.local.yml up -d --build
	docker compose -f docker-compose.test.yml up -d --build

# Wireを自動生成
docker_wire_gen:
	docker compose -f docker-compose.local.yml exec gen wire api/di/wire.go
	docker compose -f docker-compose.local.yml exec gen wire auth/di/wire.go
	docker compose -f docker-compose.local.yml exec gen wire batch/di/wire.go

# Swaggerを自動生成
docker_swag_gen:
	docker compose -f docker-compose.local.yml exec api swag init --dir=api --output=docs/swagger/api
	docker compose -f docker-compose.local.yml exec auth swag init --dir=auth --output=docs/swagger/auth

# Swaggerのモックサーバーを起動
docker_swag_mock:
	docker compose -f docker-compose.local.yml exec swagger prism mock ./docs/swagger/api/swagger.yaml --port=8000 --host=0.0.0.0

# ローカルDBに接続
docker_db_api:
	docker compose -f docker-compose.local.yml exec db_api mysql --host=localhost --user=mysql_user --password=mysql_password echo_ddd_api_local

docker_db_auth:
	docker compose -f docker-compose.local.yml exec db_auth mysql --host=localhost --user=mysql_user --password=mysql_password echo_ddd_auth_local

# Modelテスト
docker_test_model:
	docker compose -f docker-compose.test.yml exec api_test go test -v ./domain/model/...

# Daoテスト
docker_test_dao:
	docker compose -f docker-compose.test.yml exec api_test go test -v ./infra/dao/...

# E2Eテスト
docker_test_e2e:
	docker compose -f docker-compose.test.yml exec api_test go test -v ./test/e2e/api/...
