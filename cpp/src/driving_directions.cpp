#include "driving_directions.hpp"
#include "io/driver_interface.hpp"
#include <iostream>

using namespace std;

const int MAX_SPEED = 30;

Directions ComputeDrivingDirections(const CarStatus &carStatus, CarStatus const *lead, TrafficLightStatus const *trafficLight)
{
    Directions directions;
    directions.vitesseCible = 30;
    directions.doitFreiner = false;

    if (lead != nullptr)
    {
        double futurX = (double)(carStatus.vitesse.X != 0 ? carStatus.vitesse.X : MAX_SPEED) / 3.6 + carStatus.position.X;
        if(lead->position.X - futurX < 3)
        {
            if (carStatus.vitesse.X != 0)
            {
                PrintToDriver("FREINAGE D'URGENCE");
            }
            directions.vitesseCible = 0;
            return directions;
        }
    }
    else if (trafficLight != nullptr)
    {
        cout << "ok" << " " << trafficLight->couleur << endl;
        if (trafficLight->couleur == RED && trafficLight->pos.X > carStatus.position.X)
        {
            cout << "yes" << endl;
            int tempsFeuVert = trafficLight->ticker + 1;
            double distance = trafficLight->pos.X - carStatus.position.X;
            double vitesseNeeded = (distance / tempsFeuVert) * 3.6;
            if (vitesseNeeded <= MAX_SPEED)
            {
                directions.vitesseCible = vitesseNeeded;
            }
        }
    }

    return directions;
}