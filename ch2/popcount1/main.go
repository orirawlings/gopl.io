package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

func PopCountShift(x uint64) int {
	return int(x&1 +
		x>>1&1 +
		x>>2&1 +
		x>>3&1 +
		x>>4&1 +
		x>>5&1 +
		x>>6&1 +
		x>>7&1 +
		x>>8&1 +
		x>>9&1 +
		x>>10&1 +
		x>>11&1 +
		x>>12&1 +
		x>>13&1 +
		x>>14&1 +
		x>>15&1 +
		x>>16&1 +
		x>>17&1 +
		x>>18&1 +
		x>>19&1 +
		x>>20&1 +
		x>>21&1 +
		x>>22&1 +
		x>>23&1 +
		x>>24&1 +
		x>>25&1 +
		x>>26&1 +
		x>>27&1 +
		x>>28&1 +
		x>>29&1 +
		x>>30&1 +
		x>>31&1 +
		x>>32&1 +
		x>>33&1 +
		x>>34&1 +
		x>>35&1 +
		x>>36&1 +
		x>>37&1 +
		x>>38&1 +
		x>>39&1 +
		x>>40&1 +
		x>>41&1 +
		x>>42&1 +
		x>>43&1 +
		x>>44&1 +
		x>>45&1 +
		x>>46&1 +
		x>>47&1 +
		x>>48&1 +
		x>>49&1 +
		x>>50&1 +
		x>>51&1 +
		x>>52&1 +
		x>>53&1 +
		x>>54&1 +
		x>>55&1 +
		x>>56&1 +
		x>>57&1 +
		x>>58&1 +
		x>>59&1 +
		x>>60&1 +
		x>>61&1 +
		x>>62&1 +
		x>>63&1)
}

func PopCountClear(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}
