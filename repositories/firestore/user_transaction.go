package firestore

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
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

func (r *RepositoryFirestore) GetSumAllTransaction(ctx context.Context, ownerUUID, debtorUUID string) (int, error) {
	collection := r.usersTransactionCollection(ownerUUID).Doc(debtorUUID).Collection("transaction")
	query := collection.NewAggregationQuery().WithSum("value", "sum")
	results, err := query.Get(ctx)
	if err != nil {
		return 0, err
	}
	summ, ok := results["sum"]
	if !ok {
		return 0, errors.New("firestore: couldn't get alias for sum from results")
	}

	summValue := summ.(*firestorepb.Value)
	return int(summValue.GetIntegerValue()), nil
}
