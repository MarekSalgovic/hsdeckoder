package validator

import "github.com/MarekSalgovic/hsdeckoder"

type CardAPI struct {
	Id          string   `json:"id"`
	DbfId       int      `json:"dbfId"`
	Name        string   `json:"name"`
	Text        string   `json:"text, omitempty"`
	Flavor      string   `json:"flavor, omitempty"`
	Artist      string   `json:"artist, omitempty"`
	Attack      int      `json:"attack, omitempty"`
	CardClass   string   `json:"cardClass, omitempty"`
	Collectible bool     `json:"collectible, omitempty"`
	Cost        int      `json:"cost, omitempty"`
	Elite       bool     `json:"elite, omitempty"`
	Faction     string   `json:"faction, omitempty"`
	Health      int      `json:"health, omitempty"`
	Mechanics   []string `json:"mechanics, omitempty"`
	Rarity      string   `json:"rarity, omitempty"`
	Set         string   `json:"set, omitempty"`
	Type        string   `json:"type, omitempty"`
}

type CardStripped struct {
	Id        string           `json:"id"`
	DbfId     int              `json:"dbfId"`
	Name      string           `json:"name"`
	CardClass hsdeckoder.Class `json:"cardClass, omitempty"`
	Cost      int              `json:"cost, omitempty"`
	Rarity    string           `json:"rarity, omitempty"`
}

type ParsedDeck struct {
	Class hsdeckoder.Class `json:"class"`
	Cards []ParsedCard     `json:"card"`
}

type ParsedCard struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
	Cost   int    `json:"cost"`
	Rarity string `json:"rarity"`
}
