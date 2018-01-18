cmake_minimum_required(VERSION 3.0)

project(al-car)

add_executable(al-car 
    src/controller.cpp
    src/controller.hpp
    src/driving_directions.cpp
    src/driving_directions.hpp
    src/lead_analyzer.cpp
    src/lead_analyzer.hpp
    src/main.cpp
    src/traffic_lights_analyzer.cpp
    src/traffic_lights_analyzer.hpp
    src/io/broadcaster_wifi.hpp
    src/io/broadcaster_wifi.cpp
    src/io/receptor_wifi.hpp
    src/io/receptor_wifi.cpp)