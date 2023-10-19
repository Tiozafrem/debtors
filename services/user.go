package services

import (
	"context"
	"fmt"

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
	user, _ := s.repository.FindUserBytelegramId(ctx, id)

	if user != nil {
		return fmt.Errorf("telegram id is already pin")
	}

	return s.repository.AddUsers(ctx, &models.User{
		UserUUID:   userUUID,
		TelegramId: id,
	})
}

func (s *ServiceUser) ExistUser(ctx context.Context, userUUID string) bool {
	_, err := s.repository.FindUserByUUID(ctx, userUUID)
	return err == nil
}
