package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Models struct {
	Users       UserModel
	Weapons     WeaponModel
	Accessories AccessoryModel
	Armours     ArmourModel
	Shields     ShieldModel
}

func NewModels(db *dynamodb.Client, ctx context.Context) Models {
	return Models{
		Users:       UserModel{DB: db, CTX: ctx},
		Weapons:     WeaponModel{DB: db, CTX: ctx},
		Accessories: AccessoryModel{DB: db, CTX: ctx},
		Armours:     ArmourModel{DB: db, CTX: ctx},
		Shields:     ShieldModel{DB: db, CTX: ctx},
	}
}
