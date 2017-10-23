/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"log"
	"strings"
)

func extract_strings(args []string) []string {
	var strings []string

	// store the parameters in a slice
	for _, s := range args {
		strings = append(strings, s)
	}
	return strings
}

func Strcount(array_one []string, args ...string) []int {
	var result []int
	str_args := extract_strings(args)

	// check the character count and return the appropriate result
	for i := 0; i < len(array_one); i++ {
		result = append(result, strings.Count(array_one[i], str_args[0]))
	}
	return result
}

func Strjoin(delimiter string, array_one, array_two []string) []string {
	var result []string

	//check the delimiter and join the arrays appropriately
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: the shapes of the arrays must be similar to perform Strjoin operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			result = append(result, array_one[i]+delimiter+array_two[i])
		}
	}
	return result
}

func Strtoupper(array_one []string) []string {
	var result []string

	//take the args and convert every character/word of the array to uppwercase
	for i := 0; i < len(array_one); i++ {
		result = append(result, strings.ToUpper(array_one[i]))
	}
	return result
}

func Strtolower(array_one []string) []string {
	var result []string

	//take the args and covert every character/word of the array to lowercase
	for i := 0; i < len(array_one); i++ {
		result = append(result, strings.ToLower(array_one[i]))
	}
	return result
}
