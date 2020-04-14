package hsdeckoder

type Card struct {
	Id    int
	Count int
}

type Deck struct {
	Format Format
	Heroes []int
	Cards  []Card
}

type Format int

const (
	FMT_UNKNOWN  Format = 0
	FMT_WILD     Format = 1
	FMT_STANDART Format = 2
)

type Class string

const (
	NEUTRAL     Class = "NEUTRAL"
	DRUID       Class = "DRUID"
	HUNTER      Class = "HUNTER"
	MAGE        Class = "MAGE"
	PALADIN     Class = "PALADIN"
	PRIEST      Class = "PRIEST"
	ROGUE       Class = "ROGUE"
	SHAMAN      Class = "SHAMAN"
	WARLOCK     Class = "WARLOCK"
	WARRIOR     Class = "WARRIOR"
	DEMONHUNTER Class = "DEMONHUNTER"
)
