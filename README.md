# Hearthstone Deckcode Decoder - hsdeckoder

Golang package to decode hearthstone deckstring to useful data inspired by [hearthsim](https://hearthsim.info/docs/deckstrings/).

# Installation

```bash
go get github.com/MarekSalgovic/hsdeckoder
```



# Usage

Decodes deckcode string to deck struct. For more details see [format](https://hearthsim.info/docs/deckstrings/#format).
```go
deckstring := "AAECAaIHBrICyAOvBOEE5/oC/KMDDLQBywPNA9QF7gaIB90I7/ECj5cDiZsD/6UD9acDAA=="
deck, err := hsdeckoder.Decode(deckstring)
//handle error
if err != nil{
  panic(err)
}
fmt.Println(deck)
//{2 [930] [{306 1} {456 1} {559 1} {609 1} {48487 1} {53756 1} {180 2} {459 2} {461 2} 
//{724 2} {878 2} {904 2} {1117 2} {47343 2} {52111 2} {52617 2} {54015 2} {54261 2}]}
```

Deck struct contains 3 fields: 

```go
type Deck struct {
	Format	Format //constant defining format - standart/wild
	Heroes	[]int //array of hero portrait dbfIds used in deckcode
	Cards 	[]Card // array of cards
}

type Card struct {
	Id int // dbfId of card
	Count int // occurrance count of card
}
```

# hsvalidator

Package using [hearthstonejson](https://hearthstonejson.com) API to validate created deck struct.

```go
validator, err := validator.NewValidator()
//handle error
if err != nil{
  panic(err)
}
validatedDeck, err := validator.ValidateDeck(deck)
//handle error
if err != nil{
  panic(err)
}
```
Initializing validator downloads cards to runtime memory from [hearthstonejson](https://hearthstonejson.com) API.