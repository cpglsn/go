package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version = "1.0.0"
)

type Config struct {
	port        string
	environment string
}

type Application struct {
	config Config
	logger *log.Logger
}

func main() {

	var cfg Config

	flag.StringVar(&cfg.port, "port", ":4000", "API Server port")
	flag.StringVar(&cfg.environment, "environment", "dev", "API Server environment")
	flag.Parse()

	app := &Application{
		config: cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	srv := &http.Server{
		Addr:         app.config.port,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("Starting %s server on %s port", app.config.environment, app.config.port)
	err := srv.ListenAndServe()
	app.logger.Fatal(err)
}
