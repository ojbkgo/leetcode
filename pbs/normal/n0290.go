package normal

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	return wordPattern2([]byte(pattern), strings.Split(s, " "))
}

func wordPattern2(pattern []byte, s []string) bool {
	if len(pattern) != len(s) {
		return false
	}

	m := make(map[string]byte)
	m2 := make(map[byte]string)

	for i, v := range s {
		v1, ok1 := m2[pattern[i]]
		v2, ok2 := m[v]

		if !ok1 && !ok2 {
			m2[pattern[i]] = v
			m[v] = pattern[i]
		} else if ok1 && ok2 && v1 == v && pattern[i] == v2 {
			continue
		} else {
			return false
		}

	}

	return true
}

func Test_wordPattern2() {
	fmt.Println(wordPattern2([]byte("abbacc"), []string{"dog", "cat", "cat", "dog", "fish", "fish"}))
}
