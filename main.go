package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
func main() {
	input := "quick quick brown fox jumped over the lazy dog"
	rev, err := Reverse(input)
	if err != nil {
		fmt.Println(err)
	}
	doubleRev, err := Reverse(rev)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("原来: %q\n", input)
	fmt.Printf("反转: %q\n", rev)
	fmt.Printf("再反转: %q\n", doubleRev)
}
