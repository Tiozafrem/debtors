package firestore

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/tiozafrem/debtors/models"
	"google.golang.org/api/iterator"
)

func (r *RepositoryFirestore) usersTransactionCollection(userUUIDOwner string) *firestore.CollectionRef {
	return r.usersCollection().Doc(userUUIDOwner).Collection("users")
}

func (r *RepositoryFirestore) getUserSum(ctx context.Context, collection *firestore.CollectionRef) (int, error) {
	query := collection.NewAggregationQuery().WithSum("value", "sum")
	results, err := query.Get(ctx)
	if err != nil {
		return 0, err
	}
	summ, ok := results["sum"]
	if !ok {
		return 0, errors.New("firestore: couldn't get alias for sum from results")
	}
	return int((summ.(*firestorepb.Value)).GetIntegerValue()), nil
}

func (r *RepositoryFirestore) AddUsersTransaction(ctx context.Context, ownerUUID string, user *models.UserTransaction) error {
	collection := r.usersTransactionCollection(ownerUUID)

	_, err := collection.Doc(user.UserUUID).Set(ctx, user)
	return err
}

func (r *RepositoryFirestore) GetSumAllTransactionByDebtor(ctx context.Context, ownerUUID, debtorUUID string) (int, error) {
	collection := r.usersTransactionCollection(ownerUUID).Doc(debtorUUID).Collection("transaction")
	return r.getUserSum(ctx, collection)
}

func (r *RepositoryFirestore) GetSumAllDebtors(ctx context.Context, ownerUUID string) (map[string]int, error) {
	iter := r.usersTransactionCollection(ownerUUID).Documents(ctx)
	users := make(map[string]int)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		collection := doc.Ref.Collection("transaction")
		sum, err := r.getUserSum(ctx, collection)
		if err != nil {
			continue
		}
		users[doc.Ref.ID] = sum
	}

	return users, nil
}
