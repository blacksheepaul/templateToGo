package router

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/blacksheepaul/templateToGo/core/config"
	"github.com/blacksheepaul/templateToGo/core/logger"
	"github.com/blacksheepaul/templateToGo/router/middleware"
	"github.com/gin-gonic/gin"
)

var GinLogger = gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
	return fmt.Sprintf(`{"time":"%s","client":"%s","method":"%s","path":"%s","latency":"%s","status":%d,"emsg":"%s"}`+"\n",
		p.TimeStamp.Format(time.DateTime),
		p.ClientIP,
		p.Method,
		p.Path,
		p.Latency,
		p.StatusCode,
		p.ErrorMessage,
	)
})
var log logger.Logger

func Register(r *gin.Engine, cfg *config.Config, l logger.Logger) *gin.Engine {
	log = l

	r.Use(GinLogger)
	r.Use(middleware.Cors(cfg))

	api := r.Group("/api")   // api
	auth := api.Group("/v1") // api/v1
	auth.Use(middleware.Auth())

	return r
}

func LaunchServer(ctx context.Context, wg *sync.WaitGroup, r *gin.Engine, cfg *config.Config) {
	defer wg.Done()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server startup failed", err)
		}
	}()

	<-ctx.Done()
	log.Info("Server received stop signal, shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown failed", err)
	}
	log.Info("Server exited gracefully.")
}
