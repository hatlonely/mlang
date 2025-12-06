package semantic

import "errors"

// Symbol represents a symbol in the symbol table
type Symbol struct {
	Name string
	Type Type
}

// Scope represents a lexical scope
type Scope struct {
	symbols map[string]*Symbol
	parent  *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		symbols: make(map[string]*Symbol),
		parent:  parent,
	}
}

func (s *Scope) Define(name string, symbolType Type) error {
	if _, exists := s.symbols[name]; exists {
		return errors.New("symbol already defined in current scope: " + name)
	}
	s.symbols[name] = &Symbol{
		Name: name,
		Type: symbolType,
	}
	return nil
}

func (s *Scope) Lookup(name string) (*Symbol, bool) {
	if symbol, exists := s.symbols[name]; exists {
		return symbol, true
	}
	if s.parent != nil {
		return s.parent.Lookup(name)
	}
	return nil, false
}

// SymbolTable manages symbol scopes
type SymbolTable struct {
	currentScope *Scope
	globalScope  *Scope
}

func NewSymbolTable() *SymbolTable {
	global := NewScope(nil)
	st := &SymbolTable{
		currentScope: global,
		globalScope:  global,
	}
	st.initBuiltins()
	return st
}

func (st *SymbolTable) initBuiltins() {
	// 添加内置函数
	builtinFunctions := map[string]*FunctionType{
		"len": {
			ParamTypes: []Type{AnyType},
			ReturnType: IntType, // len 返回整数
			IsVariadic: false,
		},
		"abs": {
			ParamTypes: []Type{NumericBaseType},
			ReturnType: NumericBaseType, // abs 保持原类型
			IsVariadic: false,
		},
	}

	// 添加可变参数内置函数
	variadicFunctions := map[string]*FunctionType{
		"max": {
			ParamTypes:   []Type{}, // 无固定参数
			ReturnType:   NumericBaseType, // 返回与输入相同的数值类型
			IsVariadic:   true,
			VariadicType: NumericBaseType,
		},
		"min": {
			ParamTypes:   []Type{}, // 无固定参数
			ReturnType:   NumericBaseType,
			IsVariadic:   true,
			VariadicType: NumericBaseType,
		},
		"sum": {
			ParamTypes:   []Type{}, // 无固定参数
			ReturnType:   NumericBaseType,
			IsVariadic:   true,
			VariadicType: NumericBaseType,
		},
		"concat": {
			ParamTypes:   []Type{}, // 无固定参数
			ReturnType:   StringType,
			IsVariadic:   true,
			VariadicType: StringType,
		},
	}

	// 注册固定参数函数
	for name, funcType := range builtinFunctions {
		st.globalScope.Define(name, funcType)
	}

	// 注册可变参数函数
	for name, funcType := range variadicFunctions {
		st.globalScope.Define(name, funcType)
	}
}

func (st *SymbolTable) EnterScope() {
	st.currentScope = NewScope(st.currentScope)
}

func (st *SymbolTable) ExitScope() {
	if st.currentScope.parent != nil {
		st.currentScope = st.currentScope.parent
	}
}

func (st *SymbolTable) Define(name string, symbolType Type) error {
	return st.currentScope.Define(name, symbolType)
}

func (st *SymbolTable) Lookup(name string) (*Symbol, bool) {
	return st.currentScope.Lookup(name)
}

// RegisterFunction 注册一个新函数
func (st *SymbolTable) RegisterFunction(name string, paramTypes []Type, returnType Type) error {
	funcType := &FunctionType{
		ParamTypes: paramTypes,
		ReturnType: returnType,
		IsVariadic: false,
	}
	return st.globalScope.Define(name, funcType)
}

// RegisterVariadicFunction 注册一个新的可变参数函数
func (st *SymbolTable) RegisterVariadicFunction(name string, paramTypes []Type, variadicType Type, returnType Type) error {
	funcType := &FunctionType{
		ParamTypes:   paramTypes,
		ReturnType:   returnType,
		IsVariadic:   true,
		VariadicType: variadicType,
	}
	return st.globalScope.Define(name, funcType)
}

// RegisterBinaryOp 注册一个比较运算符函数
func (st *SymbolTable) RegisterBinaryOp(name string, leftType, rightType, returnType Type) error {
	funcType := &FunctionType{
		ParamTypes: []Type{leftType, rightType},
		ReturnType: returnType,
	}
	return st.globalScope.Define(name, funcType)
}

// UnregisterFunction 删除一个已注册的函数
func (st *SymbolTable) UnregisterFunction(name string) {
	delete(st.globalScope.symbols, name)
}

// ListFunctions 列出所有已注册的函数
func (st *SymbolTable) ListFunctions() map[string]*FunctionType {
	functions := make(map[string]*FunctionType)
	for name, symbol := range st.globalScope.symbols {
		if funcType, ok := symbol.Type.(*FunctionType); ok {
			functions[name] = funcType
		}
	}
	return functions
}
