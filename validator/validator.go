package validator

import (
	"encoding/json"
	"github.com/MarekSalgovic/hsdeckoder"
	"io/ioutil"
	"net/http"
)

const (
	deckSize = 30
)

const (
	apiURL = "https://api.hearthstonejson.com/v1/latest/enUS/cards.json"
)

type Validator interface {
	ValidateDeck(deck hsdeckoder.Deck) (ParsedDeck, error)
	getCard(dbfId int) (CardStripped, error)
	getClass(deck hsdeckoder.Deck) (hsdeckoder.Class, error)
}

type Validate struct {
	Cards []CardStripped
}

func NewValidator() (*Validate, error) {
	cards, err := downloadDB()
	if err != nil {
		return &Validate{}, err
	}
	return &Validate{
		Cards: cards,
	}, nil
}

func downloadDB() ([]CardStripped, error) {
	var cards []CardStripped
	res, err := http.Get(apiURL)
	if err != nil {
		return cards, err
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return cards, err
	}

	err = json.Unmarshal(resData, &cards)
	if err != nil {
		return cards, err
	}
	/*
		file, err := json.MarshalIndent(cards, "", "  ")
		if err != nil {
			return cards, ErrDatabaseWrite
		}
		err = ioutil.WriteFile(dbPath, file, 0644)
		if err != nil {
			return cards, ErrDatabaseWrite
		}*/
	return cards, nil
}

/*
func readDB() ([]CardStripped, error) {
	var cards []CardStripped
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		cards, err = downloadDB()
		if err != nil {
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
	return cards, nil
}
*/

func (v *Validate) getClass(deck hsdeckoder.Deck) (hsdeckoder.Class, error) {
	heroCard, err := v.getCard(deck.Heroes[0])
	if err != nil {
		return hsdeckoder.Class(""), err
	}
	class := heroCard.CardClass
	for _, id := range deck.Cards {
		card, err := v.getCard(id.Id)
		if err != nil {
			return hsdeckoder.Class(""), err
		}
		if card.CardClass != hsdeckoder.NEUTRAL && card.CardClass != class {
			return hsdeckoder.Class(""), ErrInvalidDeck
		}
	}
	return class, nil

}

func (v *Validate) getCard(dbfId int) (CardStripped, error) {
	for _, card := range v.Cards {
		if card.DbfId == dbfId {
			return card, nil
		}
	}
	return CardStripped{}, ErrCardNotFound
}

func (v *Validate) ValidateDeck(deck hsdeckoder.Deck) (ParsedDeck, error) {
	var parsedDeck ParsedDeck
	if len(deck.Heroes) != 1 {
		return ParsedDeck{}, ErrInvalidDeck
	}
	class, err := v.getClass(deck)
	if err != nil {
		return ParsedDeck{}, err
	}
	parsedDeck.Class = class
	var cardCount int
	for _, card := range deck.Cards {
		cardCount += card.Count
		strippedCard, err := v.getCard(card.Id)
		if err != nil {
			return ParsedDeck{}, err
		}
		parsedCard := ParsedCard{
			Id:     strippedCard.Id,
			Name:   strippedCard.Name,
			Count:  card.Count,
			Cost:   strippedCard.Cost,
			Rarity: strippedCard.Rarity,
		}
		parsedDeck.Cards = append(parsedDeck.Cards, parsedCard)
	}
	if cardCount != deckSize {
		return ParsedDeck{}, ErrInvalidDeck
	}
	return parsedDeck, nil
}
