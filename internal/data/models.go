package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Models struct {
	Users  UserModel
	Quests QuestModel
	Work   WorkModel
}

func NewModels(db *dynamodb.Client, ctx context.Context) Models {
	return Models{
		Users:  UserModel{DB: db, CTX: ctx},
		Quests: QuestModel{DB: db, CTX: ctx},
		Work:   WorkModel{DB: db, CTX: ctx},
	}
}
