package calculate

import (
	"strconv"
	"strings"

	"github.com/artemkkkkk/http-calculator/pkg/custom_errors"
)

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isOperation(s string) bool {
	switch s {
	case "+", "-", "*", "/", "(", ")":
		return true
	default:
		return false
	}
}

func splitTokens(expr string) []string {
	var units []string
	var subString = ""

	for _, val := range strings.Split(expr, "") {
		if isNumber(val) {
			subString += val

		} else if subString != "" {
			units = append(units, subString)
			subString = ""
		}
		if isOperation(val) {
			units = append(units, val)
		}
	}

	if subString != "" {
		units = append(units, subString)
	}

	return units
}

func tokenize(expr []string) []Token {
	var tokens []Token

	for _, part := range expr {
		switch part {
		case "+":
			tokens = append(tokens, Token{TokenPlus, part})
		case "-":
			tokens = append(tokens, Token{TokenMinus, part})
		case "*":
			tokens = append(tokens, Token{TokenMul, part})
		case "/":
			tokens = append(tokens, Token{TokenDiv, part})
		case "(":
			tokens = append(tokens, Token{TokenLParen, part})
		case ")":
			tokens = append(tokens, Token{TokenRParen, part})
		default:
			tokens = append(tokens, Token{TokenNumber, part})
		}
	}

	return tokens
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func infixToPostfix(tokens []Token) []Token {
	var output []Token
	var stack []Token

	for _, token := range tokens {
		switch token.Type {
		case TokenNumber:
			output = append(output, token)
		case TokenPlus, TokenMinus, TokenMul, TokenDiv:
			for len(stack) > 0 && precedence(stack[len(stack)-1].Value) >= precedence(token.Value) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case TokenLParen:
			stack = append(stack, token)
		case TokenRParen:
			for len(stack) > 0 && stack[len(stack)-1].Type != TokenLParen {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output
}

func evalPostfix(tokens []Token) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		if token.Type == TokenNumber {
			num, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, custom_errors.InvalidExpression
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token.Value {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, custom_errors.InvalidExpression
				}
				result = a / b
			default:
				return 0, custom_errors.UnknownOperator
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, custom_errors.InvalidExpression
	}

	return stack[0], nil
}

func Eval(expression string) (float64, error) {
	splitTokens := splitTokens(expression)
	tokens := tokenize(splitTokens)
	postfix := infixToPostfix(tokens)
	return evalPostfix(postfix)
}
