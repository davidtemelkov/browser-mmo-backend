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

type GeneratedAccessory struct {
	Type         string
	Name         string
	Level        int
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

func (am AccessoryModel) queryAccessoriesByTypeAndIsLegendary(accessoryType string, isLegendary bool) ([]Accessory, error) {
	filterExpression := "Type = :type AND IsLegendary = :isLegendary"
	expressionAttributeValues := map[string]types.AttributeValue{
		":type": &types.AttributeValueMemberS{
			Value: accessoryType,
		},
		":isLegendary": &types.AttributeValueMemberBOOL{
			Value: isLegendary,
		},
	}

	queryInput := &dynamodb.ScanInput{
		TableName:                 aws.String(constants.TableName),
		FilterExpression:          aws.String(filterExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	result, err := am.DB.Scan(am.CTX, queryInput)
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
			Type:        item[constants.TypeAttribute].(*types.AttributeValueMemberS).Value,
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
