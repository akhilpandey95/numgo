/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"math"
)

func Sin(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Sin(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Cos(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Cos(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Tan(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Tan(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Sinh(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Sinh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Cosh(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Cosh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Tanh(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Tanh(math.Pi*float64(parameters[i])/180))
	}
	return result
}
func Sqrt(args ...int) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Sqrt(float64(parameters[i])))
	}
	return result
}
