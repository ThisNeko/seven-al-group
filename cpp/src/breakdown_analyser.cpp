#include "breakdown_analyser.hpp"

bool analyseLead(const CarStatus &carStatus,CarStatus * lead){
	if(lead->ID > -1 && lead->panne){
		PrintToDriver("> Breakdown_Analyser : Voiture Lead en panne");
	}
	return lead->panne;
}

/*

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
*/