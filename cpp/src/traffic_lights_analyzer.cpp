#include "traffic_lights_analyzer.hpp"
#include "io/driver_interface.hpp"

TrafficLightStatus* SelectTrafficLight(std::map <int, TrafficLightStatus> &trafficLightsRegistry, const CarStatus &carStatus)
{
    TrafficLightStatus *selected = nullptr;

    for (auto it = trafficLightsRegistry.begin(); it != trafficLightsRegistry.end(); ++it)
    {
        selected = &(it->second);
    }

    if (selected == nullptr)
    {
        PrintToDriver("> TrafficLightsAnalyzer: No traffic light has been found.");
    }

    return selected;
}