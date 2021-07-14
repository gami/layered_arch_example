# Layered Architecture Template

## What's?

レイヤードアーキテクチャで作成したシンプルなWEBアプリケーションのサンプルです。

一般の事例などでよく見られるクリーンアーキテクチャに比べて
- ディレクトリ階層が極力フラットである（infra層などレイヤーのためのディレクトリがない） 
- Usecase層を省略している（単一ドメインの処理はドメインサービスで行えばよいが、複数のドメインでの処理が必要な場合はControllerに責務が置かれる）

などの特徴があります。

### pros 
- ドメインを分けることで、ドメイン外の責務を分離している

### cons
- DB -> Domain -> APIResponseなどの詰め替えが発生する

## ディレクトリ構成

```
├── domain
│   ├── error.go ... カスタムエラー
│   └── user
│       ├── reposiotry.go ... Repositoryのinterface
│       ├── service.go ... Domainサービス。Repositoryを呼び出したり、複雑な処理をおこなう。
│       └── user.go ... Domainモデル。データとそのデータの仕様に基づくロジックを持つ。
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
make setup
# migration
docker compose -f docker/docker-compose.local.yaml up
```
