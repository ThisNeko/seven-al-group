#ifndef CAR_INTERFACE_H
#define CAR_INTERFACE_H

#include "utils/communication_channel.hpp"
#include "structs/car_status.hpp"

void CarInterfaceLoop(CommunicationChannel<CarStatus> *channel);

#endif // CAR_INTERFACE_H