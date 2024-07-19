package main

import (
	"context"
	"easy-gin-template/internal/config"
	"easy-gin-template/internal/router"
	"easy-gin-template/pkg/log"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.CFG.Server.Port),
		Handler: router.SetupRouter(),
	}

	go func() {
		log.Logger.Info(fmt.Sprintf("host\thttp://%s:%d", config.CFG.Server.Host, config.CFG.Server.Port))
		log.Logger.Info(fmt.Sprintf("prefix\t%s", config.CFG.Server.Prefix))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Logger.Fatal(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal(fmt.Sprintf("Server Shutdown:%s", err))
	}
	log.Logger.Info("Server exiting")
}
