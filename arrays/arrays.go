/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

type NdArray struct {
	Size int
}

func (n NdArray) Init() []float64 {
	array := make([]float64, n.Size)
	return array
}

func extract_parameters(args []float64) []float64 {
	var parameters []float64

	// store the parameters in a list
	for _, p := range args {
		parameters = append(parameters, p)
	}
	return parameters
}

func (n NdArray) Array(data []string, details ...int) (narray *NdArray) {
	//parameters := extract_parameters(details)

	// depending on the given parameters create the appropriate array
	return narray
}

func (n NdArray) Range(start, end float64) []float64 {
	var array []float64
	for i := start; i < end; i++ {
		array = append(array, i)
	}
	return array
}

func (n NdArray) Xrange(args ...float64) []float64 {
	var i float64
	var array []float64
	number_of_args := len(args)
	parameters := extract_parameters(args)

	// now generate the array basiing on the number of arguments
	if number_of_args == 1 {
		for i = 0; i < parameters[0]; i++ {
			array = append(array, i)
		}
	} else if number_of_args == 2 {
		for i = parameters[0]; i < parameters[1]; i++ {
			array = append(array, i)
		}
	} else {
		for i = parameters[0]; i < parameters[1]; i++ {
			array = append(array, i)
		}
	}
	return array
}
