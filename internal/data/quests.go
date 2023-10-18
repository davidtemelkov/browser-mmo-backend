package data

import (
	"browser-mmo-backend/internal/constants"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base quests
type Quest struct {
	ID       string `json:",omitempty"`
	Name     string
	ImageURL string
}

type GeneratedQuest struct {
	Name     string `json:"Name"`
	ImageURL string `json:"ImageURL"`
	Time     string `json:"Time"`
	EXP      string `json:"EXP"`
	Gold     string `json:"Gold"`
}

type QuestModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (qm QuestModel) Insert(quest *Quest) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.QuestPrefix + constants.Quest,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.QuestPrefix + quest.ID,
		},
		constants.NameAttribute: &types.AttributeValueMemberS{
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

func (qm QuestModel) GetAll() ([]Quest, error) {
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(constants.TableName),
		KeyConditionExpression: aws.String("#pk = :pkval"),
		ExpressionAttributeNames: map[string]string{
			"#pk": constants.PK,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkval": &types.AttributeValueMemberS{
				Value: constants.QuestPrefix + constants.Quest,
			},
		},
	}

	result, err := qm.DB.Query(qm.CTX, queryInput)
	if err != nil {
		return nil, err
	}

	quests := make([]Quest, 0)
	for _, item := range result.Items {
		quest := Quest{}
		if err := attributevalue.UnmarshalMap(item, &quest); err != nil {
			return nil, err
		}
		quests = append(quests, quest)
	}

	return quests, nil
}
