package models

import (
	"time"
)

type UserTransaction struct {
	UserUUID       string    `firestore:"user_uuid"`
	TimeMustReturn time.Time `firestore:"time_must_return"`
}
