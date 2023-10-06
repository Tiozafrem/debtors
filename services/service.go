package services

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/tiozafrem/debtors/models"
	"golang.org/x/exp/slog"
)

type Authorization interface {
	SignIn(email, password string) (*models.Tokens, error)
	SignUp(ctx context.Context, email, password string) (*models.Tokens, error)
	RefreshToken(refreshToken string) (*models.Tokens, error)
	ParseTokenToUserUUID(ctx context.Context, token string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(ctx context.Context, app *firebase.App) *Service {
	auth, err := app.Auth(ctx)
	if err != nil {
		slog.Error("error initializing auth: %v\n", err)
	}

	return &Service{
		Authorization: NewAuthorizationService(auth),
	}
}
