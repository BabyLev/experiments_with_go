package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	var c string
	_, err := fmt.Scanf("%d%s%d", &a, &c, &b)
	if err != nil {
		fmt.Printf("Are you really stupid? Your error: [%v]\n", err)
	}
	if c == "+" {
		fmt.Println(sum(a, b))
	}
	if c == "*" {
		fmt.Println(mult(a, b))
	}
	if c == "/" {
		res, err := div(a, b)
		if err != nil {
			fmt.Printf("Div error: [%v]", err)
			return
		}
		fmt.Println(res)
	}
}

// функция, складывающая два числа
func sum(a, b int) int {
	return a + b
}

// функция, умножащая 2 числа
func mult(a, b int) int {
	return a * b
}

func div(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return float64(a) / float64(b), nil
}
