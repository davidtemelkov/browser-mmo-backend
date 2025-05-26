package items

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/gamecontent"
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
	weapons []gamecontent.Weapon,
	accessories []gamecontent.Accessory,
	shields []gamecontent.Shield,
	armours []gamecontent.Armour,
	um data.UserModel,
) (data.Item, error) {
	randIndex := rand.Intn(len(ITEM_TYPES))
	selectedItemType := ITEM_TYPES[randIndex]

	var (
		id          string
		statType    int
		baseName    string
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
		var baseWeapon gamecontent.Weapon
		// TODO: Rethink this, can store leg and non leg items seperately in memory
		var legendaryWeapons, basicWeapons []gamecontent.Weapon
		for _, w := range weapons {
			if w.IsLegendary {
				legendaryWeapons = append(legendaryWeapons, w)
			} else {
				basicWeapons = append(basicWeapons, w)
			}
		}
		if isLegendary {
			randIndex = rand.Intn(len(legendaryWeapons))
			baseWeapon = legendaryWeapons[randIndex]
		} else {
			randIndex = rand.Intn(len(basicWeapons))
			baseWeapon = basicWeapons[randIndex]
		}
		baseName = baseWeapon.BaseName
		id = baseWeapon.ID
	case constants.ShieldNotAllCaps:
		var baseShield gamecontent.Shield
		var legendaryShields, basicShields []gamecontent.Shield
		for _, s := range shields {
			if s.IsLegendary {
				legendaryShields = append(legendaryShields, s)
			} else {
				basicShields = append(basicShields, s)
			}
		}
		if isLegendary {
			randIndex = rand.Intn(len(legendaryShields))
			baseShield = legendaryShields[randIndex]
		} else {
			randIndex = rand.Intn(len(basicShields))
			baseShield = basicShields[randIndex]
		}
		baseName = baseShield.BaseName
		id = baseShield.ID
	case constants.Ring:
		var baseAccessory gamecontent.Accessory
		var legendaryAccessories, basicAccessories []gamecontent.Accessory
		for _, a := range accessories {
			if a.WhatItem == selectedItemType {
				if a.IsLegendary {
					legendaryAccessories = append(legendaryAccessories, a)
				} else {
					basicAccessories = append(basicAccessories, a)
				}
			}
		}
		if isLegendary {
			randIndex = rand.Intn(len(legendaryAccessories))
			baseAccessory = legendaryAccessories[randIndex]
		} else {
			randIndex = rand.Intn(len(basicAccessories))
			baseAccessory = basicAccessories[randIndex]
		}
		baseName = baseAccessory.BaseName
		id = baseAccessory.ID
	case constants.Amulet:
		var baseAccessory gamecontent.Accessory
		var legendaryAccessories, basicAccessories []gamecontent.Accessory
		for _, a := range accessories {
			if a.WhatItem == selectedItemType {
				if a.IsLegendary {
					legendaryAccessories = append(legendaryAccessories, a)
				} else {
					basicAccessories = append(basicAccessories, a)
				}
			}
		}
		if isLegendary {
			randIndex = rand.Intn(len(legendaryAccessories))
			baseAccessory = legendaryAccessories[randIndex]
		} else {
			randIndex = rand.Intn(len(basicAccessories))
			baseAccessory = basicAccessories[randIndex]
		}
		baseName = baseAccessory.BaseName
		id = baseAccessory.ID
	default:
		var baseArmour gamecontent.Armour
		var legendaryArmours, basicArmours []gamecontent.Armour
		for _, a := range armours {
			if a.WhatItem == selectedItemType {
				if a.IsLegendary {
					legendaryArmours = append(legendaryArmours, a)
				} else {
					basicArmours = append(basicArmours, a)
				}
			}
		}
		if isLegendary {
			randIndex = rand.Intn(len(legendaryArmours))
			baseArmour = legendaryArmours[randIndex]
		} else {
			randIndex = rand.Intn(len(basicArmours))
			baseArmour = basicArmours[randIndex]
		}
		baseName = baseArmour.BaseName
		id = baseArmour.ID
	}

	item := generateStats(user, statType, id, selectedItemType, baseName, isLegendary)

	return item, nil
}

// TODO: Balance this
func generateStats(user data.User, statType int, id, itemType, baseName string, isLegendary bool) data.Item {
	item := data.Item{
		ID:          id,
		BaseName:    baseName,
		WhatItem:    itemType,
		Name:        getItemName(isLegendary, baseName, statType),
		Lvl:         user.Lvl,
		IsLegendary: isLegendary,
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
