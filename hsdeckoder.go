package hsdeckoder

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
)

type Format int

type Card struct {
	Id int
	Count int
}


type Deck struct {
	Format	Format
	Heroes	[]int
	Cards 	[]Card
}



const (
	FMT_UNKNOWN  Format = 0
	FMT_WILD     Format = 1
	FMT_STANDART Format = 2
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
	var cards []Card
	cards = append(cards, singleCopy...)
	cards = append(cards, doubleCopy...)
	var heroes []int
	for i:=0;i<len(heroCopy);i++{
		heroes = append(heroes, heroCopy[i].Id)
	}
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
