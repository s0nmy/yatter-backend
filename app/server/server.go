package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"yatter-backend-go/app/config"
	"yatter-backend-go/app/domain/service"
	"yatter-backend-go/app/infra"
	"yatter-backend-go/app/infra/query"
	"yatter-backend-go/app/infra/transaction"
	api_auth "yatter-backend-go/app/ui/api/auth"
	"yatter-backend-go/app/ui/api/health"
	api_user "yatter-backend-go/app/ui/api/user"
	"yatter-backend-go/app/usecase/auth"
	"yatter-backend-go/app/usecase/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

const (
	requestTimeout    = 60 * time.Second
	shutdownTimeout   = 5 * time.Second
	readHeaderTimeout = 10 * time.Second
)

func Run(db *sqlx.DB) error {
	addr := ":" + strconv.Itoa(config.Port())

	// Transactor
	transactor := transaction.NewTransactor(db)

	// Repository
	userRepo := infra.NewUserRepository()

	// Domain Service
	usernameUniqueChecker := service.NewUsernameUniqueChecker(userRepo)

	// QueryService
	authQueryService := query.NewAuthQueryService(db)

	// UseCase
	userCreateUseCase := user.NewUserCreateUseCase(
		userRepo,
		usernameUniqueChecker,
		transactor,
	)
	loginUseCase := auth.NewLoginUseCase(authQueryService)

	// Handler
	userHandler := api_user.NewUserHandler(userCreateUseCase)
	authHandler := api_auth.NewAuthHandler(loginUseCase)

	// ルーターの設定
	r := chi.NewRouter()

	// ミドルウェアの設定
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(newCORS().Handler)

	// リクエストタイムアウトをコンテキストに設定
	// リクエストがタイムアウトした場合、ctx.Done()を通じて通知し、以降の処理を停止する
	r.Use(middleware.Timeout(requestTimeout))

	// v1 エンドポイント
	r.Route("/v1", func(r chi.Router) {
		// 認証関連
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
		})

		// ユーザー関連
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.SignUp)
		})

		// ヘルスチェック
		r.Route("/health", func(r chi.Router) {
			r.Get("/", health.Check)
		})
	})

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	l, err := net.Listen("tcp", addr)
	slog.Info("Serve on 127.0.0.1", "addr", addr)
	if err != nil {
		slog.Error("failed to listen", "err", err)
	}

	go func() {
		if err = srv.Serve(l); err != nil {
			slog.Error("failed to serve", "err", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", "err", err)
	}

	return nil
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	})
}
