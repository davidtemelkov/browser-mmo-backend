package data

type Item struct {
	ID           string `json:"id"`
	WhatItem     string `json:"whatItem"`
	Name         string `json:"name"`
	BaseName     string `json:"baseName"`
	Lvl          int    `json:"lvl"`
	DamageMin    int    `json:"damageMin"`
	DamageMax    int    `json:"damageMax"`
	Strength     int    `json:"strength"`
	Dexterity    int    `json:"dexterity"`
	Constitution int    `json:"constitution"`
	Intelligence int    `json:"intelligence"`
	ArmourAmount int    `json:"armourAmount"`
	BlockChance  int    `json:"blockChance"`
	IsLegendary  bool   `json:"isLegendary"`
	Price        int    `json:"price"`
}
