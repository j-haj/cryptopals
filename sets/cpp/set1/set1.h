#ifndef __SET_1_H
#define __SET_1_H

/**
 * Takes a hexadecimal input string and returns a string
 * in base64
 *
 * @param input hex input string
 *
 * @return base64 string equivalent to @p input
 */
std::string hex_to_base64(cost std::string& input);

/**
 * Takes a hex string and returns the base 10 equivalent
 *
 * @param input input value of hexstring
 *
 * @return value of hex string
 */
long hex_strtoval(const std::string& input);

#endif // __SET_1_H
