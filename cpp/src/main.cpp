#include "controller.hpp"
#include <thread>
#include "io/broadcaster_wifi.hpp"
#include "io/receptor_wifi.hpp"
#include "utils/json.hpp"

#include <iostream>
#include <utility>
#include <atomic>
#include <functional>
#include <chrono>
#include <unistd.h>

using json = nlohmann::json;

void start_controller()
{
	Controller controller;
	controller.ControllerLoop();
}

void start_wifi_broadcaster()
{
	Broadcaster_wifi broadcaster;
	usleep(1000000);
	broadcaster.BroadcasterLoop();
}

void start_wifi_receiver()
{
	Receptor_wifi receptor;
	receptor.ReceptorLoop();
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