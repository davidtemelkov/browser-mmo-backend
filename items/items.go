package items

import (
	"browser-mmo-backend/constants"
	"browser-mmo-backend/data"
	"math"
	"math/rand"
)

var ITEM_TYPES = []string{
	constants.Helmet,
	constants.Chestplate,
	constants.Amulet,
	constants.Gloves,
	constants.Boots,
	constants.WeaponNotAllCaps,
	constants.ShieldNotAllCaps,
	constants.Ring,
}

var STATS = []string{
	constants.StrengthAttribute,
	constants.DexterityAttribute,
	constants.ConstitutionAttribute,
	constants.IntelligenceAttribute,
}

func GenerateItem(
	isGuaranteedLegendary bool,
	user data.User,
	wp data.WeaponModel,
	acm data.AccessoryModel,
	sm data.ShieldModel,
	arm data.ArmourModel,
	um data.UserModel,
) (data.Item, error) {
	randIndex := rand.Intn(len(ITEM_TYPES))
	selectedItemType := ITEM_TYPES[randIndex]

	var (
		statType    int
		baseName    string
		imageURL    string
		isLegendary bool
	)

	if isGuaranteedLegendary {
		isLegendary = true
	} else {
		isLegendary = rand.Intn(25) == 0
	}

	// TODO: Use a different rand func to not get 0s and 1s, only 2,3,4
	if isLegendary {
		statType = rand.Intn(3)
	} else {
		statType = rand.Intn(4)
	}

	switch selectedItemType {
	case constants.WeaponNotAllCaps:
		var baseWeapon data.Weapon

		if isLegendary {
			legendaryWeapons, err := wp.GetAllLegendaryWeapons()
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(legendaryWeapons))
			baseWeapon = legendaryWeapons[randIndex]
		} else {
			basicWeapons, err := wp.GetAllBasicWeapons()
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(basicWeapons))
			baseWeapon = basicWeapons[randIndex]
		}

		baseName = baseWeapon.BaseName
		imageURL = baseWeapon.ImageURL
	case constants.ShieldNotAllCaps:
		var baseShield data.Shield

		if isLegendary {
			legendaryShields, err := sm.GetAllLegendaryShields()
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(legendaryShields))
			baseShield = legendaryShields[randIndex]
		} else {
			basicShields, err := sm.GetAllBasicShields()
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(basicShields))
			baseShield = basicShields[randIndex]
		}

		baseName = baseShield.BaseName
		imageURL = baseShield.ImageURL
	case constants.Ring:
		var baseAccessory data.Accessory

		if isLegendary {
			legendaryAccessories, err := acm.GetAllLegendaryAccessoriesOfType(selectedItemType)
			if err != nil {
				return data.Item{}, nil
			}
			randIndex = rand.Intn(len(legendaryAccessories))
			baseAccessory = legendaryAccessories[randIndex]
		} else {
			basicAccessories, err := acm.GetAllBasicAccessoriesOfType(selectedItemType)
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(basicAccessories))
			baseAccessory = basicAccessories[randIndex]
		}

		baseName = baseAccessory.BaseName
		imageURL = baseAccessory.ImageURL
	case constants.Amulet:
		var baseAccessory data.Accessory

		if isLegendary {
			legendaryAccessories, err := acm.GetAllLegendaryAccessoriesOfType(selectedItemType)
			if err != nil {
				return data.Item{}, nil
			}
			randIndex = rand.Intn(len(legendaryAccessories))
			baseAccessory = legendaryAccessories[randIndex]
		} else {
			basicAccessories, err := acm.GetAllBasicAccessoriesOfType(selectedItemType)
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(basicAccessories))
			baseAccessory = basicAccessories[randIndex]
		}

		baseName = baseAccessory.BaseName
		imageURL = baseAccessory.ImageURL
	default:
		var baseArmour data.Armour

		if isLegendary {
			legendaryArmours, err := arm.GetAllLegendaryArmoursOfType(selectedItemType)
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(legendaryArmours))
			baseArmour = legendaryArmours[randIndex]
		} else {
			basicArmours, err := arm.GetAllBasicArmoursOfType(selectedItemType)
			if err != nil {
				return data.Item{}, err
			}
			randIndex = rand.Intn(len(basicArmours))
			baseArmour = basicArmours[randIndex]
		}

		baseName = baseArmour.BaseName
		imageURL = baseArmour.ImageURL
	}

	item := generateStats(user, statType, selectedItemType, baseName, imageURL, isLegendary)

	return item, nil
}

// TODO: Balance this
func generateStats(user data.User, statType int, itemType, baseName, imageURL string, isLegendary bool) data.Item {
	item := data.Item{
		WhatItem:    itemType,
		Name:        getItemName(isLegendary, baseName, statType),
		Lvl:         user.Lvl,
		IsLegendary: isLegendary,
		ImageURL:    imageURL,
	}

	//TODO: Remove this after implementing different rand func
	if statType == 0 || statType == 1 {
		statType = 2
	}

	// Generate Str, Dex, Const, Int
	selectedStats := rand.Perm(len(STATS))[:statType]

	baseValue := user.Lvl * 10
	randomness := rand.Intn(10)
	totalStatValue := baseValue + randomness

	values := make([]int, statType)
	switch statType {
	case 3:
		values[0] = int(math.Round(float64(totalStatValue) * 0.3))
		values[1] = int(math.Round(float64(totalStatValue) * 0.3))
		values[2] = int(math.Round(float64(totalStatValue) * 0.3))
	case 4:
		values[0] = int(math.Round(float64(totalStatValue) * 0.1))
		values[1] = int(math.Round(float64(totalStatValue) * 0.1))
		values[2] = int(math.Round(float64(totalStatValue) * 0.1))
		values[3] = int(math.Round(float64(totalStatValue) * 0.1))
	case 2:
		values[0] = int(math.Round(float64(totalStatValue) * 0.5))
		values[1] = int(math.Round(float64(totalStatValue) * 0.5))
	}

	for i, stat := range selectedStats {
		switch stat {
		case 0:
			item.Strength = values[i]
		case 1:
			item.Dexterity = values[i]
		case 2:
			item.Constitution = values[i]
		case 3:
			item.Intelligence = values[i]
		}
	}

	// Generate item type unique stats
	switch itemType {
	case constants.WeaponNotAllCaps:
		// add dmgmin max avrg
		item.DamageMin = user.Lvl
		item.DamageMax = user.Lvl * 2
	case constants.ShieldNotAllCaps:
		item.BlockChance = user.Lvl
	case constants.Ring:
		// do nothing
	case constants.Amulet:
		// do nothing
	default:
		// add armour amount
		item.ArmourAmount = user.Lvl * 2
	}

	item.Price = user.Lvl * statType
	if isLegendary {
		item.Price *= 2
	}

	return item
}

func getItemName(isLegendary bool, baseName string, statType int) string {
	if isLegendary {
		return baseName
	}

	var suffix string

	switch statType {
	case 3:
		suffix = " Of Threesomes"
	case 4:
		suffix = " Of All Trades"
	default:
		statType = 2
		suffix = " Of Deuces"
	}

	return baseName + suffix
}
