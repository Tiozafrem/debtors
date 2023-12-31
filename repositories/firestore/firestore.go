package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/exp/slog"
)

type RepositoryFirestore struct {
	client *firestore.Client
}

func NewRepositoryFirestory(ctx context.Context, app *firebase.App) *RepositoryFirestore {
	client, err := app.Firestore(ctx)
	if err != nil {
		slog.Error("error initializing db: %v\n", err)
	}
	return &RepositoryFirestore{client: client}
}
