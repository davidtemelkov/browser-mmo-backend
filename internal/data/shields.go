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

type GeneratedShield struct {
	Name         string
	Level        int
	BlockChance  int
	IsLegendary  bool
	ImageURL     string
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Price        int
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

// Query shields by legendary status
func (sm ShieldModel) queryShieldsByIsLegendary(isLegendary bool) ([]Shield, error) {
	filterExpression := "IsLegendary = :isLegendary"
	expressionAttributeValues := map[string]types.AttributeValue{
		":isLegendary": &types.AttributeValueMemberBOOL{
			Value: isLegendary,
		},
	}

	queryInput := &dynamodb.ScanInput{
		TableName:                 aws.String(constants.TableName),
		FilterExpression:          aws.String(filterExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	result, err := sm.DB.Scan(sm.CTX, queryInput)
	if err != nil {
		return nil, err
	}

	shields := []Shield{}
	for _, item := range result.Items {
		minLevel, err := strconv.Atoi(item[constants.MinLevelAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		blockChanceMin, err := strconv.Atoi(item[constants.ShieldBlockChanceMinAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		blockChanceMax, err := strconv.Atoi(item[constants.ShieldBlockChanceMaxAttribute].(*types.AttributeValueMemberN).Value)
		if err != nil {
			return nil, err
		}

		shield := Shield{
			ID:             item[constants.SK].(*types.AttributeValueMemberS).Value,
			BaseName:       item[constants.BaseNameAttribute].(*types.AttributeValueMemberS).Value,
			MinLevel:       minLevel,
			BlockChanceMin: blockChanceMin,
			BlockChanceMax: blockChanceMax,
			IsLegendary:    item[constants.IsLegendaryAttribute].(*types.AttributeValueMemberBOOL).Value,
			ImageURL:       item[constants.ImageURLAttribute].(*types.AttributeValueMemberS).Value,
		}
		shields = append(shields, shield)
	}

	return shields, nil
}

func (sm ShieldModel) GetAllBasicShields() ([]Shield, error) {
	return sm.queryShieldsByIsLegendary(false)
}

func (sm ShieldModel) GetAllLegendaryShields() ([]Shield, error) {
	return sm.queryShieldsByIsLegendary(true)
}
