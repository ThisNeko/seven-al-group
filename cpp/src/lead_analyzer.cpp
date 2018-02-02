#include "lead_analyzer.hpp"
#include "io/driver_interface.hpp"

CarStatus* SelectLead(std::map<int, CarStatus> &carsRegistry, const CarStatus &carStatus)
{
    CarStatus *selected = nullptr;

    for (auto it = carsRegistry.begin(); it != carsRegistry.end(); ++it)
    {
        selected = &(it->second);
    }
    if (selected == nullptr)
    {
        PrintToDriver("> LeadAnalyzer: No lead has been found.");
    }
    return selected;
}