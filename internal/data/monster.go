package data

import (
	"browser-mmo-backend/internal/constants"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base Monsters
type Monster struct {
	ID       string `json:",omitempty"`
	Name     string
	ImageURL string
}

type GeneratedMonster struct {
	Name         string  `json:"name"`
	ImageURL     string  `json:"imageURL"`
	Level        int     `json:"level"`
	Constitution float32 `json:"constitution"`
	Dexterity    float32 `json:"dexterity"`
	Intelligence float32 `json:"intelligence"`
	Strength     float32 `json:"strength"`
}

type MonsterModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (mm MonsterModel) Insert(monster *Monster) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.MonsterPrefix + constants.Monster,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.MonsterPrefix + monster.ID,
		},
		constants.NameAttribute: &types.AttributeValueMemberS{
			Value: monster.Name,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: monster.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(constants.TableName),
		Item:      item,
	}

	_, err := mm.DB.PutItem(mm.CTX, putInput)
	if err != nil {
		return err
	}

	return nil
}

func (mm MonsterModel) GetAll() ([]Monster, error) {
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(constants.TableName),
		KeyConditionExpression: aws.String("#pk = :pkval"),
		ExpressionAttributeNames: map[string]string{
			"#pk": constants.PK,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkval": &types.AttributeValueMemberS{
				Value: constants.MonsterPrefix + constants.Monster,
			},
		},
	}

	result, err := mm.DB.Query(mm.CTX, queryInput)
	if err != nil {
		return nil, err
	}

	monsters := make([]Monster, 0)
	for _, item := range result.Items {
		monster := Monster{}
		if err := attributevalue.UnmarshalMap(item, &monster); err != nil {
			return nil, err
		}
		monsters = append(monsters, monster)
	}

	return monsters, nil
}
