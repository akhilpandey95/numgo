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

func TestNDArray(t *testing.T) {
	// wait til all the checks are performed
	t.Parallel()

	// the expected outputs
	output_one := 24
	output_two := 36
	input_one := []float64{1, 2, 3, 4}
	input_two := []float64{3, 3, 4}

	// create an N-D Array
	n1 := NDArray(input_one...)
	n2 := NDArray(input_two...)

	// check the values for the first array
	if len(n1.Data) != output_one {
		t.Errorf("Length of NDArray(%v) must be %v, but got %v\n", input_one, output_one, len(n1.Data))
	}

	// check the values for the second array
	if len(n2.Data) != output_two {
		t.Errorf("Length of NDArray(%v) must be %v, but got %v\n", input_two, output_two, len(n2.Data))
	}
}

func TestRange(t *testing.T) {
	// wait till all the checks are performed
	t.Parallel()

	// the expected outputs
	output_one := []float64{1, 2, 3, 4, 5}
	output_two := []float64{10, 11, 12, 13, 14}

	// create the N-D arrays
	input_one := Range(1, 6)
	input_two := Range(10, 15)

	// check the values for the first array
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Contents of Array must be %v, but got %v\n", output_one, input_one.Data)
	}

	// check the values for the second array
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Contents of Array must be %v, but got %v\n", output_two, input_two.Data)
	}
}

func TestXrange(t *testing.T) {
	// wait till all the checks are performed
	t.Parallel()

	// the expected outputs
	output_one := []float64{18, 18.1, 18.2, 18.3, 18.4, 18.5, 18.6, 18.7, 18.8, 18.9, 19, 19.1, 19.2, 19.3, 19.4, 19.5, 19.6, 19.7, 19.8, 19.9}
	output_two := []float64{50, 100, 150, 200, 250, 300}
	output_three := []float64{20, 19, 18, 17, 16, 15, 14, 13, 12, 11}

	// create the N-D array
	input_one := Xrange(18, 20, 0.1)
	input_two := Xrange(50, 350, 50)
	input_three := Xrange(20, 10)

	// check the values for the first array
	if !(reflect.DeepEqual(input_one.Data, output_one)) {
		t.Errorf("Contents of Array must be %v, but got %v\n", output_one, input_one.Data)
	}

	// check the values for the second array
	if !(reflect.DeepEqual(input_two.Data, output_two)) {
		t.Errorf("Contents of Array must be %v, but got %v\n", output_two, input_two.Data)
	}

	// check the values for teh third array
	if !(reflect.DeepEqual(input_three.Data, output_three)) {
		t.Errorf("Contents of Array must be %v, but got %v\n", output_three, input_three.Data)
	}
}
