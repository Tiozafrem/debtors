package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/tiozafrem/debtors/models"
)

func (r *RepositoryFirestore) transactionCollection(userUUIDOwner, debtorUUID string) *firestore.CollectionRef {
	return r.usersTransactionCollection(userUUIDOwner).Doc(debtorUUID).Collection("transaction")
}

func (r *RepositoryFirestore) AddTransaction(ctx context.Context, ownerUUID, debtorUUID string, transaction *models.Transaction) error {
	collection := r.transactionCollection(ownerUUID, debtorUUID)

	_, err := collection.Doc(fmt.Sprint(transaction.Date.UnixNano())).Set(ctx, transaction)
	return err
}
