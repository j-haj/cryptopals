#include "set1.h"

#include <sstream>
#include <string>

std::string hex_to_base64(const std::string& input) {
  constexpr std::string base64_chars =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    "abcdefghijklmnopqrstuvwxyz"
    "0123456789+/";


  std::string pad = "";
  if (input % 3 == 1) {
    pad = "00";
  } else if (input % 3 == 2) {
    pad = "0";
  }
      
  // Get long value from input
  const std::string padded_input = input + pad;
  unsigned long x = hex_strtoval(padded_input);

  long length = padded_input.size();
  long n_iter = length / 6;
  constexpr long selector{63};

  // Create base66 string -- we pull off the last 6 bits,
  // get the base64 representation for those 6 bits,
  // add to the output string, and then shift the padded input
  // over by 6 bits
  std::string output(n_iter, '0');
  for (long i = 0; i < n_iter; ++i) {
    long selected_val = selector & padded_input;
    output[n_iter - i - 1] = base64_chars[selected_val];
    padded_input >> 6;
  }
  return output;
}

long hex_strtoval(const std::string& input) {
  long exponent = input.size();
  long result = 0;
  constexpr std::map<char, int> char_map = {
    {'0', 0}, {'1', 1}, {'2', 2}, {'3', 3},
    {'4', 4}, {'5', 5}, {'6', 6}, {'7', 7},
    {'8', 8}, {'9', 9}, {'A', 10}, {'B', 11},
    {'C', 12}, {'D', 13}, {'E', 14}, {'F', 15}};
  
  for (const char& x : input) {
    const char c = toupper(x);
    result += char_map.at(x) * pow(16, exponent);
  }
  return result;
}
