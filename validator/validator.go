package validator

import (
	"encoding/json"
	"errors"
	"github.com/MarekSalgovic/hsdeckoder"
	"io/ioutil"
	"net/http"
	"os"
)

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

const (
	deckSize = 30
)

const (
	apiURL = "https://api.hearthstonejson.com/v1/31532/enUS/cards.json"
	dbPath = "./database.json"
)

var (
	ErrDownloadFailed = errors.New("database download failed")
	ErrDatabaseWrite  = errors.New("database write failed")
	ErrDatabaseRead = errors.New("database read failed")
	ErrCardNotFound = errors.New("card not found")
	ErrInvalidDeck = errors.New("deck invalid")
)


func stripCard(card CardAPI) CardStripped{
	return CardStripped{
		Id:        card.Id,
		DbfId:     card.DbfId,
		Name:      card.Name,
		CardClass: Class(card.CardClass),
		Cost:      card.Cost,
	}
}


func downloadDB() ([]CardStripped, error) {
	var cards []CardStripped
	res, err := http.Get(apiURL)
	if err != nil{
		return cards, ErrDownloadFailed
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return cards, ErrDownloadFailed
	}

	err = json.Unmarshal(resData,&cards)
	if err != nil{
		return cards, ErrDownloadFailed
	}
	file, err := json.MarshalIndent(cards,"","  ")
	if err != nil{
		return cards, ErrDatabaseWrite
	}
	err = ioutil.WriteFile(dbPath,file,0644)
	if err != nil{
		return cards, ErrDatabaseWrite
	}
	return cards, nil
}

func readDB() ([]CardStripped, error){
	var cards []CardStripped
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		cards, err = downloadDB()
		if err != nil{
			return cards, err
		}
		return cards, nil
	}
	jsonFile, err := os.Open(dbPath)
	if err != nil {
		return cards, ErrDatabaseRead
	}
	defer jsonFile.Close()
	bs, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return cards, ErrDatabaseRead
	}
	err = json.Unmarshal(bs, &cards)
	if err != nil {
		return cards, ErrDatabaseRead
	}
	return cards,nil
}




func getClass(deck hsdeckoder.Deck) (Class, error){
	heroCard, err := getCard(deck.Heroes[0])
	if err != nil{
		return Class(""), err
	}
	class := heroCard.CardClass
	for _, id := range deck.Cards{
		card, err := getCard(id.Id)
		if err != nil{
			return Class(""), err
		}
		if card.CardClass != NEUTRAL && card.CardClass != class{
			return Class(""), ErrInvalidDeck
		}
	}
	return class, nil

}

func getCard(dbfId int) (CardStripped, error){
	cards, err := readDB()
	if err != nil{
		return CardStripped{}, nil
	}
	for _, card := range cards{
		if card.DbfId == dbfId{
			return card, nil
		}
	}
	return CardStripped{}, ErrCardNotFound
}


func Validate(deck hsdeckoder.Deck) (Class, error){
	if len(deck.Heroes) != 1{
		return Class(""), ErrInvalidDeck
	}
	class, err := getClass(deck)
	if err != nil{
		return Class(""), err
	}
	var cardCount int
	for _, card := range deck.Cards{
		cardCount += card.Count
	}
	if cardCount != deckSize{
		return Class(""), ErrInvalidDeck
	}
	return class, nil
}



