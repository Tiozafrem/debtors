package services

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/tiozafrem/debtors/models"
	"github.com/tiozafrem/debtors/repositories/firestore"
	"golang.org/x/exp/slog"
)

type Authorization interface {
	SignIn(email, password string) (*models.Tokens, error)
	SignUp(ctx context.Context, email, password string) (*models.Tokens, error)
	RefreshToken(refreshToken string) (*models.Tokens, error)
	ParseTokenToUserUUID(ctx context.Context, token string) (string, error)
}

type User interface {
	PinTelegramId(ctx context.Context, userUUID string, id string) error
	PinUserToUser(ctx context.Context, userUUIDowner string, userUUIDchild string) error
	AddTransaction(ctx context.Context, userUUIDowner, userUUIDchild string, value int) error
	GetSumTransactionDebtor(ctx context.Context, userUUID string, debtorUUID string) (int, error)
	GetSumTransactionDebtors(ctx context.Context, userUUID string) (map[string]int, error)
	GetSumMy(ctx context.Context, userUUID string) (map[string]int, error)
}

type Service struct {
	Authorization
	User
}

func NewService(ctx context.Context, app *firebase.App,
	repository *firestore.RepositoryFirestore, apiKey string) *Service {
	auth, err := app.Auth(ctx)
	if err != nil {
		slog.Error("error initializing auth: %v\n", err)
	}

	return &Service{
		Authorization: NewAuthorizationService(auth, apiKey),
		User:          NewServiceUser(repository),
	}
}
