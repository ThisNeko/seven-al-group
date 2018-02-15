#include "car_status.hpp"
#include <iostream>
using namespace std;

using nlohmann::json;

CarStatus JSONToCarStatus(json data)
{
    CarStatus carStatus;
    carStatus.position.X = data["Position"]["X"];
    carStatus.position.Y = data["Position"]["Y"];
    carStatus.vitesse.X = data["Vitesse"]["X"];
    carStatus.vitesse.Y = data["Vitesse"]["Y"];
    carStatus.ID = data["ID"];
    carStatus.panne = data["Panne"];

    return carStatus;
}

json CarStatusToJSON(const CarStatus &carStatus)
{
    json j;
    j["Position"]["X"] = carStatus.position.X;
    j["Position"]["Y"] = carStatus.position.Y;
    j["Vitesse"]["X"] = carStatus.vitesse.X;
    j["Vitesse"]["Y"] = carStatus.vitesse.Y;
    j["ID"] = carStatus.ID;
    j["Panne"] = carStatus.panne;

    return j;
}