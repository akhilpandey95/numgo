/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numGo

type NdArray struct {
	Size int
}

func (n NdArray) Init() []int {
	array := make([]int, n.Size)
	return array
}

func (n NdArray) Range(start, end int) []int {
	var array []int
	for i := start; i < end; i++ {
		array = append(array, i)
	}
	return array
}

func (n NdArray) Xrange(args ...int) []int {
	var array []int
	var parameters []int
	number_of_args := len(args)

	// store the parameters in a slice
	for _, p := range args {
		parameters = append(parameters, p)
	}

	// now generate the array basiing on the number of arguments
	if number_of_args == 1 {
		for i := 0; i < parameters[0]; i++ {
			array = append(array, i)
		}
	} else if number_of_args == 2 {
		for i := parameters[0]; i < parameters[1]; i++ {
			array = append(array, i)
		}
	} else {
		for i := parameters[0]; i < parameters[1]; i++ {
			array = append(array, i)
		}
	}
	return array
}
