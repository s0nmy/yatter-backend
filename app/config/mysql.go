package config

import (
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
)

// accessor namespace
var MySQL _mysql

type _mysql struct{}

const (
	// 東京のタイムゾーンオフセット（秒）
	TokyoOffset = 9 * 60 * 60
)

// Read MySQL host
func (_mysql) Host() string {
	v, err := getString("MYSQL_HOST")
	if err != nil {
		slog.Error("Error fetching MySQL host", "error", err)
		panic(err)
	}
	return v
}

// Read MySQL user
func (_mysql) User() string {
	v, err := getString("MYSQL_USER")
	if err != nil {
		slog.Error("Error fetching MySQL user", "error", err)
		panic(err)
	}
	return v
}

// Read MySQL password
func (_mysql) Password() string {
	v, err := getString("MYSQL_PASSWORD")
	if err != nil {
		slog.Error("Error fetching MySQL password", "error", err)
		panic(err)
	}
	return v
}

// Read MySQL database name
func (_mysql) Database() string {
	v, err := getString("MYSQL_DATABASE")
	if err != nil {
		slog.Error("Error fetching MySQL database", "error", err)
		panic(err)
	}
	return v
}

// Read Timezone for MySQL
func (_mysql) Location() *time.Location {
	tz, err := getString("MYSQL_TZ")
	if err != nil {
		return time.FixedZone("Asia/Tokyo", TokyoOffset)
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		slog.Error("Invalid timezone", "timezone", tz, "error", err)
		panic(err)
	}
	return loc
}

// Build mysql.Config
func MySQLConfig() *mysql.Config {
	cfg := mysql.NewConfig()

	cfg.ParseTime = true
	cfg.Loc = MySQL.Location()
	if host := MySQL.Host(); host != "" {
		cfg.Net = "tcp"
		cfg.Addr = host
	}
	cfg.User = MySQL.User()
	cfg.Passwd = MySQL.Password()
	cfg.DBName = MySQL.Database()

	return cfg
}
