#include "structs/traffic_light_status.hpp"
#include <iostream>
using namespace std;


using nlohmann::json;

TrafficLightStatus JSONToTrafficLightStatus(json data)
{
    TrafficLightStatus trafficLightStatus;
    trafficLightStatus.ID = data["ID"];
    trafficLightStatus.ticker = data["Ticker"];
    trafficLightStatus.timer = data["Timer"];
    trafficLightStatus.pos.X = data["Position"]["X"];
    trafficLightStatus.pos.Y = data["Position"]["Y"];
    trafficLightStatus.couleur = data["Couleur"];

    return trafficLightStatus;
}