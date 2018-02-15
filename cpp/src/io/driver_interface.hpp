#ifndef DRIVER_INTERFACE_H
#define DRIVER_INTERFACE_H

#include <string>
#include "utils/communication_channel.hpp"
#include "structs/directions.hpp"
#include "structs/car_status.hpp"

void PrintToDriver(std::string message);

void FollowDirectionsLoop(CommunicationChannel<Directions> *chan, CarStatus *carStatus);

#endif // DRIVER_INTERFACE_H