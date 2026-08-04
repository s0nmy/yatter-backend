package main

import (
	"log/slog"
	"yatter-backend-go/app/config"
	"yatter-backend-go/app/server"
)

func main() {
	db, err := server.NewDB(config.MySQLConfig())
	if err != nil {
		slog.Error("failed to create db", "err", err)
		return
	}
	defer func() { _ = db.Close() }()

	if err = server.Run(db); err != nil {
		slog.Error("failed to run server", "err", err)
	}
}
