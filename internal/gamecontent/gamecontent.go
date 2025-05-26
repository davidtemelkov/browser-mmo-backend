package gamecontent

import (
	"browser-mmo-backend/internal/constants"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Weapon struct {
	ID          string
	BaseName    string
	MinLevel    int
	IsLegendary bool
}

// TODO: Add aromour amount to all armours
type Armour struct {
	ID          string
	WhatItem    string
	BaseName    string
	MinLevel    int
	IsLegendary bool
}

type Accessory struct {
	ID          string
	WhatItem    string
	BaseName    string
	MinLevel    int
	IsLegendary bool
}

type Shield struct {
	ID          string
	BaseName    string
	MinLevel    int
	IsLegendary bool
}

type Quest struct {
	ID   string
	Name string
}

type Monster struct {
	ID   string
	Name string
}

// TODO: Differentiate between domain and api objects
// TODO: Add exp and gold reward
type Boss struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Position     int    `json:"position"`
	Constitution int    `json:"constitution"`
	Dexterity    int    `json:"dexterity"`
	Intelligence int    `json:"intelligence"`
	Strength     int    `json:"strength"`
	Lvl          int    `json:"lvl"`
}

// GameContent holds all static game data in memory
type GameContent struct {
	Weapons     []Weapon
	Armours     []Armour
	Accessories []Accessory
	Shields     []Shield
	Quests      []Quest
	Monsters    []Monster
	Bosses      []Boss
}

// LoadGameContent loads all static content from JSON files into memory
func LoadGameContent(contentDir string) (*GameContent, error) {
	gc := &GameContent{}

	// Helper to load a JSON file into a target
	load := func(filename string, target interface{}) error {
		f, err := os.Open(filepath.Join(contentDir, filename))
		if err != nil {
			return err
		}
		defer f.Close()
		return json.NewDecoder(f).Decode(target)
	}

	if err := load("weapons.json", &gc.Weapons); err != nil {
		return nil, err
	}
	if err := load("armours.json", &gc.Armours); err != nil {
		return nil, err
	}
	if err := load("accessories.json", &gc.Accessories); err != nil {
		return nil, err
	}
	if err := load("shields.json", &gc.Shields); err != nil {
		return nil, err
	}
	if err := load("quests.json", &gc.Quests); err != nil {
		return nil, err
	}
	if err := load("monsters.json", &gc.Monsters); err != nil {
		return nil, err
	}
	if err := load("bosses.json", &gc.Bosses); err != nil {
		return nil, err
	}

	return gc, nil
}

func (gc GameContent) GetBossByPosition(position int) (Boss, error) {
	for _, boss := range gc.Bosses {
		if boss.Position == position {
			return boss, nil
		}
	}
	return Boss{}, errors.New(constants.RecordNotFoundError)
}
