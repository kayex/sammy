package text

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Query interface {
	Match(string) (i int, l int)
}

// WordQuery matches the first occurrence of W in a search text where W is
// surrounded by delimiters as defined by Delimiter.
type WordQuery struct {
	W         string
	Delimiter func(rune) bool
}

func Word(w string) WordQuery {
	if len(w) == 0 {
		panic("Cannot create WordQuery of length 0")
	}

	return WordQuery{w, WordDelimiter}
}

func (q WordQuery) Match(s string) (int, int) {
	if len(s) == 0 {
		return -1, 0
	}

	if s == q.W {
		return 0, q.Length()
	}

	_slen := len(s)
	fmt.Println(_slen)

	sr := []rune(s)
	offset := 0
	for offset < len(sr) {
		sr = sr[offset:]
		srs := string(sr)
		fmt.Println(srs)
		i := strings.Index(string(sr), q.W)
		if i < 0 {
			break
		}

		ir := utf8.RuneCountInString(s[:i])
		// Make sure that any preceding or following characters are valid
		// WordQuery delimiters.
		prev, p := at(sr, ir-1)
		next, n := at(sr, ir+q.Length())

		if p {
			if !q.Delimiter(prev) {
				offset += ir + q.Length()
				continue
			}
		}

		if n {
			if !q.Delimiter(next) {
				offset += ir + q.Length()
				continue
			}
		}

		return offset + ir, q.Length()
	}

	return -1, 0

}

func (q WordQuery) Length() int {
	return utf8.RuneCountInString(q.W)
}

type CaseInsensitiveWordQuery struct {
	WordQuery
}

func IWord(w string) CaseInsensitiveWordQuery {
	return CaseInsensitiveWordQuery{Word(w)}
}

func (q CaseInsensitiveWordQuery) Match(text string) (int, int) {
	sl := strings.ToLower(text)
	q.W = strings.ToLower(q.W)

	return q.WordQuery.Match(sl)
}

// FilenameComponentDelimiter returns a bool indicating if
// r is a reasonable filename component delimiter.
func FilenameComponentDelimiter(r rune) bool {
	switch r {
	case '_', '-':
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

// at returns the rune at index i, and a bool indicating if i exists in s.
func at(s []rune, i int) (rune, bool) {
	if i < 0 || i > len(s)-1 {
		return 0, false
	}

	r := s[i]
	return r, true
}
