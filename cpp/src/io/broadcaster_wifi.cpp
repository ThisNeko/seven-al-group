#include"broadcaster_wifi.hpp"
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <netinet/in.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <iostream>
#include <unistd.h>
#define PORT 1234

using namespace std;
Broadcaster_wifi::Broadcaster_wifi(){}


void Broadcaster_wifi::BroadcasterLoop(CarStatus *carStatus)
{
    struct sockaddr_in address;
    int sock = 0, valread;
    struct sockaddr_in serv_addr;

    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0)
    {
        printf("\n Socket creation error \n");
        return;
    }
  
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(PORT);
      
    // Convert IPv4 and IPv6 addresses from text to binary form
    if(inet_pton(AF_INET, "127.0.0.1", &serv_addr.sin_addr)<=0) 
    {
        printf("\nInvalid address/ Address not supported \n");
        return;
    }
  
    if (connect(sock, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0)
    {
        printf("\nBroadcaster: Connection Failed \n");
        return;
    }

    for(;;){
        json infos = CarStatusToJSON(*carStatus);
        json j;
        j["TypeEnum"] = "VOITURE";
        j["Info"] = infos.dump();
        std::string str = j.dump() + "\n";
        std::cout << str << std::endl;
        int rez = send(sock , str.c_str() , str.size() , 0 );
        usleep(20000);
    }
    
}
