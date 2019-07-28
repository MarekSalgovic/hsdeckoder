package validator

import (
	"encoding/json"
	"github.com/MarekSalgovic/hsdeckoder"
	"io/ioutil"
	"net/http"
	"os"
)



const (
	deckSize = 30
)

const (
	apiURL = "https://api.hearthstonejson.com/v1/31532/enUS/cards.json"
	dbPath = "./database.json"
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


func Parse(deck hsdeckoder.Deck) (ParsedDeck, error){
	var parsedDeck ParsedDeck
	for _, card := range deck.Cards{
		strippedCard, err := getCard(card.Id)
		if err != nil{
			return ParsedDeck{}, err
		}
		parsedCard := ParsedCard{
			Id:    strippedCard.Id,
			Name:  strippedCard.Name,
			Count: card.Count,
			Cost: strippedCard.Cost,
		}
		parsedDeck.Cards = append(parsedDeck.Cards, parsedCard)
	}
	return parsedDeck, nil
}


