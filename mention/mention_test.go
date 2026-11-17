package mention

import "testing"

func TestUserID(t *testing.T) {
	cases := map[string]string{
		"<@123456789>":   "123456789",
		"<@!123456789>":  "123456789",
		"hi <@42> there": "42",
		"plainname":      "",
		"<@&987>":        "",
		"":               "",
	}

	for in, want := range cases {
		if got := UserID(in); got != want {
			t.Errorf("UserID(%q) = %q, want %q", in, got, want)
		}
	}
}
