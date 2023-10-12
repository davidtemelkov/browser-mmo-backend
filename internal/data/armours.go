package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

// Base armours
type Armour struct {
	ID              string
	Type            string
	BaseName        string
	MinLevel        int
	ArmourAmountMin int
	ArmourAmountMax int
	IsLegendary     bool
	ImageURL        string
}

type ArmourModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (am ArmourModel) Insert(armour *Armour) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Armour,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.ArmourPrefix + armour.ID,
		},
		constants.TypeAttribute: &types.AttributeValueMemberS{
			Value: armour.Type,
		},
		constants.BaseNameAttribute: &types.AttributeValueMemberS{
			Value: armour.BaseName,
		},
		constants.MinLevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(armour.MinLevel),
		},
		constants.ArmourAmountMinAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(armour.ArmourAmountMin),
		},
		constants.ArmourAmountMaxAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(armour.ArmourAmountMax),
		},
		constants.IsLegendaryAttribute: &types.AttributeValueMemberBOOL{
			Value: armour.IsLegendary,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: armour.ImageURL,
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
