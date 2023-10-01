package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	firebase "firebase.google.com/go"
	"github.com/tiozafrem/debtors/handlers"
	repostiryFirestory "github.com/tiozafrem/debtors/repositories/firestore"
	"github.com/tiozafrem/debtors/services"
	"google.golang.org/api/option"
)

// @title Debtors API
// @version 1.0
// @description Api server for Debtors Application

// @host localhost:8080
// @BasePath /
func main() {
	var file_key string
	ctx := context.Background()
	basePathFile := "../../"
	filepath.Walk(basePathFile, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			slog.Error(err.Error())
			return nil
		}

		if !info.IsDir() && filepath.Dir(path) != fmt.Sprint(basePathFile, "docs") &&
			filepath.Dir(path) != fmt.Sprint(basePathFile, ".vscode") && filepath.Ext(path) == ".json" {
			file_key = path
		}

		return nil
	})
	opt := option.WithCredentialsFile(file_key)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		slog.Error("error initializing app: %v\n", err)
	}

	_ = repostiryFirestory.NewRepositoryFirestory(ctx, app)
	service := services.NewService(ctx, app)

	handler := handlers.NewHandler(service)
	srv := Server{}
	go func() {
		if err = srv.Run("8080", handler.InitRoutes()); err != nil {
			slog.Error("error occured while running http server: %s", err.Error())
		}

	}()
	slog.Info("server is run")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("Shutting Down")

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("error occured on server shutting down: %s", err.Error())
	}

}
