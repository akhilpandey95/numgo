/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"fmt"
)

type NdArray struct {
	Size int
}

type NArray struct {
	Data    []float64
	Details []float64
	Index   []float64
}

func (n NdArray) Init() []float64 {
	array := make([]float64, n.Size)
	return array
}

func extract_details(args []float64) ([]float64, float64) {
	var elements float64 = 1
	var parameters []float64

	// store the parameters in a list
	for _, p := range args {
		parameters = append(parameters, p)
		elements *= p
	}
	return parameters, elements
}

func extract_parameters(args []float64) []float64 {
	var parameters []float64

	// store the parameters in a list
	for _, p := range args {
		parameters = append(parameters, p)
	}
	return parameters
}

func NDArrayGenElements(details ...float64) []float64 {
	parameters := extract_parameters(details)
	if len(parameters) > 0 {
		dim := parameters[0]
		rem := parameters[1:]
		arr := []float64{}
		for i := 0; i < int(dim); i++ {
			arr = append(arr, NDArrayGenElements(rem...)...)
		}
		return arr
	} else {
		return []float64{0}
	}
}

func NDArray(details ...float64) (n *NArray) {
	parameters := extract_parameters(details)
	n = &NArray{
		Data:    NDArrayGenElements(details...),
		Details: parameters,
		Index:   make([]float64, len(parameters)),
	}
	return n
}

func (n *NArray) String() string {
	var array string = "numgo.NDArray("
	if len(n.Details) == 0 {
		array += "<nil>"
	}
	array += fmt.Sprintf("%v", n.Data) + ")"

	return array
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
