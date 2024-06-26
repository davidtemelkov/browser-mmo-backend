package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// TODO: Add exp and gold reward
type Boss struct {
	Position     int    `json:"position"`
	Name         string `json:"name"`
	ImageURL     string `json:"imageUrl"`
	Lvl          int    `json:"lvl"`
	Constitution int    `json:"constitution"`
	Dexterity    int    `json:"dexterity"`
	Intelligence int    `json:"intelligence"`
	Strength     int    `json:"strength"`
}

type BossModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (bm BossModel) Insert(boss *Boss) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.BossPrefix + constants.Boss,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.BossPrefix + strconv.Itoa(boss.Position),
		},
		constants.PositionAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(boss.Position),
		},
		constants.NameAttribute: &types.AttributeValueMemberS{
			Value: boss.Name,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: boss.ImageURL,
		},
		constants.LevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(boss.Lvl),
		},
		constants.ConstitutionAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(boss.Constitution),
		},
		constants.DexterityAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(boss.Dexterity),
		},
		constants.IntelligenceAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(boss.Strength),
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(constants.TableName),
		Item:      item,
	}

	_, err := bm.DB.PutItem(bm.CTX, putInput)
	if err != nil {
		return err
	}

	return nil
}

func (bm BossModel) Get(position string) (Boss, error) {
	boss := &Boss{}

	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.BossPrefix + constants.Boss,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.BossPrefix + position,
		},
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(constants.TableName),
		Key:       key,
	}

	result, err := bm.DB.GetItem(bm.CTX, input)
	if err != nil {
		return *boss, err
	}

	if len(result.Item) == 0 {
		return *boss, errors.New("boss not found")
	}

	if err := attributevalue.UnmarshalMap(result.Item, boss); err != nil {
		return *boss, err
	}

	return *boss, nil
}
