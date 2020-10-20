package text

import (
	"bytes"
	"strings"
)

type Mutation interface {
	Apply(text string) string
}

type Set struct {
	Text string
}

type Sub struct {
	Search string
	Sub    string
}

type SubQuery struct {
	Search Query
	Sub    string
}

type Append struct {
	Appendix string
}

type Prepend struct {
	Prefix string
}

func (r *Set) Apply(text string) string {
	return r.Text
}

func (s *Sub) Apply(text string) string {
	return strings.Replace(text, s.Search, s.Sub, -1)
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

func (a *Append) Apply(text string) string {
	if len(a.Appendix) == 0 {
		return text
	}

	return text + a.Appendix
}

func (p *Prepend) Apply(text string) string {
	if len(p.Prefix) == 0 {
		return text
	}

	return p.Prefix + text
}
