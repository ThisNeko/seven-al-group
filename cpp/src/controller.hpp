#ifndef CONTROLLER_H
#define CONTROLLER_H

#include <map>
#include "structs/car_status.hpp"
#include "structs/traffic_light_status.hpp"
#include "structs/directions.hpp"
#include "utils/communication_channel.hpp"

const int msBetweenTicks = 10;

class Controller {
    public:
        Controller(CommunicationChannel<CarStatus> *chanReceiverCar, 
                   CommunicationChannel<TrafficLightStatus> *chanReceiverTrafficLight, CommunicationChannel<Directions> *chanDirections,
                   CarStatus *carStatus)
        :chanReceiverCar(chanReceiverCar), chanReceiverTrafficLight(chanReceiverTrafficLight),
        chanDirections(chanDirections), m_carStatus(carStatus){}
        void ControllerLoop();
    private:
        CarStatus *m_carStatus;
        std::map<int, CarStatus> m_carsRegistry;
        std::map<int, TrafficLightStatus> m_trafficLightsRegistry;
        std::map<int, int> m_timeouts;
        CommunicationChannel<CarStatus> *chanReceiverCar;
        CommunicationChannel<TrafficLightStatus> *chanReceiverTrafficLight;
        CommunicationChannel<Directions> *chanDirections;
};

#endif // CONTROLLER_H