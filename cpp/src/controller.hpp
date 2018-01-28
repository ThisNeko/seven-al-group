#ifndef CONTROLLER_H
#define CONTROLLER_H

#include <map>
#include "structs/car_status.hpp"
#include "structs/traffic_light_status.hpp"
#include "utils/communication_channel.hpp"

class Controller {
    public:
        Controller(CommunicationChannel<CarStatus> *chanBroadcaster, CommunicationChannel<CarStatus> *chanReceiver/* , CommunicationChannel chanCar */)
        :chanBroadcaster(chanBroadcaster), chanReceiver(chanReceiver){}
        void ControllerLoop();
    private:
        CarStatus m_carStatus;
        std::map<int, CarStatus> m_carsRegistry;
        std::map<int, int> m_timeouts;
        std::map<int, TrafficLightStatus> m_trafficLightsRegistry;
        CommunicationChannel<CarStatus> *chanBroadcaster;
        CommunicationChannel<CarStatus> *chanReceiver;
        // CommunicationChannel chanCar;
};

#endif // CONTROLLER_H