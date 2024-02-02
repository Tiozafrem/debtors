package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	firebase "firebase.google.com/go"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/tiozafrem/debtors/handlers/bot"
	repostiryFirestory "github.com/tiozafrem/debtors/repositories/firestore"
	"github.com/tiozafrem/debtors/services"
	"google.golang.org/api/option"
)

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

	apiKey := os.Getenv("FIREBASE_API_KEY")
	if apiKey == "" {
		slog.Error("env FIREBASE_API_KEY is null")
		return
	}

	rp := repostiryFirestory.NewRepositoryFirestory(ctx, app)
	service := services.NewService(ctx, app, rp, apiKey)

	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	handler := bot.NewHandler(service, dispatcher)
	_ = handler

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	updater.Idle()
}
