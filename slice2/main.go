package main

import (
	"errors"
	"fmt"
)

var (
	ErrWrongNum = errors.New("wrong num, must be > 0")
)

func main() {
	res, err := rotate([]int{1, 2, 3, 4, 5}, 1)
	if err != nil {
		fmt.Printf("func rotate error: [%v]", err)
		return
	}
	fmt.Println(res)
}

func rotate(arr []int, num int) ([]int, error) {
	if num < 0 {
		return nil, ErrWrongNum
	}
	num = num % len(arr)

	sl := make([]int, len(arr)) // len, cap?

	for i := 0; i < len(arr); i++ {
		if i+num >= len(arr) {
			sl[i+num-len(arr)] = arr[i]
			continue
		}
		sl[i+num] = arr[i]
	}

	return sl, nil
}
