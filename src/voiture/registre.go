package voiture

type Registre struct{
	updateVoiture chan StatusVoiture
	updateFeu chan Feu
	getAllVoiture chan chan map[int]StatusVoiture
	getAllFeux chan chan map[int]Feu
}

func RegistreLoop(reg Registre, mods *ModuleDispatcher){
	voitures := make(map[int]StatusVoiture)
	for{
		select {
		case message := <- reg.updateVoiture:
			voitures[message.ID] = message
			mods.Notify()
		case response := <- reg.getAllVoiture:
			tmp := make(map[int]StatusVoiture)
			for k,v := range voitures{
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
	go RegistreLoop(registre,mods)
	return registre
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