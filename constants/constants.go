package constants

// General
const (
	EmailRX            = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	AWSRegion          = "eu-central-1"
	TimeFormat         = "2006-01-02T15:04:05"
	TimeFormatJustDate = "2006-01-02T00:00:00"
)

// DB constants
const (
	TableName = "browser_mmo"
	PK        = "PK"
	SK        = "SK"
)

// User DB contsants
const (
	UserPrefix               = "USER#"
	EmailAttribute           = "Email"
	UsernameAttribute        = "Username"
	PasswordHashAttribute    = "PasswordHash"
	CreatedOnAttribute       = "CreatedOn"
	ImageURLAttribute        = "ImageURL"
	LevelAttribute           = "Lvl"
	GoldAttribute            = "Gold"
	TimeAttribute            = "Time"
	EXPAttribute             = "EXP"
	BigDPointsAttribute      = "BigDPoints"
	StrengthAttribute        = "Strength"
	DexterityAttribute       = "Dexterity"
	ConstitutionAttribute    = "Constitution"
	IntelligenceAttribute    = "Intelligence"
	EquippedItemsAttribute   = "EquippedItems"
	WeaponShopAttribute      = "WeaponShop"
	MagicShopAttribute       = "MagicShop"
	MountAttribute           = "Mount"
	MountImageURLAttribute   = "MountImageURL"
	InventoryAttribute       = "Inventory"
	IsQuestingAttribute      = "IsQuesting"
	IsWorkingAttribute       = "IsWorking"
	QuestsAttribute          = "Quests"
	CurrentQuestAttribute    = "CurrentQuest"
	QuestingUntilAttribute   = "QuestingUntil"
	WorkingUntilAttribute    = "WorkingUntil"
	WorkRewardAttribute      = "WorkReward"
	WorkDurationAttribute    = "WorkDuration"
	LastPlayedDateAttribute  = "LastPlayedDate"
	DailyQuestCountAttribute = "DailyQuestCount"
	DungeonAttribute         = "Dungeon"
)

// Item DB constants
const (
	ItemPrefix           = "ITEM#"
	NameAttribute        = "Name"
	BaseNameAttribute    = "BaseName"
	MinLevelAttribute    = "MinLevel"
	IsLegendaryAttribute = "IsLegendary"
	PriceAttribute       = "Price"
	WhatItemAttribute    = "WhatItem"
	Ring                 = "Ring"
	Amulet               = "Amulet"
	Gloves               = "Gloves"
	Boots                = "Boots"
	Helmet               = "Helmet"
	Chestplate           = "Chestplate"
)

// Weapon DB constants
const (
	WeaponPrefix       = "WEAPON#"
	Weapon             = "WEAPON"
	WeaponNotAllCaps   = "Weapon"
	DamageMinAttribute = "DamageMin"
	DamageMaxAttribute = "DamageMax"
)

// Accessory DB constants
const (
	AccessoryPrefix = "ACCESSORY#"
	Accessory       = "ACCESSORY"
)

// Armour DB constants
const (
	ArmourPrefix          = "ARMOUR#"
	Armour                = "ARMOUR"
	ArmourAmountAttribute = "ArmourAmount"
)

// Shield DB constants
const (
	ShieldPrefix         = "SHIELD#"
	Shield               = "SHIELD"
	ShieldNotAllCaps     = "Shield"
	BlockChanceAttribute = "BlockChance"
)

// Quest DB constants
const (
	QuestPrefix = "QUEST#"
	Quest       = "QUEST"
)

// Monster DB constants
const (
	MonsterPrefix = "MONSTER#"
	Monster       = "MONSTER"
)

// Boss DB constants
const (
	BossPrefix        = "BOSS#"
	Boss              = "BOSS"
	PositionAttribute = "Position"
)
