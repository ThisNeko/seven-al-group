package voiture

import "time"

type Registre struct{
	updateVoiture chan StatusVoiture
	updateFeu chan Feu
	getAllVoiture chan chan map[int]StatusVoiture
	getAllFeux chan chan map[int]Feu
}

func RegistreLoop(reg Registre, mods *ModuleDispatcher, voitures map[int]StatusVoiture , feux map[int]Feu, timeouts map[int]time.Time){
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
		}
	}
}

func NewRegistre(mods *ModuleDispatcher) Registre{
	registre := Registre{
		make(chan StatusVoiture),
		make(chan Feu),
		make(chan chan map[int]StatusVoiture),
		make(chan chan map [int]Feu),
	}
	voitures := make(map[int]StatusVoiture)
	feux := make(map[int]Feu)
	timeouts := make(map[int]time.Time)
	go RegistreLoop(registre,mods,voitures,feux,timeouts)
	go TimeoutLoop(mods,voitures,timeouts)
	return registre
}
func TimeoutLoop(mods *ModuleDispatcher, voitures map[int]StatusVoiture, timeouts map[int]time.Time) {
	for{
		<- time.After(time.Second/2)
		for i,t := range timeouts {
			now := time.Now()
			if now.Sub(t) > time.Second/2 {
				voiture := voitures[i]
				voiture.Panne = true
				voitures[i] = voiture
				mods.Notify()
			}
		}
	}
}

func (reg *Registre) UpdateVoiture(mess StatusVoiture){
	reg.updateVoiture <- mess
}

func (reg *Registre) GetAllVoiture() map[int]StatusVoiture {
	response := make(chan map[int]StatusVoiture)
	reg.getAllVoiture <- response
	return <-response
}

func (reg *Registre) UpdateFeu(mess Feu){
	reg.updateFeu <- mess
}

func (reg *Registre) GetAllFeux() map[int]Feu {
	response := make(chan map[int]Feu)
	reg.getAllFeux <- response
	return <-response
}