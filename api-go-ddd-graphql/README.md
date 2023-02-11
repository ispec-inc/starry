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
