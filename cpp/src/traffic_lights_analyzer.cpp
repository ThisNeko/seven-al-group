#include "traffic_lights_analyzer.hpp"
#include "io/driver_interface.hpp"

TrafficLightStatus* SelectTrafficLight(const std::list <TrafficLightStatus> trafficLightsRegistry, const CarStatus &carStatus)
{
    TrafficLightStatus *selected = nullptr;

    if (selected == nullptr)
    {
        PrintToDriver("> TrafficLightsAnalyzer: No traffic light has been found.");
    }

    return selected;
}