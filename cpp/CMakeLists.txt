cmake_minimum_required(VERSION 3.1)

set (CMAKE_CXX_STANDARD 11)
set (CMAKE_CXX_STANDARD_REQUIRED ON)
ENABLE_LANGUAGE(CXX)

set (THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)

project(al-car)

include_directories(src)

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
    src/structs/car_status.hpp
    src/structs/car_status.cpp
    src/structs/traffic_light_status.hpp
    src/structs/directions.hpp
    src/io/broadcaster_wifi.hpp
    src/io/broadcaster_wifi.cpp
    src/io/receptor_wifi.hpp
    src/io/receptor_wifi.cpp
    src/io/car_interface.hpp
    src/io/car_interface.cpp
    src/io/driver_interface.hpp
    src/io/driver_interface.cpp
    src/utils/communication_channel.hpp
    src/utils/json.hpp)

target_link_libraries(al-car Threads::Threads)