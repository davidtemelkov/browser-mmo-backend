package data

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/quests"
	"browser-mmo-backend/internal/utils"
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type QuestModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (qm QuestModel) SetQuests(email string, generatedQuests []quests.GeneratedQuest) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + email,
		},
	}

	generatedQuestsAttribute := map[string]types.AttributeValue{}
	for i, quest := range generatedQuests {
		questKey := "Quest" + strconv.Itoa(i)
		generatedQuestsAttribute[questKey] = &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{
				constants.NameAttribute: &types.AttributeValueMemberS{
					Value: quest.Name,
				},
				constants.IDAttribute: &types.AttributeValueMemberS{
					Value: quest.ID,
				},
				constants.TimeAttribute: &types.AttributeValueMemberS{
					Value: quest.Time,
				},
				constants.EXPAttribute: &types.AttributeValueMemberN{
					Value: quest.EXP,
				},
				constants.GoldAttribute: &types.AttributeValueMemberN{
					Value: quest.Gold,
				},
			},
		}
	}

	updateExpression := "SET " + constants.QuestsAttribute + " = :quests"
	expressionAttributeValues := map[string]types.AttributeValue{
		":quests": &types.AttributeValueMemberM{
			Value: generatedQuestsAttribute,
		},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := qm.DB.UpdateItem(qm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}

func (qm QuestModel) SetCurrentQuest(email string, quest map[string]quests.GeneratedQuest) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + email,
		},
	}

	var timeStr string
	currentQuestAttribute := map[string]types.AttributeValue{}
	for _, quest := range quest {
		questKey := "CurrentQuest"
		currentQuestAttribute[questKey] = &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{
				constants.NameAttribute: &types.AttributeValueMemberS{
					Value: quest.Name,
				},
				constants.IDAttribute: &types.AttributeValueMemberS{
					Value: quest.ID,
				},
				constants.TimeAttribute: &types.AttributeValueMemberS{
					Value: quest.Time,
				},
				constants.EXPAttribute: &types.AttributeValueMemberN{
					Value: quest.EXP,
				},
				constants.GoldAttribute: &types.AttributeValueMemberN{
					Value: quest.Gold,
				},
			},
		}
		timeStr = quest.Time
	}

	//TODO: Rework this later, quest.Time can be an integer
	minutes, err := strconv.Atoi(strings.Split(timeStr, " ")[0])
	if err != nil {
		return err
	}

	questingUntilTime := time.Now().UTC().Add(time.Minute * time.Duration(minutes))
	questingUntilFormatted := questingUntilTime.Format(constants.TimeFormat)

	updateExpression := "SET " + constants.CurrentQuestAttribute + " = :currentQuest, " + constants.IsQuestingAttribute + " = :isQuesting, " + constants.QuestingUntilAttribute + " = :questingUntil"
	expressionAttributeValues := map[string]types.AttributeValue{
		":currentQuest": &types.AttributeValueMemberM{
			Value: currentQuestAttribute,
		},
		":isQuesting":    &types.AttributeValueMemberBOOL{Value: true},
		":questingUntil": &types.AttributeValueMemberS{Value: questingUntilFormatted},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err = qm.DB.UpdateItem(qm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}

func (qm QuestModel) CancelCurrentQuest(user *User) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
	}

	emptyQuest := map[string]types.AttributeValue{}
	questKey := "CurrentQuest"
	emptyQuest[questKey] = &types.AttributeValueMemberM{
		Value: map[string]types.AttributeValue{
			constants.NameAttribute: &types.AttributeValueMemberS{
				Value: "Empty Quest 0",
			},
			constants.IDAttribute: &types.AttributeValueMemberS{
				Value: "",
			},
			constants.TimeAttribute: &types.AttributeValueMemberS{
				Value: "",
			},
			constants.EXPAttribute: &types.AttributeValueMemberN{
				Value: "0",
			},
			constants.GoldAttribute: &types.AttributeValueMemberN{
				Value: "0",
			},
		},
	}

	updateExpression := "SET " +
		constants.CurrentQuestAttribute + " = :emptyQuest, " +
		constants.IsQuestingAttribute + " = :isQuesting, " +
		constants.QuestingUntilAttribute + " = :questingUntil, " +
		constants.LastPlayedDateAttribute + " = :lastPlayedDate, " +
		constants.DailyQuestCountAttribute + " = :dailyQuestCount"

	expressionAttributeValues := map[string]types.AttributeValue{
		":emptyQuest":      &types.AttributeValueMemberM{Value: emptyQuest},
		":isQuesting":      &types.AttributeValueMemberBOOL{Value: false},
		":questingUntil":   &types.AttributeValueMemberS{Value: ""},
		":lastPlayedDate":  &types.AttributeValueMemberS{Value: utils.GetCurrentDate()},
		":dailyQuestCount": &types.AttributeValueMemberN{Value: strconv.Itoa(user.DailyQuestCount + 1)},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := qm.DB.UpdateItem(qm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}

func (qm QuestModel) CollectCurrentQuestRewards(user *User) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
	}

	emptyQuest := map[string]types.AttributeValue{}
	questKey := "CurrentQuest"
	emptyQuest[questKey] = &types.AttributeValueMemberM{
		Value: map[string]types.AttributeValue{
			constants.NameAttribute: &types.AttributeValueMemberS{
				Value: "Empty Quest 0",
			},
			constants.IDAttribute: &types.AttributeValueMemberS{
				Value: "",
			},
			constants.TimeAttribute: &types.AttributeValueMemberS{
				Value: "",
			},
			constants.EXPAttribute: &types.AttributeValueMemberN{
				Value: "0",
			},
			constants.GoldAttribute: &types.AttributeValueMemberN{
				Value: "0",
			},
		},
	}

	goldReward, _ := strconv.Atoi(user.CurrentQuest["CurrentQuest"].Gold)
	EXPReward, _ := strconv.Atoi(user.CurrentQuest["CurrentQuest"].EXP)

	user.Gold += goldReward
	user.EXP += EXPReward

	updateExpression := "SET " + constants.CurrentQuestAttribute + " = :emptyQuest, " +
		constants.IsQuestingAttribute + " = :isQuesting, " +
		constants.QuestingUntilAttribute + " = :questingUntil, " +
		constants.GoldAttribute + " = :newGold, " +
		constants.EXPAttribute + " = :newEXP, " +
		constants.LastPlayedDateAttribute + " = :lastPlayedDate, " +
		constants.DailyQuestCountAttribute + " = :dailyQuestCount"

	expressionAttributeValues := map[string]types.AttributeValue{
		":emptyQuest": &types.AttributeValueMemberM{
			Value: emptyQuest,
		},
		":isQuesting":      &types.AttributeValueMemberBOOL{Value: false},
		":questingUntil":   &types.AttributeValueMemberS{Value: ""},
		":newGold":         &types.AttributeValueMemberN{Value: strconv.Itoa(user.Gold)},
		":newEXP":          &types.AttributeValueMemberN{Value: strconv.Itoa(user.EXP)},
		":lastPlayedDate":  &types.AttributeValueMemberS{Value: utils.GetCurrentDate()},
		":dailyQuestCount": &types.AttributeValueMemberN{Value: strconv.Itoa(user.DailyQuestCount + 1)},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := qm.DB.UpdateItem(qm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}

func (qm QuestModel) ResetQuests(user *User) error {
	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
	}

	updateExpression := "SET " +
		constants.LastPlayedDateAttribute + " = :lastPlayedDate, " +
		constants.DailyQuestCountAttribute + " = :dailyQuestCount"

	expressionAttributeValues := map[string]types.AttributeValue{
		":lastPlayedDate":  &types.AttributeValueMemberS{Value: utils.GetCurrentDate()},
		":dailyQuestCount": &types.AttributeValueMemberN{Value: "0"},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.TableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := qm.DB.UpdateItem(qm.CTX, input)
	if err != nil {
		return err
	}

	return nil
}
