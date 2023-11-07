package services

import (
	"context"
	"fmt"
	"time"

	"github.com/tiozafrem/debtors/models"
)

func (s *ServiceUser) AddTransaction(ctx context.Context, userUUIDowner, userUUIDchild string, value int) error {
	if userUUIDchild == userUUIDowner {
		return fmt.Errorf("user owner == user child")
	}

	if !s.ExistUser(ctx, userUUIDchild) {
		return fmt.Errorf("child uuid not found")
	}

	return s.repository.AddTransaction(ctx, userUUIDowner, userUUIDchild, &models.Transaction{
		Date:  time.Now(),
		Value: value,
	})
}
