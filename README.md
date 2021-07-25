# Layered Architecture Template

## What's?

レイヤードアーキテクチャによるAPIサーバーのサンプルです。
言語はGoです。

レイヤーで責務を分けることで、テストを容易にし、かつそれぞれのパッケージをシンプルに保っています。


## レイヤー構成と所属するパッケージ
### 1st
外部とのアダプタ
- cmd
- adapter
  - mysql
  - aws
### 2nd
リクエストやSQLを処理する
- controller
- repository
### 3rd
Coreしか呼んではいけない
- usecase
### Core
どこからimportしてもよい
- domain 
### レイヤーに属さない
- di
- config
 pkg

## ディレクトリ構成

```
├── domain
│   ├── error.go ... カスタムエラー
│   └── user
│       ├── reposiotry.go ... Repositoryのinterface
│       ├── service.go ... Domainサービス。Repositoryを呼び出したり、複雑な処理をおこなう。
│       └── user.go ... Domainモデル。データとそのデータの仕様に基づくロジックを持つ。
├── usecase ... 複数のサービスにまたがる複雑な処理を受け持つ
├── Makefile
├── README.md
├── client ... AWS/HTTP/Slack/Twitterなどの外部サービスのAPIを呼び出す
├── cmd
│   ├── server ... APIサーバエントリポイント
│   └── worker ... Workerエントリポイント
├── config
│   ├── config.go ... 設定
│   └── env 環境依存ファイル
│       └── *.toml
├── controller ... WEBサーバのHandler/Router
│   ├── controller.go ... Controllerのベース構造体（Handlerを単純にマージする）
│   ├── rooter.go ... Chi Rooter/ Middleware
│   └── user.go ... Handler
├── di ... DI
├── docker ... Docker関連
├── gen
│   ├── openapi ... oapi-codegenの生成済ファイル
│   └── schema　... SqlBoilerの生成済ファイル
├── go.mod
├── mysql
│   └── db.go ... DB接続
├── pkg　... ユーティリティ関数群
├── repository
│   ├── build ... DBレコードとドメインモデルの変換を行う
│   ├── testdata
│   │   └── fixture ... DBテスト用
│   └── user.go ... Repository実体
├── resources
│   ├── openapi ... OpenAPIスキーマ
│   └── sql ... DBスキーマ
└── scripts ... スクリプト
```
## 環境構築

```
direnv
make setup
docker compose -f docker/docker-compose.local.yaml run app make setup

# migration
docker compose -f docker/docker-compose.local.yaml up
```

