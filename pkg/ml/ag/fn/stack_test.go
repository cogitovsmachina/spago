// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"github.com/nlpodyssey/spago/pkg/mat"
	"gonum.org/v1/gonum/floats"
	"testing"
)

func TestStack_Forward(t *testing.T) {
	x1 := &variable{
		value:        mat.NewVecDense([]float64{0.1, 0.2, 0.3, 0.5}),
		grad:         nil,
		requiresGrad: true,
	}
	x2 := &variable{
		value:        mat.NewVecDense([]float64{0.4, 0.5, 0.6, 0.4}),
		grad:         nil,
		requiresGrad: true,
	}
	x3 := &variable{
		value:        mat.NewVecDense([]float64{0.8, 0.9, 0.7, 0.6}),
		grad:         nil,
		requiresGrad: true,
	}

	f := NewStack([]Operand{x1, x2, x3})
	y := f.Forward()

	if !floats.EqualApprox(y.Data(), []float64{0.1, 0.2, 0.3, 0.5, 0.4, 0.5, 0.6, 0.4, 0.8, 0.9, 0.7, 0.6}, 1.0e-6) {
		t.Error("The output doesn't match the expected values")
	}

	if y.Rows() != 3 && y.Columns() != 4 {
		t.Error("The output size doesn't match the expected values")
	}

	f.Backward(mat.NewDense(3, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		4.0, 5.0, 6.0, 0.5,
		7.0, 8.0, 9.0, -0.3,
	}))

	if !floats.EqualApprox(x1.grad.Data(), []float64{1.0, 2.0, 3.0, 4.0}, 1.0e-6) {
		t.Error("The x1-gradients don't match the expected values")
	}

	if !floats.EqualApprox(x2.grad.Data(), []float64{4.0, 5.0, 6.0, 0.5}, 1.0e-6) {
		t.Error("The x2-gradients don't match the expected values")
	}

	if !floats.EqualApprox(x3.grad.Data(), []float64{7.0, 8.0, 9.0, -0.3}, 1.0e-6) {
		t.Error("The x3-gradients don't match the expected values")
	}
}
