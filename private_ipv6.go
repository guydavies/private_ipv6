package main

// create a random IPv6 private CIDR (/64)
//
// Usage: ./private_ipv6.py
//
// Author: Guy Davies
// Email: aguydavies@gmail.com

import (
	"fmt"
	"strings"
	"math"
	"math/rand"
)

func main() {
	// construct random private IPv6 CIDR block
	//
	// First construct the first group which must begin with 'fd' and append a
	// random byte
	//
	// Then iterate three times to construct three more random hex blocks of 4
	// hex digits (two bytes) each
	var hex_digits_per_group int = 4
	var hex_first_byte string = "fd"
	var hex_two_bytes []string
	var addr_all_groups []string
	hex_two_bytes = append(hex_two_bytes, hex_first_byte)
	var hex_byte string = generate_hex(2)
	hex_two_bytes = append(hex_two_bytes, hex_byte)
	var hex_two_byte_str string = strings.Join(hex_two_bytes, "")
	addr_all_groups = append(addr_all_groups, hex_two_byte_str)

	for i := 0; i < 3; i++ {
		addr_all_groups = append(addr_all_groups, generate_hex(hex_digits_per_group))
	}

	addr_all_groups = append(addr_all_groups, ":/64")
	var ipv6_address string = strings.Join(addr_all_groups[:], ":")
	fmt.Printf(ipv6_address)
	fmt.Println("")
}

func prepend_to_length(new_hex_number string, hex_digits_per_group int) string {
	// check the length of the produced string and prepend 0s to expected length
	if len(new_hex_number) < hex_digits_per_group {
		var hex_strings = [2]string{"0", new_hex_number}
		new_hex_number = strings.Join(hex_strings[:], "")
		var padded_hex_number string = prepend_to_length(new_hex_number, hex_digits_per_group)
		return padded_hex_number
	} else {
		return new_hex_number
	}
}

func generate_hex(hex_digits_per_group int) string {
	// function to generate a single hex number of length hex_digits_per_group
	// lots of rather ugly type conversions in order to be able to manipulate the objects
	var rand_hex_str string = fmt.Sprintf("%x", (rand.Intn(int(math.Round(math.Pow(2, (4.0*float64(hex_digits_per_group))))) - 1)))
	var padded_hex_number string = prepend_to_length(rand_hex_str, hex_digits_per_group)
	return padded_hex_number
}
