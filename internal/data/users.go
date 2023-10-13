package data

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/validator"
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name          string            `json:"name" dynamodbav:"Username"`
	Email         string            `json:"email" dynamodbav:"Email"`
	Password      Password          `json:"-"`
	PasswordHash  string            `json:"-" dynamodbav:"PasswordHash"`
	CreatedOn     string            `json:"createdOn,omitempty" dynamodbav:"CreatedOn"`
	ImageURL      string            `json:"imageURL" dynamodbav:"ImageURL"`
	Level         int               `json:"level" dynamodbav:"Level"`
	Gold          int               `json:"gold" dynamodbav:"Gold"`
	EXP           int               `json:"EXP" dynamodbav:"EXP"`
	BigDPoints    int               `json:"bigDPoints" dynamodbav:"BigDPoints"`
	Strength      int               `json:"strength" dynamodbav:"Strength"`
	Dexterity     int               `json:"dexterity" dynamodbav:"Dexterity"`
	Constitution  int               `json:"constitution" dynamodbav:"Constitution"`
	Intelligence  int               `json:"intelligence" dynamodbav:"Intelligence"`
	Items         map[string]string `json:"items" dynamodbav:"Items"`
	WeaponShop    map[string]string `json:"weaponShop" dynamodbav:"WeaponShop"`
	MagicShop     map[string]string `json:"magicShop" dynamodbav:"MagicShop"`
	Mount         string            `json:"mount" dynamodbav:"Mount"`
	MountImageURL string            `json:"mountImageURL" dynamodbav:"MountImageURL"`
	Inventory     map[string]string `json:"inventory" dynamodbav:"Inventory"`
	IsQuesting    bool              `json:"isQuesting" dynamodbav:"IsQuesting"`
	IsWorking     bool              `json:"isWorking" dynamodbav:"IsWorking"`
}

type Password struct {
	plaintext *string
	hash      []byte
}

func (p *Password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}

func (p *Password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", constants.RequiredFieldError.Error())
	v.Check(validator.Matches(email, validator.EmailRX), "email", constants.EmailFormatError.Error())
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", constants.RequiredFieldError.Error())
	v.Check(len(password) >= 8, "password", constants.PasswordMinLengthError.Error())
	v.Check(len(password) <= 72, "password", constants.PasswordMaxLengthError.Error())
}

func ValidateRegisterInput(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", constants.RequiredFieldError.Error())
	v.Check(len(user.Name) >= 4, "name", constants.UserNameMinLengthError.Error())
	v.Check(len(user.Name) <= 50, "name", constants.UserNameMaxLengthError.Error())

	ValidateEmail(v, user.Email)
	ValidatePasswordPlaintext(v, *user.Password.plaintext)
}

func ValidateLoginInput(v *validator.Validator, email, password string) {
	ValidateEmail(v, email)
	ValidatePasswordPlaintext(v, password)
}

type UserModel struct {
	DB  *dynamodb.Client
	CTX context.Context
}

func (um UserModel) Insert(user *User) error {
	item := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + user.Email,
		},
		constants.UsernameAttribute: &types.AttributeValueMemberS{
			Value: user.Name,
		},
		constants.EmailAttribute: &types.AttributeValueMemberS{
			Value: user.Email,
		},
		constants.PasswordHashAttribute: &types.AttributeValueMemberS{
			Value: string(user.Password.hash),
		},
		constants.CreatedOnAttribute: &types.AttributeValueMemberS{
			Value: user.CreatedOn,
		},
		constants.ImageURLAttribute: &types.AttributeValueMemberS{
			Value: user.ImageURL,
		}, constants.LevelAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.Level),
		},
		constants.GoldAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.Gold),
		},
		constants.EXPAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.EXP),
		},
		constants.BigDPointsAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.BigDPoints),
		},
		constants.StrengthAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.Strength),
		},
		constants.DexterityAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.Dexterity),
		},
		constants.ConstitutionAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.Constitution),
		},
		constants.IntelligenceAttribute: &types.AttributeValueMemberN{
			Value: strconv.Itoa(user.Intelligence),
		},
		constants.ItemsAttribute: &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{},
		},
		constants.WeaponShopAttribute: &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{},
		},
		constants.MagicShopAttribute: &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{},
		},
		constants.InventoryAttribute: &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{},
		},
		constants.MountAttribute: &types.AttributeValueMemberS{
			Value: user.Mount,
		},
		constants.MountImageURLAttribute: &types.AttributeValueMemberS{
			Value: user.MountImageURL,
		},
		constants.IsQuestingAttribute: &types.AttributeValueMemberBOOL{
			Value: user.IsQuesting,
		},
		constants.IsWorkingAttribute: &types.AttributeValueMemberBOOL{
			Value: user.IsWorking,
		},
	}

	for key, value := range user.Items {
		item[constants.ItemsAttribute].(*types.AttributeValueMemberM).Value[key] = &types.AttributeValueMemberS{
			Value: value,
		}
	}

	for key, value := range user.WeaponShop {
		item[constants.WeaponShopAttribute].(*types.AttributeValueMemberM).Value[key] = &types.AttributeValueMemberS{
			Value: value,
		}
	}

	for key, value := range user.MagicShop {
		item[constants.MagicShopAttribute].(*types.AttributeValueMemberM).Value[key] = &types.AttributeValueMemberS{
			Value: value,
		}
	}

	for key, value := range user.Inventory {
		item[constants.InventoryAttribute].(*types.AttributeValueMemberM).Value[key] = &types.AttributeValueMemberS{
			Value: value,
		}
	}

	putInput := &dynamodb.PutItemInput{
		TableName:           aws.String(constants.TableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	}

	_, err := um.DB.PutItem(um.CTX, putInput)
	if err != nil {
		if strings.Contains(err.Error(), "The conditional request failed") {
			return constants.DuplicateEmailError
		}
		return err
	}

	return nil
}

func (um UserModel) Get(email string) (*User, error) {
	if email == "" {
		return nil, constants.UserNotFoundError
	}

	key := map[string]types.AttributeValue{
		constants.PK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + email,
		},
		constants.SK: &types.AttributeValueMemberS{
			Value: constants.UserPrefix + email,
		},
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(constants.TableName),
		Key:       key,
	}

	result, err := um.DB.GetItem(um.CTX, input)
	if err != nil {
		return nil, err
	}

	if len(result.Item) == 0 {
		return nil, constants.UserNotFoundError
	}

	user := &User{}
	if err := attributevalue.UnmarshalMap(result.Item, user); err != nil {
		return nil, err
	}

	user.Password = Password{
		hash: []byte(user.PasswordHash),
	}

	return user, nil
}

func (um UserModel) CanLoginUser(password string, user *User) (bool, error) {
	passwordIsCorrect, err := user.Password.Matches(password)
	if err != nil || !passwordIsCorrect {
		return false, err
	}

	return true, nil
}
