package firestore

import (
	"context"
	"fmt"

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

func (r *RepositoryFirestore) FindUserByUUID(ctx context.Context, uuid string) (*models.User, error) {
	var user *models.User
	doc, err := r.usersCollection().Doc(uuid).Get(ctx)
	doc.DataTo(&user)
	return user, err
}

func (r *RepositoryFirestore) FindUserBytelegramId(ctx context.Context, id string) (*models.User, error) {
	var user *models.User
	docs, err := r.usersCollection().Where("telegram_id", "==", id).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	if len(docs) < 1 {
		return nil, fmt.Errorf("not found")
	}
	err = docs[0].DataTo(&user)
	return user, err
}
