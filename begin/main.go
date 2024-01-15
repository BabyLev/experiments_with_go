package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "In programming, the term 'bug' can refer to an error, flaw, or fault in a computer program. It's intriguing how often bugs can disrupt a system, yet provide valuable lessons. Bugs, bugs, and more bugs!"
	fmt.Println(countWords(str))
}

// programming - 3
// in - 4
// ...

func countWords(str string) map[string]int {
	mapWords := make(map[string]int)

	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, "'", "", -1)
	str = strings.Replace(str, "!", "", -1)
	sl := strings.Split(str, " ")

	for i := 0; i < len(sl); i++ {
		if _, ok := mapWords[sl[i]]; ok {
			mapWords[sl[i]]++
		} else {
			mapWords[sl[i]] = 1
		}
	}
	return mapWords
}
