package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/tiozafrem/debtors/models"
)

func (r *RepositoryFirestore) usersTransactionCollection(userUUIDOwner string) *firestore.CollectionRef {
	return r.usersCollection().Doc(userUUIDOwner).Collection("users")
}

func (r *RepositoryFirestore) AddUsersTransaction(ctx context.Context, ownerUUID string, user *models.UserTransaction) error {
	collection := r.usersTransactionCollection(ownerUUID)

	_, err := collection.Doc(user.UserUUID).Set(ctx, user)
	return err
}
