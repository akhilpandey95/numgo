/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"log"
)

func extract_parameters(args []float64) []float64 {
	var parameters []float64

	// store the parameters in a slice
	for _, p := range args {
		parameters = append(parameters, p)
	}
	return parameters
}

func Equal(array_one, array_two []float64) []bool {
	var result []bool

	// check the args and return the appropriate function result
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: Shapes of two arrays must be similar to perform Equal() operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			if array_one[i] == array_two[i] {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}
	return result
}

func Not_Equal(array_one, array_two []float64) []bool {
	var result []bool

	// check the args and return the appropriate function result
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: Shapes of two arrays must be similar to perform Not_Equal operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			if array_one[i] != array_two[i] {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}
	return result
}

func Greater(array_one, array_two []float64) []float64 {
	var result []float64

	// check the args and return the appropriate function result
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: Shapes of two arrays must be similar to perform Greater operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			if array_one[i] >= array_two[i] {
				result = append(result, array_one[i])
			} else {
				result = append(result, array_two[i])
			}
		}
	}
	return result
}

func Greater_B(array_one, array_two []float64) []bool {
	var result []bool

	// check the args and return the appropriate function result
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: Shapes of two arrays must be similar to perform Greater operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			if array_one[i] >= array_two[i] {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}
	return result
}

func Lesser(array_one, array_two []float64) []float64 {
	var result []float64

	//check the args and return the appropriate function result
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: Shapes of two arrays must be similar to perform Lesser operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			if array_one[i] <= array_two[i] {
				result = append(result, array_one[i])
			} else {
				result = append(result, array_two[i])
			}
		}
	}
	return result
}

func Lesser_B(array_one, array_two []float64) []bool {
	var result []bool

	//check the args and return the appropriate function result
	if len(array_one) != len(array_two) {
		log.Fatal("[ERROR]: Shapes of two arrays must be similar to perform Lesser operation")
	} else {
		for i := 0; i < len(array_one); i++ {
			if array_one[i] <= array_two[i] {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}
	return result
}
