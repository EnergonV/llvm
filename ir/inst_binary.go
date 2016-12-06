// generated by gen.go using 'go generate'; DO NOT EDIT.

// === [ Binary instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#binary-operations

package ir

import (
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ add ] -----------------------------------------------------------------

// InstAdd represents an addition instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#add-instruction
type InstAdd struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y value.Value) *InstAdd {
	return &InstAdd{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstAdd) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAdd) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstAdd) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstAdd) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstAdd) String() string {
	return fmt.Sprintf("%s = add %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstAdd) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstAdd) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the add instruction.
func (inst *InstAdd) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the add instruction.
func (inst *InstAdd) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the add instruction.
func (inst *InstAdd) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the add instruction.
func (inst *InstAdd) SetY(y value.Value) {
	inst.y = y
}

// --- [ fadd ] ----------------------------------------------------------------

// InstFAdd represents a floating-point addition instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fadd-instruction
type InstFAdd struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewFAdd returns a new fadd instruction based on the given operands.
func NewFAdd(x, y value.Value) *InstFAdd {
	return &InstFAdd{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstFAdd) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFAdd) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstFAdd) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstFAdd) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFAdd) String() string {
	return fmt.Sprintf("%s = fadd %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstFAdd) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFAdd) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the fadd instruction.
func (inst *InstFAdd) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the fadd instruction.
func (inst *InstFAdd) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the fadd instruction.
func (inst *InstFAdd) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the fadd instruction.
func (inst *InstFAdd) SetY(y value.Value) {
	inst.y = y
}

// --- [ sub ] -----------------------------------------------------------------

// InstSub represents a subtraction instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#sub-instruction
type InstSub struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewSub returns a new sub instruction based on the given operands.
func NewSub(x, y value.Value) *InstSub {
	return &InstSub{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstSub) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSub) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstSub) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstSub) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstSub) String() string {
	return fmt.Sprintf("%s = sub %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstSub) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstSub) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the sub instruction.
func (inst *InstSub) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the sub instruction.
func (inst *InstSub) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the sub instruction.
func (inst *InstSub) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the sub instruction.
func (inst *InstSub) SetY(y value.Value) {
	inst.y = y
}

// --- [ fsub ] ----------------------------------------------------------------

// InstFSub represents a floating-point subtraction instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fsub-instruction
type InstFSub struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewFSub returns a new fsub instruction based on the given operands.
func NewFSub(x, y value.Value) *InstFSub {
	return &InstFSub{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstFSub) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFSub) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstFSub) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstFSub) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFSub) String() string {
	return fmt.Sprintf("%s = fsub %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstFSub) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFSub) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the fsub instruction.
func (inst *InstFSub) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the fsub instruction.
func (inst *InstFSub) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the fsub instruction.
func (inst *InstFSub) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the fsub instruction.
func (inst *InstFSub) SetY(y value.Value) {
	inst.y = y
}

// --- [ mul ] -----------------------------------------------------------------

// InstMul represents a multiplication instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#mul-instruction
type InstMul struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y value.Value) *InstMul {
	return &InstMul{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstMul) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstMul) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstMul) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstMul) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstMul) String() string {
	return fmt.Sprintf("%s = mul %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstMul) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstMul) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the mul instruction.
func (inst *InstMul) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the mul instruction.
func (inst *InstMul) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the mul instruction.
func (inst *InstMul) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the mul instruction.
func (inst *InstMul) SetY(y value.Value) {
	inst.y = y
}

// --- [ fmul ] ----------------------------------------------------------------

// InstFMul represents a floating-point multiplication instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fmul-instruction
type InstFMul struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewFMul returns a new fmul instruction based on the given operands.
func NewFMul(x, y value.Value) *InstFMul {
	return &InstFMul{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstFMul) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFMul) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstFMul) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstFMul) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFMul) String() string {
	return fmt.Sprintf("%s = fmul %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstFMul) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFMul) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the fmul instruction.
func (inst *InstFMul) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the fmul instruction.
func (inst *InstFMul) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the fmul instruction.
func (inst *InstFMul) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the fmul instruction.
func (inst *InstFMul) SetY(y value.Value) {
	inst.y = y
}

// --- [ udiv ] ----------------------------------------------------------------

// InstUDiv represents an unsigned division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#udiv-instruction
type InstUDiv struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewUDiv returns a new udiv instruction based on the given operands.
func NewUDiv(x, y value.Value) *InstUDiv {
	return &InstUDiv{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstUDiv) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstUDiv) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstUDiv) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstUDiv) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstUDiv) String() string {
	return fmt.Sprintf("%s = udiv %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstUDiv) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstUDiv) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the udiv instruction.
func (inst *InstUDiv) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the udiv instruction.
func (inst *InstUDiv) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the udiv instruction.
func (inst *InstUDiv) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the udiv instruction.
func (inst *InstUDiv) SetY(y value.Value) {
	inst.y = y
}

// --- [ sdiv ] ----------------------------------------------------------------

// InstSDiv represents a signed division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#sdiv-instruction
type InstSDiv struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewSDiv returns a new sdiv instruction based on the given operands.
func NewSDiv(x, y value.Value) *InstSDiv {
	return &InstSDiv{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstSDiv) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSDiv) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstSDiv) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstSDiv) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstSDiv) String() string {
	return fmt.Sprintf("%s = sdiv %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstSDiv) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstSDiv) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the sdiv instruction.
func (inst *InstSDiv) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the sdiv instruction.
func (inst *InstSDiv) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the sdiv instruction.
func (inst *InstSDiv) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the sdiv instruction.
func (inst *InstSDiv) SetY(y value.Value) {
	inst.y = y
}

// --- [ fdiv ] ----------------------------------------------------------------

// InstFDiv represents a floating-point division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fdiv-instruction
type InstFDiv struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewFDiv returns a new fdiv instruction based on the given operands.
func NewFDiv(x, y value.Value) *InstFDiv {
	return &InstFDiv{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstFDiv) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFDiv) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstFDiv) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstFDiv) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFDiv) String() string {
	return fmt.Sprintf("%s = fdiv %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstFDiv) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFDiv) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the fdiv instruction.
func (inst *InstFDiv) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the fdiv instruction.
func (inst *InstFDiv) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the fdiv instruction.
func (inst *InstFDiv) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the fdiv instruction.
func (inst *InstFDiv) SetY(y value.Value) {
	inst.y = y
}

// --- [ urem ] ----------------------------------------------------------------

// InstURem represents an unsigned remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#urem-instruction
type InstURem struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewURem returns a new urem instruction based on the given operands.
func NewURem(x, y value.Value) *InstURem {
	return &InstURem{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstURem) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstURem) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstURem) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstURem) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstURem) String() string {
	return fmt.Sprintf("%s = urem %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstURem) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstURem) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the urem instruction.
func (inst *InstURem) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the urem instruction.
func (inst *InstURem) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the urem instruction.
func (inst *InstURem) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the urem instruction.
func (inst *InstURem) SetY(y value.Value) {
	inst.y = y
}

// --- [ srem ] ----------------------------------------------------------------

// InstSRem represents a signed remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#srem-instruction
type InstSRem struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewSRem returns a new srem instruction based on the given operands.
func NewSRem(x, y value.Value) *InstSRem {
	return &InstSRem{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstSRem) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSRem) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstSRem) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstSRem) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstSRem) String() string {
	return fmt.Sprintf("%s = srem %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstSRem) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstSRem) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the srem instruction.
func (inst *InstSRem) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the srem instruction.
func (inst *InstSRem) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the srem instruction.
func (inst *InstSRem) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the srem instruction.
func (inst *InstSRem) SetY(y value.Value) {
	inst.y = y
}

// --- [ frem ] ----------------------------------------------------------------

// InstFRem represents a floating-point remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#frem-instruction
type InstFRem struct {
	// Parent basic block.
	parent *BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewFRem returns a new frem instruction based on the given operands.
func NewFRem(x, y value.Value) *InstFRem {
	return &InstFRem{x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstFRem) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFRem) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstFRem) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstFRem) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFRem) String() string {
	return fmt.Sprintf("%s = frem %s %s, %s",
		inst.Ident(),
		inst.Type(),
		inst.x.Ident(),
		inst.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstFRem) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFRem) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// X returns the x operand of the frem instruction.
func (inst *InstFRem) X() value.Value {
	return inst.x
}

// SetX sets the x operand of the frem instruction.
func (inst *InstFRem) SetX(x value.Value) {
	inst.x = x
}

// Y returns the y operand of the frem instruction.
func (inst *InstFRem) Y() value.Value {
	return inst.y
}

// SetY sets the y operand of the frem instruction.
func (inst *InstFRem) SetY(y value.Value) {
	inst.y = y
}
