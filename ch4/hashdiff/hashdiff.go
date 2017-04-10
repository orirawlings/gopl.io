package hashdiff

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Return the number of bits that differ between 256-bit hashes
func Bits(a, b [32]byte) int {
	var sum int
	for i := 0; i < 32; i++ {
		sum += int(pc[a[i]^b[i]])
	}
	return sum
}
