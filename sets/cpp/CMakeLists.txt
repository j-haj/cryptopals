cmake_minimum_required(VERSION 3.0.2)
project(cryptopals)

set(SRC_DIR "${PROJECT_SOURCE_DIR}/src")
set(TEST_DIR "${PROJECT_SOURCE_DIR}/test")
set(CMAKE_CXX_STANDARD 11)

# Look for Google Test library
enable_testing()
find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})

# Include headers
include_directories("${PROJECT_SOURCE_DIR}/src")

set(SOURCES
  ${SRC_DIR}/set1/set1.cc)

set(HEADERS
  ${SRC_DIR}/set1/set1.h)

set(TEST_HEADERS
  ${TEST_DIR}/set1/set1_test.h)

# Create executables
add_executable(test_crypto ${TEST_DIR}/test.cc ${SOURCES} ${HEADERS} ${TEST_HEADERS})
target_link_libraries(test_crypto ${GTEST_BOTH_LIBRARIES})
