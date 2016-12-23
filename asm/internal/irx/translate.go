// Translates AST values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Index global variables.
//       - Store preliminary content type.
//    3. Index function.
//       - Store type.
//    4. Fix type definitions.
//    5. Fix globals.
//    6. Fix functions.
//
// Per function.
//
//    1. Index function parameters.
//    2. Index basic blocks.
//    3. Index local variables produced by instructions.
//       - Store preliminary type.
//    4. Fix basic blocks.

package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Modules ] =============================================================

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	m := NewModule()

	// Index type definitions.
	for _, old := range module.Types {
		name := old.Name
		if _, ok := m.types[name]; ok {
			panic(fmt.Errorf("type name %q already present; old `%v`, new `%v`", name, m.types[name], old))
		}
		typ := &types.NamedType{
			Name: name,
		}
		m.Types = append(m.Types, typ)
		m.types[name] = typ
	}

	// Index global variables.
	for _, old := range module.Globals {
		name := old.Name
		if _, ok := m.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, m.globals[name], old))
		}
		global := &ir.Global{
			Name: name,
		}
		// TODO: Verify if it is needed to store preliminary content type of
		// globals; i.e. validate type resolution for circularly defined globals.
		//
		//// Store preliminary content type.
		// content := m.irType(old.Content)
		//global.Content = content
		//global.Typ = types.NewPointer(content)
		m.Globals = append(m.Globals, global)
		m.globals[name] = global
	}

	// Index functions.
	for _, old := range module.Funcs {
		name := old.Name
		if _, ok := m.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, m.globals[name], old))
		}
		// Store type.
		oldSig := m.irType(old.Sig)
		sig, ok := oldSig.(*types.FuncType)
		if !ok {
			panic(fmt.Errorf("invalid function signature type, expected *types.FuncType, got %T", oldSig))
		}
		var params []*ir.Param
		for _, param := range sig.Params {
			params = append(params, ir.NewParam(param.Name, param.Typ))
		}
		typ := types.NewPointer(sig)
		f := &ir.Function{
			Parent: m.Module,
			Name:   name,
			Typ:    typ,
			Sig:    sig,
			Params: params,
		}
		m.Funcs = append(m.Funcs, f)
		m.globals[name] = f
	}

	// Fix type definitions.
	for _, typ := range module.Types {
		m.typeDef(typ)
	}

	// Fix globals.
	for _, global := range module.Globals {
		m.globalDecl(global)
	}

	// Fix functions.
	for _, f := range module.Funcs {
		m.funcDecl(f)
	}

	if len(m.errs) > 0 {
		// TODO: Return a list of all errors.
		return nil, m.errs[0]
	}
	return m.Module, nil
}

// === [ Type definitions ] ====================================================

// typeDef translates the given type definition to LLVM IR, emitting code to m.
func (m *Module) typeDef(old *ast.NamedType) {
	typ := m.getType(old.Name)
	def := m.irType(old.Def)
	typ.Def = def
}

// === [ Global variables ] ====================================================

// globalDecl translates the given global variable declaration to LLVM IR,
// emitting code to m.
func (m *Module) globalDecl(old *ast.Global) {
	v := m.getGlobal(old.Name)
	global, ok := v.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid global type; expected *ir.Global, got %T", v))
	}
	if old.Init != nil {
		init := m.irConstant(old.Init)
		// TODO: Verify that two circularly referential globals both get the
		// correct type; more specifically that neither get global.Content == nil
		// after resolution.
		global.Content = init.Type()
		global.Init = init
	} else {
		global.Content = m.irType(old.Content)
	}
	global.Typ = types.NewPointer(global.Content)
	global.IsConst = old.Immutable
}

// === [ Functions ] ===========================================================

// funcDecl translates the given function declaration to LLVM IR, emitting code
// to m.
func (m *Module) funcDecl(oldFunc *ast.Function) {
	// Reset locals.
	m.locals = make(map[string]value.Named)

	v := m.getGlobal(oldFunc.Name)
	f, ok := v.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid function type; expected *ir.Function, got %T", v))
	}

	// Index function parameters.
	for _, param := range f.Params {
		name := param.Name
		if _, ok := m.locals[name]; ok {
			panic(fmt.Errorf("local identifier %q already present; old `%v`, new `%v`", name, m.locals[name], param))
		}
		m.locals[name] = param
	}

	// Index basic blocks.
	for _, old := range oldFunc.Blocks {
		name := old.Name
		if _, ok := m.locals[name]; ok {
			panic(fmt.Errorf("local identifier %q already present; old `%v`, new `%v`", name, m.locals[name], old))
		}
		block := &ir.BasicBlock{
			Name: name,
		}
		f.Blocks = append(f.Blocks, block)
		m.locals[name] = block
	}

	// Index local variables produced by instructions.
	for i := 0; i < len(oldFunc.Blocks); i++ {
		oldBlock := oldFunc.Blocks[i]
		block := f.Blocks[i]
		for _, oldInst := range oldBlock.Insts {
			var inst ir.Instruction
			switch oldInst := oldInst.(type) {
			// Binary instructions
			case *ast.InstAdd:
				inst = &ir.InstAdd{Name: oldInst.Name}
			case *ast.InstFAdd:
				inst = &ir.InstFAdd{Name: oldInst.Name}
			case *ast.InstSub:
				inst = &ir.InstSub{Name: oldInst.Name}
			case *ast.InstFSub:
				inst = &ir.InstFSub{Name: oldInst.Name}
			case *ast.InstMul:
				inst = &ir.InstMul{Name: oldInst.Name}
			case *ast.InstFMul:
				inst = &ir.InstFMul{Name: oldInst.Name}
			case *ast.InstUDiv:
				inst = &ir.InstUDiv{Name: oldInst.Name}
			case *ast.InstSDiv:
				inst = &ir.InstSDiv{Name: oldInst.Name}
			case *ast.InstFDiv:
				inst = &ir.InstFDiv{Name: oldInst.Name}
			case *ast.InstURem:
				inst = &ir.InstURem{Name: oldInst.Name}
			case *ast.InstSRem:
				inst = &ir.InstSRem{Name: oldInst.Name}
			case *ast.InstFRem:
				inst = &ir.InstFRem{Name: oldInst.Name}

			// Bitwise instructions
			case *ast.InstShl:
				inst = &ir.InstShl{Name: oldInst.Name}
			case *ast.InstLShr:
				inst = &ir.InstLShr{Name: oldInst.Name}
			case *ast.InstAShr:
				inst = &ir.InstAShr{Name: oldInst.Name}
			case *ast.InstAnd:
				inst = &ir.InstAnd{Name: oldInst.Name}
			case *ast.InstOr:
				inst = &ir.InstOr{Name: oldInst.Name}
			case *ast.InstXor:
				inst = &ir.InstXor{Name: oldInst.Name}

			// Memory instructions
			case *ast.InstAlloca:
				inst = &ir.InstAlloca{Name: oldInst.Name}
			case *ast.InstLoad:
				inst = &ir.InstLoad{Name: oldInst.Name}
			case *ast.InstStore:
				// Store instructions produce no value, and are thus not assigned
				// names.
				inst = &ir.InstStore{}
			case *ast.InstGetElementPtr:
				inst = &ir.InstGetElementPtr{Name: oldInst.Name}

			// Conversion instructions
			case *ast.InstTrunc:
				inst = &ir.InstTrunc{Name: oldInst.Name}
			case *ast.InstZExt:
				inst = &ir.InstZExt{Name: oldInst.Name}
			case *ast.InstSExt:
				inst = &ir.InstSExt{Name: oldInst.Name}
			case *ast.InstFPTrunc:
				inst = &ir.InstFPTrunc{Name: oldInst.Name}
			case *ast.InstFPExt:
				inst = &ir.InstFPExt{Name: oldInst.Name}
			case *ast.InstFPToUI:
				inst = &ir.InstFPToUI{Name: oldInst.Name}
			case *ast.InstFPToSI:
				inst = &ir.InstFPToSI{Name: oldInst.Name}
			case *ast.InstUIToFP:
				inst = &ir.InstUIToFP{Name: oldInst.Name}
			case *ast.InstSIToFP:
				inst = &ir.InstSIToFP{Name: oldInst.Name}
			case *ast.InstPtrToInt:
				inst = &ir.InstPtrToInt{Name: oldInst.Name}
			case *ast.InstIntToPtr:
				inst = &ir.InstIntToPtr{Name: oldInst.Name}
			case *ast.InstBitCast:
				inst = &ir.InstBitCast{Name: oldInst.Name}
			case *ast.InstAddrSpaceCast:
				inst = &ir.InstAddrSpaceCast{Name: oldInst.Name}

			// Other instructions
			case *ast.InstICmp:
				inst = &ir.InstICmp{Name: oldInst.Name}
			case *ast.InstFCmp:
				inst = &ir.InstFCmp{Name: oldInst.Name}
			case *ast.InstPhi:
				inst = &ir.InstPhi{Name: oldInst.Name}
			case *ast.InstSelect:
				inst = &ir.InstSelect{Name: oldInst.Name}
			case *ast.InstCall:
				inst = &ir.InstCall{Name: oldInst.Name}

			default:
				panic(fmt.Errorf("support for instruction %T not yet implemented", oldInst))
			}
			block.Insts = append(block.Insts, inst)

			// TODO: Validate if it is required to store a preliminary type of
			// instructions prior to local variable resolution.
			//
			// What happens if a getelementptr instruction refers to the value
			// produced by an instruction which cannot be calculated prior to its
			// operands being resolved?
			//
			//// Store preliminary type.

			// Index local variable.
			if inst, ok := inst.(value.Named); ok {
				// Ignore local value if of type void.
				if oldInst, ok := oldInst.(*ast.InstCall); ok {
					if _, ok := oldInst.Type.(*ast.VoidType); ok {
						continue
					}
				}
				m.locals[inst.GetName()] = inst
			}
		}
	}

	// Fix basic blocks.
	for i := 0; i < len(oldFunc.Blocks); i++ {
		oldBlock := oldFunc.Blocks[i]
		block := f.Blocks[i]
		m.basicBlock(oldBlock, block)
	}
}

// === [ Identifiers ] =========================================================

// === [ Types ] ===============================================================

// === [ Values ] ==============================================================

// === [ Constants ] ===========================================================

// --- [ Binary expressions ] --------------------------------------------------

// --- [ Bitwise expressions ] -------------------------------------------------

// --- [ Memory expressions ] --------------------------------------------------

// --- [ Conversion expressions ] ----------------------------------------------

// --- [ Other expressions ] ---------------------------------------------------

// === [ Basic blocks ] ========================================================

// basicBlock translates the given basic block to LLVM IR, emitting code to m.
func (m *Module) basicBlock(oldBlock *ast.BasicBlock, block *ir.BasicBlock) {
	// Fix instructions.
	for i := 0; i < len(oldBlock.Insts); i++ {
		oldInst := oldBlock.Insts[i]
		v := block.Insts[i]
		switch oldInst := oldInst.(type) {
		// Binary instructions
		case *ast.InstAdd:
			inst, ok := v.(*ir.InstAdd)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAdd, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFAdd:
			inst, ok := v.(*ir.InstFAdd)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFAdd, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstSub:
			inst, ok := v.(*ir.InstSub)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSub, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFSub:
			inst, ok := v.(*ir.InstFSub)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFSub, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstMul:
			inst, ok := v.(*ir.InstMul)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstMul, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFMul:
			inst, ok := v.(*ir.InstFMul)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFMul, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstUDiv:
			inst, ok := v.(*ir.InstUDiv)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstUDiv, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstSDiv:
			inst, ok := v.(*ir.InstSDiv)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSDiv, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFDiv:
			inst, ok := v.(*ir.InstFDiv)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFDiv, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstURem:
			inst, ok := v.(*ir.InstURem)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstURem, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstSRem:
			inst, ok := v.(*ir.InstSRem)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSRem, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFRem:
			inst, ok := v.(*ir.InstFRem)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFRem, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)

		// Bitwise instructions
		case *ast.InstShl:
			inst, ok := v.(*ir.InstShl)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstShl, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstLShr:
			inst, ok := v.(*ir.InstLShr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstLShr, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstAShr:
			inst, ok := v.(*ir.InstAShr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAShr, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstAnd:
			inst, ok := v.(*ir.InstAnd)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAnd, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstOr:
			inst, ok := v.(*ir.InstOr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstOr, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstXor:
			inst, ok := v.(*ir.InstXor)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstXor, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)

		// Memory instructions
		case *ast.InstAlloca:
			inst, ok := v.(*ir.InstAlloca)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAlloca, got %T", v))
			}
			elem := m.irType(oldInst.Elem)
			typ := types.NewPointer(elem)
			inst.Typ = typ
			inst.Elem = elem
			if oldInst.NElems != nil {
				inst.NElems = m.irValue(oldInst.NElems)
			}
		case *ast.InstLoad:
			inst, ok := v.(*ir.InstLoad)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstLoad, got %T", v))
			}
			src := m.irValue(oldInst.Src)
			srcType, ok := src.Type().(*types.PointerType)
			if !ok {
				panic(fmt.Errorf("invalid source type; expected *types.PointerType, got %T", src.Type()))
			}
			typ := srcType.Elem
			if got, want := typ, m.irType(oldInst.Elem); !got.Equal(want) {
				m.errs = append(m.errs, errors.Errorf("source element type mismatch; expected `%v`, got `%v`", want, got))
			}
			inst.Typ = typ
			inst.Src = src
		case *ast.InstStore:
			inst, ok := v.(*ir.InstStore)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstStore, got %T", v))
			}
			inst.Src = m.irValue(oldInst.Src)
			inst.Dst = m.irValue(oldInst.Dst)
		case *ast.InstGetElementPtr:
			inst, ok := v.(*ir.InstGetElementPtr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstGetElementPtr, got %T", v))
			}
			_ = inst
			panic("not yet implemented")

		// Conversion instructions
		case *ast.InstTrunc:
			inst, ok := v.(*ir.InstTrunc)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstTrunc, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstZExt:
			inst, ok := v.(*ir.InstZExt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstZExt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstSExt:
			inst, ok := v.(*ir.InstSExt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSExt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPTrunc:
			inst, ok := v.(*ir.InstFPTrunc)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPTrunc, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPExt:
			inst, ok := v.(*ir.InstFPExt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPExt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPToUI:
			inst, ok := v.(*ir.InstFPToUI)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPToUI, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPToSI:
			inst, ok := v.(*ir.InstFPToSI)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPToSI, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstUIToFP:
			inst, ok := v.(*ir.InstUIToFP)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstUIToFP, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstSIToFP:
			inst, ok := v.(*ir.InstSIToFP)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSIToFP, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstPtrToInt:
			inst, ok := v.(*ir.InstPtrToInt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstPtrToInt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstIntToPtr:
			inst, ok := v.(*ir.InstIntToPtr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstIntToPtr, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstBitCast:
			inst, ok := v.(*ir.InstBitCast)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstBitCast, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstAddrSpaceCast:
			inst, ok := v.(*ir.InstAddrSpaceCast)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAddrSpaceCast, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)

		// Other instructions
		case *ast.InstICmp:
			inst, ok := v.(*ir.InstICmp)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstICmp, got %T", v))
			}
			_ = inst
			panic("not yet implemented")
		case *ast.InstFCmp:
			inst, ok := v.(*ir.InstFCmp)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFCmp, got %T", v))
			}
			_ = inst
			panic("not yet implemented")
		case *ast.InstPhi:
			inst, ok := v.(*ir.InstPhi)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstPhi, got %T", v))
			}
			_ = inst
			panic("not yet implemented")
		case *ast.InstSelect:
			inst, ok := v.(*ir.InstSelect)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSelect, got %T", v))
			}
			_ = inst
			panic("not yet implemented")
		case *ast.InstCall:
			inst, ok := v.(*ir.InstCall)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstCall, got %T", v))
			}
			v := m.irValue(oldInst.Callee)
			callee, ok := v.(value.Named)
			if !ok {
				panic(fmt.Errorf("invalid callee type; expected value.Named, got %T", v))
			}
			typ, ok := callee.Type().(*types.PointerType)
			if !ok {
				panic(fmt.Errorf("invalid callee type, expected *types.PointerType, got %T", callee.Type()))
			}
			sig, ok := typ.Elem.(*types.FuncType)
			if !ok {
				panic(fmt.Errorf("invalid callee signature type, expected *types.FuncType, got %T", typ.Elem))
			}
			inst.Callee = callee
			inst.Sig = sig
			// TODO: Validate oldInst.Type against inst.Sig.
			for _, oldArg := range oldInst.Args {
				arg := m.irValue(oldArg)
				inst.Args = append(inst.Args, arg)
			}

		default:
			panic(fmt.Errorf("support for instruction %T not yet implemented", oldInst))
		}
	}

	// Fix terminator.
	switch oldTerm := oldBlock.Term.(type) {
	case *ast.TermRet:
		term := &ir.TermRet{
			Parent: block,
		}
		if oldTerm.X != nil {
			term.X = m.irValue(oldTerm.X)
		}
		block.Term = term
	case *ast.TermBr:
		panic("not yet implemented")
	case *ast.TermCondBr:
		panic("not yet implemented")
	case *ast.TermSwitch:
		panic("not yet implemented")
	case *ast.TermUnreachable:
		panic("not yet implemented")
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", oldTerm))
	}
}

// === [ Instructions ] ========================================================

// --- [ Binary instructions ] -------------------------------------------------

// --- [ Bitwise instructions ] ------------------------------------------------

// --- [ Memory instructions ] -------------------------------------------------

// --- [ Conversion instructions ] ---------------------------------------------

// --- [ Other instructions ] --------------------------------------------------

// === [ Terminators ] =========================================================
