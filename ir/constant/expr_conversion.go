// generated by gen.go using 'go generate'; DO NOT EDIT.

// === [ Conversion expressions ] ==============================================
//
// References:
//    http://llvm.org/docs/LangRef.html#conversion-operations

package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ trunc ] ---------------------------------------------------------------

// ExprTrunc represents a truncation expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#trunc-instruction
type ExprTrunc struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewTrunc returns a new trunc expression based on the given source constant and target type.
func NewTrunc(from Constant, to types.Type) *ExprTrunc {
	return &ExprTrunc{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprTrunc) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprTrunc) Ident() string {
	from := expr.From()
	return fmt.Sprintf("trunc (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprTrunc) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprTrunc) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the trunc expression.
func (expr *ExprTrunc) From() Constant {
	return expr.from
}

// --- [ zext ] ----------------------------------------------------------------

// ExprZExt represents a zero extension expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#zext-instruction
type ExprZExt struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewZExt returns a new zext expression based on the given source constant and target type.
func NewZExt(from Constant, to types.Type) *ExprZExt {
	return &ExprZExt{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprZExt) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprZExt) Ident() string {
	from := expr.From()
	return fmt.Sprintf("zext (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprZExt) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprZExt) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the zext expression.
func (expr *ExprZExt) From() Constant {
	return expr.from
}

// --- [ sext ] ----------------------------------------------------------------

// ExprSExt represents a sign extension expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#sext-instruction
type ExprSExt struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewSExt returns a new sext expression based on the given source constant and target type.
func NewSExt(from Constant, to types.Type) *ExprSExt {
	return &ExprSExt{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprSExt) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprSExt) Ident() string {
	from := expr.From()
	return fmt.Sprintf("sext (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSExt) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprSExt) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the sext expression.
func (expr *ExprSExt) From() Constant {
	return expr.from
}

// --- [ fptrunc ] -------------------------------------------------------------

// ExprFPTrunc represents a floating-point truncation expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#fptrunc-instruction
type ExprFPTrunc struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewFPTrunc returns a new fptrunc expression based on the given source constant and target type.
func NewFPTrunc(from Constant, to types.Type) *ExprFPTrunc {
	return &ExprFPTrunc{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprFPTrunc) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprFPTrunc) Ident() string {
	from := expr.From()
	return fmt.Sprintf("fptrunc (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPTrunc) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprFPTrunc) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the fptrunc expression.
func (expr *ExprFPTrunc) From() Constant {
	return expr.from
}

// --- [ fpext ] ---------------------------------------------------------------

// ExprFPExt represents a floating-point extension expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#fpext-instruction
type ExprFPExt struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewFPExt returns a new fpext expression based on the given source constant and target type.
func NewFPExt(from Constant, to types.Type) *ExprFPExt {
	return &ExprFPExt{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprFPExt) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprFPExt) Ident() string {
	from := expr.From()
	return fmt.Sprintf("fpext (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPExt) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprFPExt) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the fpext expression.
func (expr *ExprFPExt) From() Constant {
	return expr.from
}

// --- [ fptoui ] --------------------------------------------------------------

// ExprFPToUI represents a floating-point to unsigned integer conversion expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#fptoui-instruction
type ExprFPToUI struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewFPToUI returns a new fptoui expression based on the given source constant and target type.
func NewFPToUI(from Constant, to types.Type) *ExprFPToUI {
	return &ExprFPToUI{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprFPToUI) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprFPToUI) Ident() string {
	from := expr.From()
	return fmt.Sprintf("fptoui (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToUI) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprFPToUI) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the fptoui expression.
func (expr *ExprFPToUI) From() Constant {
	return expr.from
}

// --- [ fptosi ] --------------------------------------------------------------

// ExprFPToSI represents a floating-point to signed integer conversion expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#fptosi-instruction
type ExprFPToSI struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewFPToSI returns a new fptosi expression based on the given source constant and target type.
func NewFPToSI(from Constant, to types.Type) *ExprFPToSI {
	return &ExprFPToSI{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprFPToSI) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprFPToSI) Ident() string {
	from := expr.From()
	return fmt.Sprintf("fptosi (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToSI) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprFPToSI) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the fptosi expression.
func (expr *ExprFPToSI) From() Constant {
	return expr.from
}

// --- [ uitofp ] --------------------------------------------------------------

// ExprUIToFP represents an unsigned integer to floating-point conversion expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#uitofp-instruction
type ExprUIToFP struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewUIToFP returns a new uitofp expression based on the given source constant and target type.
func NewUIToFP(from Constant, to types.Type) *ExprUIToFP {
	return &ExprUIToFP{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprUIToFP) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprUIToFP) Ident() string {
	from := expr.From()
	return fmt.Sprintf("uitofp (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUIToFP) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprUIToFP) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the uitofp expression.
func (expr *ExprUIToFP) From() Constant {
	return expr.from
}

// --- [ sitofp ] --------------------------------------------------------------

// ExprSIToFP represents a signed integer to floating-point conversion expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#sitofp-instruction
type ExprSIToFP struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewSIToFP returns a new sitofp expression based on the given source constant and target type.
func NewSIToFP(from Constant, to types.Type) *ExprSIToFP {
	return &ExprSIToFP{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprSIToFP) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprSIToFP) Ident() string {
	from := expr.From()
	return fmt.Sprintf("sitofp (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSIToFP) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprSIToFP) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the sitofp expression.
func (expr *ExprSIToFP) From() Constant {
	return expr.from
}

// --- [ ptrtoint ] ------------------------------------------------------------

// ExprPtrToInt represents a pointer to integer conversion expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#ptrtoint-instruction
type ExprPtrToInt struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewPtrToInt returns a new ptrtoint expression based on the given source constant and target type.
func NewPtrToInt(from Constant, to types.Type) *ExprPtrToInt {
	return &ExprPtrToInt{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprPtrToInt) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprPtrToInt) Ident() string {
	from := expr.From()
	return fmt.Sprintf("ptrtoint (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprPtrToInt) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprPtrToInt) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the ptrtoint expression.
func (expr *ExprPtrToInt) From() Constant {
	return expr.from
}

// --- [ inttoptr ] ------------------------------------------------------------

// ExprIntToPtr represents an integer to pointer conversion expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#inttoptr-instruction
type ExprIntToPtr struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewIntToPtr returns a new inttoptr expression based on the given source constant and target type.
func NewIntToPtr(from Constant, to types.Type) *ExprIntToPtr {
	return &ExprIntToPtr{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprIntToPtr) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprIntToPtr) Ident() string {
	from := expr.From()
	return fmt.Sprintf("inttoptr (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprIntToPtr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprIntToPtr) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the inttoptr expression.
func (expr *ExprIntToPtr) From() Constant {
	return expr.from
}

// --- [ bitcast ] -------------------------------------------------------------

// ExprBitCast represents a bitcast expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#bitcast-instruction
type ExprBitCast struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewBitCast returns a new bitcast expression based on the given source constant and target type.
func NewBitCast(from Constant, to types.Type) *ExprBitCast {
	return &ExprBitCast{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprBitCast) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprBitCast) Ident() string {
	from := expr.From()
	return fmt.Sprintf("bitcast (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprBitCast) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprBitCast) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the bitcast expression.
func (expr *ExprBitCast) From() Constant {
	return expr.from
}

// --- [ addrspacecast ] -------------------------------------------------------

// ExprAddrSpaceCast represents an address space cast expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#addrspacecast-instruction
type ExprAddrSpaceCast struct {
	// Constant before conversion.
	from Constant
	// Type after conversion.
	to types.Type
}

// NewAddrSpaceCast returns a new addrspacecast expression based on the given source constant and target type.
func NewAddrSpaceCast(from Constant, to types.Type) *ExprAddrSpaceCast {
	return &ExprAddrSpaceCast{from: from, to: to}
}

// Type returns the type of the constant expression.
func (expr *ExprAddrSpaceCast) Type() types.Type {
	return expr.to
}

// Ident returns the string representation of the constant expression.
func (expr *ExprAddrSpaceCast) Ident() string {
	from := expr.From()
	return fmt.Sprintf("addrspacecast (%s %s to %s)",
		from.Type(),
		from.Ident(),
		expr.Type())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAddrSpaceCast) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprAddrSpaceCast) Simplify() Constant {
	panic("not yet implemented")
}

// From returns the constant before conversion of the addrspacecast expression.
func (expr *ExprAddrSpaceCast) From() Constant {
	return expr.from
}
