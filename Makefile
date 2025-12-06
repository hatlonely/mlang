# Makefile for mlang ANTLR grammar compilation

# Variables
GRAMMAR_FILE = mlang.g4
PYTHON_OUTPUT_DIR = gen/py
GOLANG_OUTPUT_DIR = gen/go
ANTLR = antlr

# Default target
all: python golang

# Generate Python code from ANTLR grammar
python: $(PYTHON_OUTPUT_DIR)/mlangLexer.py $(PYTHON_OUTPUT_DIR)/mlangParser.py

$(PYTHON_OUTPUT_DIR)/mlangLexer.py $(PYTHON_OUTPUT_DIR)/mlangParser.py: $(GRAMMAR_FILE) | $(PYTHON_OUTPUT_DIR)
	$(ANTLR) -Dlanguage=Python3 -o $(PYTHON_OUTPUT_DIR) $(GRAMMAR_FILE)
	touch $(PYTHON_OUTPUT_DIR)/__init__.py

# Generate Golang code from ANTLR grammar
golang: $(GOLANG_OUTPUT_DIR)/mlang_lexer.go $(GOLANG_OUTPUT_DIR)/mlang_parser.go

$(GOLANG_OUTPUT_DIR)/mlang_lexer.go $(GOLANG_OUTPUT_DIR)/mlang_parser.go: $(GRAMMAR_FILE) | $(GOLANG_OUTPUT_DIR)
	$(ANTLR) -Dlanguage=Go -o $(GOLANG_OUTPUT_DIR) $(GRAMMAR_FILE)

# Create output directories if they don't exist
$(PYTHON_OUTPUT_DIR):
	mkdir -p $(PYTHON_OUTPUT_DIR)

$(GOLANG_OUTPUT_DIR):
	mkdir -p $(GOLANG_OUTPUT_DIR)

# Clean generated files
clean:
	rm -rf $(PYTHON_OUTPUT_DIR)/*.py
	rm -rf $(PYTHON_OUTPUT_DIR)/*.tokens
	rm -rf $(PYTHON_OUTPUT_DIR)/*.interp
	rm -rf $(GOLANG_OUTPUT_DIR)/*.go
	rm -rf $(GOLANG_OUTPUT_DIR)/*.tokens
	rm -rf $(GOLANG_OUTPUT_DIR)/*.interp

# Clean Python files only
clean-python:
	rm -rf $(PYTHON_OUTPUT_DIR)/*.py
	rm -rf $(PYTHON_OUTPUT_DIR)/*.tokens
	rm -rf $(PYTHON_OUTPUT_DIR)/*.interp

# Clean Golang files only
clean-golang:
	rm -rf $(GOLANG_OUTPUT_DIR)/*.go
	rm -rf $(GOLANG_OUTPUT_DIR)/*.tokens
	rm -rf $(GOLANG_OUTPUT_DIR)/*.interp

# Check ANTLR installation
check-antlr:
	@which $(ANTLR) > /dev/null || (echo "Error: ANTLR not found. Install with: brew install antlr" && exit 1)
	@echo "ANTLR found: $$(which $(ANTLR))"
	@$(ANTLR) 2>&1 | head -1

# Test the generated parser (requires Python antlr4-python3-runtime)
test: python
	@echo "Testing generated parser..."
	cd tests && python run_tests.py

# Help
help:
	@echo "Available targets:"
	@echo "  all           - Generate Python and Golang code (default)"
	@echo "  python        - Generate Python code from ANTLR grammar"
	@echo "  golang        - Generate Golang code from ANTLR grammar"
	@echo "  check-antlr   - Check ANTLR installation"
	@echo "  clean         - Remove all generated files"
	@echo "  clean-python  - Remove generated Python files only"
	@echo "  clean-golang  - Remove generated Golang files only"
	@echo "  test          - Run tests on generated parser"
	@echo "  help          - Show this help message"

.PHONY: all python golang clean clean-python clean-golang check-antlr test help