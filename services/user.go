package services

import (
	"context"

	"github.com/tiozafrem/debtors/models"
	"github.com/tiozafrem/debtors/repositories/firestore"
)

type ServiceUser struct {
	repository *firestore.RepositoryFirestore
}

func NewServiceUser(repository *firestore.RepositoryFirestore) *ServiceUser {
	return &ServiceUser{
		repository: repository,
	}
}

func (s *ServiceUser) PinTelegramId(ctx context.Context, userUUID string, id string) error {
	return s.repository.AddUsers(ctx, &models.User{
		UserUUID:   userUUID,
		TelegramId: id,
	})
}
