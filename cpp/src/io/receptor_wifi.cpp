#include"receptor_wifi.hpp"
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <netinet/in.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#define PORT 1234

Receptor_wifi::Receptor_wifi(){}

  
void Receptor_wifi::receptor()
{
    struct sockaddr_in address;
    int sock = 0, valread;
    struct sockaddr_in serv_addr;
   // char *hello = "Hello from receptor";
    char *hello = "{\"TypeEnum\":\"VOITURE\",\"Info\":\"{\\\"ID\\\":9113953410437231233,\\\"Vitesse\\\":{\\\"X\\\":80,\\\"Y\\\":0},\\\"Position\\\":{\\\"X\\\":-20,\\\"Y\\\":0},\\\"Panne\\\":false}\"}\n";

    char buffer[1024] = {0};

    printf("coucou\n");
    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0)
    {
        printf("\n Socket creation error \n");
        return;
    }
  
    memset(&serv_addr, '0', sizeof(serv_addr));
  
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
    	valread = read( sock , buffer, 1024);
    	printf("Receptor : %s\n",buffer );
        sleep(1);
    }

}