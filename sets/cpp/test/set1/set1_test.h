#ifndef __SET1_TEST_H
#define __SET1_TEST_H

#include <random>
#include <string>
#include <vector>

#include <gtest/gtest.h>

#include "set1/set1.h"

namespace constants {
  constexpr int kNumberOfTests = 1000;
}

TEST(Set1, RandomByte) {
  // Generate a random byte, convert it to a hex string, then parse hex string
  // and compare resulting byte array with initial value
  std::random_device r;
  std::default_random_engine e(r());
  std::uniform_int_distribution<int> dist(0, 255);

  for (int i = 0; i < constants::kNumberOfTests; ++i) {
    int random_byte = dist(e);
    std::stringstream ss;
    ss << std::hex << random_byte;
    std::vector<uint8_t> res = hex_to_bytes(ss.str());
    ASSERT_EQ(res[0], random_byte);
  }
}


#endif // __SET1_TEST_H
