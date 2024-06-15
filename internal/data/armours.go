package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

// Base armours
type GeneratedArmour struct {
	Type         string
	Name         string
	Level        int
	ArmourAmount int
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	IsLegendary  bool
	ImageURL     string
	Price        int
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

// Query armours by type and legendary status
func (am ArmourModel) queryArmoursByTypeAndIsLegendary(armourType string, isLegendary bool) ([]Armour, error) {
	filterExpression := "Type = :type AND IsLegendary = :isLegendary"
	expressionAttributeValues := map[string]types.AttributeValue{
		":type": &types.AttributeValueMemberS{
			Value: armourType,
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

	armours := []Armour{}
	for _, item := range result.Items {
		minLevel, err := strconv.Atoi(item[constants.MinLevelAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		armourAmountMin, err := strconv.Atoi(item[constants.ArmourAmountMinAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		armourAmountMax, err := strconv.Atoi(item[constants.ArmourAmountMaxAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		armour := Armour{
			ID:              item[constants.SK].(*types.AttributeValueMemberS).Value,
			Type:            item[constants.TypeAttribute].(*types.AttributeValueMemberS).Value,
			BaseName:        item[constants.BaseNameAttribute].(*types.AttributeValueMemberS).Value,
			MinLevel:        minLevel,
			ArmourAmountMin: armourAmountMin,
			ArmourAmountMax: armourAmountMax,
			IsLegendary:     item[constants.IsLegendaryAttribute].(*types.AttributeValueMemberBOOL).Value,
			ImageURL:        item[constants.ImageURLAttribute].(*types.AttributeValueMemberS).Value,
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
