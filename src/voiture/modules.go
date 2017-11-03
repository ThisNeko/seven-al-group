package voiture

import (
	"time"
)

type ModuleDispatcher []ModuleNotifier

func NewModuleDispatcher() ModuleDispatcher{
	return make(ModuleDispatcher,0)
}

func (mods *ModuleDispatcher) AddModule(mod ModuleNotifier){
	*mods = append(*mods, mod)
}

func (mods *ModuleDispatcher) Notify(){
	for _,mod := range *mods {
		select {
		case <- mod:
		default:
		}
		mod <- struct {}{}
	}
}

type ModuleNotifier chan struct{}

func moduleFrein(frein ModuleNotifier, reg *Registre, stat *Status, conducteur Conducteur){
	for {
		<- frein
		voitures := reg.GetAllVoiture()
		status := stat.Get()
		futurePos := Position{
			status.Position.X+status.Vitesse.X,
			status.Position.Y+status.Vitesse.Y,
		}
		for _,v := range voitures{
			if status.Position.X > v.Position.X {continue}
			futurePos2 := Position{
				v.Position.X+v.Vitesse.X,
				v.Position.Y+v.Vitesse.Y,
			}
			//if futurePos.Distance(futurePos2) <= 10 {
			//	conducteur.AlerteFrein()
			//	<- time.After(time.Second)
			//}
			if futurePos2.X - futurePos.X <= 10 {
				conducteur.AlerteFrein()
				<- time.After(time.Second)
			}
		}
	}
}

func NewModuleFrein(reg *Registre, stat *Status, conducteur Conducteur) ModuleNotifier{
	frein := make(ModuleNotifier,1)
	go moduleFrein(frein, reg, stat, conducteur)
	return frein
}