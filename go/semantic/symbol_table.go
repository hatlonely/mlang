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
	// 支持函数重载：函数名 -> 重载列表
	functionOverloads map[string][]*Symbol
	// 支持 binaryOp 重载：运算符名 -> 重载列表
	binaryOpOverloads map[string][]*Symbol
	parent           *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		symbols:           make(map[string]*Symbol),
		functionOverloads: make(map[string][]*Symbol),
		binaryOpOverloads: make(map[string][]*Symbol),
		parent:            parent,
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

// DefineFunction 定义一个函数，支持重载
func (s *Scope) DefineFunction(name string, funcType *FunctionType) error {
	symbol := &Symbol{
		Name: name,
		Type: funcType,
	}
	
	// 检查是否已存在相同签名的重载
	if overloads, exists := s.functionOverloads[name]; exists {
		for _, existing := range overloads {
			if existingFunc, ok := existing.Type.(*FunctionType); ok {
				if functionsSignatureEqual(funcType, existingFunc) {
					return errors.New("function with same signature already defined: " + name)
				}
			}
		}
		s.functionOverloads[name] = append(s.functionOverloads[name], symbol)
	} else {
		s.functionOverloads[name] = []*Symbol{symbol}
	}
	
	// 如果这是第一个重载，也在普通符号表中注册（向后兼容）
	if len(s.functionOverloads[name]) == 1 {
		s.symbols[name] = symbol
	}
	
	return nil
}

// LookupFunctionOverloads 查找函数的所有重载
func (s *Scope) LookupFunctionOverloads(name string) ([]*Symbol, bool) {
	if overloads, exists := s.functionOverloads[name]; exists {
		return overloads, true
	}
	if s.parent != nil {
		return s.parent.LookupFunctionOverloads(name)
	}
	return nil, false
}

// DefineBinaryOp 定义一个新的 binary operator 重载
func (s *Scope) DefineBinaryOp(name string, binaryOpType *BinaryOpType) error {
	symbol := &Symbol{
		Name: name,
		Type: binaryOpType,
	}
	
	// 检查是否已存在相同签名的重载
	if overloads, exists := s.binaryOpOverloads[name]; exists {
		for _, existing := range overloads {
			if existingBinOp, ok := existing.Type.(*BinaryOpType); ok {
				if binaryOpSignatureEqual(binaryOpType, existingBinOp) {
					return errors.New("binary operator with same signature already defined: " + name)
				}
			}
		}
		s.binaryOpOverloads[name] = append(s.binaryOpOverloads[name], symbol)
	} else {
		s.binaryOpOverloads[name] = []*Symbol{symbol}
	}
	
	// 如果这是第一个重载，也在普通符号表中注册（向后兼容）
	if len(s.binaryOpOverloads[name]) == 1 {
		// 为了向后兼容，也创建一个 FunctionType 表示
		funcType := &FunctionType{
			ParamTypes: []Type{binaryOpType.LeftType, binaryOpType.RightType},
			ReturnType: binaryOpType.ReturnType,
		}
		s.symbols[name] = &Symbol{
			Name: name,
			Type: funcType,
		}
	}
	
	return nil
}

// LookupBinaryOpOverloads 查找 binary operator 的所有重载
func (s *Scope) LookupBinaryOpOverloads(name string) ([]*Symbol, bool) {
	if overloads, exists := s.binaryOpOverloads[name]; exists {
		return overloads, true
	}
	if s.parent != nil {
		return s.parent.LookupBinaryOpOverloads(name)
	}
	return nil, false
}

// functionsSignatureEqual 检查两个函数签名是否相等
func functionsSignatureEqual(f1, f2 *FunctionType) bool {
	if len(f1.ParamTypes) != len(f2.ParamTypes) {
		return false
	}
	if f1.IsVariadic != f2.IsVariadic {
		return false
	}
	
	for i, p1 := range f1.ParamTypes {
		if !p1.Equals(f2.ParamTypes[i]) {
			return false
		}
	}
	
	if f1.IsVariadic && !f1.VariadicType.Equals(f2.VariadicType) {
		return false
	}
	
	return true
}

// binaryOpSignatureEqual 检查两个 binary operator 签名是否相等
func binaryOpSignatureEqual(b1, b2 *BinaryOpType) bool {
	return b1.LeftType.Equals(b2.LeftType) && 
		   b1.RightType.Equals(b2.RightType) && 
		   b1.ReturnType.Equals(b2.ReturnType)
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
	// 不添加任何内置函数，所有函数都通过用户注册
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
	return st.globalScope.DefineFunction(name, funcType)
}

// RegisterVariadicFunction 注册一个新的可变参数函数
func (st *SymbolTable) RegisterVariadicFunction(name string, paramTypes []Type, variadicType Type, returnType Type) error {
	funcType := &FunctionType{
		ParamTypes:   paramTypes,
		ReturnType:   returnType,
		IsVariadic:   true,
		VariadicType: variadicType,
	}
	return st.globalScope.DefineFunction(name, funcType)
}

// RegisterBinaryOp 注册一个比较运算符函数
func (st *SymbolTable) RegisterBinaryOp(name string, leftType, rightType, returnType Type) error {
	binaryOpType := &BinaryOpType{
		LeftType:   leftType,
		RightType:  rightType,
		ReturnType: returnType,
	}
	return st.globalScope.DefineBinaryOp(name, binaryOpType)
}

// LookupFunctionOverloads 查找函数的所有重载
func (st *SymbolTable) LookupFunctionOverloads(name string) ([]*Symbol, bool) {
	return st.globalScope.LookupFunctionOverloads(name)
}

// LookupBinaryOpOverloads 查找 binaryOp 的所有重载
func (st *SymbolTable) LookupBinaryOpOverloads(name string) ([]*Symbol, bool) {
	return st.globalScope.LookupBinaryOpOverloads(name)
}

// ResolveFunctionOverload 根据参数类型解析最佳的函数重载
func (st *SymbolTable) ResolveFunctionOverload(name string, argTypes []Type) (*Symbol, error) {
	overloads, exists := st.LookupFunctionOverloads(name)
	if !exists {
		return nil, errors.New("function not found: " + name)
	}
	
	var bestMatch *Symbol
	var bestScore int = -1
	
	for _, overload := range overloads {
		funcType, ok := overload.Type.(*FunctionType)
		if !ok {
			continue
		}
		
		score := st.calculateMatchScore(funcType, argTypes)
		if score > bestScore {
			bestScore = score
			bestMatch = overload
		}
	}
	
	if bestMatch == nil {
		return nil, errors.New("no matching overload found for function: " + name)
	}
	
	return bestMatch, nil
}

// calculateMatchScore 计算函数签名与参数类型的匹配分数
func (st *SymbolTable) calculateMatchScore(funcType *FunctionType, argTypes []Type) int {
	requiredParams := len(funcType.ParamTypes)
	totalArgs := len(argTypes)
	
	// 检查参数数量是否匹配
	if !funcType.IsVariadic {
		if totalArgs != requiredParams {
			return -1 // 不匹配
		}
	} else {
		if totalArgs < requiredParams {
			return -1 // 参数不足
		}
	}
	
	score := 0
	
	// 计算固定参数的匹配分数
	for i := 0; i < requiredParams && i < totalArgs; i++ {
		paramType := funcType.ParamTypes[i]
		argType := argTypes[i]
		
		if argType.Equals(paramType) {
			score += 100 // 完全匹配
		} else if st.canImplicitlyCast(argType, paramType) {
			score += 50 // 可以隐式转换
		} else {
			return -1 // 不兼容
		}
	}
	
	// 计算可变参数的匹配分数，同时考虑类型统一性
	if funcType.IsVariadic {
		variadicType := funcType.VariadicType
		variadicArgs := argTypes[requiredParams:]
		
		// 检查所有可变参数是否能统一到目标类型
		if st.canUnifyToType(variadicArgs, variadicType) {
			// 计算统一后的匹配分数
			for i := requiredParams; i < totalArgs; i++ {
				argType := argTypes[i]
				
				if argType.Equals(variadicType) {
					score += 100 // 完全匹配
				} else if st.canImplicitlyCast(argType, variadicType) {
					score += 50 // 可以隐式转换
				} else {
					return -1 // 不兼容
				}
			}
		} else {
			return -1 // 无法统一类型
		}
	}
	
	return score
}

// canUnifyToType 检查是否可以将所有参数统一到目标类型
func (st *SymbolTable) canUnifyToType(argTypes []Type, targetType Type) bool {
	for _, argType := range argTypes {
		if !argType.Equals(targetType) && !st.canImplicitlyCast(argType, targetType) {
			return false
		}
	}
	return true
}

// ResolveBinaryOpOverload 根据参数类型解析最佳的 binaryOp 重载
func (st *SymbolTable) ResolveBinaryOpOverload(name string, leftType, rightType Type) (*Symbol, error) {
	overloads, exists := st.LookupBinaryOpOverloads(name)
	if !exists {
		return nil, errors.New("binary operator not found: " + name)
	}
	
	var bestMatch *Symbol
	var bestScore int = -1
	
	for _, overload := range overloads {
		binaryOpType, ok := overload.Type.(*BinaryOpType)
		if !ok {
			continue
		}
		
		score := st.calculateBinaryOpMatchScore(binaryOpType, leftType, rightType)
		if score > bestScore {
			bestScore = score
			bestMatch = overload
		}
	}
	
	if bestMatch == nil {
		return nil, errors.New("no matching overload found for binary operator: " + name)
	}
	
	return bestMatch, nil
}

// calculateBinaryOpMatchScore 计算 binaryOp 签名与参数类型的匹配分数
func (st *SymbolTable) calculateBinaryOpMatchScore(binaryOpType *BinaryOpType, leftType, rightType Type) int {
	score := 0
	
	// 计算左操作数的匹配分数
	if leftType.Equals(binaryOpType.LeftType) {
		score += 100 // 完全匹配
	} else if st.canImplicitlyCast(leftType, binaryOpType.LeftType) {
		score += 50 // 可以隐式转换
	} else {
		return -1 // 不兼容
	}
	
	// 计算右操作数的匹配分数
	if rightType.Equals(binaryOpType.RightType) {
		score += 100 // 完全匹配
	} else if st.canImplicitlyCast(rightType, binaryOpType.RightType) {
		score += 50 // 可以隐式转换
	} else {
		return -1 // 不兼容
	}
	
	return score
}

// canImplicitlyCast 检查是否可以进行隐式类型转换
func (st *SymbolTable) canImplicitlyCast(from, to Type) bool {
	if from.Equals(to) {
		return true
	}
	
	// int -> float 隐式转换
	if from.Equals(IntType) && to.Equals(FloatType) {
		return true
	}
	
	// 数值类型与NumericBaseType的兼容性
	if to.Equals(NumericBaseType) {
		return from.Equals(IntType) || from.Equals(FloatType)
	}
	
	// 检查类型兼容性
	return from.IsCompatibleWith(to)
}

// UnregisterFunction 删除一个已注册的函数
func (st *SymbolTable) UnregisterFunction(name string) {
	delete(st.globalScope.symbols, name)
	delete(st.globalScope.functionOverloads, name)
}

// RegisterVariable registers a variable in the global scope
func (st *SymbolTable) RegisterVariable(name string, varType Type) error {
	return st.globalScope.Define(name, varType)
}

// RegisterPropertyVariable registers a property variable with getter and setter
func (st *SymbolTable) RegisterPropertyVariable(name string, baseType Type, getter, setter string) error {
	propertyType := &PropertyType{
		BaseType: baseType,
		Getter:   getter,
		Setter:   setter,
	}
	return st.globalScope.Define(name, propertyType)
}

// RegisterStructType 注册一个结构体类型
func (st *SymbolTable) RegisterStructType(name string, fields map[string]Type) error {
	structType := &StructType{
		Name:   name,
		Fields: fields,
	}
	return st.globalScope.Define(name, structType)
}

// RegisterStructVariable 注册一个结构体类型的变量
func (st *SymbolTable) RegisterStructVariable(varName string, structTypeName string) error {
	// 查找结构体类型
	symbol, exists := st.globalScope.Lookup(structTypeName)
	if !exists {
		return errors.New("struct type not found: " + structTypeName)
	}
	
	structType, ok := symbol.Type.(*StructType)
	if !ok {
		return errors.New(structTypeName + " is not a struct type")
	}
	
	return st.globalScope.Define(varName, structType)
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

// ListStructTypes 列出所有已注册的结构体类型
func (st *SymbolTable) ListStructTypes() map[string]*StructType {
	structs := make(map[string]*StructType)
	for name, symbol := range st.globalScope.symbols {
		if structType, ok := symbol.Type.(*StructType); ok {
			structs[name] = structType
		}
	}
	return structs
}
