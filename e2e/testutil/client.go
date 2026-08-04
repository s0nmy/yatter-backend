package testutil

import (
	"net"
	"net/url"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jmoiron/sqlx"
)

type TestClient struct {
	*resty.Client
	db   *sqlx.DB
	port int
}

func NewTestClient(t *testing.T) *TestClient {
	ctx := t.Context()
	mysqlContainer := setupMySQL(ctx, t)
	t.Cleanup(func() {
		downMySQL(mysqlContainer)
	})

	// /app/config/mysql.go の ParseTime オプションに合わせる
	// /mysql/conf.d/my.cnf の character-set-server と collation-server と default-time-zone に合わせる
	dsn, err := mysqlContainer.ConnectionString(ctx, "parseTime=true", "charset=utf8mb4", "collation=utf8mb4_bin", "loc="+url.QueryEscape("Asia/Tokyo"))
	if err != nil {
		t.Fatalf("MySQLコンテナの接続情報を取得できませんでした: %s", err)
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("MySQLに接続できませんでした: %s", err)
	}
	t.Cleanup(func() {
		err = db.Close()
		if err != nil {
			t.Errorf("データベースの接続を閉じることができませんでした: %v", err)
		}
	})

	// 利用可能なポートを動的に取得
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("利用可能なポートを取得できませんでした: %s", err)
	}

	addr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		t.Fatalf("TCPアドレスの取得に失敗しました")
	}
	port := addr.Port

	if err = listener.Close(); err != nil {
		t.Fatalf("リスナーのクローズに失敗しました: %s", err)
	}

	client := resty.New()
	client.SetBaseURL("http://" + net.JoinHostPort("localhost", strconv.Itoa(port)) + "/v1")

	return &TestClient{
		Client: client,
		db:     db,
		port:   port,
	}
}

func (c *TestClient) Cleanup(t *testing.T) {
	if err := c.db.Close(); err != nil {
		t.Errorf("Failed to close database connection: %v", err)
	}
}

// DB returns the database connection
func (c *TestClient) DB() *sqlx.DB {
	return c.db
}

func (c *TestClient) Port() int {
	return c.port
}
