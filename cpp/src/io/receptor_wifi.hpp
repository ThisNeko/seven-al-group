#ifndef RECEPTOR_WIFI_HPP
#define RECEPTOR_WIFI_HPP

#include "utils/communication_channel.hpp"
#include "structs/car_status.hpp"
#include "structs/traffic_light_status.hpp"

class Receptor_wifi
{
public:
	Receptor_wifi();
	void ReceptorLoop(CommunicationChannel<CarStatus> *chanCar, CommunicationChannel<TrafficLightStatus> *chanTrafficLight, int ignoreId);

};




#endif