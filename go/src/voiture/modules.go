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


func modulePanne(panne ModuleNotifier, reg *Data, conducteur Conducteur){
	for{
		<- panne
		lead := reg.GetLead()
		if lead.ID > -1 && lead.Panne {
			conducteur.PanneLead()
		}
		<- time.After(time.Second)
	}
}



func NewModulePanne(reg *Data, conducteur Conducteur) ModuleNotifier{
	panne := make(ModuleNotifier,1)
	go modulePanne(panne, reg, conducteur)
	return panne
}

func moduleFrein(frein ModuleNotifier, reg *Data, conducteur Conducteur){
	for {
		<- frein
		voitures := reg.GetAllVoiture()
		status := reg.GetStatus()
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

func NewModuleFrein(reg *Data, conducteur Conducteur) ModuleNotifier{
	frein := make(ModuleNotifier,1)
	go moduleFrein(frein, reg, conducteur)
	return frein
}

func moduleFeu(feu ModuleNotifier, reg *Data, conducteur Conducteur) {
	for {
		<- time.After(time.Second)
		<- feu
		feux := reg.GetAllFeux()
		status := reg.GetStatus()

		for _,f := range feux{
			vitesseVoiture := status.Vitesse
			tempsEcoule:=f.Ticker
			//log.Println("tempsRestant",f.Timer-tempsEcoule)
			tempsFeuTotal := float64(f.Timer)
			//log.Println("tempsFeuTotal",f.Timer)
			X1 := f.Position.X
			Y1 := f.Position.Y
			X2 := status.Position.X
			Y2 := status.Position.Y
			//distance en metre
			distanceFeuVoiture := math.Sqrt(math.Pow(X2-X1,2)+math.Pow(Y2-Y1,2))
			//log.Println("distanceFeuVoiture = ",distanceFeuVoiture)
			//log.Println("La couleur du feu est = ",f.Couleur.String())
			//temps en s
			tempsVoitureArriveFeu := (distanceFeuVoiture)/(vitesseVoiture.X/3.6)
			//log.Println("tempsVoitureFeu = ",tempsVoitureFeu)
			//log.Println("tempsVoitureArriveFeu = ",tempsVoitureArriveFeu)
			modulo := math.Mod(tempsVoitureArriveFeu+float64(tempsEcoule),2*tempsFeuTotal)
			//log.Println("modulo = ",modulo)
			if f.Couleur.String() == "RED"{
				if modulo <= tempsFeuTotal{
					//mod := math.Mod(tempsVoitureArriveFeu,tempsFeuTotal)
					nouvelleVitesse := int(distanceFeuVoiture)/int((tempsVoitureArriveFeu + (tempsFeuTotal- modulo)))
					conducteur.VitesseFeu(float64(nouvelleVitesse)*3.6,f)

				}else{
					//log.Println("la couleur est rouge et la voiture va arriver au vert donc on fait rien")
				}
			}else{
				if modulo <= tempsFeuTotal{
					//log.Println("La couleur est GREEN et on fait rien")
				}else{
					//log.Println("la couleur est VERTE et la voiture va arriver au rouge")
					nouvelleVitesse := int(distanceFeuVoiture)/int(tempsVoitureArriveFeu+ (2*tempsFeuTotal-(modulo+1)))
					conducteur.VitesseFeu(float64(nouvelleVitesse)*3.6,f)
				}

			}
			//log.Println()
		}
	}
}
//test :
/*
distance feu-voiture : 200m
vitesse voiture : 80km/h => 22.2m/s
temps voiture pour aller au feu a 22.2m/s = 9sec => arrive au rouge

resultat :
25 km/h = 6.9m/s ; le couleur GREEN et va changer de couleur dans 7 sec
temps = 28.98sec
resultat attendu = 17sec
v = d/t = 200/17 = 11.76m/s = 42.35km/h

 */
func NewModuleFeu(reg *Data, conducteur Conducteur) ModuleNotifier{
	feu := make(ModuleNotifier,1)
	go moduleFeu(feu, reg, conducteur)
	return feu
}