#include "lead_analyzer.hpp"
#include "io/driver_interface.hpp"
#include <iostream>

using namespace std;

CarStatus* SelectLead(std::map<int, CarStatus> &carsRegistry, const CarStatus &carStatus)
{
    CarStatus *selected = nullptr;

    for (auto it = carsRegistry.begin(); it != carsRegistry.end(); ++it)
    {
        //selected = nullptr;
        if(it->second.position.X>carStatus.position.X && it->second.position.Y == carStatus.position.Y 
            && it->second.panne==false){
            selected = &(it->second);
        }
        cout << "Panne: " << it->second.panne << endl;

    }
    //selected = nullptr;
    if (selected == nullptr)
    {
        PrintToDriver("> LeadAnalyzer: No lead has been found.");
    } else {
        PrintToDriver("> LeadAnalyzer: Following a Lead !");
    }
    return selected;
}