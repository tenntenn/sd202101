package list14

import (
	"io"
	"io/ioutil"
	"unicode"
	"unicode/utf8"
)

func UpperCount(r io.Reader) (count int, _ error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		b = b[size:]
		if r != utf8.RuneError && unicode.IsUpper(r) {
			count++
		}
	}
	return count, nil
}
