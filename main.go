package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/hysem/charts/config"
	"github.com/hysem/charts/container"
	"github.com/hysem/charts/routes"
	"github.com/hysem/charts/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Setup
	config.Load()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	routes.Set(e)
	templates.Set(e)
	container.Init()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Start server
	go func() {
		serverConfig := config.Current.Server
		serverAddress := ""
		if serverConfig.Address != "localhost" &&
			serverConfig.Address != "" &&
			serverConfig.Address != "0.0.0.0" &&
			serverConfig.Address != "172.0.0.1" {
		} else {
			serverAddress = serverConfig.Address
		}
		sAddress := fmt.Sprintf("%s:%d", serverAddress, serverConfig.Port)
		if err := e.Start(sAddress); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
