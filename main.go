
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "time"

    "github.com/CatKap/sixPsyh/config"
    "github.com/CatKap/sixPsyh/internal/server"
    "github.com/CatKap/sixPsyh/pkg/loger"
)

func main() {
    // load config
    cfg := config.LoadFromEnv()

    // init loger
    log := loger.New(cfg.Env)

    srv, err := server.New(cfg, log)
    if err != nil {
        log.Fatalf("failed to create server: %v", err)
    }

    // run server in goroutine
    go func() {
        log.Infof("starting server on %s", cfg.Address)
        if err := srv.Start(); err != nil {
            log.Fatalf("server exited: %v", err)
        }
    }()

    // graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    log.Info("shutdown signal received")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Errorf("shutdown error: %v", err)
    } else {
        log.Info("server stopped gracefully")
    }
}
