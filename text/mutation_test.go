package text

import "testing"

func TestSubWord_Apply(t *testing.T) {
	cases := []struct {
		str string
		sw  ReplaceQuery
		exp string
	}{
		{
			str: "foo bar",
			sw:  ReplaceQuery{NewWord("bar"), "boo"},
			exp: "foo boo",
		},
		{
			str: "foo åäö",
			sw:  ReplaceQuery{NewWord("foo"), "bar"},
			exp: "bar åäö",
		},
		{
			str: "foo åäö",
			sw:  ReplaceQuery{NewWord("åäö"), "bar"},
			exp: "foo bar",
		},
		{
			str: "foo barbaz",
			sw:  ReplaceQuery{NewWord("foo"), "long replacement"},
			exp: "long replacement barbaz",
		},
	}

	for _, c := range cases {
		act := c.sw.Apply(c.str)

		if act != c.exp {
			t.Errorf("Expected SubQuery{%#v, %q}.Apply(%q) to return %q, got %q", c.sw.Search, c.sw.Replacement, c.str, c.exp, act)
		}
	}
}
