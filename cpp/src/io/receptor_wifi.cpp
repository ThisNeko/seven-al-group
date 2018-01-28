#include "receptor_wifi.hpp"
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <netinet/in.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <iostream>
#include "utils/json.hpp"
#include <algorithm>

// for convenience
using json = nlohmann::json;
#define PORT 1234

Receptor_wifi::Receptor_wifi(){}

  
void Receptor_wifi::ReceptorLoop(CommunicationChannel<CarStatus> *chan)
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
        printf("\nConnection Failed \n");
        return;
    }

    for(;;){
        char buffer[1024] = {0};
    	valread = read( sock , buffer, 1024);
        if (json::accept(buffer)){
            std::string str(buffer);
            auto j = json::parse(str);
            if (j["TypeEnum"] == "VOITURE")
            {
                str = j["Info"].dump();
                str.erase(std::remove(str.begin(), str.end(), '\\'), str.end());
                str.erase(0, 1);
                str.erase(str.size() - 1);
                j = json::parse(str);
                std::cout << j.dump() << std::endl;
                CarStatus s = JSONToCarStatus(j);
                chan->put(s);
            }
        }
    }

}