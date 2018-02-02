#include "car_status.hpp"

using nlohmann::json;

CarStatus JSONToCarStatus(json data){
    CarStatus carStatus;
    int i = data["Position"]["X"];
    carStatus.position.X = data["Position"]["X"];
    carStatus.position.Y = data["Position"]["Y"];
    carStatus.vitesse.X = data["Vitesse"]["X"];
    carStatus.vitesse.Y = data["Vitesse"]["Y"];
    carStatus.ID = data["ID"];
    carStatus.panne = data["Panne"];

    return carStatus;
}