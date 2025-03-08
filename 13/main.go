package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	intToAdd := []int{}
	romanSlice := strings.Split(s, "")
	i := 0
	for i < len(romanSlice) {
		if romanSlice[i] == "I" && i+1 < len(romanSlice) && romanSlice[i+1] == "V" {
			intToAdd = append(intToAdd, 4)
			i = i + 2
		} else if romanSlice[i] == "I" && i+1 < len(romanSlice) && romanSlice[i+1] == "X" {
			intToAdd = append(intToAdd, 9)
			i = i + 2
		} else if romanSlice[i] == "X" && i+1 < len(romanSlice) && romanSlice[i+1] == "L" {
			intToAdd = append(intToAdd, 40)
			i = i + 2
		} else if romanSlice[i] == "X" && i+1 < len(romanSlice) && romanSlice[i+1] == "C" {
			intToAdd = append(intToAdd, 90)
			i = i + 2
		} else if romanSlice[i] == "C" && i+1 < len(romanSlice) && romanSlice[i+1] == "D" {
			intToAdd = append(intToAdd, 400)
			i = i + 2
		} else if romanSlice[i] == "C" && i+1 < len(romanSlice) && romanSlice[i+1] == "M" {
			intToAdd = append(intToAdd, 900)
			i = i + 2
		} else {
			intToAdd = append(intToAdd, romanMap[romanSlice[i]])
			i++
		}
	}
	result := 0
	for j := range intToAdd {
		result = result + intToAdd[j]
	}
	return result
}

func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("III"))

}
