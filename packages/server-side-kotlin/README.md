# Server Side Kotlin
このディレクトリでは、サーバーサイドKotlinのテンプレートを提供しています。
要最低限のGraphQLのサーバーが動く部分までを提供しています。ドメイン駆動設計を行う際のディレクトリ設計等は、各自で行ってください。

## 事前準備
- OpenJDK 21のインストール

## 起動
```
./gradlew build # 起動
./gradlew run # サーバーの実行
```

http://localhost:8080/graphiql にアクセスすると、GraphQLのPlaygroundが表示されます。
