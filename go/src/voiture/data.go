package voiture

import (
	"time"
	"math"
)

type Data struct{
	updateVoiture chan StatusVoiture
	updateFeu     chan Feu
	getAllVoiture chan chan map[int]StatusVoiture
	getAllFeux    chan chan map[int]Feu
	updateStatus  chan Materiel
	getStatus     chan chan StatusVoiture
	getLead       chan chan StatusVoiture
}


type Position struct{
	X float64
	Y float64
}

func (pos Position) Distance(pos2 Position)float64{
	return math.Sqrt(math.Pow(pos.X-pos2.X,2)+math.Pow(pos.Y-pos2.Y,2))
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

func RegisterLoop(reg *Data, mods *ModuleDispatcher, voitures map[int]StatusVoiture , feux map[int]Feu, timeouts map[int]time.Time, lead *int){
	for{
		select {
		case message := <- reg.updateVoiture:
			voitures[message.ID] = message
			timeouts[message.ID] = time.Now()
			mods.Notify()
		case message := <- reg.updateFeu:
			feux[message.ID] = message
			mods.Notify()
		case response := <- reg.getAllVoiture:
			tmp := make(map[int]StatusVoiture)
			for k,v := range voitures{
				tmp[k] = v
			}
			response <- tmp
		case response := <- reg.getAllFeux:
			tmp := make(map[int]Feu)
			for k,v := range feux{
				tmp[k] = v
			}
			response <- tmp
		case response := <- reg.getLead:
			if *lead == -1 {
				var l StatusVoiture
				l.ID = -1
				response <- l
			}else{
				response <- voitures[*lead]
			}
		}

	}
}

func StatusLoop(status *StatusVoiture, stat *Data, mods *ModuleDispatcher, conn *connection){

	for{
		select {
		case mat := <- stat.updateStatus:
			*status = StatusVoiture{
				ID:mat.ID,
				Vitesse:Vitesse{mat.Vitesse,0},
				Position:Position{mat.Position.X,mat.Position.Y},
				Panne:mat.Panne,
			}
			mods.Notify()
			info := InfoVoiture{
				mat.ID,
				Vitesse{mat.Vitesse,0},
				Position{mat.Position.X,mat.Position.Y},
				mat.Panne,
				//time.Now().UTC().UnixNano(),
			}
			conn.Broadcast(info)
		case response := <- stat.getStatus:
			response <- *status
		}
	}
}

func TimeoutLoop(mods *ModuleDispatcher, data *Data, timeouts map[int]time.Time) {
	for{
		<- time.After(time.Second/2)
		for i,t := range timeouts {
			now := time.Now()
			if now.Sub(t) > time.Second/2 {
				voitures := data.GetAllVoiture()
				voiture := voitures[i]
				voiture.Panne = true
				voitures[i] = voiture
				mods.Notify()
			}
		}
	}
}

func LeadLoop(status *StatusVoiture, data *Data, lead *int) {
	for{
		<- time.After(50*time.Millisecond)
		voitures := data.GetAllVoiture()
		*lead = findLead(*status,voitures)
	}
}

func (data *Data) Start(mods *ModuleDispatcher, conn *connection){
	voitures := make(map[int]StatusVoiture)
	feux := make(map[int]Feu)
	timeouts := make(map[int]time.Time)
	var status StatusVoiture
	lead := -1
	go RegisterLoop(data,mods,voitures,feux,timeouts,&lead)
	go TimeoutLoop(mods,data,timeouts)
	go StatusLoop(&status, data, mods, conn)
	go LeadLoop(&status,data,&lead)
}


func NewData() Data{
	registre := Data{
		make(chan StatusVoiture),
		make(chan Feu),
		make(chan chan map[int]StatusVoiture),
		make(chan chan map [int]Feu),
		make(chan Materiel),
		make(chan chan StatusVoiture),
		make(chan chan StatusVoiture),
	}
	return registre
}


func (reg *Data) UpdateVoiture(mess StatusVoiture){
	reg.updateVoiture <- mess
}

func (reg *Data) GetAllVoiture() map[int]StatusVoiture {
	response := make(chan map[int]StatusVoiture)
	reg.getAllVoiture <- response
	return <-response
}

func (reg *Data) UpdateFeu(mess Feu){
	reg.updateFeu <- mess
}

func (reg *Data) GetAllFeux() map[int]Feu {
	response := make(chan map[int]Feu)
	reg.getAllFeux <- response
	return <-response
}



func (status *Data) UpdateStatus(mat Materiel){
	status.updateStatus <- mat
}

func (status *Data) GetStatus() StatusVoiture{
	response := make(chan StatusVoiture)
	status.getStatus <- response
	return <- response
}


func (data *Data) GetLead() StatusVoiture{
	response := make(chan StatusVoiture)
	data.getLead <- response
	return <- response
}

func findLead(status StatusVoiture, voitures map[int]StatusVoiture) int {
	maxX := status.Position.X + 100
	minX := status.Position.X
	minY := status.Position.Y - 1
	maxY := status.Position.Y + 1
	nearestX := maxX
	lead := -1
	for _,v := range voitures{
		x := v.Position.X
		y := v.Position.Y
		if x > minX && x < nearestX && y > minY && y < maxY {
			lead = v.ID
			nearestX = x
		}
	}
	return lead
}