package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/currency/app/config"
	"projects/currency/app/db"
	"projects/currency/app/server"
	"time"
)

// @title Swagger CurrencyRates
// @version 1.0
// @description This is a currency rates server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:13
// @BasePath /currency/{date}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	ctx := context.Background()

	cfg := config.Get()

	sqlDB, err := db.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t",log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	s := server.Init(ctx, cfg, sqlDB,errorLog)

	httpServer := &http.Server{
		Addr:         cfg.Port,
		Handler:      s,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Running http server on %s\n", cfg.Port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			infoLog.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	httpServer.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

	return nil
}


