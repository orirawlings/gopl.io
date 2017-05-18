package reverse

import (
	"bytes"
	"testing"
)

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"Hello, 世界", "界世 ,olleH"},
		{"界世 ,olleH", "Hello, 世界"},
		{"", ""},
	}
	for _, c := range cases {
		i, o := []byte(c.in), []byte(c.out)
		ReverseRunes(i)
		if bytes.Compare(i, o) != 0 {
			t.Errorf("%q != %q", string(i), string(o))
		}
	}
}
