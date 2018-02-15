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
    m_carStatus.position.Y = 0;
    m_carStatus.vitesse.X = 80;
    int count = 10;
    for (;;)
    {

        for(std::map<int,int>::iterator it = m_timeouts.begin(); it != m_timeouts.end(); ++it) {            
            if(m_timeouts[it->first]>0){
                m_timeouts[it->first]-=1;
                PrintToDriver("> Test\n");
            }
        }

        if(count>10){
            count--;
        }
        if(count==0){
            m_carStatus.position.Y = 1;
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
        
        CarStatus *selectedLead = SelectLead(m_carsRegistry, m_carStatus);
        TrafficLightStatus *selectedTrafficLight = SelectTrafficLight(m_trafficLightsRegistry, m_carStatus);
        if(selectedLead!=nullptr){

            bool lead_panne = analyseLead(m_carStatus,selectedLead,m_timeouts[selectedLead->ID]);
        }

        Directions directions;

        if (!ComputeDrivingDirections(m_carStatus, selectedLead, selectedTrafficLight))
        {
           // PrintToDriver("> Controller: No directions to give. Drive as you want!\n");
        }
        else
        {
            //PrintToDriver("> Controller: Directions found but not printable yet. Come back later!\n");
        }

        sleep(1); 
    }  
}