/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type NArray struct {
	Data    []float64
	Details []float64
	Index   []float64
}

func Init(length float64) (n *NArray) {
	n = &NArray{
		Data:    make([]float64, 0),
		Details: []float64{1, length},
		Index:   make([]float64, 2),
	}
	return n
}

func Round(number, decimal, precision float64) float64 {
	var result float64
	var zeros float64 = math.Pow(10, precision)
	var new_number float64 = zeros * number
	_, mod_decimal := math.Modf(new_number)
	if mod_decimal >= decimal {
		result = math.Ceil(new_number)
	} else {
		result = math.Floor(new_number)
	}
	result /= zeros
	return result
}

func RowColFormula(details []float64) float64 {
	var result float64 = 1
	for x := 0; x < len(details); x++ {
		result *= details[x]
	}
	return result
}

func GetIndex(n *NArray, details ...float64) float64 {
	var result float64 = 0
	parameters := extract_parameters(details)
	if len(parameters) == len(n.Details) {
		for x := 0; x < len(n.Details); x++ {
			result += parameters[x] * RowColFormula(n.Details[x+1:])
		}
		return result
	} else {
		log.Fatal("incorrect shape given as index")
		return -1
	}
}

func extract_parameters(args []float64) []float64 {
	var parameters []float64

	// store the parameters in a list
	for _, p := range args {
		parameters = append(parameters, p)
	}
	return parameters
}

func NDArrayGenElements(value float64, details ...float64) []float64 {
	parameters := extract_parameters(details)
	if len(parameters) > 0 {
		dim := parameters[0]
		rem := parameters[1:]
		arr := []float64{}
		for i := 0; i < int(dim); i++ {
			arr = append(arr, NDArrayGenElements(value, rem...)...)
		}
		return arr
	} else {
		return []float64{value}
	}
}

func NDArray(details ...float64) (n *NArray) {
	parameters := extract_parameters(details)
	n = &NArray{
		Data:    NDArrayGenElements(0, details...),
		Details: parameters,
		Index:   make([]float64, len(parameters)),
	}
	return n
}

func (n *NArray) String() (array string) {
	array = "NDArray>\n("
	if len(n.Details) == 0 {
		return "NDArray>(<nil>)"
	} else {
		var i float64 = 0
		var details []float64
		var index_number float64 = n.Details[len(n.Details)-1]
		var element float64 = index_number

		// store the rows and cols number for adding [ and ] symbols
		// at appropriate places
		for index := len(n.Details) - 1; index > 0; index-- {
			details = append(details, index_number*n.Details[index-1])
			index_number = index_number * n.Details[index-1]
		}
		for i = 0; i+element <= float64(len(n.Data)); i += element {
			var array_symbol string = ""
			// add the '[' symbol at the start of every new dimension
			for _, d := range details {
				if int(i)%int(d) == 0 {
					array_symbol += "["
				}
			}
			array += strings.Repeat(" ", len(n.Details)-len(array_symbol)-1) + array_symbol

			// add the data between the '[' ']' symbols
			array += fmt.Sprint(n.Data[int(i):int(i+element)])

			array_symbol = ""

			// add the ']' symbol at the end of every new dimension
			for _, d := range details {
				if int((i+element))%int(d) == 0 {
					array_symbol += "]"
				}
			}
			array += array_symbol + strings.Repeat(" ", (len(n.Details)-len(array_symbol))-1)

			// add proper spacing for displaying arrays in presentable manner
			if i+element != float64(len(n.Data)) {
				if n.Details[0] == 1 && len(n.Details) == 2 {
					array += ""
				} else {
					if len(array_symbol) > 0 {
						array += "\n"
					}
					array += "\n"
				}
			}
		}
	}
	return array + ")"
}

func (n *NArray) GetElement(details ...float64) float64 {
	var result float64
	if int(GetIndex(n, details...)) < len(n.Data) {
		result = GetIndex(n, details...)
		return n.Data[int(result)]
	} else {
		log.Fatal("Incorrect Shape passed as an argument")
		return -1
	}
}

func (n *NArray) SetElement(details []float64, value float64) *NArray {
	var result float64
	if int(GetIndex(n, details...)) < len(n.Data) {
		result = GetIndex(n, details...)
		n.Data[int(result)] = value
	} else {
		log.Fatal("Incorrect Shape passed as an argument")
	}
	return n
}

func (n *NArray) FillElements(values []float64) *NArray {
	if len(n.Data) > len(values) {
		for i := 0; i < len(values); i++ {
			n.Data[i] = values[i]
		}
	} else {
		log.Fatal("Cannot accomodate the NArray with the given size of values")
	}
	return n
}

func Range(start, end float64) (n *NArray) {
	n = Init(end - start)
	for i := start; i < end; i++ {
		n.Data = append(n.Data, i)
	}
	return n
}

func Xrange(args ...float64) (n *NArray) {
	var i float64
	number_of_args := len(args)
	parameters := extract_parameters(args)

	// now generate the array basiing on the number of arguments
	if number_of_args == 1 {
		n = Init(parameters[0])
		for i = 0; i < parameters[0]; i++ {
			n.Data = append(n.Data, i)
		}
	} else if number_of_args == 2 {
		n = Init(parameters[1] - parameters[0])
		for i = parameters[0]; i < parameters[1]; i++ {
			n.Data = append(n.Data, i)
		}
	} else if number_of_args == 3 {
		n = Init((parameters[1] - parameters[0]) / parameters[2])
		for i = parameters[0]; i < parameters[1]; i += parameters[2] {
			n.Data = append(n.Data, Round(i, 0.5, 1))
		}
	} else if number_of_args == 4 {
		n = Init((parameters[1] - parameters[0]) / parameters[2])
		for i = parameters[0]; i < parameters[1]; i += parameters[2] {
			n.Data = append(n.Data, Round(i, 0.5, parameters[3]))
		}
	} else {
		n = Init((parameters[1] - parameters[0]) / parameters[2])
		for i = parameters[0]; i < parameters[1]; i += parameters[2] {
			n.Data = append(n.Data, Round(i, 0.5, 1))
		}
	}
	return n
}
