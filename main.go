package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputA := "1010"
	inputB := "1011"
	output := addBinary(inputA, inputB)
	fmt.Println(output)
}

func addBinary(a, b string) string {
	first, second := a, b
	if len(first) < len(second) {
		first, second = second, first
	}
	result := make([]string, len(first)+1)

	temp := 0
	for i, j := len(first)-1, len(second)-1; i >= 0; i, j = i-1, j-1 {
		bin1, err := strconv.Atoi(string(first[i]))
		if err != nil {
			panic(err)
		}
		bin2 := 0
		if j >= 0 {
			bin2, err = strconv.Atoi(string(second[j]))
			if err != nil {
				panic(err)
			}
		}

		element := "0"
		countOnes := bin1 + bin2 + temp
		switch countOnes {
		case 1:
			element = "1"
			temp = 0
		case 2:
			temp = 1
		case 3:
			element = "1"
		}

		result[i+1] = element
	}
	if temp == 1 {
		result[0] = "1"
		temp = 0
		return strings.Join(result, "")
	}

	return strings.Join(result[1:], "")
}
