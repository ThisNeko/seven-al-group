#include "io/driver_interface.hpp"

#include <iostream>

using namespace std;

void PrintToDriver(std::string message)
{
    // cout << ">>> MESSAGE AU CONDUCTEUR: " << message << endl;
}

void FollowDirectionsLoop(CommunicationChannel<Directions> *chan, CarStatus *carStatus)
{
    while(true)
    {
        Directions d = chan->get();

        // cout << d->vitesse.X << endl;

        // if (d.vitesseCible == 0)
        // {
        //     carStatus->vitesse.X = 0;
        // }
        if (d.vitesseCible < carStatus->vitesse.X)
        {
            cout << ">>> MESSAGE AU CONDUCTEUR: DECCELERE A " << d.vitesseCible << " km/h" << endl;
            carStatus->vitesse.X = d.vitesseCible;
        }
        else if (d.vitesseCible > carStatus->vitesse.X)
        {
            cout << ">>> MESSAGE AU CONDUCTEUR: ACCELERE A " << d.vitesseCible << " km/h" << endl;
            carStatus->vitesse.X = d.vitesseCible;
        }
    }
}