package items

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"math/rand"
)

var ITEM_TYPES = []string{
	constants.Helmet,
	constants.Chestplate,
	constants.Amulet,
	constants.Gloves,
	constants.Boots,
	constants.Weapon,
	constants.Shield,
	constants.Ring,
}

func GenerateItem(wp data.WeaponModel, acm data.AccessoryModel, sm data.ShieldModel, arm data.ArmourModel) error {
	randIndex := rand.Intn(len(ITEM_TYPES))
	selectedItemType := ITEM_TYPES[randIndex]
	var statType int
	var suffix string

	// TODO: Other legendary chance logic
	isLegendary := rand.Intn(25) == 0

	if !isLegendary {
		statType = rand.Intn(4)

		switch statType {
		case 3:
			suffix = " Of Threesomes"
		case 4:
			suffix = " Of All Trades"
		default:
			statType = 2
			suffix = " Of Deuces"
		}
	}

	switch selectedItemType {
	case constants.Weapon:
		var baseWeapon data.Weapon

		if isLegendary {
			legendaryWeapons, err := wp.GetAllLegendaryWeapons()
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(legendaryWeapons))
			baseWeapon = legendaryWeapons[randIndex]
		} else {
			basicWeapons, err := wp.GetAllBasicWeapons()
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(basicWeapons))
			baseWeapon = basicWeapons[randIndex]
		}

		// TODO: Add different generation logic based upon isLegendary and user stats
		generatedWeapon := data.GeneratedWeapon{
			Name:          baseWeapon.BaseName + suffix,
			Level:         1,
			DamageMin:     1,
			DamageMax:     1,
			DamageAverage: 1,
			Strength:      0,
			Dexterity:     0,
			Constitution:  0,
			Intelligence:  0,
			IsLegendary:   isLegendary,
			ImageURL:      baseWeapon.ImageURL,
			Price:         1,
		}

		// TODO: Add generated item to user's inventory/store

	case constants.Shield:
		var baseShield data.Shield

		if isLegendary {
			legendaryShields, err := sm.GetAllLegendaryShields()
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(legendaryShields))
			baseShield = legendaryShields[randIndex]
		} else {
			basicShields, err := sm.GetAllBasicShields()
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(basicShields))
			baseShield = basicShields[randIndex]
		}

		// TODO: Add different generation logic based upon isLegendary and user stats
		generatedShield := data.GeneratedShield{
			Name:         baseShield.BaseName + suffix,
			Level:        1,
			BlockChance:  1,
			Strength:     0,
			Dexterity:    0,
			Constitution: 0,
			Intelligence: 0,
			IsLegendary:  isLegendary,
			ImageURL:     baseShield.ImageURL,
			Price:        1,
		}

		// TODO: Add generated item to user's inventory/store

	case constants.Accessory:
		var baseAccessory data.Accessory

		if isLegendary {
			legendaryAccessories, err := acm.GetAllLegendaryAccessoriesOfType(selectedItemType)
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(legendaryAccessories))
			baseAccessory = legendaryAccessories[randIndex]
		} else {
			basicAccessories, err := acm.GetAllBasicAccessoriesOfType(selectedItemType)
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(basicAccessories))
			baseAccessory = basicAccessories[randIndex]
		}

		// TODO: Add different generation logic based upon isLegendary and user stats
		generatedAccessory := data.GeneratedAccessory{
			Type:         selectedItemType,
			Name:         baseAccessory.BaseName + suffix,
			Level:        1,
			Strength:     0,
			Dexterity:    0,
			Constitution: 0,
			Intelligence: 0,
			IsLegendary:  isLegendary,
			ImageURL:     baseAccessory.ImageURL,
			Price:        1,
		}

		// TODO: Add generated item to user's inventory/store

	default:
		var baseArmour data.Armour

		if isLegendary {
			legendaryArmours, err := arm.GetAllLegendaryArmoursOfType(selectedItemType)
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(legendaryArmours))
			baseArmour = legendaryArmours[randIndex]
		} else {
			basicArmours, err := arm.GetAllBasicArmoursOfType(selectedItemType)
			if err != nil {
				return err
			}
			randIndex = rand.Intn(len(basicArmours))
			baseArmour = basicArmours[randIndex]
		}

		// TODO: Add different generation logic based upon isLegendary and user stats
		generatedArmour := data.GeneratedAccessory{
			Type:         selectedItemType,
			Name:         baseArmour.BaseName + suffix,
			Level:        1,
			Strength:     0,
			Dexterity:    0,
			Constitution: 0,
			Intelligence: 0,
			IsLegendary:  isLegendary,
			ImageURL:     baseArmour.ImageURL,
			Price:        1,
		}

		// TODO: Add generated item to user's inventory/store
	}

	return nil
}
