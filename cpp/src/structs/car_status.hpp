#ifndef CAR_STATUS_H
#define CAR_STATUS_H

#include "utils/json.hpp"

using nlohmann::json;

typedef struct {
    double X = 0;
    double Y = 0;
} Position;

typedef struct {
    double X = 0;
    double Y = 0;
} Vitesse;

typedef struct CarStatus{
    Position position;
    Vitesse vitesse;
    bool panne;
    int ID;
} CarStatus;

CarStatus JSONToCarStatus(json data);
json CarStatusToJSON(const CarStatus &carStatus);

#endif // CAR_STATUS_H