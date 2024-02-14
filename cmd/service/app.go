package main

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/junkd0g/go-client-cintworks"

	"github.com/junkd0g/estimates/internal/config"
	internallogger "github.com/junkd0g/estimates/internal/logger"
	"github.com/junkd0g/estimates/internal/service"
	transporthttp "github.com/junkd0g/estimates/internal/transport/http"
)

func main() {
	ctx := context.Background()
	logger, err := internallogger.NewLogger()
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Starting the application")

	logger.Info(ctx, "Creating config")
	config, err := config.GetAppConfig("./assets/config.yaml")
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Creating the Cint API client")
	cintAPI, err := client.New(config.CintAPI.URL, config.CintAPI.Key, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Creating the service")
	srv, err := service.New(cintAPI)
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Creating the HTTP server")
	httpServer, err := transporthttp.New(logger, srv, &transporthttp.ServerConfig{
		Port:    config.Server.Port,
		TimeOut: config.Server.TimeOut,
	})
	if err != nil {
		panic(fmt.Errorf("creating_http_server %w", err))
	}

	err = httpServer.Start()
	if err != nil {
		panic(fmt.Errorf("starting_http_server %w", err))
	}
}
