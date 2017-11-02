package feu


type Couleur int

const (
	GREEN Couleur = 1 + iota
	RED
	YELLOW
)

var couleurs = [...]string {
	"GREEN",
	"RED",
	"YELLOW",
}

type Feu struct {
	ID       int
	Position Position
	Couleur  Couleur
	Ticker int
}

type Position struct{
	X float64
	Y float64
}


func (couleur Couleur) String() string {
	return couleurs[couleur - 1]
}



func (feu *Feu) Update(updated Feu){
	(*feu)= Feu{
		ID:updated.ID,
		Position:Position{updated.Position.X,updated.Position.Y},
		Couleur:updated.Couleur,
		Ticker:updated.Ticker,
	}
	//if t := updated.ticker {

	//}
}

func newFeu() Feu{
	feu := Feu{}
	return feu
}
