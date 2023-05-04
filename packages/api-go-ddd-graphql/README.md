# API
Go + GraphQL + DDDで設計されているWebAPIです。

## ドキュメント
- [ソフトウェアアーキテクチャ](./docs/software-architecture.md)

## 起動方法
```
./hack/run-local.sh go run cmd/api/main.go
```

## 環境構築の前提
このAPIは、以下がローカルで起動していることを前提に作られています。
- MySQL 5.7が、開発用に3306ポート、テスト用に3307ポートで動いている
- Redis 6.2が、開発用に6379ポートで動いている。

## 環境変数のExport
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
```
