#ifndef LEAD_ANALYZER_H
#define LEAD_ANALYZER_H

#include <list>
#include "structs/car_status.hpp"

CarStatus* SelectLead(const std::list<CarStatus> &carsRegistry, const CarStatus &carStatus);

#endif // LEAD_ANALYZER_H