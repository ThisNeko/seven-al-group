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
    Controller controller;
	controller.ControllerLoop();

	threadController.join();

	return 0;
}