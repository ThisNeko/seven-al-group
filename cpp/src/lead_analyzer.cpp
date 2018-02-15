#include "lead_analyzer.hpp"
#include "io/driver_interface.hpp"
#include <iostream>

using namespace std;

CarStatus* SelectLead(std::map<int, CarStatus> &carsRegistry, const CarStatus &carStatus)
{
    CarStatus *selected = nullptr;

    for (auto it = carsRegistry.begin(); it != carsRegistry.end(); ++it)
    {
        if(it->second.position.X > carStatus.position.X && it->second.position.Y == carStatus.position.Y ){
            selected = &(it->second);
        }

    }
    if (selected == nullptr)
    {
        PrintToDriver("> LeadAnalyzer: No lead has been found.");
    } else {
        PrintToDriver("> LeadAnalyzer: Following a Lead !");
    }

    if(selected != nullptr && selected->panne==true)
    {
        PrintToDriver("> Lead en panne changer de voie !");
    }
    return selected;
}