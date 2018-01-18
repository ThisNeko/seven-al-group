#ifndef CONTROLLER_H
#define CONTROLLER_H

#include <list>
#include "structs/car_status.hpp"
#include "structs/traffic_light_status.hpp"

class Controller {
    public:
        Controller(/* dependencies (wifi, car interface, driver interface)*/);
        void ControllerLoop();
    private:
        CarStatus m_carStatus;
        std::list<CarStatus> m_carsRegistry;
        std::list<TrafficLightStatus> m_trafficLightsRegistry;
};

#endif // CONTROLLER_H