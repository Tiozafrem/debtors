package models

type User struct {
	UserUUID   string `firestore:"user_uuid"`
	TelegramId string `firestore:"telegram_id"`
}
