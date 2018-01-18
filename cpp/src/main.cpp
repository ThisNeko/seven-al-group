#include "controller.hpp"
#include <thread>

void start_controller()
{
	Controller controller;
	controller.ControllerLoop();
}

void start_wifi_broadcaster()
{

}

void start_wifi_receiver()
{

}

void start_car_interface()
{

}

int main(){
	std::thread threadController(start_controller);
    std::thread threadWifiBroadcaster(start_wifi_broadcaster);
	std::thread threadWifiReceiver(start_wifi_receiver);
	std::thread threadCarInterface(start_car_interface);

	threadController.join();
	threadWifiBroadcaster.join();
	threadWifiReceiver.join();
	threadCarInterface.join();

	return 0;
}