package byteline

import (
	"testing"
	"unicode/utf8"
)

func TestUTF8EndsWithSpace(t *testing.T) {
	if !UTF8EndsWithSpace([]byte("你好 ")) {
		t.Fatal("1")
	}
	if UTF8EndsWithSpace([]byte("你好")) {
		t.Fatal("2")
	}
	if !UTF8EndsWithSpace([]byte("你好\t")) {
		t.Fatal("3")
	}
	if UTF8EndsWithSpace([]byte{}) {
		t.Fatal("4")
	}
}

func TestUTF8IndexOfRune(t *testing.T) {
	out := UTF8IndexOfRune([]byte("你好]"), ']')
	if out != utf8.RuneLen('你')+utf8.RuneLen('好') {
		t.Fatal("1", out)
	}

	out = UTF8IndexOfRune([]byte(""), ']')
	if out != -1 {
		t.Fatal("2")
	}
}
