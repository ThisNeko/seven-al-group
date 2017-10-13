cmake_minimum_required(VERSION 3.8)
project(seven_al_group)

set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES main.cpp donnees/StatusVoiture.cpp donnees/StatusVoiture.h modules/FreinageUrgence.cpp modules/FreinageUrgence.h interfaces/InterfaceMateriel.cpp interfaces/InterfaceMateriel.h interfaces/InterfaceConducteur.cpp interfaces/InterfaceConducteur.h communication/Broadcasteur.cpp communication/Broadcasteur.h communication/Recepteur.cpp communication/Recepteur.h donnees/Registre.cpp donnees/Registre.h)
add_executable(seven_al_group ${SOURCE_FILES})