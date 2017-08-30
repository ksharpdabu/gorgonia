package tensor

import "github.com/pkg/errors"

/*
GENERATED FILE. DO NOT EDIT
*/

func (t *Dense) Add(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if adder, ok := e.(Adder); ok {
		var ret Tensor
		if ret, err = adder.Add(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Add()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Add")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Add()")
}

func (t *Dense) Sub(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if suber, ok := e.(Suber); ok {
		var ret Tensor
		if ret, err = suber.Sub(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Sub()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Sub")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Sub()")
}

func (t *Dense) Mul(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if muler, ok := e.(Muler); ok {
		var ret Tensor
		if ret, err = muler.Mul(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Mul()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Mul")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Mul()")
}

func (t *Dense) Div(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if diver, ok := e.(Diver); ok {
		var ret Tensor
		if ret, err = diver.Div(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Div()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Div")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Div()")
}

func (t *Dense) Pow(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if power, ok := e.(Power); ok {
		var ret Tensor
		if ret, err = power.Pow(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Pow()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Pow")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Pow()")
}

func (t *Dense) Mod(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if moder, ok := e.(Moder); ok {
		var ret Tensor
		if ret, err = moder.Mod(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Mod()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Mod")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Mod()")
}

func (t *Dense) AddScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if adder, ok := e.(Adder); ok {
		var ret Tensor
		if ret, err = adder.AddScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do AddScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "AddScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support AddScalar()")
}

func (t *Dense) SubScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if suber, ok := e.(Suber); ok {
		var ret Tensor
		if ret, err = suber.SubScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do SubScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "SubScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support SubScalar()")
}

func (t *Dense) MulScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if muler, ok := e.(Muler); ok {
		var ret Tensor
		if ret, err = muler.MulScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do MulScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "MulScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support MulScalar()")
}

func (t *Dense) DivScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if diver, ok := e.(Diver); ok {
		var ret Tensor
		if ret, err = diver.DivScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do DivScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "DivScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support DivScalar()")
}

func (t *Dense) PowScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if power, ok := e.(Power); ok {
		var ret Tensor
		if ret, err = power.PowScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do PowScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "PowScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support PowScalar()")
}

func (t *Dense) ModScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if moder, ok := e.(Moder); ok {
		var ret Tensor
		if ret, err = moder.ModScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do ModScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "ModScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support ModScalar()")
}
