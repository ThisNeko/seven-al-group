package voiture

import "math"

func (pos Position) Distance(pos2 Position)float64{
	return math.Sqrt(math.Pow(pos.X-pos2.X,2)+math.Pow(pos.Y-pos2.Y,2))
}

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
	Panne bool
}

type Status struct{
	update chan Materiel
	get chan chan StatusVoiture
}

func (status *Status) Update(mat Materiel){
	status.update <- mat
}

func (status *Status) Get() StatusVoiture{
	response := make(chan StatusVoiture)
	status.get <- response
	return <- response
}

func StatusLoop(stat Status, mods *ModuleDispatcher, conn *connection){
	var status StatusVoiture
	for{
		select {
		case mat := <- stat.update:
			status = StatusVoiture{
				ID:mat.ID,
				Vitesse:Vitesse{mat.Vitesse,0},
				Position:Position{mat.Position.X,mat.Position.Y},
				Panne:mat.Panne,
			}
			mods.Notify()
			conn.Broadcast(status)
		case response := <- stat.get:
			response <- status
		}
	}
}

func NewStatus(mods *ModuleDispatcher, conn *connection) Status{
	stat := Status{
		make(chan Materiel),
		make(chan chan StatusVoiture),
	}
	go StatusLoop(stat, mods, conn)
	return stat
}