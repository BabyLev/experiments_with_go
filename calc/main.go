package main

import "fmt"

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
}

// функция, складывающая два числа
func sum(a, b int) int {
	return a + b
}

// функция, умножащая 2 числа
func mult(a, b int) int {
	return a * b
}
