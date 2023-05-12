# ソフトウェアアーキテクチャ

## アーキテクチャ図
ドメイン駆動設計を採用しています。



```mermaid
graph LR

web --> controller
controller -->  use case (uc)
use case (uc) --> domain
```

- [web](api/app/web)は、APIサーバーの中でwebに関する関心事を持つ層です。
- [controller](api/app/controller)は、webからの入力をucが解釈できる形式に変換する責務を持つ層です。
- [use case](api/app/uc)は、アプリケーションビジネルルールの責務を持つ層であり、domainを呼び出してアプリケーションの処理を組み立てます。
- [domain](api/app/domain)は、エンタープライズビジネスルールの責務を持つ層です。


## アーキテクチャ思想
このプロダクトはビジネスロジックが複雑なので、その複雑さを管理するために、`domain`にビジネスルールが凝集するように設計をし、`uc`は`domain`に実装されているロジックを組み合わせてアプリケーションの処理を行うことに専念します。
