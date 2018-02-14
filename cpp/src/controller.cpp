#include <unistd.h>

#include "controller.hpp"
#include "driving_directions.hpp"
#include "lead_analyzer.hpp"
#include "traffic_lights_analyzer.hpp"
#include "structs/directions.hpp"
#include "structs/traffic_light_status.hpp"
#include "io/driver_interface.hpp"
#include "breakdown_analyser.hpp"

void Controller::ControllerLoop()
{   
    m_carStatus.position.X = 91;
    m_carStatus.vitesse.X = 80;
    for (;;)
    {
        while (!chanReceiverCar->isEmpty())
        {
            CarStatus carStatus = chanReceiverCar->get();
            m_carsRegistry[carStatus.ID] = carStatus;
        }

        while (!chanReceiverTrafficLight->isEmpty())
        {
            TrafficLightStatus trafficLightStatus = chanReceiverTrafficLight->get();
            m_trafficLightsRegistry[trafficLightStatus.ID] = trafficLightStatus;
        }
        
        CarStatus *selectedLead = SelectLead(m_carsRegistry, m_carStatus);
        TrafficLightStatus *selectedTrafficLight = SelectTrafficLight(m_trafficLightsRegistry, m_carStatus);
        if(selectedLead!=nullptr){
            bool lead_panne = analyseLead(m_carStatus,selectedLead);
        }

        Directions directions;

        if (!ComputeDrivingDirections(m_carStatus, selectedLead, selectedTrafficLight))
        {
            PrintToDriver("> Controller: No directions to give. Drive as you want!\n");
        }
        else
        {
            PrintToDriver("> Controller: Directions found but not printable yet. Come back later!\n");
        }

        sleep(1); 
    }  
}