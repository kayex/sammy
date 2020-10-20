package sammy

import (
	"fmt"

	"github.com/kayex/sammy/text"
)

type Transformer func(string) string

func ExtendMajor(s string) string {
	os := s
	fmt.Println(os)
	for _, n := range notes() {
		sub := fmt.Sprintf("%smaj", n)
		q := &text.SubQuery{
			Search: text.CaseInsensitiveWordQuery{
				WordQuery: text.WordQuery{
					W:         n,
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Sub: sub,
		}

		s = q.Apply(s)
	}

	return s
}

func ExtendMinor(s string) string {
	for _, n := range notes() {
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
