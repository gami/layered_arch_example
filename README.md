# Layered Architecture Template

## What's?

レイヤードアーキテクチャによるGo製APIサーバーのサンプルです。

レイヤーで責務を分けることで、テストを容易にし、かつそれぞれのパッケージや構造体をシンプルに保つことができます。各パッケージには使用可能なパッケージのルールがあり、そのルールを守ることで依存の方向もシンプルかつ最小限になります。

一方でトレードオフとして、データベーススキーマから自動生成されたモデル、ドメインモデル、OpenAPIスキーマから自動生成されたモデルとの間で合計2回の変換処理が発生し、記述量が若干増える懸念があります。

## レイヤー構成と所属するパッケージ
### 構成
<img src="http://www.plantuml.com/plantuml/png/VPBDJiCm48JlUOezmY51_1mHwdj44LevMuZmJx0tI17gksCiWnDdrQDdfljDfpq5Hi-BqOrr8u76bVmzV3S0CweFV2F04MScdpG0vSpiB5a6iuPF7RNB9glCUCZrblZkdNaUKlZIR4WFEv9obhtJMe0jWVnhyPCxMIP_HaNGHrjXe9j_wNQecatsx54-wsbsOMBdrromz7lSzSiK-KeswsRH-fhKeHdiTtZQSKPdSB8YHY9dH4qkNwk6_xBBhm9j-tBOjw-4b99tbaHuxZhejxkaN7dcUz8wHHJlPJVVYquh6UK9aIDlyPSivW1TToMT_l6PM7smEP4T5wD_0000" width="50%" />

### 各レイヤー
各パッケージは、上のレイヤーに依存してはいけない。上のレイヤーを使いたい場合はinterfaceを使用する。
usecase、domainサービスを使う場合もinterfaceを使用する。
#### 1stレイヤ
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
  - 単純な処理の場合
    - domain内のserviceにそのまま委譲する
      - errorのWrapも不要
  - 複雑な処理の場合
    - トランザクションを使用してもよい
  - 他のusecaseを使用してはいけない
  - domainモデル/サービス以外に依存してはいけない
- usecase/form
  - 更新系や検索などの複雑な入力をcontrollerから受け取るための構造体
    - 単純な入力であれば引数でよい
  - domainサービスにそのまま渡してはいけない
    - repositoryに検索クエリとかフィルターみたいなことをしたい場合は、domain内にモデルを作る
#### 4th
- domain 
  - <各ドメイン集約>
    - ドメインごとにパッケージをわける
    - 集約内に複数のドメインモデルを持つ
      - それぞれのドメインの粒度はアプリケーションによる
        - 1サービス/1リポジトリの制約があるので、大きすぎると
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
      - 相互参照できないので、他の集約に依存する場合は単一方向になるようにする
    - 他の集約のドメインサービスを使うことはできない
      - 集約にまたがる処理はユースケースで吸収する
    - エラーなど、集約外のものは自由に使ってよい
  - failure
    - アプリケーションエラーを管理する
      - controllerで40X系のエラーに変換される
      - アプリケーションエラーはどのレイヤーで返しても良い 
  - logger
  - 上位のどのレイヤーにも直接依存してはいけない
#### レイヤーに属さないパッケージ
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
│   ├── tx.go トランザクションinterface
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
├── .golangci.yaml ... 推奨設定で修正したgolangci-lintの設定ファイル
├── go.mod
└── go.sumb
```

## 処理の追加フロー

### DBスキーマにフィールドを追加し、その値をAPIで返す （参照のみ）
- DBスキーマに追加しマイグレーションする
  - `make gen_model`を実行し、SQLBoilerのモデルを再生成する
- domainモデルにフィールドを追加する
- repositoryに必要な処理を追加する
  - repository/buildのSQLBoilerのモデルとdomainモデルの変換処理を修正する
- OpenAPIスキーマに追加する
  - `make gen_openapi`を実行し、OpenAPIの型を再生成する
- controllerに必要な処理を追加する
  - repository/buildのdomainモデルとのOpenAPIの型の変換処理を修正する

### APIのエンドポイントを既存コントローラに追加する
- 必要に応じてDBスキーマに追加しマイグレーションする
  - `make gen_model`を実行し、SQLBoilerのモデルを再生成する
- domainモデルにフィールドを追加する
- domainサービスにメソッドを追加する
- domainのrepositoryインターフェースにメソッドを追加する
- di/repository.goがビルドエラーになるので、インターフェースを満たすように修正する
- OpenAPIスキーマに追加する
  - `make gen_openapi`を実行し、OpenAPIの型を再生成する
- di/controller.goがビルドエラーになるので、インターフェースを満たすように修正する
- controllerのusecaseインターフェースにメソッドを追加する
- di/usecase.goがビルドエラーになるので、インターフェースを満たすように修正する
  - usecaseのserviceインターフェースにメソッドを追加する 

### APIのエンドポイントを新設する
- 必要に応じてDBスキーマに追加しマイグレーションする
  - `make gen_model`を実行し、SQLBoilerのモデルを再生成する
- domainに新しい集約を追加する
  - domain/xxx/xxx.go（ドメインモデル）を追加  
  - repositoryインターフェースを追加
  - domain/xxx/xxx_service.go（ドメインモデル）を追加 
  - diにドメインサービスの初期化を追加
- repositoryを追加
  - 実装を追加
    - xxx_repositoryを追加
    - repository/buildに必要な処理を追加 
  - diにrepositoryの初期化を追加
- OpenAPIスキーマに追加する
  - `make gen_openapi`を実行し、OpenAPIの型を再生成する
- controllerを追加
  - xxx_controller.goを追加
    - controller/buildに追加
    - diで初期化できるようにする
  - controller.goの構造体に新設したcontrollerを埋め込む
  - diでcontrollerの初期化処理を修正
  - usecaseのインターフェースを追加
- usecaseに追加する
  - xxx_usecase.goを追加
  - diでusecaseの初期化処理を修正
  - serviceのインターフェースを追加

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
  - OpenAPI3に対応している
  - スキーマ駆動開発ができる
  - コンポーネントを型として自動生成できるので、型構造が会わなくなった場合にビルドエラーにできる
  - スキーマバリデーションできる
- chi
  - oapi-codegenが対応している
  - net/httpに完全準拠している
  - middlewareがまあまあある
- SQLBoiler
  - DBスキーマからモデルを生成できる
    - スキーマと実装が合わなくなったらビルドエラーにできる
    - facebook/entも可能だが、あちらはmigrationも同時に行う必要があり、使いづらい。
  - EagerLoadingしたい場合は外部キー制約が必須（よい）
  - SQLライクなインターフェース
  - パフォーマンスが良い
  - nullが独自型
  - bulk insertできない
