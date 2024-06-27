package data

import (
	"browser-mmo-backend/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base armours
type Armour struct {
	ID          string
	WhatItem    string
	BaseName    string
	MinLevel    int
	IsLegendary bool
	ImageURL    string
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
		constants.WhatItemAttribute: &types.AttributeValueMemberS{
			Value: armour.WhatItem,
		},
		constants.BaseNameAttribute: &types.AttributeValueMemberS{
			Value: armour.BaseName,
		},
		constants.MinLevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(armour.MinLevel),
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

func (am ArmourModel) queryArmoursByTypeAndIsLegendary(armourType string, isLegendary bool) ([]Armour, error) {
	keyConditionExpression := "#pk = :pk AND begins_with(#sk, :sk)"
	expressionAttributeNames := map[string]string{
		"#pk": constants.PK,
		"#sk": constants.SK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Armour,
		},
		":sk": &types.AttributeValueMemberS{
			Value: constants.ArmourPrefix,
		},
		":whatItem": &types.AttributeValueMemberS{
			Value: armourType,
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

	armours := []Armour{}
	for _, item := range result.Items {
		minLevel, err := strconv.Atoi(item[constants.MinLevelAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		armour := Armour{
			ID:          item[constants.SK].(*types.AttributeValueMemberS).Value,
			WhatItem:    item[constants.WhatItemAttribute].(*types.AttributeValueMemberS).Value,
			BaseName:    item[constants.BaseNameAttribute].(*types.AttributeValueMemberS).Value,
			MinLevel:    minLevel,
			IsLegendary: item[constants.IsLegendaryAttribute].(*types.AttributeValueMemberBOOL).Value,
			ImageURL:    item[constants.ImageURLAttribute].(*types.AttributeValueMemberS).Value,
		}
		armours = append(armours, armour)
	}

	return armours, nil
}

func (am ArmourModel) GetAllBasicArmoursOfType(armourType string) ([]Armour, error) {
	return am.queryArmoursByTypeAndIsLegendary(armourType, false)
}

func (am ArmourModel) GetAllLegendaryArmoursOfType(armourType string) ([]Armour, error) {
	return am.queryArmoursByTypeAndIsLegendary(constants.Helmet, true)
}
