package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

     _ "github.com/labstack/echo/v4"

	"github.com/anuraj2023/bank-account-management-be/internal/api"
	"github.com/anuraj2023/bank-account-management-be/internal/config"
	"github.com/anuraj2023/bank-account-management-be/internal/repository"
	"github.com/anuraj2023/bank-account-management-be/pkg/immudb"
)

// @title Swagger - Bank Account Management APIs
// @version 1.0
// @description This projects deals with creating and fetching bank accounts
// @host https://bank-account-management-be.onrender.com/
// @BasePath /
func main() {

    // Loading environment variables and store in config 
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    // Initializing the immudb client 
    immudbClient, err := immudb.NewClient(cfg.ImmuDbUrl, cfg.ImmuDbApiKey)
    if err != nil {
        log.Fatalf("Error creating immudb client: %v", err)
    }

    repo := repository.NewAccountRepository(immudbClient)
    server := api.NewServer(cfg, repo)

    // Starting server in a separate goroutine
    go func() {
        if err := server.Start(cfg.ServerPort); err != nil {
            log.Printf("Server stopped with error: %v", err)
        }
    }()

    // Creating channel to listen for OS interrupt signal (Ctrl+C)
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Gracefully shutting down the server
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown error: %v", err)
    }

    log.Println("Server has exited gracefully")
}
