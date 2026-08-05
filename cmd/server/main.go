package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"

	"github.com/Golukpal/url-shortener/internal/app"
	"github.com/Golukpal/url-shortener/internal/config"
	"github.com/Golukpal/url-shortener/internal/db"
	"github.com/Golukpal/url-shortener/internal/logger"
	"github.com/Golukpal/url-shortener/internal/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log, err := logger.New()
	if err != nil {
		panic(err)
	}

	dbPool, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("db connection failed")
	}
	defer dbPool.Close()

	application := app.New(cfg, log, (*pgxpool.Pool)(unsafe.Pointer(dbPool)))

	r := router.Setup(application)
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Info("server started")

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)
	<-stop

	log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)

	log.Info("server stopped")
}
