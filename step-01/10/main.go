// Challenge: Calculator Edge Cases
// Task: given a simple calculator implement edge case and input verification
//		that outputs error messages and gracefully shuts down.

package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(operand string) (*float64, error) {
	parsedOp, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operand: %w", err)
	}

	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (*float64, error) {
	ops := strings.Fields(expr)
	nops := len(ops)
	if nops != 3 {
		return nil, fmt.Errorf("expected 3 expressions, got %d", nops)
	}

	left, err := parseOperand(ops[0])
	if err != nil {
		return nil, err
	}

	right, err := parseOperand(ops[2])
	if err != nil {
		return nil, err
	}

	f, ok := operators[ops[1]]
	if !ok {
		return nil, fmt.Errorf("unknown operator: %s", ops[1])
	}

	result := f(*left, *right)
	return &result, nil
}

func main() {
	expr := flag.String("expr", "", "The expression to calculate on, separated by spaces.")
	flag.Parse()

	result, err := calculate(*expr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s = %.2f\n", *expr, *result)
}
