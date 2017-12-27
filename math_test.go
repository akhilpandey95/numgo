/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
	"math"
	"reflect"
	"testing"
)

func TestSin(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{1}
	output_two := []float64{0}

	// the function inputs
	input_one := Sin("deg", Xrange(90, 91, 1))
	input_two := Sin("rad", Xrange(0, 1, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestCos(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{1}

	// the function inputs
	input_one := Cos("deg", Xrange(90, 91, 1))
	input_two := Cos("rad", Xrange(0, 1, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(Round(input_one.Data[0], 0.5, 1), output_one[0])) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestTan(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{1.6331239353195392e+16}
	output_two := []float64{0}

	// the function inputs
	input_one := Tan("deg", Xrange(90, 91, 1))
	input_two := Tan("deg", Xrange(0, 1, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestCot(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{math.Inf(0)}
	output_two := []float64{-0.5012027833801532}

	// the function inputs
	input_one := Cot("deg", Xrange(0, 1, 1))
	input_two := Cot("rad", Xrange(90, 91, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestAsin(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{0.01745417873758517}

	// the function inputs
	input_one := Asin(Xrange(0, 1, 1))
	input_two := Asin(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestAcos(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{1.5707963267948966}
	output_two := []float64{1.5533421480573113}

	// the function inputs
	input_one := Acos(Xrange(0, 1, 1))
	input_two := Acos(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestAtan(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{0.017451520651465824}

	// the function inputs
	input_one := Atan(Xrange(0, 1, 1))
	input_two := Atan(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestSinh(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{0.01745417862959511}

	// the function inputs
	input_one := Sinh(Xrange(0, 1, 1))
	input_two := Sinh(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestCosh(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{1}
	output_two := []float64{1.0001523125762564}

	// the function inputs
	input_one := Cosh(Xrange(0, 1, 1))
	input_two := Cosh(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestTanh(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{0.017451520543541533}

	// the function inputs
	input_one := Tanh(Xrange(0, 1, 1))
	input_two := Tanh(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestAsinh(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{0.01745240654522972}

	// the function inputs
	input_one := Asinh(Xrange(0, 1, 1))
	input_two := Asinh(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestAcosh(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{math.NaN()}
	output_two := []float64{math.NaN()}

	// the function inputs
	input_one := Acosh(Xrange(0, 1, 1))
	input_two := Acosh(Xrange(1, 2, 1))

	// check with the values of the first output
	if input_one.Data[0] == output_one[0] {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if input_two.Data[0] == output_two[0] {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestAtanh(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{0}
	output_two := []float64{0.01745506503622958}

	// the function inputs
	input_one := Atanh(Xrange(0, 1, 1))
	input_two := Atanh(Xrange(1, 2, 1))

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	// the expected outputs
	output_one := []float64{5, 6, 7, 8}
	output_two := []float64{123, 198}

	// the function inputs
	input_one := Sqrt(25, 36, 49, 64)
	input_two := Sqrt(15129, 39204)

	// check with the values of the first output
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Expected output was %v, but got %v\n", output_one, input_one.Data)
	}

	// check with the values of the second output
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Expected output was %v, but got %v\n", output_two, input_two.Data)
	}
}
