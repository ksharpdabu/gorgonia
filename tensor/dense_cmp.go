package tensor

import "github.com/pkg/errors"

/*
GENERATED FILE. DO NOT EDIT
*/

func (t *Dense) Gt(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if gter, ok := e.(Gter); ok {
		var ret Tensor
		if ret, err = gter.Gt(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Gt()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Gt")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Gt()")
}

func (t *Dense) Gte(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if gteer, ok := e.(Gteer); ok {
		var ret Tensor
		if ret, err = gteer.Gte(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Gte()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Gte")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Gte()")
}

func (t *Dense) Lt(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if lter, ok := e.(Lter); ok {
		var ret Tensor
		if ret, err = lter.Lt(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Lt()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Lt")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Lt()")
}

func (t *Dense) Lte(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if lteer, ok := e.(Lteer); ok {
		var ret Tensor
		if ret, err = lteer.Lte(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Lte()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Lte")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Lte()")
}

func (t *Dense) ElEq(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if eleqer, ok := e.(ElEqer); ok {
		var ret Tensor
		if ret, err = eleqer.ElEq(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Eq()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Eq")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Eq()")
}

func (t *Dense) ElNe(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e

	if eleqer, ok := e.(ElEqer); ok {
		var ret Tensor
		if ret, err = eleqer.ElNe(t, other, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do Ne()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Ne")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Ne()")
}

func (t *Dense) GtScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if gter, ok := e.(Gter); ok {
		var ret Tensor
		if ret, err = gter.GtScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do GtScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "GtScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support GtScalar()")
}

func (t *Dense) GteScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if gteer, ok := e.(Gteer); ok {
		var ret Tensor
		if ret, err = gteer.GteScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do GteScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "GteScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support GteScalar()")
}

func (t *Dense) LtScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if lter, ok := e.(Lter); ok {
		var ret Tensor
		if ret, err = lter.LtScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do LtScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "LtScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support LtScalar()")
}

func (t *Dense) LteScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if lteer, ok := e.(Lteer); ok {
		var ret Tensor
		if ret, err = lteer.LteScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do LteScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "LteScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support LteScalar()")
}

func (t *Dense) ElEqScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if eleqer, ok := e.(ElEqer); ok {
		var ret Tensor
		if ret, err = eleqer.EqScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do EqScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "EqScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support EqScalar()")
}

func (t *Dense) ElNeScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	e := t.e
	if eleqer, ok := e.(ElEqer); ok {
		var ret Tensor
		if ret, err = eleqer.NeScalar(t, other, leftTensor, opts...); err != nil {
			err = errors.Wrapf(err, "Unable to do NeScalar()")
			return
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "NeScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support NeScalar()")
}
