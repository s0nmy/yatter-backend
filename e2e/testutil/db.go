package testutil

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func setupMySQL(ctx context.Context, t *testing.T) *mysql.MySQLContainer {
	t.Helper()

	// migrations を読み込む
	entries, err := os.ReadDir(filepath.Join("..", "..", "migrations"))
	if err != nil {
		slog.Error("マイグレーションファイルを読み込めませんでした", "error", err)
		panic(err)
	}

	scripts := make([]string, 0, len(entries))
	for _, entry := range entries {
		scripts = append(scripts, filepath.Join("..", "..", "migrations", entry.Name()))
	}

	mysqlContainer, err := mysql.Run(
		ctx,
		"mysql:5.7",
		mysql.WithScripts(scripts...),
		mysql.WithDatabase("test"),
		mysql.WithUsername("root"),
		mysql.WithPassword("password"),
		mysql.WithConfigFile(filepath.Join("..", "..", "mysql", "conf.d", "my.cnf")),
	)
	if err != nil {
		slog.Error("MySQLコンテナを起動できませんでした", "error", err)
		panic(err)
	}

	return mysqlContainer
}

func downMySQL(mysqlContainer *mysql.MySQLContainer) {
	err := mysqlContainer.Terminate(context.Background())
	if err != nil {
		slog.Error("MySQLコンテナを停止できませんでした", "error", err)
		panic(err)
	}
}
