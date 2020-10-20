package sammy

import (
	"fmt"

	"github.com/kayex/sammy/text"
)

type Transformer func(string) string

func ExtendMajor(s string) string {
	fmt.Printf("Extending major for %v\n", s)

	before := s

	for i, n := range notes() {
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

		fmt.Println(i)
		fmt.Println(s)
		fmt.Println(before)
	}
	fmt.Printf("Last iteration: %s\n", s)

	return s
}

func ExtendMinor(s string) string {
	for _, n := range notes() {
		q := &text.SubQuery{
			Search: text.CaseInsensitiveWordQuery{
				WordQuery: text.WordQuery{
					W:         fmt.Sprintf("%vm", n),
					Delimiter: text.FilenameComponentDelimiter,
				},
			},
			Sub: fmt.Sprintf("%vmin", n),
		}

		s = q.Apply(s)
	}

	return s
}
