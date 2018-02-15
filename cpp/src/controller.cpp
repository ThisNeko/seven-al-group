#include <unistd.h>
#include <iostream>
using namespace std;

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
    for (;;)
    {
        for(std::map<int,int>::iterator it = m_timeouts.begin(); it != m_timeouts.end(); ++it) {            
            if(m_timeouts[it->first]>0){
                m_timeouts[it->first]-=1;
                PrintToDriver("> Test\n");
            }
        }

        while (!chanReceiverCar->isEmpty())
        {
            CarStatus carStatus = chanReceiverCar->get();
            m_carsRegistry[carStatus.ID] = carStatus;
            m_timeouts[carStatus.ID] = 4;
        }
        
        while (!chanReceiverTrafficLight->isEmpty())
        {
            TrafficLightStatus trafficLightStatus = chanReceiverTrafficLight->get();
            m_trafficLightsRegistry[trafficLightStatus.ID] = trafficLightStatus;
        }
        
        CarStatus *selectedLead = SelectLead(m_carsRegistry, *m_carStatus);
        TrafficLightStatus *selectedTrafficLight = SelectTrafficLight(m_trafficLightsRegistry, *m_carStatus);
        if(selectedLead!=nullptr){
            bool lead_panne = analyseLead(*m_carStatus,selectedLead,m_timeouts[selectedLead->ID]);
        }

        Directions directions = ComputeDrivingDirections(*m_carStatus, selectedLead, selectedTrafficLight);
        chanDirections->put(directions);

        usleep(10000); 
        m_carStatus->position.X += (double)m_carStatus->vitesse.X / (3.6 * 100);
        
    }  
}