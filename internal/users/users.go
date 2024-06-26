package users

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/utils"
	"time"
)

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageURL string `json:"imageURL"`
}

func GetInitialUser(input UserInput) *data.User {
	return &data.User{
		Name:         input.Name,
		Email:        input.Email,
		CreatedOn:    time.Now().UTC().Format(constants.TimeFormat),
		ImageURL:     input.ImageURL,
		Lvl:          1,
		Gold:         100,
		EXP:          0,
		BigDPoints:   0,
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		EquippedItems: map[string]data.Item{
			"Helmet":     {Name: "Empty Helmet"},
			"Chestplate": {Name: "Empty Chestplate"},
			"Amulet":     {Name: "Empty Amulet"},
			"Gloves":     {Name: "Empty Gloves"},
			"Boots":      {Name: "Empty Boots"},
			"Weapon":     {Name: "Empty Weapon"},
			"Shield":     {Name: "Empty Shield"},
			"Ring":       {Name: "Empty Ring"},
		},
		WeaponShop: map[string]data.Item{
			"Item1": {Name: "Empty Item 1"},
			"Item2": {Name: "Empty Item 2"},
			"Item3": {Name: "Empty Item 3"},
			"Item4": {Name: "Empty Item 4"},
			"Item5": {Name: "Empty Item 5"},
			"Item6": {Name: "Empty Item 6"}},
		MagicShop: map[string]data.Item{
			"Item1": {Name: "Empty Item 1"},
			"Item2": {Name: "Empty Item 2"},
			"Item3": {Name: "Empty Item 3"},
			"Item4": {Name: "Empty Item 4"},
			"Item5": {Name: "Empty Item 5"},
			"Item6": {Name: "Empty Item 6"}},
		Inventory: map[string]data.Item{
			"Item1":  {Name: "Empty Item 1"},
			"Item2":  {Name: "Empty Item 2"},
			"Item3":  {Name: "Empty Item 3"},
			"Item4":  {Name: "Empty Item 4"},
			"Item5":  {Name: "Empty Item 5"},
			"Item6":  {Name: "Empty Item 6"},
			"Item7":  {Name: "Empty Item 7"},
			"Item8":  {Name: "Empty Item 8"},
			"Item9":  {Name: "Empty Item 9"},
			"Item10": {Name: "Empty Item 10"},
			"Item11": {Name: "Empty Item 11"},
			"Item12": {Name: "Empty Item 12"},
			"Item13": {Name: "Empty Item 13"},
			"Item14": {Name: "Empty Item 14"},
			"Item15": {Name: "Empty Item 15"}},
		Mount:         "",
		MountImageURL: "",
		IsQuesting:    false,
		IsWorking:     false,
		Quests: map[string]data.GeneratedQuest{
			"Quest0": {Name: "Empty Quest 0", ImageURL: "", Time: "", EXP: "", Gold: ""},
			"Quest1": {Name: "Empty Quest 1", ImageURL: "", Time: "", EXP: "", Gold: ""},
			"Quest2": {Name: "Empty Quest 2", ImageURL: "", Time: "", EXP: "", Gold: ""},
		},
		CurrentQuest: map[string]data.GeneratedQuest{
			"CurrentQuest": {Name: "Empty Quest 0", ImageURL: "", Time: "", EXP: "", Gold: ""}},
		QuestingUntil:   "",
		WorkingUntil:    "",
		WorkReward:      0,
		WorkDuration:    0,
		LastPlayedDate:  utils.GetCurrentDate(),
		DailyQuestCount: 0,
		ArmourAmount:    0,
		Dungeon:         1,
	}
}

func CalculateTotalStats(user data.User) data.User {
	// TODO: User really should be keeping these, maybe on item equip and stat upgrade
	user.TotalStrength += user.Strength
	user.TotalDexterity += user.Dexterity
	user.TotalConstitution += user.Constitution
	user.TotalIntelligence += user.Intelligence

	for _, item := range user.EquippedItems {
		user.TotalStrength += item.Strength
		user.TotalDexterity += item.Dexterity
		user.TotalConstitution += item.Constitution
		user.TotalIntelligence += item.Intelligence
		user.BlockChance += item.BlockChance
		user.ArmourAmount += item.ArmourAmount
		user.DamageMin += item.DamageMin
		user.DamageMax += item.DamageMax
		user.DamageAverage += item.DamageAverage
	}

	return user
}
