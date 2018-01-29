#ifndef LEAD_ANALYZER_H
#define LEAD_ANALYZER_H

#include <map>
#include "structs/car_status.hpp"

CarStatus* SelectLead(std::map<int, CarStatus> &carsRegistry, const CarStatus &carStatus);

#endif // LEAD_ANALYZER_H