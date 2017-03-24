package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkDrawMandelbrotFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotFloat64(ioutil.Discard)
	}
}

func BenchmarkDrawMandelbrotComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotComplex64(ioutil.Discard)
	}
}

func BenchmarkDrawMandelbrotComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotComplex128(ioutil.Discard)
	}
}

func BenchmarkDrawMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotBigFloat(ioutil.Discard)
	}
}

// TODO: big.Rat based implementation is too slow to complete
/*
func BenchmarkDrawMandelbrotBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotBigRat(ioutil.Discard)
	}
}
*/
