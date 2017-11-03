package voiture

import ("time"
		"math"
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
			if futurePos2.X - futurePos.X <= 10{
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

		/*if len(feux) > 0 {
			temps:=feux[0].Ticker
			println(temps)
			X1 := feux[0].Position.X
			Y1 := feux[0].Position.Y
			println(X1)
			X2 := status.Position.X
			Y2 := status.Position.Y
			println(X2)
			distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
			println(distance)
			vitesse := distance/float64(temps)
			println(vitesse)
			conducteur.VitesseFeu(vitesse)
		}*/
		for _,f := range feux{
			time:=f.Timer - f.Ticker
			var vitesse float64

			if f.Couleur.String()=="RED"{
				temps := time
				X1 := f.Position.X
				Y1 := f.Position.Y
				X2 := status.Position.X
				Y2 := status.Position.Y
				distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
				vitesseT := (distance/1000)/(float64(temps)/3600)

				if vitesseT > 50{//si je ne peux pas avoir le prochain feu vert alors j'aurais celui d'après
					temps := time + (f.Timer*2)
					X1 := f.Position.X
					Y1 := f.Position.Y
					X2 := status.Position.X
					Y2 := status.Position.Y
					distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
					vitesse = (distance/1000)/(float64(temps)/3600)
				} else{
					vitesse = vitesseT
				}

			} else if f.Couleur.String()=="GREEN"{//je calcul la vitesse pour le prochain feu vert
				temps:= time
				X1 := f.Position.X
				Y1 := f.Position.Y
				X2 := status.Position.X
				Y2 := status.Position.Y
				distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
				vitesseT := (distance/1000)/(float64(temps)/3600)

				if vitesseT > 50{//si je ne peux pas avoir le prochain feu vert alors j'aurais celui d'après
					temps := time + f.Timer
					X1 := f.Position.X
					Y1 := f.Position.Y
					X2 := status.Position.X
					Y2 := status.Position.Y
					distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
					vitesse = (distance/1000)/(float64(temps)/3600)
				} else{
					vitesse = vitesseT
				}
			}

			conducteur.VitesseFeu(-vitesse,f)
		}

		<- time.After(time.Second)
	}
}

func NewModuleFeu(reg *Registre, stat *Status, conducteur Conducteur) ModuleNotifier{
	feu := make(ModuleNotifier,1)
	go moduleFeu(feu, reg, stat, conducteur)
	return feu
}