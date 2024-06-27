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
	Quests      QuestModel
	Work        WorkModel
	Monsters    MonsterModel
	Bosses      BossModel
}

func NewModels(db *dynamodb.Client, ctx context.Context) Models {
	return Models{
		Users:       UserModel{DB: db, CTX: ctx},
		Weapons:     WeaponModel{DB: db, CTX: ctx},
		Accessories: AccessoryModel{DB: db, CTX: ctx},
		Armours:     ArmourModel{DB: db, CTX: ctx},
		Shields:     ShieldModel{DB: db, CTX: ctx},
		Quests:      QuestModel{DB: db, CTX: ctx},
		Work:        WorkModel{DB: db, CTX: ctx},
		Monsters:    MonsterModel{DB: db, CTX: ctx},
		Bosses:      BossModel{DB: db, CTX: ctx},
	}
}
