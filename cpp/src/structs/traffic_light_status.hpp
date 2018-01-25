#ifndef TRAFFIC_LIGHT_STATUS_H
#define TRAFFIC_LIGHT_STATUS_H

#include "car_status.hpp"

typedef enum couleur
{
	RED,
	GREEN,
	YELLOW
	
} Couleur;

typedef struct{
	int ID;
	Position pos;
	Couleur couleur;
	int ticker;
	int timer;

} TrafficLightStatus;

#endif // TRAFFIC_LIGHT_STATUS_H