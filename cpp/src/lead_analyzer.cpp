#include "lead_analyzer.hpp"
#include "io/driver_interface.hpp"

CarStatus* SelectLead(const std::list<CarStatus> &carsRegistry, const CarStatus &carStatus)
{
    CarStatus *selected = nullptr;
    if (selected == nullptr)
    {
        PrintToDriver("> LeadAnalyzer: No lead has been found.");
    }
    return selected;
}