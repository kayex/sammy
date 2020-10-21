package sammy

import (
	"testing"
)

func TestExtendMajor(t *testing.T) {
	cases := []struct {
		name string
		exp  string
	}{
		{"K4YEX_Dreamy_Synth_A_120.wav", "K4YEX_Dreamy_Synth_Amaj_120.wav"},
		{"K4YEX_Dreamy_Synth_B_120.wav", "K4YEX_Dreamy_Synth_Bmaj_120.wav"},
		{"K4YEX_Dreamy_Synth_C_120.wav", "K4YEX_Dreamy_Synth_Cmaj_120.wav"},
		{"K4YEX_Dreamy_Synth_A#_120.wav", "K4YEX_Dreamy_Synth_A#maj_120.wav"},
		{"K4YEX_Dreamy_Synth_Amaj_120.wav", "K4YEX_Dreamy_Synth_Amaj_120.wav"},
		{"Z:\\Sound Library\\User Library\\Splice Sounds - Shigeto繁登Custom_Sounds繁登Found_Sounds繁登_SpringinProvence.wav", "Z:\\Sound Library\\User Library\\Splice Sounds - Shigeto繁登Custom_Sounds繁登Found_Sounds繁登_SpringinProvence.wav"},
		{"繁", "繁"},
	}

	for _, c := range cases {
		actual := ExtendMajor(c.name)

		if actual != c.exp {
			t.Errorf("Expected ExtendMajor(%v) to be %v, was %v", c.name, c.exp, actual)
		}
	}
}

func TestExtendMinor(t *testing.T) {
	cases := []struct {
		name string
		exp  string
	}{
		{"K4YEX_Dreamy_Synth_Am_120.wav", "K4YEX_Dreamy_Synth_Amin_120.wav"},
		{"K4YEX_Dreamy_Synth_Bm_120.wav", "K4YEX_Dreamy_Synth_Bmin_120.wav"},
		{"K4YEX_Dreamy_Synth_Cm_120.wav", "K4YEX_Dreamy_Synth_Cmin_120.wav"},
		{"K4YEX_Dreamy_Synth_A#m_120.wav", "K4YEX_Dreamy_Synth_A#min_120.wav"},
		{"K4YEX_Dreamy_Synth_Amin_120.wav", "K4YEX_Dreamy_Synth_Amin_120.wav"},
		{"C:\\Users\\jv\\Development\\sammy\\samples\\bass\\test_Am_2.wav", "C:\\Users\\jv\\Development\\sammy\\samples\\bass\\test_Amin_2.wav"},
	}

	for _, c := range cases {
		actual := ExtendMinor(c.name)

		if actual != c.exp {
			t.Errorf("Expected ExtendMinor(%v) to be %v, was %v", c.name, c.exp, actual)
		}
	}
}
