/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
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
