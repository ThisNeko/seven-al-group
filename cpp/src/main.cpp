#include "controller.hpp"
#include <thread>
#include "io/broadcaster_wifi.hpp"
#include "io/receptor_wifi.hpp"
#include "io/car_interface.hpp"
#include "io/driver_interface.hpp"
#include "utils/communication_channel.hpp"
#include "utils/json.hpp"
#include "structs/directions.hpp"

#include <iostream>
#include <utility>
#include <atomic>
#include <functional>
#include <chrono>
#include <unistd.h>
#include <string>

using namespace std;

using json = nlohmann::json;

void start_controller(CommunicationChannel<CarStatus> *chanControllerReceiverCar,
					  CommunicationChannel<TrafficLightStatus> *chanControllerReceiverTrafficLight, CommunicationChannel<Directions> *chanControllerCarInterface,
					  CarStatus *carStatus)
{
	Controller controller(chanControllerReceiverCar, chanControllerReceiverTrafficLight, chanControllerCarInterface, carStatus);
	controller.ControllerLoop();
}

void start_wifi_broadcaster(CarStatus *carStatus)
{
	Broadcaster_wifi broadcaster;
	usleep(1000000);
	broadcaster.BroadcasterLoop(carStatus);
}

void start_wifi_receiver(CommunicationChannel<CarStatus> *chanControllerReceiverCar,
						 CommunicationChannel<TrafficLightStatus> *chanControllerReceiverTrafficLight, int ignoreId)
{
	Receptor_wifi receptor;
	receptor.ReceptorLoop(chanControllerReceiverCar, chanControllerReceiverTrafficLight, ignoreId);
}

void start_car_interface(CarStatus *carStatus)
{
	CarInterfaceLoop(carStatus);
}

void start_follow_directions(CommunicationChannel<Directions> *chanControllerFollowDirections,
						 CarStatus *carStatus)
{
	FollowDirectionsLoop(chanControllerFollowDirections, carStatus);
}

void print_status(CarStatus *carStatus)
{
	return;
	while(true)
	{
		cout << "X: " << carStatus->position.X << ", vitesse " << carStatus->vitesse.X << endl;
		sleep(1);
	}
}

int main(int argc, char *argv[]){
	CommunicationChannel<CarStatus> *chanControllerReceiverCar 						= new CommunicationChannel<CarStatus>;
	CommunicationChannel<TrafficLightStatus> *chanControllerReceiverTrafficLight	= new CommunicationChannel<TrafficLightStatus>;
	CommunicationChannel<Directions> *chanControllerFollowDirections 				= new CommunicationChannel<Directions>;

	CarStatus *carStatus = new CarStatus;
	carStatus->ID = 5;
	carStatus->vitesse.X = 0;
	carStatus->vitesse.Y = 0;
	carStatus->position.X = std::stod(argv[1]);
	carStatus->position.Y = std::stod(argv[2]);

	std::thread threadController(start_controller, chanControllerReceiverCar, chanControllerReceiverTrafficLight,
								 chanControllerFollowDirections, carStatus);
    std::thread threadWifiBroadcaster	(start_wifi_broadcaster, carStatus);
	std::thread threadWifiReceiver		(start_wifi_receiver, chanControllerReceiverCar, chanControllerReceiverTrafficLight, carStatus->ID);
	std::thread threadCarInterface		(start_car_interface, carStatus);
	std::thread threadFollowDirections	(start_follow_directions, chanControllerFollowDirections, carStatus);
	std::thread printStatus	(print_status, carStatus);

	threadController.join();
	threadWifiBroadcaster.join();
	threadWifiReceiver.join();
	threadCarInterface.join();
	threadFollowDirections.join();
	printStatus.join();

	return 0;
}