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
		Level:        1,
		Gold:         100,
		EXP:          0,
		BigDPoints:   0,
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		Items: map[string]data.Item{
			"Helmet":     {ID: "0", Name: "Empty Helmet"},
			"Chestplate": {ID: "0", Name: "Empty Chestplate"},
			"Amulet":     {ID: "0", Name: "Empty Amulet"},
			"Gloves":     {ID: "0", Name: "Empty Gloves"},
			"Boots":      {ID: "0", Name: "Empty Boots"},
			"Weapon":     {ID: "0", Name: "Empty Weapon"},
			"Shield":     {ID: "0", Name: "Empty Shield"},
			"Ring":       {ID: "0", Name: "Empty Ring"},
		},
		WeaponShop: map[string]data.Item{
			"Item1": {ID: "0", Name: "Empty Item 1"},
			"Item2": {ID: "0", Name: "Empty Item 2"},
			"Item3": {ID: "0", Name: "Empty Item 3"},
			"Item4": {ID: "0", Name: "Empty Item 4"},
			"Item5": {ID: "0", Name: "Empty Item 5"},
			"Item6": {ID: "0", Name: "Empty Item 6"}}, //TODO: Fill this in with random items later
		MagicShop: map[string]data.Item{
			"Item1": {ID: "0", Name: "Empty Item 1"},
			"Item2": {ID: "0", Name: "Empty Item 2"},
			"Item3": {ID: "0", Name: "Empty Item 3"},
			"Item4": {ID: "0", Name: "Empty Item 4"},
			"Item5": {ID: "0", Name: "Empty Item 5"},
			"Item6": {ID: "0", Name: "Empty Item 6"}}, //TODO: Fill this in with random items later
		Inventory: map[string]data.Item{
			"Item1":  {ID: "0", Name: "Empty Item 1"},
			"Item2":  {ID: "0", Name: "Empty Item 2"},
			"Item3":  {ID: "0", Name: "Empty Item 3"},
			"Item4":  {ID: "0", Name: "Empty Item 4"},
			"Item5":  {ID: "0", Name: "Empty Item 5"},
			"Item6":  {ID: "0", Name: "Empty Item 6"},
			"Item7":  {ID: "0", Name: "Empty Item 7"},
			"Item8":  {ID: "0", Name: "Empty Item 8"},
			"Item9":  {ID: "0", Name: "Empty Item 9"},
			"Item10": {ID: "0", Name: "Empty Item 10"},
			"Item11": {ID: "0", Name: "Empty Item 11"},
			"Item12": {ID: "0", Name: "Empty Item 12"},
			"Item13": {ID: "0", Name: "Empty Item 13"},
			"Item14": {ID: "0", Name: "Empty Item 14"},
			"Item15": {ID: "0", Name: "Empty Item 15"}}, //TODO: Fill this in with random items later
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
	}
}
