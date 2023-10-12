package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base accessories
type Accessory struct {
	ID          string
	Type        string
	BaseName    string
	MinLevel    int
	IsLegendary bool
	ImageURL    string
}

type AccessoryModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (am AccessoryModel) Insert(accessory *Accessory) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Accessory,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.AccessoryPrefix + accessory.ID,
		},
		constants.TypeAttribute: &types.AttributeValueMemberS{
			Value: accessory.Type,
		},
		constants.BaseNameAttribute: &types.AttributeValueMemberS{
			Value: accessory.BaseName,
		},
		constants.MinLevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(accessory.MinLevel),
		},
		constants.IsLegendaryAttribute: &types.AttributeValueMemberBOOL{
			Value: accessory.IsLegendary,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: accessory.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(constants.TableName),
		Item:      item,
	}

	_, err := am.DB.PutItem(am.CTX, putInput)
	if err != nil {
		return err
	}

	return nil
}
