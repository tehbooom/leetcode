package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	lengthOfPrefix := len(strs[0])
	shortestWord := 0

	for i, word := range strs {
		lengthOfWord := len(word)
		if lengthOfWord < lengthOfPrefix {
			shortestWord = i
			lengthOfPrefix = lengthOfWord
		}
	}

	prefix := strings.Split(strs[shortestWord], "")
	prefixFound := false

	for _, word := range strs {
		splitWord := strings.Split(word, "")
		for j := range lengthOfPrefix {
			if splitWord[j] == prefix[j] {
				prefixFound = true
				continue
			} else if j > 0 {
				prefix = prefix[:j]
				lengthOfPrefix = j
				break
			} else {
				prefixFound = false
				return ""
			}
		}
	}

	if prefixFound {
		return strings.Join(prefix, "")
	}

	return ""
}

func main() {
	fmt.Printf("Prefix: %s\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("Prefix: %s\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("Prefix: %s\n", longestCommonPrefix([]string{"a"}))
	fmt.Printf("Prefix: %s\n", longestCommonPrefix([]string{"reflower", "flow", "flight"}))

}
