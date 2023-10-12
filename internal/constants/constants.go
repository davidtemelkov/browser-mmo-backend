package constants

// General
const (
	EmailRX    = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	AWSRegion  = "eu-central-1"
	TimeFormat = "2006-01-02T15:04:05"
)

// DB constants
const (
	TableName = "browser_mmo"
	PK        = "PK"
	SK        = "SK"
)

// User DB contsants
const (
	UserPrefix             = "USER#"
	EmailAttribute         = "Email"
	UsernameAttribute      = "Username"
	PasswordHashAttribute  = "PasswordHash"
	CreatedOnAttribute     = "CreatedOn"
	ImageURLAttribute      = "ImageURL"
	LevelAttribute         = "Level"
	GoldAttribute          = "Gold"
	EXPAttribute           = "EXP"
	StrengthAttribute      = "Strength"
	DexterityAttribute     = "Dexterity"
	ConstitutionAttribute  = "Constitution"
	IntelligenceAttribute  = "Intelligence"
	ItemsAttribute         = "Items"
	WeaponShopAttribute    = "WeaponShop"
	MagicShopAttribute     = "MagicShop"
	MountAttribute         = "Mount"
	MountImageURLAttribute = "MountImageURL"
)

// Item DB constants
const (
	ItemPrefix           = "ITEM#"
	BaseNameAttribute    = "BaseName"
	MinLevelAttribute    = "MinLevel"
	IsLegendaryAttribute = "IsLegendary"
	TypeAttribute        = "Type"
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
	DamageMinAttribute = "DMGMin"
	DamageMaxAttribute = "DMGMax"
)

// Accessory DB constants
const (
	AccessoryPrefix = "ACCESSORY#"
	Accessory       = "ACCESSORY"
)

// Armour DB constants
const (
	ArmourPrefix             = "ARMOUR#"
	Armour                   = "ARMOUR"
	ArmourAmountMinAttribute = "ArmourAmountMin"
	ArmourAmountMaxAttribute = "ArmourAmountMax"
)

// Shield DB constants
const (
	ShieldPrefix                  = "SHIELD#"
	Shield                        = "SHIELD"
	ShieldBlockChanceMinAttribute = "ShieldMinBlockChance"
	ShieldBlockChanceMaxAttribute = "ShieldMaxBlockChance"
)
