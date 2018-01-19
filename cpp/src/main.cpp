#include "controller.hpp"
#include <thread>
#include "io/broadcaster_wifi.hpp"
#include "io/receptor_wifi.hpp"
#include <iostream>

#include <utility>
#include <atomic>
#include <functional>
#include <chrono>

void start_controller()
{
	Controller controller;
	controller.ControllerLoop();
}

void test(int n){
	for(;;){
		n++;
	}
}

void f1(Broadcaster_wifi * b)
{
	b->broadcast();
}

void f2(Receptor_wifi * r)
{
	r->receptor();
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

	//std::thread t3(test,0);

	//threadController.join();
	//threadWifiBroadcaster.join();
	//threadWifiReceiver.join();
	//threadCarInterface.join();


	static Broadcaster_wifi broadcaster;
	static Receptor_wifi receptor;

	std::thread t1(f1,&broadcaster);
	std::thread t2(f2,&receptor);

	//t1.join();
	//t2.join();
	threadController.join();
	threadWifiBroadcaster.join();
	threadWifiReceiver.join();
	threadCarInterface.join();
	t1.join();
	t2.join();




	//return 0;
}