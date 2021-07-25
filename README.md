# Layered Architecture Template

## What's?

レイヤードアーキテクチャによるAPIサーバーのサンプルです。
言語はGoです。

レイヤーで責務を分けることで、テストを容易にし、かつそれぞれのパッケージや構造体をシンプルに保っています。

## レイヤー構成と所属するパッケージ
### 各レイヤー
各パッケージは、上のレイヤーに依存してはいけない。上のレイヤーを使いたい場合はinterfaceを使用する。下のレイヤーまた同じレイヤーの横のパッケージに
#### 1st
- cmd
  - エントリーポイント。
  - controllerを使用する
- adapter
  - mysql
    - データベース接続をおこなう。
    - 他のレイヤーに依存しない
  - aws 
    - AWSSDKを呼び出す。
    - 他のレイヤーに依存しない
#### 2nd
- controller
  - リクエストを処理し、interfaceを経由してusecaseを呼び出す  
  - domainモデルとの変換処理をcontroller/buildにおく
  - 基本的にはusecaseをそのまま呼び出す
  - 1つの処理の中でusecaseを複数呼び出してもよい
    - 例えばdomainモデルの形式とレスポンスの形式が異なる場合に、複数の参照系usecaseからそれぞれのdomainモデルを取得し、buildパッケージ内のbuilderを使ってモデルをマージする
- infra
  - repository
    - RDBやKVSにアクセスし、データを永続化する
    - 1つのレポジトリで複数のテーブルを管理してもよい
    - トランザクションを使用してもよい
    - domainモデルとの変換処理をrepository/buildにおく
    - 同レイヤーの他のパッケージを呼び出してはいけない
  - messenger
    - キューを処理する
  - client
    - 外部サービスのAPIクライアント
  - storarge
    - S3などのストレージ
  - mail
    - メール送信
#### 3rd
- usecase
  - domainを組み合わせてユースケースを実現する
    - ユースケースは再利用されない
  - 単純な処理の場合
    - domain内のserviceにそのまま委譲する
    　　- errorのWrapも不要
  - 複雑な処理の場合
    - トランザクションを使用してもよい
  - 他のusecaseを使用してはいけない
  - domainモデル/サービス以外に依存してはいけない
- usecase/form
  - 更新系や検索などの複雑な入力をcontrollerから受け取るための構造体
  - domainにそのまま渡してはいけない
    - repositoryに検索クエリとかフィルターみたいなことをしたい場合は、repositoryのためのこうぞdomainに
#### 4th
- domain 
  - <各ドメイン集約>
    - 複数のドメインモデルを持つ
      - ドメインモデルはデータとロジック
      - サービスが肥大化しないように、ドメイン自体の振る舞いはなるべくドメインモデルのメソッドにする
    - 1つのドメインサービスを持つ
      - ドメインモデルやリポジトリを使用して、1つの処理を実現する
        - 処理は再利用されうる
      - トランザクションを使用してもよい
      - interfaceを経由して1つのリポジトリを使う
    - 他の集約のドメインモデルを構造体のフィールドとして持つことはできない
      - IDしか持ってはいけない
      - ロジックの引数で受け取ることはOK
    - 他の集約のドメインサービスを使うことはできない
      - 集約にまたがる処理はユースケースで吸収する
  - failure
    - アプリケーションエラーを管理する
      - controllerで40X系のエラーに変換される
  - logger
#### レイヤーに属さない
- di
  - パッケージ依存の複雑さを吸収する
  - cmd/test以外から呼んではいけない
  - 何をimportしてもよい
  - 生成自動化したい
- config
  - 設定情報
  - アプリケーション内のどこから呼んでもよい
- pkg
  - utility的な関数群
  - アプリケーションに依存しないもの
    - 依存する場合はdomainにおく

### ディレクトリ構成

```
├── cmd
│   ├── server ... APIサーバエントリポイント
│   └── worker ... Workerエントリポイント
├── controller
│   ├── rooter ... rooter/middleare
│   ├── build ... req/resとドメインモデルを変換する
│   ├── controller.go ... Controllerのベース構造体（Handlerを単純にマージする）
│   └── user_controller.go
├── usecase
│   └── form
├── domain
│   ├── failure
│   └── user
│       ├── reposiotry.go
│       ├── messenger.go
│       ├── service.go
│       └── user.go
├── infra
│   └─ repository
│   │  ├── build ... DBレコードとドメインモデルの変換を行う
│   │  ├── testdata
│   │  │   └── fixture ... DBテスト用
│   │  └── user.go ... Repository実体
│   └─ messenger
├── adapter ... AWS/HTTP/Slack/Twitterなどの外部サービスのAPIを呼び出す
│   ├── aws
│   │   ├── config.go
│   │   └── sqs.go
│   ├── mysql
│   └── db.go ... DB接続
├── di ... DI
├── pkg　... ユーティリティ関数群
├── config
│   ├── config.go ... 設定
│   └── env 環境依存ファイル
│       └── *.toml
├── docker ... Docker関連
├── gen
│   ├── resources
│   │    ├── openapi ... OpenAPIスキーマ
│   │    └── sql ... DBスキーマ
│   ├── openapi ... oapi-codegenの生成済ファイル
│   └── schema　... SqlBoilerの生成済ファイル
├── scripts ... スクリプト
├── Makefile
├── README.md
├── go.mod
└── go.sumb
```

## 処理の追加フロー

## 環境構築

```
direnv
make setup
docker compose -f docker/docker-compose.local.yaml run app make setup

# migration
docker compose -f docker/docker-compose.local.yaml up
```

## 使用しているサードパーティライブラリ
- oapi-codegen
- chi
- SQLBoiler