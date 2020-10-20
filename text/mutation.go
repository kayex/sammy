package text

import (
	"bytes"
)

type Mutation interface {
	Apply(text string) string
}

type SubQuery struct {
	Search Query
	Sub    string
}

func (s *SubQuery) Apply(text string) string {
	for i, ln := s.Search.Match(text); i >= 0; i, ln = s.Search.Match(text) {
		tr := []rune(text)

		beginning := tr[:i]
		end := tr[i+ln:]

		var buf bytes.Buffer

		buf.WriteString(string(beginning))
		buf.WriteString(s.Sub)
		buf.WriteString(string(end))

		text = buf.String()
	}

	return text
}
