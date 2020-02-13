package interview

import . "github.com/ojbkgo/leetcode-common"

func replaceSpace(s string) string {
	b := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b = append(b, []byte("%20")...)
		} else {
			b = append(b, s[i])
		}
	}

	return string(b)
}

func Test_replaceSpace() {
	Dump(replaceSpace("we are famaly"))
}