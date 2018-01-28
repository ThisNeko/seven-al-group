#include <unistd.h>

#include "controller.hpp"
#include "driving_directions.hpp"
#include "lead_analyzer.hpp"
#include "traffic_lights_analyzer.hpp"
#include "structs/directions.hpp"
#include "io/driver_interface.hpp"

void Controller::ControllerLoop()
{   
    for (;;)
    {
        while (!chanReceiver->isEmpty())
        {
            CarStatus carStatus = chanReceiver->get();
            m_carsRegistry[carStatus.ID] = carStatus;
            // m_carsRegistry.put(chanReceiver->get());
        }
        
        CarStatus *selectedLead = SelectLead(m_carsRegistry, m_carStatus);
        TrafficLightStatus *selectedTrafficLight = SelectTrafficLight(m_trafficLightsRegistry, m_carStatus);

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