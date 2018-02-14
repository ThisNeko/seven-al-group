#ifndef BREAKDOWN_ANALYSER_H
#define BREAKDOWN_ANALYSER_H

#include <list>
#include "structs/car_status.hpp"
#include "structs/directions.hpp"
#include "io/driver_interface.hpp"



bool analyseLead(const CarStatus &carStatus,CarStatus * lead,int timeout);

#endif // DRIVING_DIRECTIONS_H