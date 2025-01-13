package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/blacksheepaul/templateToGo/core/config"
	log "github.com/blacksheepaul/templateToGo/core/logger"
	"github.com/blacksheepaul/templateToGo/model"
	"github.com/blacksheepaul/templateToGo/router"
	"github.com/blacksheepaul/templateToGo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig("config.yml")
	logger := log.SetZapLogger(*cfg)

	r := router.Register(gin.New(), cfg, logger)
	service.InitService(logger, cfg)

	model.InitDao(cfg, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go router.LaunchServer(ctx, &wg, r, cfg)

	byebye := make(chan os.Signal, 1) // Listen for system signal，such as SIGINT, SIGTERM
	signal.Notify(byebye, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Server started, press Ctrl+C to stop.")

	fmt.Println("Program is running ...")
	logger.Info("Program is running, waiting for termination signal...")
	someonesaidbye := <-byebye // waiting for signal
	logger.Info("Received signal: %s, shutting down...", someonesaidbye)

	cancel() // tell other goroutines to stop
	logger.Info("Program exited gracefully.")
	fmt.Println("Program exited gracefully.")
}
