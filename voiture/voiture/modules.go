package voiture

import ("time"
		"math")

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
			if futurePos.Distance(futurePos2) <= 10 {
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
//Pour le moment on admet qu'il n'y a qu'un seul feu
func moduleFeu(feu ModuleNotifier, reg *Registre, stat *Status, conducteur Conducteur){
	for {
		<- feu
		feux := reg.GetAllFeux()
		status := stat.Get()
		if len(feux) > 0 {
			temps:=feux[1].Ticker
			X1 := feux[1].Position.X
			Y1 := feux[1].Position.Y
			X2 := status.Position.X
			Y2 := status.Position.Y
			distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
			vitesse := distance/float64(temps)
			conducteur.VitesseFeu(vitesse)
		}
	}
}

func NewModuleFeu(reg *Registre, stat *Status, conducteur Conducteur) ModuleNotifier{
	feu := make(ModuleNotifier,1)
	go moduleFeu(feu, reg, stat, conducteur)
	return feu
}