#ifndef BREAKDOWN_ANALYSER_H
#define BREAKDOWN_ANALYSER_H

#include <list>
#include "structs/car_status.hpp"
#include "structs/directions.hpp"
#include "io/driver_interface.hpp"
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

bool analyseLead(const CarStatus &carStatus,CarStatus * lead);

#endif // DRIVING_DIRECTIONS_H