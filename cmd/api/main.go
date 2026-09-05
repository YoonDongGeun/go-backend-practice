package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YoonDongGeun/go-backend-practice/internal/config"
	"github.com/YoonDongGeun/go-backend-practice/internal/db"
	"github.com/YoonDongGeun/go-backend-practice/internal/server"
	"github.com/YoonDongGeun/go-backend-practice/internal/store"
	"github.com/YoonDongGeun/go-backend-practice/internal/user"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pgPool, err := db.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Error("failed to connect postgres", "error", err)
		os.Exit(1)
	}
	defer pgPool.Close()

	redisClient := db.NewRedisClient(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	defer redisClient.Close()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		logger.Error("failed to connect redis", "error", err)
		os.Exit(1)
	}

	queries := store.New(pgPool)
	userService := user.NewService(queries)
	userHandler := user.NewHandler(userService)

	httpServer := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      server.NewRouter(userHandler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Info("api server listening", "addr", cfg.HTTPAddr, "env", cfg.AppEnv)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("api server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("api server shutdown failed", "error", err)
		os.Exit(1)
	}
}
