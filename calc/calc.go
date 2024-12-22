package calc

import (
	"errors"
	"strconv"
	"unicode"
)

// Проверка, является ли символ оператором
func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

// Приоритет операторов
func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

// Преобразование инфиксного выражения в обратную польскую нотацию
func toPolishNotation(expr string) ([]string, error) {
	var polishNotation []string
	var operators []rune
	i := 0
	length := len(expr)

	for i < length {
		ch := rune(expr[i])

		// Если цифра или точка, парсим число
		if unicode.IsDigit(ch) || ch == '.' {
			start := i
			dotCount := 0
			for i < length && (unicode.IsDigit(rune(expr[i])) || rune(expr[i]) == '.') {
				if rune(expr[i]) == '.' {
					dotCount++
					if dotCount > 1 {
						return nil, errors.New("неправильный формат числа")
					}
				}
				i++
			}
			polishNotation = append(polishNotation, expr[start:i])
			continue
		}

		// Если открывающая скобка, кладем в стек
		if ch == '(' {
			operators = append(operators, ch)
		} else if ch == ')' {
			// Выгружаем из стека до открывающей скобки
			found := false
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				if top == '(' {
					found = true
					break
				}
				polishNotation = append(polishNotation, string(top))
			}
			if !found {
				return nil, errors.New("несоответствие скобок")
			}
		} else if isOperator(ch) {
			// Операторы
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				if isOperator(top) && precedence(top) >= precedence(ch) {
					operators = operators[:len(operators)-1]
					polishNotation = append(polishNotation, string(top))
				} else {
					break
				}
			}
			operators = append(operators, ch)
		} else {
			return nil, errors.New("неизвестный символ")
		}
		i++
	}

	// Выгружаем оставшиеся операторы из стека
	for len(operators) > 0 {
		top := operators[len(operators)-1]
		if top == '(' || top == ')' {
			return nil, errors.New("несоответствие скобок")
		}
		polishNotation = append(polishNotation, string(top))
		operators = operators[:len(operators)-1]
	}

	return polishNotation, nil
}

// Вычисление выражения в обратной польской нотации (ОПН)
func evaluatePolishNotation(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		// Если оператор
		if len(token) == 1 && (token == "+" || token == "-" || token == "*" || token == "/") {
			if len(stack) < 2 {
				return 0, errors.New("недостаточно операндов")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("деление на ноль")
				}
				result = a / b
			}
			stack = append(stack, result)
		} else {
			// Парсим число
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("неправильный операнд")
			}
			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("неправильное выражение")
	}

	return stack[0], nil
}

// Основная функция вычисления выражения
func Calc(expression string) (float64, error) {
	tokens, err := toPolishNotation(expression)
	if err != nil {
		return 0, err
	}

	result, err := evaluatePolishNotation(tokens)
	if err != nil {
		return 0, err
	}

	return result, nil
}
