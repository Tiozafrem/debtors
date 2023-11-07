package models

import (
	"time"
)

type Transaction struct {
	Date  time.Time `firestore:"date"`
	Value int       `firestore:"value"`
}
