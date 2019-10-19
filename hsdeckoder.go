package hsdeckoder

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
)



var (
	ErrInvalidCode = errors.New("deckcode invalid")
)



func parseHeader(bs []byte) ([]byte, Format, error) {

	byteZero, c := binary.Uvarint(bs)
	if byteZero == 0 && c <= 0 {

		return bs, FMT_UNKNOWN, ErrInvalidCode
	}
	if byteZero != 0 || c != 1 {

		return bs, FMT_UNKNOWN, ErrInvalidCode
	}
	bs = bs[c:]
	version, c := binary.Uvarint(bs)
	if version != 1 || c != 1 {

		return bs, FMT_UNKNOWN, ErrInvalidCode
	}
	if version == 0 && c <= 0 {

		return bs, FMT_UNKNOWN, ErrInvalidCode
	}
	bs = bs[c:]
	format, c := binary.Uvarint(bs)
	if (format != uint64(FMT_WILD) && format != uint64(FMT_STANDART)) || c != 1 {
		return bs, FMT_UNKNOWN, ErrInvalidCode
	}
	if format == 0 && c <= 0 {

		return bs, FMT_UNKNOWN, ErrInvalidCode
	}
	bs = bs[c:]
	return bs, Format(format), nil
}


func parseNCopyCards(bs []byte) ([]byte, []Card, error) {
	byteRemainder := len(bs)
	var cards []Card
	for byteRemainder>0{
		cardId, c := binary.Uvarint(bs)
		if cardId == 0 && c <= 0 {
			return bs, []Card{}, ErrInvalidCode
		}
		byteRemainder -= c
		cardCount, c := binary.Uvarint(bs)
		if cardCount == 0 && c <= 0 {
			return bs, []Card{}, ErrInvalidCode
		}
		cards = append(cards, Card{Id: int(cardId), Count: int(cardCount)})
	}
	return bs, cards, nil

}



func parseBodyHelper(bs []byte, count int) ([]byte, []Card, error){
	var cards []Card
	uniqueCount, c := binary.Uvarint(bs)
	if uniqueCount == 0 && c <= 0 {
		return bs, []Card{}, ErrInvalidCode
	}
	bs = bs[c:]
	for i := uint64(0); i < uniqueCount; i++ {
		card, c := binary.Uvarint(bs)
		if card == 0 && c <= 0 {
			return bs, []Card{}, ErrInvalidCode
		}
		cards = append(cards, Card{Id: int(card), Count: count})
		bs = bs[c:]
	}
	return bs, cards, nil
}


func parseBody(bs []byte, d Deck) (Deck, error) {

	bs, heroCopy, err := parseBodyHelper(bs,0)

	if err != nil{
		return Deck{},ErrInvalidCode
	}

	bs, singleCopy, err := parseBodyHelper(bs, 1)
	if err != nil{
		return Deck{},ErrInvalidCode
	}
	bs, doubleCopy, err := parseBodyHelper(bs,2)

	if err != nil{
		return Deck{},ErrInvalidCode
	}
	bs, nCopy, err := parseNCopyCards(bs)

	if err != nil{
		return Deck{},ErrInvalidCode
	}
	var heroes []int
	for _, hero:= range heroCopy{
		heroes = append(heroes, hero.Id)
	}

	var cards []Card
	cards = append(cards, singleCopy...)
	cards = append(cards, doubleCopy...)
	cards = append(cards, nCopy...)
	d.Heroes = heroes
	d.Cards = cards
	return d, nil
}

func Decode(dc string) (Deck, error) {
	var deck Deck
	bs, err := base64.StdEncoding.DecodeString(dc)
	if err != nil {
		return Deck{}, ErrInvalidCode
	}
	bs, format, err := parseHeader(bs)
	if err != nil {
		return Deck{}, err
	}
	deck.Format = format
	deck, err = parseBody(bs, deck)
	if err != nil {
		return Deck{}, nil
	}

	return deck, nil
}



