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

func (s *ServiceUser) GetSumTransactionDebtor(ctx context.Context, userUUID string, debtorUUID string) (int, error) {
	value, err := s.repository.GetSumAllTransactionByDebtor(ctx, userUUID, debtorUUID)
	return value, err
}

func (s *ServiceUser) GetSumTransactionDebtors(ctx context.Context, userUUID string) (map[string]int, error) {
	value, err := s.repository.GetSumAllDebtors(ctx, userUUID)
	return value, err
}

func (s *ServiceUser) GetSumMy(ctx context.Context, userUUID string) (map[string]int, error) {

	return s.repository.GetAllMy(ctx, userUUID)
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

func (s *ServiceUser) GetUUIDByTelegramId(ctx context.Context, id string) (string, error) {
	user, err := s.repository.FindUserBytelegramId(ctx, id)

	if user == nil {
		return "", err
	}

	return user.UserUUID, err
}
func (s *ServiceUser) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.repository.GetUsers(ctx)
}
func (s *ServiceUser) ExistUser(ctx context.Context, userUUID string) bool {
	_, err := s.repository.FindUserByUUID(ctx, userUUID)
	return err == nil
}
