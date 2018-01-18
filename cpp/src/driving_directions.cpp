#include "driving_directions.hpp"

bool ComputeDrivingDirections(Directions &directions, const CarStatus &carStatus, CarStatus const *lead, TrafficLightStatus const *trafficLight)
{
    if (lead == nullptr && trafficLight == nullptr)
    {
        return false;
    }

    return true;
}