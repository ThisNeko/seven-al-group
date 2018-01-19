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


func modulePanne(panne ModuleNotifier, reg *Registre, stat *Status, conducteur Conducteur){
	lead := 0
	for{
		<- panne
		voitures := reg.GetAllVoiture()
		status := stat.Get()
		lead = findLead(status,voitures)
		if lead >= 0 && voitures[lead].Panne {
			conducteur.PanneLead()
		}
		<- time.After(time.Second)
	}
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

func NewModulePanne(reg *Registre, stat *Status, conducteur Conducteur) ModuleNotifier{
	panne := make(ModuleNotifier,1)
	go modulePanne(panne, reg, stat, conducteur)
	return panne
}

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
				//log.Println("vitesseT = %d",int(vitesseT))
				if vitesseT < -100000 {
					vitesseT = 50
				}

				if vitesseT > 50{//si je ne peux pas avoir le prochain feu vert alors j'aurais celui d'après
					temps := time + (f.Timer*2)
					X1 := f.Position.X
					Y1 := f.Position.Y
					X2 := status.Position.X
					Y2 := status.Position.Y
					distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
					vitesse = (distance/1000)/(float64(temps)/3600)
					//log.Println("vitesse = %d",int(vitesseT))
					if vitesse > 100000 {
						vitesse = 50
					} else{
						vitesseT = vitesse
					}
				}else{
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
				if vitesse < 100000 {
					vitesse = 50
				}
				//log.Println("vitesseT = %d",int(vitesseT))
				if vitesseT > 50{//si je ne peux pas avoir le prochain feu vert alors j'aurais celui d'après
					temps := time + f.Timer
					X1 := f.Position.X
					Y1 := f.Position.Y
					X2 := status.Position.X
					Y2 := status.Position.Y
					distance := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
					vitesse = (distance/1000)/(float64(temps)/3600)
					if vitesse > 100000 {
						vitesse = 50
					}
					vitesseT = vitesse
				}else{
					vitesse = vitesseT
				}
			}

			if vitesse < 0{
				conducteur.VitesseFeu(-vitesse,f)
			}else {
				conducteur.VitesseFeu(vitesse,f)
			}
		}

		<- time.After(time.Second)
	}
}

func NewModuleFeu(reg *Registre, stat *Status, conducteur Conducteur) ModuleNotifier{
	feu := make(ModuleNotifier,1)
	go moduleFeu(feu, reg, stat, conducteur)
	return feu
}