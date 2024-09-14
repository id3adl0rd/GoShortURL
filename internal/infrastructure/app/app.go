package app

import (
	"context"
	cache2 "github.com/id3adl0rd/cache"
	"github.com/sirupsen/logrus"
	"go-short-me/internal/infrastructure/cache"
	"go-short-me/internal/infrastructure/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
}

// TODO Move to .env file
const (
	addr             = "localhost:8080"
	writeTimeout     = 15 * time.Second
	readTimeout      = 15 * time.Second
	gracefulShutdown = 10 * time.Second
)

func (app *App) Start() {
	logrus.Println("Starting up...")

	cache.Lru = cache2.NewLru(100)

	go func() {
		r := route.NewRoute()

		app.server = &http.Server{
			Handler:      r,
			Addr:         addr,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		}

		logrus.Fatal(app.server.ListenAndServe())
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill,
		syscall.SIGTERM, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGQUIT)
	<-sc

	logrus.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdown)
	defer cancel()
	app.server.Shutdown(ctx)
}
