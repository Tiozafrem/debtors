package firestore

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/tiozafrem/debtors/models"
	"google.golang.org/api/iterator"
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

func (r *RepositoryFirestore) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	var user models.User
	iter := r.usersCollection().Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		err = doc.DataTo(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *RepositoryFirestore) GetUsersNotMy(ctx context.Context, uuid string) ([]models.User, error) {
	var users []models.User
	var user models.User
	myUsers, err := r.GetUsersMy(ctx, uuid)
	if err != nil {
		return nil, err
	}
	iter := r.usersCollection().Where("user_uuid", "!=", uuid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		err = doc.DataTo(&user)
		if err != nil {
			return nil, err
		}

		if slices.Contains(myUsers, user) {
			continue
		}

		users = append(users, user)
	}
	return users, nil
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

func (r *RepositoryFirestore) GetAllMy(ctx context.Context, myUUID string) (map[string]int, error) {
	iter := r.client.CollectionGroup("users").Where("user_uuid", "==", myUUID).Documents(ctx)
	users := make(map[string]int)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(strings.Split(doc.Ref.Path, "/")) <= 7 {
			continue
		}
		collection := doc.Ref.Collection("transaction")
		sum, err := r.getUserSum(ctx, collection)
		if err != nil {
			continue
		}
		users[doc.Ref.Parent.Parent.ID] = sum

	}
	return users, nil
}
