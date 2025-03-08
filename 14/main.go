package main

import (
	"fmt"
	"slices"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	firstStringSet := []string{}
	prefix := []string{}
	lengthOfWords := []int{}
	for longetWord := range strs {
		lengthOfWords = append(lengthOfWords, len(strs[longetWord]))
	}
	for _, character := range strings.Split(strs[0], "") {
		firstStringSet = append(firstStringSet, character)
	}
	for i := 1; i < slices.Max(lengthOfWords); i++ {
		word := strs[i]
		for j := range word {
			if string(word[j]) == firstStringSet[j] {
				prefix = append(prefix, string(word[j]))
			}
		}
	}

	return strings.Join(prefix, "")
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

}
