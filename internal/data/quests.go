package data

import (
	"browser-mmo-backend/internal/constants"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base quests
type Quest struct {
	ID       string
	Name     string
	ImageURL string
}

type QuestModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (qm QuestModel) Insert(quest *Quest) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.Questrefix + constants.Quest,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.Questrefix + quest.ID,
		},
		constants.BaseNameAttribute: &types.AttributeValueMemberS{
			Value: quest.Name,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: quest.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(constants.TableName),
		Item:      item,
	}

	_, err := qm.DB.PutItem(qm.CTX, putInput)
	if err != nil {
		return err
	}

	return nil
}
