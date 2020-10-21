package sammy

import (
	"fmt"

	"github.com/kayex/sammy/text"
)

type Normalizer func(string) string

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
	for _, n := range keys() {
		q := &text.SubQuery{
			Search: text.CaseInsensitiveWordQuery{
				WordQuery: text.WordQuery{
					W:         n,
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Sub: major(n),
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
			Sub: minor(n),
		}

		s = q.Apply(s)
	}

	return s
}
