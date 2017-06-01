#include "set1.h"

#include <cctype>
#include <map>
#include <sstream>
#include <stdexcept>

std::string hex_to_base64(const std::string& input) {
  std::string base64_chars =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    "abcdefghijklmnopqrstuvwxyz"
    "0123456789+/";


  std::string pad = "";
  if (input.size() % 3 == 1) {
    pad = "00";
  } else if (input.size() % 3 == 2) {
    pad = "0";
  }
      
  // Get long value from input
  const std::string padded_input = input + pad;
  std::vector<uint8_t> bytes = hex_to_bytes(padded_input);

  long length = padded_input.size();
  long n_iter = length / 6;
  constexpr long selector{63};

  // Create base66 string -- we pull off the last 6 bits,
  // get the base64 representation for those 6 bits,
  // add to the output string, and then shift the padded input
  // over by 6 bits
  std::string output(n_iter, '0');
  for (long i = 0; i < n_iter; ++i) {
    long selected_val = selector & bytes[i];
    output[n_iter - i - 1] = base64_chars[selected_val];
    //padded_input = padded_input >> 6;
  }
  return output;
}

std::vector<uint8_t> hex_to_bytes(const std::string& s) {
  if (s.size() % 2 == 1) {
    throw std::runtime_error("Cannot convert an odd number of hex values to bytes\n"); 
  }
  size_t length = s.size() / 2;
  std::vector<uint8_t> output(length);
  std::istringstream ss(s);
  for (size_t i = s.size() - 1; i > 0; i -= 2) {
    std::istringstream ss(s.substr(i - 1, 2));
    uint8_t b;
    ss >> std::hex >> b;
    output[(i+1)/2 - 1] = b;
  }
  return output;
}

