package validator


type CardAPI struct {
	Id          string   `json:"id"`
	DbfId       int      `json:"dbfId"`
	Name        string   `json:"name"`
	Text        string   `json:"text"`
	Flavor      string	 `json:"flavor"`
	Artist      string   `json:"artist"`
	Attack      int  	 `json:"attack"`
	CardClass   string   `json:"cardClass"`
	Collectible bool  	 `json:"collectible"`
	Cost        int      `json:"cost"`
	Elite       bool     `json:"elite"`
	Faction     string   `json:"faction"`
	Health      int      `json:"health"`
	Mechanics   []string `json:"mechanics"`
	Rarity		string	 `json:"rarity"`
	Set			string	 `json:"set"`
	Type		string	 `json:"type"`
}

type CardStripped struct {
	Id          string   `json:"id"`
	DbfId       int      `json:"dbfId"`
	Name        string   `json:"name"`
	CardClass   Class    `json:"cardClass"`
	Cost        int      `json:"cost"`
}


type ParsedDeck struct{
	Cards []ParsedCard `json:"card"`
}

type ParsedCard struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Count int `json:"count"`
}


type Class string

const (
	NEUTRAL Class = "NEUTRAL"
	DRUID Class = "DRUID"
	HUNTER Class = "HUNTER"
	MAGE Class = "MAGE"
	PALADIN Class = "PALADIN"
	PRIEST Class = "PRIEST"
	ROGUE Class = "ROGUE"
	SHAMAN Class = "SHAMAN"
	WARLOCK Class = "WARLOCK"
	WARRIOR Class = "WARRIOR"
)