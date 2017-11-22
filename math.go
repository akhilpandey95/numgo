/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"math"
	"strings"
)

func Sin(measure string, n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, math.Sin(math.Pi*float64(parameters[i])/180))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, math.Sin(float64(parameters[i])))
		}
	}
	return result
}

func Cos(measure string, n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, math.Cos(math.Pi*float64(parameters[i])/180))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, math.Cos(float64(parameters[i])))
		}
	}
	return result
}

func Tan(measure string, n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, math.Tan(math.Pi*float64(parameters[i])/180))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, math.Tan(float64(parameters[i])))
		}
	}
	return result
}

func Cosec(measure string, n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, (1 / math.Sin(math.Pi*float64(parameters[i])/180)))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, (1 / math.Sin(float64(parameters[i]))))
		}
	}
	return result
}

func Sec(measure string, n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, (1 / math.Cos(math.Pi*float64(parameters[i])/180)))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, (1 / math.Cos(float64(parameters[i]))))
		}
	}
	return result
}

func Cot(measure string, n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, (1 / math.Tan(math.Pi*float64(parameters[i])/180)))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result.Data = append(result.Data, (1 / math.Tan(float64(parameters[i]))))
		}
	}
	return result
}

func Asin(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Asin(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Acos(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Acos(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Atan(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Atan(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Sinh(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Sinh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Cosh(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Cosh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Tanh(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Tanh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Asinh(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Asinh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Acosh(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Acosh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Atanh(n *NArray) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: n.Details,
		Index:   n.Index,
	}
	parameters := extract_parameters(n.Data)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Atanh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Sqrt(args ...float64) *NArray {
	var result = &NArray{
		Data:    make([]float64, 0),
		Details: []float64{1, float64(len(args))},
		Index:   make([]float64, 2),
	}
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result.Data = append(result.Data, math.Sqrt(float64(parameters[i])))
	}
	return result
}
