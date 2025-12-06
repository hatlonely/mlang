# Makefile for mlang ANTLR grammar compilation

# Variables
GRAMMAR_FILE = mlang.g4
OUTPUT_DIR = gen/py
ANTLR = antlr

# Default target
all: python

# Generate Python code from ANTLR grammar
python: $(OUTPUT_DIR)/mlangLexer.py $(OUTPUT_DIR)/mlangParser.py

$(OUTPUT_DIR)/mlangLexer.py $(OUTPUT_DIR)/mlangParser.py: $(GRAMMAR_FILE) | $(OUTPUT_DIR)
	$(ANTLR) -Dlanguage=Python3 -o $(OUTPUT_DIR) $(GRAMMAR_FILE)
	touch $(OUTPUT_DIR)/__init__.py

# Create output directory if it doesn't exist
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Clean generated files
clean:
	rm -rf $(OUTPUT_DIR)/*.py
	rm -rf $(OUTPUT_DIR)/*.tokens
	rm -rf $(OUTPUT_DIR)/*.interp

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
	@echo "  all         - Generate Python code (default)"
	@echo "  python      - Generate Python code from ANTLR grammar"
	@echo "  check-antlr - Check ANTLR installation"
	@echo "  clean       - Remove generated Python files"
	@echo "  test        - Run tests on generated parser"
	@echo "  help        - Show this help message"

.PHONY: all python clean check-antlr test help