package data

import (
	"browser-mmo-backend/constants"
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type WorkModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (wm WorkModel) Set(user *User, hours int, workReward int) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
	}

	workUntilTime := time.Now().Add(time.Hour * time.Duration(hours))
	workUntilFormatted := workUntilTime.Format(constants.TimeFormat)

	workDuration := hours * 60

	updateExpression := "SET " + constants.IsWorkingAttribute + " = :isWorking, " +
		constants.WorkDurationAttribute + " = :workDuration, " +
		constants.WorkingUntilAttribute + " = :workingUntil, " +
		constants.WorkRewardAttribute + " = :workReward"
	expressionAttributeValues := map[string]types.AttributeValue{
		":isWorking":    &types.AttributeValueMemberBOOL{Value: true},
		":workDuration": &types.AttributeValueMemberN{Value: strconv.Itoa(workDuration)},
		":workingUntil": &types.AttributeValueMemberS{Value: workUntilFormatted},
		":workReward":   &types.AttributeValueMemberN{Value: strconv.Itoa(workReward)},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := wm.DB.UpdateItem(wm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}

func (wm WorkModel) CollectRewards(user *User) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
	}

	newGold := user.Gold + user.WorkReward

	updateExpression := "SET " + constants.IsWorkingAttribute + " = :isWorking, " +
		constants.WorkDurationAttribute + " = :workDuration, " +
		constants.WorkingUntilAttribute + " = :workingUntil, " +
		constants.GoldAttribute + " = :newGold, " +
		constants.WorkRewardAttribute + " = :workReward"
	expressionAttributeValues := map[string]types.AttributeValue{
		":isWorking":    &types.AttributeValueMemberBOOL{Value: false},
		":workDuration": &types.AttributeValueMemberN{Value: "0"},
		":workingUntil": &types.AttributeValueMemberS{Value: ""},
		":newGold":      &types.AttributeValueMemberN{Value: strconv.Itoa(newGold)},
		":workReward":   &types.AttributeValueMemberN{Value: "0"},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := wm.DB.UpdateItem(wm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}

func (wm WorkModel) Cancel(user *User) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
	}

	updateExpression := "SET " + constants.IsWorkingAttribute + " = :isWorking, " +
		constants.WorkDurationAttribute + " = :workDuration, " +
		constants.WorkingUntilAttribute + " = :workingUntil, " +
		constants.WorkRewardAttribute + " = :workReward"
	expressionAttributeValues := map[string]types.AttributeValue{
		":isWorking":    &types.AttributeValueMemberBOOL{Value: false},
		":workDuration": &types.AttributeValueMemberN{Value: "0"},
		":workingUntil": &types.AttributeValueMemberS{Value: ""},
		":workReward":   &types.AttributeValueMemberN{Value: "0"},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := wm.DB.UpdateItem(wm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}
