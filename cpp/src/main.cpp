#include "controller.hpp"
#include <thread>
#include "io/broadcaster_wifi.hpp"
#include "io/receptor_wifi.hpp"
#include "utils/communication_channel.hpp"
#include "utils/json.hpp"

#include <iostream>
#include <utility>
#include <atomic>
#include <functional>
#include <chrono>
#include <unistd.h>

using json = nlohmann::json;

void start_controller(CommunicationChannel<CarStatus> *chanControllerBroadcaster, CommunicationChannel<CarStatus> *chanControllerReceiver/* , CommunicationChannel chanControllerCar */)
{
	Controller controller(chanControllerBroadcaster, chanControllerReceiver);
	controller.ControllerLoop();
}

void start_wifi_broadcaster(CommunicationChannel<CarStatus> *chanControllerBroadcaster)
{
	Broadcaster_wifi broadcaster;
	usleep(1000000);
	broadcaster.BroadcasterLoop(chanControllerBroadcaster);
}

void start_wifi_receiver(CommunicationChannel<CarStatus> *chanControllerReceiver)
{
	Receptor_wifi receptor;
	receptor.ReceptorLoop(chanControllerReceiver);
}

void start_car_interface()
{

}

int main(){
	CommunicationChannel<CarStatus> *chanControllerBroadcaster = new CommunicationChannel<CarStatus>();
	CommunicationChannel<CarStatus> *chanControllerReceiver = new CommunicationChannel<CarStatus>;
	// chanControllerCar;
	std::thread threadController(start_controller, chanControllerBroadcaster, chanControllerReceiver);
    std::thread threadWifiBroadcaster(start_wifi_broadcaster, chanControllerBroadcaster);
	std::thread threadWifiReceiver(start_wifi_receiver, chanControllerReceiver);
	std::thread threadCarInterface(start_car_interface);

	threadController.join();
	threadWifiBroadcaster.join();
	threadWifiReceiver.join();
	threadCarInterface.join();

	return 0;
}