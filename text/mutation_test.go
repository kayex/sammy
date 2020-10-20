package text

import "testing"

func TestSubWord_Apply(t *testing.T) {
	cases := []struct {
		str string
		sw  SubQuery
		exp string
	}{
		{
			str: "foo bar",
			sw:  SubQuery{Word("bar"), "boo"},
			exp: "foo boo",
		},
		{
			str: "foo åäö",
			sw:  SubQuery{Word("foo"), "bar"},
			exp: "bar åäö",
		},
		{
			str: "foo åäö",
			sw:  SubQuery{Word("åäö"), "bar"},
			exp: "foo bar",
		},
		{
			str: "foo barbaz",
			sw:  SubQuery{Word("foo"), "long replacement"},
			exp: "long replacement barbaz",
		},
	}

	for _, c := range cases {
		act := c.sw.Apply(c.str)

		if act != c.exp {
			t.Errorf("Expected SubQuery{%#v, %q}.Apply(%q) to return %q, got %q", c.sw.Search, c.sw.Sub, c.str, c.exp, act)
		}
	}
}
