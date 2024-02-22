cmake_minimum_required(VERSION 3.27)
project(pybinddemo)

set(CMAKE_CXX_STANDARD 20)

add_subdirectory(pybind11)

add_executable(pybinddemo main.cpp)

target_link_libraries(pybinddemo PRIVATE pybind11::embed)

if(Python_EXECUTABLE)
    message(STATUS "Python used: ${Python_EXECUTABLE}")
else()
    message(FATAL_ERROR "Python not found.")
endif()
