# ECS

ecspressoとその周辺のコードを置く場所

## deploy
ECSタスクの更新とECSサービスへの更新を行う
環境変数`TAG`でイメージのタグを指定する

```bash
TAG=xxxxxxxxxxx ecspresso deploy --config env/$env/$service/ecspresso.yaml
```

## run
ECSタスクの更新とECSタスクを実行する
`tag=`でイメージのタグを指定する

```bash
TAG=xxxxxxxxxxx ecspresso run --config env/$env/$service/ecspresso.yaml
```

## 踏み台
ECSのportforward機能を使って、RDSのポートをローカルに転送します

```
$ task portforward env=$env
```
https://zenn.dev/ispec_inc/articles/ecspresso-bastion
