// generated by gen.go using 'go generate'; DO NOT EDIT.

// === [ Bitwise binary instructions ] =========================================
//
// References:
//    http://llvm.org/docs/LangRef.html#bitwise-binary-operations

package irx

import (
	"github.com/llir/llvm/ir/instruction"
	"github.com/mewkiz/pkg/errutil"
)

// --- [ shl ] -----------------------------------------------------------------

// NewShL returns a new shl instruction based on the given operand type and
// operands.
func NewShLInst(typ, xVal, yVal interface{}) (*instruction.ShL, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewShL(x, y)
}

// --- [ lshr ] ----------------------------------------------------------------

// NewLShR returns a new lshr instruction based on the given operand type and
// operands.
func NewLShRInst(typ, xVal, yVal interface{}) (*instruction.LShR, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewLShR(x, y)
}

// --- [ ashr ] ----------------------------------------------------------------

// NewAShR returns a new ashr instruction based on the given operand type and
// operands.
func NewAShRInst(typ, xVal, yVal interface{}) (*instruction.AShR, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewAShR(x, y)
}

// --- [ and ] -----------------------------------------------------------------

// NewAnd returns a new and instruction based on the given operand type and
// operands.
func NewAndInst(typ, xVal, yVal interface{}) (*instruction.And, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewAnd(x, y)
}

// --- [ or ] ------------------------------------------------------------------

// NewOr returns a new or instruction based on the given operand type and
// operands.
func NewOrInst(typ, xVal, yVal interface{}) (*instruction.Or, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewOr(x, y)
}

// --- [ xor ] -----------------------------------------------------------------

// NewXor returns a new xor instruction based on the given operand type and
// operands.
func NewXorInst(typ, xVal, yVal interface{}) (*instruction.Xor, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewXor(x, y)
}