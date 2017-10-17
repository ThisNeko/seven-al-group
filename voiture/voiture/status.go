package voiture

type Position struct{
	X float64
	Y float64
}

type Vitesse struct{
	X float64
	Y float64
}

type StatusVoiture struct{
	ID int
	Vitesse Vitesse
	Position Position
	Frein bool
}

type Status struct{
	mods *ModuleDispatcher
	status StatusVoiture
}

func (status *Status) Update(mat Materiel){
	(*status).status = StatusVoiture{
		ID:mat.ID,
		Vitesse:Vitesse{mat.Vitesse,0},
		Position:Position{mat.Position.X,mat.Position.Y},
		Frein:mat.Frein,
	}
}

func NewStatus(mods *ModuleDispatcher) Status{
	stat := Status{}
	stat.mods = mods
	return stat
}