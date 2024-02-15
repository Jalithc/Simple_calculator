package main

import "fmt"

func main() {
	num1 := 100
	num2 := 35
	addResult, err := calc(num1, num2, "+")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	subResult, err := calc(num1, num2, "-")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	multResult, err := calc(num1, num2, "*")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	divResult, err := calc(num1, num2, "/")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Addition result:", addResult)
	fmt.Println("Subtraction result:", subResult)
	fmt.Println("Multiplication result:", multResult)
	fmt.Println("Division result:", divResult)
}

func calc(num1, num2 int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		result = num1 / num2
	default:
		return 0, fmt.Errorf("invalid operator")
	}
	return result, nil
}
