#ifndef CAR_INTERFACE_H
#define CAR_INTERFACE_H

#include "utils/communication_channel.hpp"
#include "structs/car_status.hpp"
#include "structs/directions.hpp"

void CarInterfaceLoop(CarStatus *carStatus);

#endif // CAR_INTERFACE_H