# yatter-backend-go

## Table of Contents

- [yatter-backend-go](#yatter-backend-go)
  - [Table of Contents](#table-of-contents)
  - [Library](#library)
  - [Development Environment](#development-environment)
  - [Commands](#commands)
    - [Available Commands](#available-commands)
      - [Build Commands](#build-commands)
      - [Development Environment](#development-environment-1)
      - [Testing and Linting](#testing-and-linting)
        - [Lint](#lint)
      - [Database Migration](#database-migration)
    - [Examples](#examples)
    - [Requirements](#requirements)
    - [Start](#start)
    - [Shutdown](#shutdown)
    - [Log](#log)
    - [Hot Reload](#hot-reload)
    - [Swagger UI](#swagger-ui)
      - [Test](#test)
      - [Authentication](#authentication)
    - [Migrate](#migrate)
      - [Create migration file](#create-migration-file)
      - [Execute migration](#execute-migration)
      - [Troubleshooting](#troubleshooting)
  - [Code](#code)
    - [Architecture](#architecture)
  - [Mock](#mock)

## Library

- HTTP
  - chi（[ドキュメント](https://pkg.go.dev/github.com/go-chi/chi/v5)）
- DB
  - sqlx（[ドキュメント](https://pkg.go.dev/github.com/jmoiron/sqlx)）

## Development Environment

開発環境を docker-compose で構築しています。

## Commands

開発で使用する主要なコマンドは Makefile に記載しています。
下記のコマンドでどの様なコマンドがあるか確認できます。

```bash
make help
```

### Available Commands

#### Build Commands

- `make build` - アプリケーションをビルド
- `make build-linux` - Linux 用にアプリケーションをビルド
- `make clean` - ビルド成果物を削除

#### Development Environment

- `make dev-up` - 開発環境を起動
- `make dev-down` - 開発環境を停止
- `make dev-logs` - 開発環境のログを表示

#### Testing and Linting

- `make test` - テストを実行
- `make test-coverage` - テストカバレッジを確認
- `make lint` - リンターを実行
- `make format` - フォーマットを実行

##### Lint

[golangci-lint](https://golangci-lint.run/) v2 を使用しています。`make lint` は Go があれば実行でき、コンテナの起動は不要です。設定は `.golangci.yml` にあり、プリセットは standard です。

現在有効なルール(standard)

- errcheck - エラー戻り値をチェックしていない箇所を検出する(未チェックのエラーはバグの原因になり得る)
- govet - 疑わしいコードを検出する(`go vet` と同様の検査)
- ineffassign - 代入したが使っていない変数(無駄な代入)を検出する
- staticcheck - バグ・パフォーマンス・スタイルに関する staticcheck のルールを適用する
- unused - 未使用の定数・変数・関数・型を検出する


#### Database Migration

- `make migrate-up` - マイグレーションを実行
- `make migrate-down` - マイグレーションをロールバック
- `make migrate-create name=<table_name>` - 新しいマイグレーションファイルを作成
- `make migrate-version` - 現在のマイグレーション状態を表示
- `make migrate-force version=<version>` - 特定のバージョンに強制的にマイグレーション

### Examples

```bash
# 開発環境の起動
make dev-up

# テストの実行
make test

# 新しいマイグレーションファイルの作成
make migrate-create name=users

# マイグレーションの実行
make migrate-up
```

各コマンドの詳細な説明は `make help` で確認できます。

### Requirements

- Go
- docker / docker-compose

### Start

```
docker compose up -d
```

### Shutdown

```
docker compose down
```

### Log

```
# ログの確認
docker compose logs

# ストリーミング
docker compose logs -f

# webサーバonly
docker compose logs web
docker compose logs -f web
```

### Hot Reload

[air](https://github.com/cosmtrek/air)によるホットリロードをサポートしており、コードを編集・保存すると自動で反映されます。
読み込まれない場合は`docker compose restart`を実行してください。

### Swagger UI

API 仕様を Swagger UI で確認できます。

開発環境を立ち上げ、Web ブラウザで[localhost:8081](http://localhost:8081)にアクセスしてください。

#### Test

各 API 定義の"Try it out"から API の動作確認を行うことができます。

#### Authentication

鍵マークのついたエンドポイントは認証付きエンドポイントです。

`Authentication`という HTTP ヘッダに`username ${ユーザー名}`を指定する単純な仕様です。

動作確認の際には画面上部の"Authorize"からヘッダの値の設定を行ってください。

### Migrate

データベースマイグレーションには`golang-migrate`を使用しています。

#### Create migration file

```sh
make migrate-create name={table_name}
```

マイグレーションファイルは以下の 2 つが自動生成されます：

- 更新用（up）：テーブルの作成やカラムの追加など、スキーマを更新する定義
  - 保存先：`migrations/000000_create_{table_name}_table.up.sql`
- 切り戻し用（down）：更新を取り消すための定義
  - 保存先：`migrations/000000_create_{table_name}_table.down.sql`

#### Execute migration

データベースの更新（マイグレーションの適用）を行う場合：

```sh
make migrate-up
```

変更を元に戻す（ロールバック）場合：

```sh
make migrate-down
```

#### Troubleshooting

`error: Dirty database version 3. Fix and force version.` のようなエラーが発生した場合、以下の手順で対処できます：

1. 現在のマイグレーションの状態を確認：

```sh
make migrate-version
```

2. 正常に動作しているバージョンを指定して戻します：

```sh
# 例：バージョン1に戻す場合
make migrate-force version=1
```

3. その後、SQL を修正して再度マイグレーションを実行：

```sh
make migrate-up
```

このエラーは通常、以下のような場合に発生します：

- マイグレーションの実行が途中で失敗した
- データベースの状態とマイグレーションファイルの状態が一致していない

注意：`force`コマンドを使用する前に、現在のデータベースの状態を必ず確認してください。

## Code

### Architecture

[ソフトウェアアーキテクチャ](doc/app-arch.md)

## Mock

### Prism

#### Branch

- `prism`

以下コマンドで prism ブランチにチェックアウトして下さい。

```bash
git checkout prism
```

#### Commands

- `docker compose up -d` - コンテナの立ち上げ
- `docker compose down` - コンテナの停止

##### Examples

```bash
# Mockの起動
docker compose up -d

# Mockの停止
docker compose down
```
