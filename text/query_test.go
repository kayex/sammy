package text

import (
	"testing"
)

func TestWordQuery_Match(t *testing.T) {
	cases := []struct {
		s  string
		q  WordQuery
		i  int
		ln int
	}{
		{
			s:  "foo",
			q:  Word("foo"),
			i:  0,
			ln: 3,
		},
		{
			s:  "foo bar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
		{
			s:  "foo barbaz",
			q:  Word("bar"),
			i:  -1,
			ln: 0,
		},
		{
			s:  "foobar baz",
			q:  Word("bar"),
			i:  -1,
			ln: 0,
		},
		{
			s:  "foo bar",
			q:  Word("FOO"),
			i:  -1,
			ln: 0,
		},
		{
			s:  "FOO BAR",
			q:  Word("foo"),
			i:  -1,
			ln: 0,
		},
		{
			s:  "foo\tbar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
		{
			s:  "foo\nbar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
		{
			s:  "foo\vbar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
		{
			s:  "foo\fbar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
		{
			s:  "foo\rbar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
		{
			s:  "åäö bar",
			q:  Word("bar"),
			i:  4,
			ln: 3,
		},
	}

	for _, c := range cases {
		i, ln := c.q.Match(c.s)

		if i != c.i || ln != c.ln {
			t.Errorf("Expected Word(%q).Match(%q) to return (%v, %v) got (%v, %v)", c.q.W, c.s, c.i, c.ln, i, ln)
		}
	}
}

func TestCaseInsensitiveWordQuery_Match(t *testing.T) {
	cases := []struct {
		s  string
		q  CaseInsensitiveWordQuery
		i  int
		ln int
	}{
		{
			s:  "foo",
			q:  IWord("foo"),
			i:  0,
			ln: 3,
		},
		{
			s:  "FOO",
			q:  IWord("foo"),
			i:  0,
			ln: 3,
		},
	}

	for _, c := range cases {
		i, ln := c.q.Match(c.s)

		if i != c.i || ln != c.ln {
			t.Errorf("Expected IWord(%#v).Match(%q) to return (%v, %v), got (%v, %v)", c.q, c.s, c.i, c.ln, i, ln)
		}
	}
}

// BenchmarkWord_MatchNotExist_6_587 benchmarks a single Word query of length
// 6 against a search text of length 587, where the sought string does not
// exist in the search text.
//
// This benchmark gives a good indication of the average performance of an
// unsuccessful search.
func BenchmarkWord_MatchNotExist_6_587(b *testing.B) {
	w := Word("foobar")

	txt := `Lorem ipsum dolor sit amet, an cum vero soleat concludaturque, te purto vero reprimique vis.
	Ignota mediocritatem ut sea. Cetero deserunt pericula te vel. Omnis legendos no per.
	Sale illum pertinax no sed, est posse putent minimum no. Pri et vitae mentitum eligendi,
	no ius reque fugit libris, eos ad quaeque pericula mediocrem. Habemus corpora an mea,
	inermis partiendo per et, at nemore dolorem iudicabit eos. At est mucius docendi. Sed et nisl facilisi.
	Idque suavitate argumentum eu eam, vis putant insolens dissentiunt id. Dictas labitur in mei, duo omnium assentior scripserit cu omnium`

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		w.Match(txt)
	}
}

// BenchmarkWord_MatchPartials_5_587 benchmarks a single Word query of length
// 6 against a search text of length 587 with partials of length 5.
//
// This benchmark gives a good indication of the worst case performance
// of an unsuccessful search.
func BenchmarkWord_MatchPartials_6_587(b *testing.B) {
	w := Word("foobar")

	txt := `fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba fooba fooba
		fooba fooba fooba fooba fooba fooba fooba fooba`

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		w.Match(txt)
	}
}

// BenchmarkWord_MatchExist_6_587 benchmarks a single Word query of length
// 6 against a search text of length 587, where the sought string is at the
// very end of the search text.
func BenchmarkWord_MatchExist_6_587(b *testing.B) {
	w := Word("foobar")

	txt := `Lorem ipsum dolor sit amet, an cum vero soleat concludaturque, te purto vero reprimique vis.
	Ignota mediocritatem ut sea. Cetero deserunt pericula te vel. Omnis legendos no per.
	Sale illum pertinax no sed, est posse putent minimum no. Pri et vitae mentitum eligendi,
	no ius reque fugit libris, eos ad quaeque pericula mediocrem. Habemus corpora an mea,
	inermis partiendo per et, at nemore dolorem iudicabit eos. At est mucius docendi. Sed et nisl facilisi.
	Idque suavitate argumentum eu eam, vis putant insolens dissentiunt id. Dictas labitur in mei, duo omnium assentior scripserit cu foobar`

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		w.Match(txt)
	}
}
