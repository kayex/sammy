package sammy

import (
	"fmt"

	"github.com/kayex/sammy/text"
)

type Transformer func(string) string

func NormalizeAccidentals(s string) string {
	accidentals := flats
	replacements := make(map[string]string)

	for _, a := range accidentals {
		m := flatMirror[a]
		replacements[a] = m
		replacements[major(a)] = major(m)
		replacements[minor(a)] = minor(m)
	}

	for search, replace := range replacements {
		q := &text.SubQuery{
			Search: text.CaseInsensitiveWordQuery{
				WordQuery: text.WordQuery{
					W:         search,
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Sub: replace,
		}

		s = q.Apply(s)
	}

	return s
}

func ExtendMajor(s string) string {
	os := s
	fmt.Println(os)
	for _, n := range keys() {
		q := &text.SubQuery{
			Search: text.CaseInsensitiveWordQuery{
				WordQuery: text.WordQuery{
					W:         n,
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Sub: fmt.Sprintf("%smaj", n),
		}

		s = q.Apply(s)
	}

	return s
}

func ExtendMinor(s string) string {
	for _, n := range keys() {
		q := &text.SubQuery{
			Search: text.CaseInsensitiveWordQuery{
				WordQuery: text.WordQuery{
					W:         fmt.Sprintf("%sm", n),
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Sub: fmt.Sprintf("%smin", n),
		}

		s = q.Apply(s)
	}

	return s
}
