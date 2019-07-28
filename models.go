package hsdeckoder




type Card struct {
	Id int
	Count int
}


type Deck struct {
	Format	Format
	Heroes	[]int
	Cards 	[]Card
}

type Format int

const (
	FMT_UNKNOWN  Format = 0
	FMT_WILD     Format = 1
	FMT_STANDART Format = 2
)

