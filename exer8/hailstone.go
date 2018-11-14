package exer8

// ---------------- Benchmarking ----------------
// BenchmarkHailSeqAppend-8         1000000              1906 ns/op
// BenchmarkHailSeqAllocate-8       2000000               814 ns/op
// HailSeqAllocate runs the set of four calls in 814ns, which is about 1000ns faster than
// HailSeqAppend. HailSeqAllocate is run 2000000 and HailSeqAppend is run 1000000 times.
// ---------------- Benchmarking ----------------

// ---------------- This is for testing ----------------
// import (
// 	"fmt"
// )

// func main() {
// 	n := uint(1)
// 	fmt.Println("hail sequence appended: ", HailstoneSequenceAppend(n))
// 	fmt.Println("hail sequence allocated: ", HailstoneSequenceAllocate(n))
// }
// ---------------- This is for testing ----------------

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions

// Hailstone function returns the next value in the hailstone sequence
func Hailstone(n uint) (x uint) {
	if n%2 == 0 {
		x = n / 2
	} else {
		x = n*3 + 1
	}
	return
}

// HailstoneSequenceAppend generates the hail sequence by appending values
func HailstoneSequenceAppend(n uint) (hailSeq []uint) {
	for ; n != 1; n = Hailstone(n) {
		hailSeq = append(hailSeq, n)
	}
	hailSeq = append(hailSeq, n)
	return
}

// HailstoneSequenceAllocate generates the hail sequence by allocating space
func HailstoneSequenceAllocate(n uint) []uint {
	arrLen := 0
	tempN := n
	for ; tempN != 1; tempN = Hailstone(tempN) {
		arrLen++
	}

	hailSeq := make([]uint, arrLen+1)

	for i := 0; i <= arrLen; i++ {
		hailSeq[i] = n
		n = Hailstone(n)
	}

	return hailSeq
}
