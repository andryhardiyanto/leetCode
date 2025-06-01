package main

import "fmt"

func main() {
	s := "aec"
	t := "abcde"

	if s == "" {
		fmt.Println("true")
	}

	temp := make([]string, 0)
	index := 0

	for i := 0; i < len(s); i++ {
		temp = append(temp, string(s[i]))
	}

	for i := 0; i < len(t); i++ {
		if string(t[i]) == temp[index] {
			index++
		}
		if index == len(temp) {
			break
		}
	}

	if index == len(temp) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
