package data

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Maybe when generating a drop for a player
// type Weapon struct {
// 	ID           string
// 	Name         string
// 	Level        int
// 	DamageMin    int
// 	DamageMax    int
// 	DamageMean   int
// 	IsLegendary  bool
// 	ImageURL     string
// 	Strength     int
// 	Dexterity    int
// 	Constitution int
// 	Intelligence int
// 	Price        int
// }

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
	ID            string `json:"id"`
	Name          string `json:"name"`
	Level         string `json:"level"`
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
