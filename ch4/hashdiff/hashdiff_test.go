package hashdiff

import (
	"encoding/hex"
	"testing"
)

func decode(t testing.TB, s string) [32]byte {
	d, err := hex.DecodeString(s)
	if err != nil {
		t.Fatalf("Error decoding hash [%s]. %s", s, err)
	}
	if len(d) != 32 {
		t.Fatalf("Expected hash of 32 bytes, was %d for [%s]", len(d), s)
	}
	var r [32]byte
	copy(r[:], d)
	return r
}

func TestBits(t *testing.T) {
	data := map[[2]string]int{
		[2]string{"971916a768aeb90d3a0c0b9c92146b83a9bf13fc6a84b0b46bd79c677c7968a0", "971916a768aeb90d3a0c0b9c92146b83a9bf13fc6a84b0b46bd79c677c7968a0"}: 0,
		[2]string{"971916a768aeb90d3a0c0b9c92146b83a9bf13fc6a84b0b46bd79c677c7968a0", "971916a768aeb90d3a0c0b9c92146b83a9bf13fc6a84b0b46bd79c677c7968a1"}: 1,
		[2]string{"971916a768aeb90d3a0c0b9c92146b83a9bf13fc6a84b0b46bd79c677c7968a0", "c6b305ddb211cbbbb8d31626b8cd8013ef67d9ec458d037196a950b0f80f01ed"}: 136,
	}
	for p, e := range data {
		if a := Bits(decode(t, p[0]), decode(t, p[1])); a != e {
			t.Errorf("Expected %d bits of diff between hashes %v, was %d.", e, p, a)
		}
	}
}

func BenchmarkBits(b *testing.B) {
	h0 := decode(b, "971916a768aeb90d3a0c0b9c92146b83a9bf13fc6a84b0b46bd79c677c7968a0")
	h1 := decode(b, "c6b305ddb211cbbbb8d31626b8cd8013ef67d9ec458d037196a950b0f80f01ed")
	for i := 0; i < b.N; i++ {
		Bits(h0, h1)
	}
}
