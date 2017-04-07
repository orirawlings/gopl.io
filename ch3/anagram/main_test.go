package main

import (
	"testing"
)

type pair struct{ a, b string }

func TestAnagram(t *testing.T) {
	data := map[pair]bool{
		pair{"こんにちは", "hello"}:               false,
		pair{"こんにちは, hello", "hello, こんにちは"}: true,
		pair{"hello, こんにちは", "こんにちは, hello"}: true,
		pair{"golang", "lognag"}:             true,
		pair{"anagram", "nagaram"}:           true,
		pair{"hello", "hello"}:               false,
	}
	for args, expected := range data {
		if v := anagram(args.a, args.b); v != expected {
			t.Errorf("anagram(%v, %v) is %v. Expected %v.", args.a, args.b, v, expected)
		}
	}
}
