# echo-ddd
Echo＆DDD

## 環境構築
- コンテナを起動
```
make docket_all
```

## API
- 確認用API
```
GET: http://localhost:8001/example/key1/get_example
```
- response
```json
{
    "types": "get_example",
    "status": 200,
    "items": {
        "example_key": "key1",
        "example_name": "Name1",
        "message": "get example completed"
    }
}
```

## Batch
- 確認用Batch
```
docker compose -f docker-compose.local.yml exec batch go run batch/main.go --command=ListExample
```

## Test
- テストを実行
```
docker compose -f docker-compose.test.yml exec test go test -v ./infra/dao/...
docker compose -f docker-compose.test.yml exec test go test -v ./domain/model/...
docker compose -f docker-compose.test.yml exec test go test -v ./test/e2e/...
```

## Swagger
- Swaggerを自動生成
```
make docker_swag_gen
```
- Mockサーバーを起動
```
make docker_swag_mock
```

## DI
- Wireを自動生成
```
make docker_wire_gen
```
