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
	getCard(dbfId int) (CardAPI, error)
	getClass(deck hsdeckoder.Deck) (hsdeckoder.Class, error)
}

type Validate struct {
	Cards []CardAPI
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

func downloadDB() ([]CardAPI, error) {
	var cards []CardAPI
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
	return cards, nil
}


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

func (v *Validate) getCard(dbfId int) (CardAPI, error) {
	for _, card := range v.Cards {
		if card.DbfId == dbfId {
			return card, nil
		}
	}
	return CardAPI{}, ErrCardNotFound
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
		cardAPI, err := v.getCard(card.Id)
		if err != nil {
			return ParsedDeck{}, err
		}
		parsedDeck.Cards = append(parsedDeck.Cards, cardAPI)
	}
	if cardCount != deckSize {
		return ParsedDeck{}, ErrInvalidDeck
	}
	return parsedDeck, nil
}
