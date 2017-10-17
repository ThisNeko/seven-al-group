package voiture

import (
	"math/rand"
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

func moduleFrein(frein ModuleNotifier){
	for {
		<- frein
		if rand.Intn(100) < 2 {
			AlerteFrein()
		}
	}
}

func NewModuleFrein() ModuleNotifier{
	frein := make(ModuleNotifier,1)
	go moduleFrein(frein)
	return frein
}