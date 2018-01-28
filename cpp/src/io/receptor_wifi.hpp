#ifndef RECEPTOR_WIFI_HPP
#define RECEPTOR_WIFI_HPP

#include "utils/communication_channel.hpp"
#include "structs/car_status.hpp"

class Receptor_wifi
{
public:
	Receptor_wifi();
	void ReceptorLoop(CommunicationChannel<CarStatus> *chan);

};




#endif