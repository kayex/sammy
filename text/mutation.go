package text

import (
	"bytes"
)

type Mutation interface {
	Apply(text string) string
}

type ReplaceQuery struct {
	Search Query
	Replacement    string
}

func (s *ReplaceQuery) Apply(text string) string {
	for {
		i, ln := s.Search.Match(text)
		if i < 0 {
			return text
		}

		tr := []rune(text)

		beginning := tr[:i]
		end := tr[i+ln:]

		var buf bytes.Buffer

		buf.WriteString(string(beginning))
		buf.WriteString(s.Replacement)
		buf.WriteString(string(end))

		text = buf.String()
	}
}
