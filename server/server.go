package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"github.com/octo-5/karrot-api/config"
	"github.com/octo-5/karrot-api/graphql"
	mid "github.com/octo-5/karrot-api/middleware"
)

func RunServer(cfg *config.Config, db *xorm.Engine) {
	router := gin.New()

	if cfg.IsProduction {
		router.Use(mid.RequestLogging())
	}

	router.Use(mid.RecoverPanic())
	router.Use(mid.Config(cfg))
	//router.Use(mid.Cache(db))
	router.Use(mid.Database(db))
	router.Use(mid.Dataloader(db))
	router.Use(mid.AuthUser(db))

	router.GET("/heartbeat", heartbeat())
	router.POST("/query", graphql.Handler(db))

	if !cfg.IsProduction {
		router.GET("/playground", graphql.Playground("karrot-api", "/query"))
	}

	httpSrv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.APIPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			if err.Error() == "http: Server closed" {
				log.Infoln("http server closed")
			} else {
				log.WithError(err).Fatalln("unable to start the http server")
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.WithError(err).Fatalln("unable to shutdown the http server gracefully")
	}
}

func heartbeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
