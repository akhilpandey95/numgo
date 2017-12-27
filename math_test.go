/* This Source Code Form is subject to the terms of the MPL
*  License. If a copy of the same was not distributed with this
*  file, You can obtain one at
*  https://github.com/AkhilHector/pubundsci/blob/master/LICENSE.
 */

package numgo

import (
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
