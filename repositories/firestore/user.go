package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/tiozafrem/debtors/models"
)

func (r *RepositoryFirestore) usersCollection() *firestore.CollectionRef {
	return r.client.Collection("users")
}

func (r *RepositoryFirestore) AddUsers(ctx context.Context, user *models.User) error {
	collection := r.usersCollection()

	_, err := collection.Doc(user.UserUUID).Set(ctx, user)
	return err
}
