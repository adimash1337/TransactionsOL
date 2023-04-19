package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/logger"
	"awesomeProject/internal/service"
	"awesomeProject/internal/storage"
	"awesomeProject/internal/transport/http"
	"awesomeProject/internal/transport/http/handlers"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

// @title Transactions microservice API
// @version	1.0
// @description Microservice with transactions
// @termsOfService http://swagger.io/terms/

// @contact.name Dinmikhamed Amurzakov
// @contact.email dimash.amurzakov@gmail.com

// @host localhost:9000

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Fatal(fmt.Sprintf("Service shut down: %s", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)

	cfg, err := config.New()
	if err != nil {
		logger.Logger().Println(err)
		log.Fatal(err)
	}

	str, err := storage.New(ctx, cfg)
	if err != nil {
		logger.Logger().Println(err)
		log.Fatal(err)
	}
	svc, err := service.NewManager(str)
	h := handlers.NewManager(cfg, svc)
	server := http.NewServer(cfg, h)
	return server.StartHTTpServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		logger.Logger().Println(<-osC)
		c()
	}()
}
