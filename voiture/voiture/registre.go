package voiture

type Registre struct{
	update chan StatusVoiture
	getAll chan chan map[int]StatusVoiture
}

func RegistreLoop(reg Registre, mods *ModuleDispatcher){
	voitures := make(map[int]StatusVoiture)
	for{
		select {
		case message := <- reg.update:
			voitures[message.ID] = message
			mods.Notify()
		case response := <- reg.getAll:
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
		make(chan chan map[int]StatusVoiture),
	}
	go RegistreLoop(registre,mods)
	return registre
}

func (reg *Registre) Update(mess StatusVoiture){
	reg.update <- mess
}

func (reg *Registre) GetAll() map[int]StatusVoiture {
	response := make(chan map[int]StatusVoiture)
	reg.getAll <- response
	return <-response
}