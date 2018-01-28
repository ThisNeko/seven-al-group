#ifndef TRAFFIC_LIGHTS_ANALYZER_H
#define TRAFFIC_LIGHTS_ANALYZER_H

#include <map>
#include "structs/traffic_light_status.hpp"
#include "structs/car_status.hpp"

TrafficLightStatus* SelectTrafficLight(const std::map <int, TrafficLightStatus> trafficLightsRegistry, const CarStatus &carStatus);

#endif // TRAFFIC_LIGHTS_ANALYZER_H