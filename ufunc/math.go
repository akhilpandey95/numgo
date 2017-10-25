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

func Sin(measure string, args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result = append(result, math.Sin(math.Pi*float64(parameters[i])/180))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result = append(result, math.Sin(float64(parameters[i])))
		}
	}
	return result
}

func Cos(measure string, args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result = append(result, math.Cos(math.Pi*float64(parameters[i])/180))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result = append(result, math.Cos(float64(parameters[i])))
		}
	}
	return result
}

func Tan(measure string, args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result = append(result, math.Tan(math.Pi*float64(parameters[i])/180))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result = append(result, math.Tan(float64(parameters[i])))
		}
	}
	return result
}

func Cosec(measure string, args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result = append(result, (1 / math.Sin(math.Pi*float64(parameters[i])/180)))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result = append(result, (1 / math.Sin(float64(parameters[i]))))
		}
	}
	return result
}

func Sec(measure string, args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result = append(result, (1 / math.Cos(math.Pi*float64(parameters[i])/180)))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result = append(result, (1 / math.Cos(float64(parameters[i]))))
		}
	}
	return result
}

func Cot(measure string, args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)
	measure = strings.ToLower(measure)

	// check the args and return the appropriate function result
	if strings.HasPrefix(measure, "d") || strings.Contains(measure, "deg") {
		for i := 0; i < len(parameters); i++ {
			result = append(result, (1 / math.Tan(math.Pi*float64(parameters[i])/180)))
		}
	} else {
		for i := 0; i < len(parameters); i++ {
			result = append(result, (1 / math.Tan(float64(parameters[i]))))
		}
	}
	return result
}

func Asin(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Asin(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Acos(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Acos(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Atan(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Atan(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Sinh(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Sinh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Cosh(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Cosh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Tanh(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Tanh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Asinh(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Asinh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Acosh(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Acosh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Atanh(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Atanh(math.Pi*float64(parameters[i])/180))
	}
	return result
}

func Sqrt(args ...float64) []float64 {
	var result []float64
	parameters := extract_parameters(args)

	// check the args and return the appropriate function result
	for i := 0; i < len(parameters); i++ {
		result = append(result, math.Sqrt(float64(parameters[i])))
	}
	return result
}
