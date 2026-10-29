cmake_minimum_required(VERSION 3.35)
project(HTTPSecurityLab CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

find_package(Boost REQUIRED)

# add_library(security_core)

# target_sources(security_core
#     PUBLIC
#     FILE_SET CXX_MODULES
#     FILES
#     core/include/security/http/cookies/CookiesInterception.cppm
# )

add_executable(HTTPSecurityLab)
target_sources(HTTPSecurityLab PRIVATE core/src/main.cpp)
