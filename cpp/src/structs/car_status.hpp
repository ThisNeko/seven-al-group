#ifndef CAR_STATUS_H
#define CAR_STATUS_H

#include "utils/json.hpp"

using nlohmann::json;

typedef struct {
    int X;
    int Y;
} Position;

typedef struct {
    int X = 0;
    int Y = 0;
} Vitesse;

typedef struct CarStatus{
    Position position;
    Vitesse vitesse;
    bool panne;
    int ID;
} CarStatus;

CarStatus JSONToCarStatus(json data);
json CarStatusToJSON(CarStatus carStatus);

#endif // CAR_STATUS_H