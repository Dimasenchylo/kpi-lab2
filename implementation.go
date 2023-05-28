package lab2

import (
	"fmt"
	"strings"
)

func PostfixToPrefix(postfix string) (string, error) {
    stack := make([]string, 0)

    if postfix == "" {
        return "", fmt.Errorf("invalid expression")
    }

    for _, token := range strings.Split(postfix, " ") {
        switch {
		case IsOperator(token):
            if len(stack) < 2 {
                return "", fmt.Errorf("not enough operands")
            }
            number1 := stack[len(stack)-1]
			number2 := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            prefix := fmt.Sprintf("%s %s %s", token, number2, number1)
            stack = append(stack, prefix)
        default:
            stack = append(stack, token)
        }
    }

    if len(stack) != 1 {
        return "", fmt.Errorf("invalid expression")
    }

    return stack[0], nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "/" || token == "*" || token == "^";
}

