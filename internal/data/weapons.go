package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type GeneratedWeapon struct {
	Name          string
	Lvl           int
	DamageMin     int
	DamageMax     int
	DamageAverage int
	IsLegendary   bool
	ImageURL      string
	Strength      int
	Dexterity     int
	Constitution  int
	Intelligence  int
	Price         int
}

// Base weapons
type Weapon struct {
	ID                 string
	BaseName           string
	MinLevel           int
	DamageLowRangeMin  int
	DamageLowRangeMax  int
	DamageHighRangeMin int
	DamageHighRangeMax int
	IsLegendary        bool
	ImageURL           string
}

// Temporary placement
type Item struct {
	WhatItem      string `json:"whatItem"`
	Name          string `json:"name"`
	Lvl           int    `json:"lvl"`
	DamageMin     int    `json:"damageMin"`
	DamageMax     int    `json:"damageMax"`
	DamageAverage int    `json:"damageAverage"`
	Strength      int    `json:"strength"`
	Dexterity     int    `json:"dexterity"`
	Constitution  int    `json:"constitution"`
	Intelligence  int    `json:"intelligence"`
	ArmourAmount  int    `json:"armourAmount"`
	BlockChance   int    `json:"blockChance"`
	IsLegendary   bool   `json:"isLegendary"`
	ImageURL      string `json:"imageURL"`
	Price         int    `json:"price"`
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
		constants.DamageLowRangeMinAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(weapon.DamageLowRangeMin),
		},
		constants.DamageLowRangeMaxAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(weapon.DamageLowRangeMax),
		},
		constants.DamageHighRangeMinAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(weapon.DamageHighRangeMin),
		},
		constants.DamageHighRangeMaxAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(weapon.DamageLowRangeMax),
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

		damageLowRangeMin, err := strconv.Atoi(item[constants.DamageLowRangeMinAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		damageLowRangeMax, err := strconv.Atoi(item[constants.DamageLowRangeMaxAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		damageHighRangeMin, err := strconv.Atoi(item[constants.DamageHighRangeMinAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		damageHighRangeMax, err := strconv.Atoi(item[constants.DamageHighRangeMaxAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		weapon := Weapon{
			ID:                 item[constants.SK].(*types.AttributeValueMemberS).Value,
			BaseName:           item[constants.BaseNameAttribute].(*types.AttributeValueMemberS).Value,
			MinLevel:           minLevel,
			DamageLowRangeMin:  damageLowRangeMin,
			DamageLowRangeMax:  damageLowRangeMax,
			DamageHighRangeMin: damageHighRangeMin,
			DamageHighRangeMax: damageHighRangeMax,
			IsLegendary:        item[constants.IsLegendaryAttribute].(*types.AttributeValueMemberBOOL).Value,
			ImageURL:           item[constants.ImageURLAttribute].(*types.AttributeValueMemberS).Value,
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
