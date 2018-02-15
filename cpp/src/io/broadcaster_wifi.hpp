#ifndef BROADCASTER_WIFI_HPP
#define BROADCASTER_WIFI_HPP

#include "structs/car_status.hpp"

class Broadcaster_wifi
{
public:
	Broadcaster_wifi();
	void BroadcasterLoop(CarStatus *carStatus);

};



#endif