#ifndef __SET_1_H
#define __SET_1_H

#include <vector>
#include <string>

/**
 * Takes a hexadecimal input string and returns a string
 * in base64
 *
 * @param input hex input string
 *
 * @return base64 string equivalent to @p input
 */
std::string hex_to_base64(const std::string& input);

/**
 * Takes a hex string and returns the base 10 equivalent
 *
 * @param input input value of hexstring
 *
 * @return an array of bytes
 */
std::vector<uint8_t> hex_strtoval(const std::string& input);

#endif // __SET_1_H
