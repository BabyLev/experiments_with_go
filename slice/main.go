package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceA := []int{3, 4, 6, 8, 2}
	sliceB := []int{1, 2, 5, 7}
	fmt.Println(mergeSl(sliceA, sliceB))
}

func mergeSl(sl1 []int, sl2 []int) []int {
	var res []int
	var slLen int = len(sl1)
	if len(sl1) < len(sl2) {
		slLen = len(sl2)
	}
	for i := 0; i < slLen; i++ {

		if i < len(sl1) && sl1[i]%2 == 0 {
			res = append(res, sl1[i])
		}
		if i < len(sl2) && sl2[i]%2 != 0 {
			res = append(res, sl2[i])
		}
	}

	sort.Ints(res)
	return res
}
