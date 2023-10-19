package services

import (
	"context"
	"fmt"
	"time"

	"github.com/tiozafrem/debtors/models"
)

func (s *ServiceUser) PinUserToUser(ctx context.Context, userUUIDowner string, userUUIDchild string) error {
	if userUUIDchild == userUUIDowner {
		return fmt.Errorf("user owner == user child")
	}

	if !s.ExistUser(ctx, userUUIDchild) {
		return fmt.Errorf("child uuid not found")
	}

	return s.repository.AddUsersTransaction(ctx, userUUIDowner, &models.UserTransaction{
		UserUUID:       userUUIDchild,
		TimeMustReturn: time.Now(),
	})
}
