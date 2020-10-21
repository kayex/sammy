package sammy

import (
	"fmt"

	"github.com/kayex/sammy/text"
)

type Normalizer func(string) string

func NormalizeAccidentals(s string) string {
	replacements := make(map[string]string)

	for _, a := range flats {
		m := flatMirror[a]
		replacements[a] = m
		replacements[major(a)] = major(m)
		replacements[minor(a)] = minor(m)
	}

	for search, replace := range replacements {
		q := &text.ReplaceQuery{
			Search: text.CaseInsensitiveWord{
				Word: text.Word{
					W:         search,
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Replacement: replace,
		}

		s = q.Apply(s)
	}

	return s
}

func ExtendMajor(s string) string {
	for _, k := range keys() {
		q := &text.ReplaceQuery{
			Search: text.CaseInsensitiveWord{
				Word: text.Word{
					W:         k,
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Replacement: major(k),
		}

		s = q.Apply(s)
	}

	return s
}

func ExtendMinor(s string) string {
	for _, k := range keys() {
		q := &text.ReplaceQuery{
			Search: text.CaseInsensitiveWord{
				Word: text.Word{
					W:         fmt.Sprintf("%sm", k),
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Replacement: minor(k),
		}

		s = q.Apply(s)
	}

	return s
}
