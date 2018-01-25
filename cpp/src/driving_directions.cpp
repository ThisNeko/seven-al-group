#include "driving_directions.hpp"
#include "io/driver_interface.hpp"

bool ComputeDrivingDirections(const CarStatus &carStatus, CarStatus const *lead, TrafficLightStatus const *trafficLight)
{
    if (lead == nullptr && trafficLight == nullptr)
    {
        PrintToDriver("> DrivingDirections: Without any lead nor traffic light, no Driving Directions could be generated.");
        return false;
    }

    Position currentPos = carStatus.position;
    Vitesse vitesse = carStatus.vitesse;

    Position futurPos;
    futurPos.X=currentPos.X+vitesse.X;
    futurPos.Y=currentPos.Y+vitesse.Y;

    //if(currentPos.X>lead->X){}
    Position currentPoslead= lead->position;
    Vitesse vitesseLead = lead->vitesse;

    Position futurPoslead;
    futurPoslead.X=currentPoslead.X+vitesseLead.X;
    futurPoslead.Y=currentPoslead.Y+vitesseLead.Y;

    if(futurPoslead.X - futurPos.X <= 10){
    	PrintToDriver("Frein !");
    }

    return true;
}