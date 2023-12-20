package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(sum(a, b))
}

func sum(a, b int) int {
	return a + b
}
