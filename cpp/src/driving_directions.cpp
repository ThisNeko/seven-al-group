#include "driving_directions.hpp"
#include "io/driver_interface.hpp"

bool ComputeDrivingDirections(Directions &directions, const CarStatus &carStatus, CarStatus const *lead, TrafficLightStatus const *trafficLight)
{
    if (lead == nullptr && trafficLight == nullptr)
    {
        PrintToDriver("> DrivingDirections: Without any lead nor traffic light, no Driving Directions could be generated.");
        return false;
    }

    return true;
}