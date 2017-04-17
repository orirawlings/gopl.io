package main

import "testing"

func TestTrimspace(t *testing.T) {
	data := [][2]string{
		[2]string{"Hello,世界", "Hello,世界"},
		[2]string{"Hello, 世界", "Hello, 世界"},
		[2]string{"Hello,  世界", "Hello, 世界"},
		[2]string{"Hello,   世界", "Hello, 世界"},
		[2]string{" Hello,世界", " Hello,世界"},
		[2]string{"  Hello,世界", " Hello,世界"},
		[2]string{"Hello,世界 ", "Hello,世界 "},
		[2]string{"Hello,世界  ", "Hello,世界 "},
		[2]string{"Hello,\n世界", "Hello,\n世界"},
		[2]string{"Hello,\n\n世界", "Hello, 世界"},
		[2]string{"\nHello,世界", "\nHello,世界"},
		[2]string{"\n\nHello,世界", " Hello,世界"},
		[2]string{"Hello,世界\n", "Hello,世界\n"},
		[2]string{"Hello,世界\n\n", "Hello,世界 "},
		[2]string{"Hello,\t世界", "Hello,\t世界"},
		[2]string{"Hello,\t\t世界", "Hello, 世界"},
		[2]string{"\tHello,世界", "\tHello,世界"},
		[2]string{"\t\tHello,世界", " Hello,世界"},
		[2]string{"Hello,世界\t", "Hello,世界\t"},
		[2]string{"Hello,世界\t\t", "Hello,世界 "},
		[2]string{"Hello,\v世界", "Hello,\v世界"},
		[2]string{"Hello,\v\v世界", "Hello, 世界"},
		[2]string{"\vHello,世界", "\vHello,世界"},
		[2]string{"\v\vHello,世界", " Hello,世界"},
		[2]string{"Hello,世界\v", "Hello,世界\v"},
		[2]string{"Hello,世界\v\v", "Hello,世界 "},
		[2]string{"Hello,\f世界", "Hello,\f世界"},
		[2]string{"Hello,\f\f世界", "Hello, 世界"},
		[2]string{"\fHello,世界", "\fHello,世界"},
		[2]string{"\f\fHello,世界", " Hello,世界"},
		[2]string{"Hello,世界\f", "Hello,世界\f"},
		[2]string{"Hello,世界\f\f", "Hello,世界 "},
		[2]string{"Hello,\r世界", "Hello,\r世界"},
		[2]string{"Hello,\r\r世界", "Hello, 世界"},
		[2]string{"\rHello,世界", "\rHello,世界"},
		[2]string{"\r\rHello,世界", " Hello,世界"},
		[2]string{"Hello,世界\r", "Hello,世界\r"},
		[2]string{"Hello,世界\r\r", "Hello,世界 "},
		[2]string{"Hello,\u0085世界", "Hello,\u0085世界"},
		[2]string{"Hello,\u0085\u0085世界", "Hello, 世界"},
		[2]string{"\u0085Hello,世界", "\u0085Hello,世界"},
		[2]string{"\u0085\u0085Hello,世界", " Hello,世界"},
		[2]string{"Hello,世界\u0085", "Hello,世界\u0085"},
		[2]string{"Hello,世界\u0085\u0085", "Hello,世界 "},
		[2]string{"Hello,\u00a0世界", "Hello,\u00a0世界"},
		[2]string{"Hello,\u00a0\u00a0世界", "Hello, 世界"},
		[2]string{"\u00a0Hello,世界", "\u00a0Hello,世界"},
		[2]string{"\u00a0\u00a0Hello,世界", " Hello,世界"},
		[2]string{"Hello,世界\u00a0", "Hello,世界\u00a0"},
		[2]string{"Hello,世界\u00a0\u00a0", "Hello,世界 "},
	}
	for _, d := range data {
		actual := string(trimspace([]byte(d[0])))
		if actual != d[1] {
			t.Errorf("%q is not expected %q for string %q", actual, d[1], d[0])
		}
	}
}
