package data

import (
	"browser-mmo-backend/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Base weapons
type Weapon struct {
	ID          string
	BaseName    string
	MinLevel    int
	IsLegendary bool
	ImageURL    string
}

type WeaponModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (wm WeaponModel) Insert(weapon *Weapon) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Weapon,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.WeaponPrefix + weapon.ID,
		},
		constants.BaseNameAttribute: &types.AttributeValueMemberS{
			Value: weapon.BaseName,
		},
		constants.MinLevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(weapon.MinLevel),
		},
		constants.IsLegendaryAttribute: &types.AttributeValueMemberBOOL{
			Value: weapon.IsLegendary,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: weapon.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(constants.TableName),
		Item:      item,
	}

	_, err := wm.DB.PutItem(wm.CTX, putInput)
	if err != nil {
		return err
	}

	return nil
}

func (wm WeaponModel) queryWeaponsByIsLegendary(isLegendary bool) ([]Weapon, error) {
	keyConditionExpression := "#pk = :pk AND begins_with(#sk, :sk)"
	expressionAttributeNames := map[string]string{
		"#pk": constants.PK,
		"#sk": constants.SK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: constants.ItemPrefix + constants.Weapon,
		},
		":sk": &types.AttributeValueMemberS{
			Value: constants.WeaponPrefix,
		},
		":isLegendary": &types.AttributeValueMemberBOOL{
			Value: isLegendary,
		},
	}

	filterExpression := "IsLegendary = :isLegendary"

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(constants.TableName),
		KeyConditionExpression:    aws.String(keyConditionExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
		FilterExpression:          aws.String(filterExpression),
	}

	result, err := wm.DB.Query(wm.CTX, queryInput)
	if err != nil {
		return nil, err
	}

	weapons := []Weapon{}
	for _, item := range result.Items {
		minLevel, err := strconv.Atoi(item[constants.MinLevelAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		weapon := Weapon{
			ID:          item[constants.SK].(*types.AttributeValueMemberS).Value,
			BaseName:    item[constants.BaseNameAttribute].(*types.AttributeValueMemberS).Value,
			MinLevel:    minLevel,
			IsLegendary: item[constants.IsLegendaryAttribute].(*types.AttributeValueMemberBOOL).Value,
			ImageURL:    item[constants.ImageURLAttribute].(*types.AttributeValueMemberS).Value,
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

func (wm WeaponModel) GetAllBasicWeapons() ([]Weapon, error) {
	return wm.queryWeaponsByIsLegendary(false)
}

func (wm WeaponModel) GetAllLegendaryWeapons() ([]Weapon, error) {
	return wm.queryWeaponsByIsLegendary(true)
}
