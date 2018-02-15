#ifndef DRIVING_DIRECTIONS_H
#define DRIVING_DIRECTIONS_H

#include <list>
#include "structs/car_status.hpp"
#include "structs/traffic_light_status.hpp"
#include "structs/directions.hpp"

Directions ComputeDrivingDirections(const CarStatus &carStatus, CarStatus const *lead, TrafficLightStatus const *trafficLight);

#endif // DRIVING_DIRECTIONS_H