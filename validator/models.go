package validator

import "github.com/MarekSalgovic/hsdeckoder"

type CardAPI struct {
	Id          string   `json:"id"`
	DbfId       int      `json:"dbfId"`
	Name        string   `json:"name"`
	Text        string   `json:"text"`
	Flavor      string   `json:"flavor"`
	Artist      string   `json:"artist"`
	Attack      int      `json:"attack"`
	CardClass   string   `json:"cardClass"`
	Collectible bool     `json:"collectible"`
	Cost        int      `json:"cost"`
	Elite       bool     `json:"elite"`
	Faction     string   `json:"faction"`
	Health      int      `json:"health"`
	Mechanics   []string `json:"mechanics"`
	Rarity      string   `json:"rarity"`
	Set         string   `json:"set"`
	Type        string   `json:"type"`
}

type CardStripped struct {
	Id        string           `json:"id"`
	DbfId     int              `json:"dbfId"`
	Name      string           `json:"name"`
	CardClass hsdeckoder.Class `json:"cardClass"`
	Cost      int              `json:"cost"`
}

type ParsedDeck struct {
	Class hsdeckoder.Class
	Cards []ParsedCard `json:"card"`
}

type ParsedCard struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Cost  int    `json:"cost"`
}
