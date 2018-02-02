#ifndef TRAFFIC_LIGHT_STATUS_H
#define TRAFFIC_LIGHT_STATUS_H

#include "car_status.hpp"

typedef enum couleur
{
	GREEN = 1,
	RED,
	YELLOW
	
} Couleur;

typedef struct{
	int ID;
	Position pos;
	Couleur couleur;
	int ticker;
	int timer;

} TrafficLightStatus;

TrafficLightStatus JSONToTrafficLightStatus(json data);

#endif // TRAFFIC_LIGHT_STATUS_H