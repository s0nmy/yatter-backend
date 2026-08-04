BINARY := yatter-backend-go
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
PATH := $(PATH):${MAKEFILE_DIR}bin
SHELL := env PATH="$(PATH)" /bin/bash

GOARCH = amd64

.PHONY: help
help: ## ヘルプメッセージを表示
	@echo "使用方法: make [ターゲット]"
	@echo ""
	@echo "ターゲット:"
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "  %-20s %s\n", $$1, $$NF }' $(MAKEFILE_LIST)

.PHONY: build
build: ## アプリケーションをビルド
	CGO_ENABLED=0 go build -o build/${BINARY}

.PHONY: build-linux
build-linux: ## Linux用にアプリケーションをビルド
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o build/${BINARY}-linux-${GOARCH} .

.PHONY: test
test: ## テストを実行
	go test $(shell go list ${MAKEFILE_DIR}/...)

.PHONY: test-coverage
test-coverage: ## テストカバレッジを確認
	go test -coverprofile=coverage.out $(shell go list ${MAKEFILE_DIR}/...)
	go tool cover -html=coverage.out

GOLANGCI_LINT_VERSION := v2.10.1

.PHONY: lint
lint: ## リンターを実行
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...

.PHONY: format
format: ## フォーマットを実行
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) fmt

.PHONY: clean
clean: ## ビルド成果物を削除
	rm -rf build/
	rm -f coverage.out

.PHONY: dev-up
dev-up: ## 開発環境を起動
	docker compose up -d

.PHONY: dev-down
dev-down: ## 開発環境を停止
	docker compose down

.PHONY: dev-logs
dev-logs: ## 開発環境のログを表示
	docker compose logs -f

.PHONY: migrate-up
migrate-up: ## データベースマイグレーションを実行
	$(MAKEFILE_DIR)scripts/migrate.sh up

.PHONY: migrate-down
migrate-down: ## データベースマイグレーションをロールバック
	$(MAKEFILE_DIR)scripts/migrate.sh down

.PHONY: migrate-create
migrate-create: ## 新しいマイグレーションファイルを作成 (使用方法: make migrate-create name=テーブル名)
	@if [ -z "$(name)" ]; then \
		echo "エラー: nameパラメータが必要です。使用方法: make migrate-create name=テーブル名"; \
		exit 1; \
	fi
	$(MAKEFILE_DIR)scripts/create_migration_file.sh create_$(name)_table

.PHONY: migrate-version
migrate-version: ## 現在のマイグレーション状態を表示
	$(MAKEFILE_DIR)scripts/migrate.sh version

.PHONY: migrate-force
migrate-force: ## 特定のバージョンに強制的にマイグレーション (使用方法: make migrate-force version=<version_number>)
	@if [ -z "$(version)" ]; then \
		echo "エラー: versionパラメータが必要です。使用方法: make migrate-force version=<version_number>"; \
		exit 1; \
	fi
	$(MAKEFILE_DIR)scripts/migrate.sh force $(version)

.PHONY:	build mod test lint vet clean dev-up dev-down dev-logs migrate-up migrate-down migrate-create migrate-version migrate-force
