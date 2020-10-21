package text

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Query interface {
	Match(string) (i int, l int)
}

// Word matches the first occurrence of W in a search text where W is
// surrounded by delimiters that satisfy Delimiter.
type Word struct {
	W         string
	Delimiter func(rune) bool
}

func NewWord(w string) Word {
	if len(w) == 0 {
		panic("Cannot create WordQuery of length 0")
	}

	return Word{w, WordDelimiter}
}

func (q Word) Match(s string) (int, int) {
	if len(s) == 0 {
		return -1, 0
	}

	if s == q.W {
		return 0, q.Length()
	}

	sr := []rune(s)

	offset := 0
	for offset < len(sr) {
		so := string(sr[offset:])
		i := strings.Index(so, q.W)
		if i < 0 {
			break
		}

		// Make sure that any preceding or following characters are valid delimiters.
		prev, p := atByte(s, offset+i-1)
		next, n := atByte(s, offset+i+len(q.W))
		if (p && !q.Delimiter(rune(prev))) || (n && !q.Delimiter(rune(next))) {
			offset += i + len(q.W)
			continue
		}

		// Get true rune index in s.
		ti := utf8.RuneCountInString(s[:offset+i])

		return ti, q.Length()
	}

	return -1, 0

}

func (q Word) Length() int {
	return utf8.RuneCountInString(q.W)
}

type CaseInsensitiveWord struct {
	Word
}

func IWord(w string) CaseInsensitiveWord {
	return CaseInsensitiveWord{NewWord(w)}
}

func (q CaseInsensitiveWord) Match(text string) (int, int) {
	sl := strings.ToLower(text)
	q.W = strings.ToLower(q.W)

	return q.Word.Match(sl)
}

// FilenameComponentDelimiter returns a bool indicating if
// r is a reasonable filename component delimiter.
func FilenameComponentDelimiter(r rune) bool {
	switch r {
	case '_', '-', '.':
		return true
	}

	return false
}

// WordDelimiter return as bool indicating if r is a word delimiter.
func WordDelimiter(r rune) bool {
	if unicode.IsSpace(r) {
		return true
	}

	switch r {
	case ',', '.':
		return true
	}

	return false
}

// atRune returns the rune at rune index i, and a bool indicating if i exists in s.
func atRune(s []rune, i int) (rune, bool) {
	if i < 0 || i > len(s)-1 {
		return 0, false
	}

	r := s[i]
	return r, true
}

// atByte returns the byte at index i, and a bool indicating if i exists in s.
func atByte(s string, i int) (byte, bool) {
	if i < 0 || i > len(s)-1 {
		return 0, false
	}

	r := s[i]
	return r, true
}
