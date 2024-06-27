package data

type Item struct {
	WhatItem     string `json:"whatItem"`
	Name         string `json:"name"`
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
	ImageURL     string `json:"imageURL"`
	Price        int    `json:"price"`
}
