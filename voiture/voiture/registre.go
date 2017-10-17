package voiture

type Registre struct{
	mods *ModuleDispatcher
	voitures map[int]VoitureMessage
}

func NewRegistre(mods *ModuleDispatcher) Registre{
	return Registre{
		mods,
		make(map[int]VoitureMessage),
	}
}

func (reg *Registre) Update(mess VoitureMessage){
	reg.voitures[mess.ID] = mess
	reg.mods.Notify()
}