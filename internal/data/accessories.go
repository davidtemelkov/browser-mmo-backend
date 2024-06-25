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
	WhatItem    string
	BaseName    string
	MinLevel    int
	IsLegendary bool
	ImageURL    string
}

type GeneratedAccessory struct {
	Type         string
	Name         string
	Lvl          int
	IsLegendary  bool
	ImageURL     string
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Price        int
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
		constants.WhatItemAttribute: &types.AttributeValueMemberS{
			Value: accessory.WhatItem,
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

func (am AccessoryModel) queryAccessoriesByTypeAndIsLegendary(accessoryType string, isLegendary bool) ([]Accessory, error) {
	keyConditionExpression := "#pk = :pk AND begins_with(#sk, :sk)"
	expressionAttributeNames := map[string]string{
		"#pk": constants.PK,
		"#sk": constants.SK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Accessory,
		},
		":sk": &types.AttributeValueMemberS{
			Value: constants.AccessoryPrefix,
		},
		":whatItem": &types.AttributeValueMemberS{
			Value: accessoryType,
		},
		":isLegendary": &types.AttributeValueMemberBOOL{
			Value: isLegendary,
		},
	}

	filterExpression := "WhatItem = :whatItem AND IsLegendary = :isLegendary"

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(constants.TableName),
		KeyConditionExpression:    aws.String(keyConditionExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
		FilterExpression:          aws.String(filterExpression),
	}

	result, err := am.DB.Query(am.CTX, queryInput)
	if err != nil {
		return nil, err
	}

	accessories := []Accessory{}
	for _, item := range result.Items {
		minLevel, err := strconv.Atoi(item[constants.MinLevelAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		accessory := Accessory{
			ID:          item[constants.SK].(*types.AttributeValueMemberS).Value,
			WhatItem:    item[constants.WhatItemAttribute].(*types.AttributeValueMemberS).Value,
			BaseName:    item[constants.BaseNameAttribute].(*types.AttributeValueMemberS).Value,
			MinLevel:    minLevel,
			IsLegendary: item[constants.IsLegendaryAttribute].(*types.AttributeValueMemberBOOL).Value,
			ImageURL:    item[constants.ImageURLAttribute].(*types.AttributeValueMemberS).Value,
		}
		accessories = append(accessories, accessory)
	}

	return accessories, nil
}

func (am AccessoryModel) GetAllBasicAccessoriesOfType(accessoryType string) ([]Accessory, error) {
	return am.queryAccessoriesByTypeAndIsLegendary(accessoryType, false)
}

func (am AccessoryModel) GetAllLegendaryAccessoriesOfType(accessoryType string) ([]Accessory, error) {
	return am.queryAccessoriesByTypeAndIsLegendary(accessoryType, true)
}
