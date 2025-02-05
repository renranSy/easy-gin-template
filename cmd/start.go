package cmd

import (
	"context"
	"easy-gin-template/cmd/wire"
	"easy-gin-template/internal/config"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartCmd() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "workdir",
				Aliases:     []string{"d"},
				Usage:       "Working directory",
				DefaultText: "configs",
				Value:       "configs",
			},
		},
		Action: func(c *cli.Context) error {
			mods := wire.SetupInjector(context.Background())
			srv := &http.Server{
				Addr:    fmt.Sprintf(":%d", config.CFG.Server.Port),
				Handler: mods.Mods.SetupRouter(),
			}

			go func() {
				mods.Log.Info(fmt.Sprintf("host\thttp://%s:%d", config.CFG.Server.Host, config.CFG.Server.Port))
				mods.Log.Info(fmt.Sprintf("prefix\t%s", config.CFG.Server.Prefix))
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					mods.Log.Fatal(fmt.Sprintf("listen: %s\n", err))
				}
			}()

			go func() {
				err := mods.Casbinx.ResetPolicy()
				if err != nil {
					mods.Log.Fatal(err.Error())
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit
			mods.Log.Info("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				mods.Log.Fatal(fmt.Sprintf("Server Shutdown:%s", err))
			}
			mods.Log.Info("Server exiting")
			return nil
		},
	}
}
