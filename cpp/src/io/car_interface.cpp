#include "io/car_interface.hpp"
#include <iostream>

using namespace std;

void CarInterfaceLoop(CarStatus *carStatus)
{
    while(true)
    {
        string in;
        cin >> in;
        if (in == "pon")
        {
            carStatus->panne = true;
        }
        else if (in == "poff")
        {
            carStatus->panne = false;
        }
        else if (in == "u")
        {
            carStatus->position.Y--;
        }
        else if (in == "d")
        {
            carStatus->position.Y++;
        }
    }
}