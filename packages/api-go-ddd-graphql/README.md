# API
Go + GraphQL + DDDで設計されているWebAPIです。

## ドキュメント
- [ソフトウェアアーキテクチャ](./docs/software-architecture.md)

## 起動方法
```
./hack/run-local.sh go run cmd/api/main.go
```

## 環境変数のExport
hackディレクトリのシェルスクリプトを使わない場合、以下のように環境変数をエクスポートすれば任意のgoのコマンドが実行できます。
```
export ROUTER_TIMEOUT=10s
export ROUTER_ALLOW_ORIGINS=http://localhost:3000
export MYSQL_USER=root
export MYSQL_PASSWORD=password
export MYSQL_DATABASE=starry
export MYSQL_HOST=localhost
export MYSQL_PORT=13306
export MYSQL_SHOW_ALL_LOG=true
export REDIS_MSGBS_HOST=localhost
export REDIS_MSGBS_PORT=16379

go test ./app/infra/reader
```
