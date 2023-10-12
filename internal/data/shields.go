package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base shields
type Shield struct {
	ID             string
	BaseName       string
	MinLevel       int
	BlockChanceMin int
	BlockChanceMax int
	IsLegendary    bool
	ImageURL       string
}

type ShieldModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (sm ShieldModel) Insert(shield *Shield) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Shield,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.Shield + shield.ID,
		},
		constants.BaseNameAttribute: &types.AttributeValueMemberS{
			Value: shield.BaseName,
		},
		constants.MinLevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(shield.MinLevel),
		},
		constants.ShieldBlockChanceMinAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(shield.BlockChanceMin),
		},
		constants.ShieldBlockChanceMaxAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(shield.BlockChanceMax),
		},
		constants.IsLegendaryAttribute: &types.AttributeValueMemberBOOL{
			Value: shield.IsLegendary,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: shield.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(constants.TableName),
		Item:      item,
	}

	_, err := sm.DB.PutItem(sm.CTX, putInput)
	if err != nil {
		return err
	}

	return nil
}
