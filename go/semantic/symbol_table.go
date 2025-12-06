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
			ReturnType: NumberType,
		},
		"max": {
			ParamTypes: []Type{NumberType, NumberType},
			ReturnType: NumberType,
		},
		"min": {
			ParamTypes: []Type{NumberType, NumberType},
			ReturnType: NumberType,
		},
		"abs": {
			ParamTypes: []Type{NumberType},
			ReturnType: NumberType,
		},
	}

	for name, funcType := range builtinFunctions {
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