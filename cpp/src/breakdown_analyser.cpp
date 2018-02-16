#include "breakdown_analyser.hpp"
#include "controller.hpp"
#include <iostream>

using namespace std;

int lastTimePrint = 0;

bool analyseLead(const CarStatus &carStatus,CarStatus * lead,int timeout){
	if(lead->ID > -1 && (lead->panne || timeout==0))
	{
		lastTimePrint += msBetweenTicks;
		lead->panne=true;
		if (lastTimePrint >= 1000)
		{
			cout << "Voiture Lead en panne, Il faut changer de voie" << endl;
			lastTimePrint = 0;
		}
	}
	return lead->panne;
}

