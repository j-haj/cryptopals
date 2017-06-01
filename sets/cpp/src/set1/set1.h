#ifndef __SET_1_H
#define __SET_1_H

#include <string>
#include <vector>

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
 * Takes a hex string and returns an @p std::array<uint8_t> of bytes. Note that
 * the input string must have an even number of hex values
 *
 * @param s input string of hexadecimal characters
 *
 * @return an array of bytes
 */
std::vector<uint8_t> hex_to_bytes(const std::string& s);

#endif // __SET_1_H
