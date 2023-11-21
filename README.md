<div align="center">
  <img
    src="https://raw.githubusercontent.com/ispec-inc/starry/master/.github/logo/logo_transparent.png"
    alt="starry"
    height="250"
    width="250"
  />
</div>

# Starry
Starryは、Web APIやフロントエンドなどのアプリケーションを開発する際に有用なテンプレート群です。

# テンプレートの利用方法
Starryのコマンド経由で、テンプレートのコードをpullすることができます。

```bash
# Install
$ brew install ispec-inc/tap/starry

# Usage
$ starry
```

# テンプレートの開発

## ADRの生成
JavaScriptランタイムとして[bun](https://github.com/oven-sh/bun)のインストールが必要です。

```
make adr
```
生成されるファイルの名前は日付+タグ+タイトルとなります。(タグについては`ge=general`, `fe=front-end`, `be=back-end`となっています)