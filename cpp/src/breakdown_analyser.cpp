#include "breakdown_analyser.hpp"

bool analyseLead(const CarStatus &carStatus,CarStatus * lead,int timeout){
	
	if(lead->ID > -1 && lead->panne || timeout==0){
		lead->panne=true;
		PrintToDriver("> Breakdown_Analyser : Voiture Lead en panne, Il faut changer de voie");
	}
	return lead->panne;
}

