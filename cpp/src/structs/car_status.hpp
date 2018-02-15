#ifndef CAR_STATUS_H
#define CAR_STATUS_H

#include "utils/json.hpp"

using nlohmann::json;

typedef struct {
    int X = 0;
    int Y = 0;
} Position;

typedef struct {
    int X = 0;
    int Y = 0;
} Vitesse;

typedef struct {
    int X = 0;
    int Y = 0;
} Acceleration;

typedef struct CarStatus{
    Position position;
    Vitesse vitesse;
    bool panne;
    int ID;

    // CarStatus(json data){
        
    // }
} CarStatus;

CarStatus JSONToCarStatus(json data);

#endif // CAR_STATUS_H