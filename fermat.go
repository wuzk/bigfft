package bigfft

import (
	"math/big"
)

// Arithmetic modulo 2^n+1.

// A fermat of length w+1 represents a number modulo 2^(w*_W) + 1. The last
// word is zero or one. A number has at most two representatives satisfying the
// 0-1 last word constraint.
type fermat nat

func (n fermat) String() string { return nat(n).String() }

func (z fermat) norm() {
	n := len(z) - 1
	c := z[n]
	if c == 0 {
		return
	}
	if z[0] >= c {
		z[n] = 0
		z[0] -= c
		return
	}
	// z[0] < z[n].
	subVW(z, z, c) // Substract c
	if c > 1 {
		z[n] -= c - 1
		c = 1
	}
	// Add back c.
	if z[n] == 1 {
		z[n] = 0
		return
	} else {
		addVW(z, z, 1)
	}
}

// Shift computes (x << k) mod (2^n+1).
func (z fermat) Shift(x fermat, k int) {
	if len(z) != len(x) {
		println(len(z), len(x))
		panic("len(z) != len(x) in Shift")
	}
	n := len(x) - 1
	// Shift by n*_W is Neg.
	k %= 2 * n * _W
	if k < 0 {
		k += 2 * n * _W
	}
	neg := false
	if k >= n*_W {
		k -= n * _W
		neg = true
	}
	kw, kb := k/_W, k%_W
	// Shift left by kw words.
	// x = a·2^(n-k) + b
	// x<<k = (b<<k) - a
	copy(z[kw:], x[:n-kw])
	for i := 0; i < kw; i++ {
		z[i] = 0
	}
	z[n] = 1 // Add (-1)
	b := subVV(z[:kw+1], z[:kw+1], x[n-kw:])
	if z[kw+1] > 0 {
		z[kw+1] -= b
	} else {
		subVW(z[kw+1:], z[kw+1:], b)
	}
	if z[0] < ^big.Word(0) {
		z[0]++
	} else {
		addVW(z, z, 1) // Add back 1.
	}
	// Shift left by kb bits
	shlVU(z, z, uint(kb))
	if neg {
		z.Neg()
	}
	z.norm()
}

// Neg computes the opposite of x mod 2^n+1.
func (x fermat) Neg() {
	n := len(x) - 1
	c := x[n]
	// 2^n + 1 - x = ^x + 2.
	for i, a := range x {
		x[i] = ^a
	}
	x[n] = 0
	if x[0] <= ^big.Word(0)-(c+2) {
		x[0] += c + 2
	} else {
		x[n] = addVW(x[:n], x[:n], c+2)
	}
	x.norm()
	return
}

// Add computes addition mod 2^n+1.
func (z fermat) Add(x, y fermat) fermat {
	if len(z) != len(x) {
		panic("Add: len(z) != len(x)")
	}
	addVV(z, x, y) // there cannot be a carry here.
	z.norm()
	return z
}

func (z fermat) Mul(x, y fermat) fermat {
	var xi, yi, zi big.Int
	xi.SetBits(x)
	yi.SetBits(y)
	zi.SetBits(z)
	z = zi.Mul(&xi, &yi).Bits()
	n := len(x) - 1
	// len(z) is at most 2n+1.
	if len(z) > 2*n+1 {
		panic("len(z) > 2n+1")
	}
	i := len(z) - (n + 1) // i <= n
	c := subVV(z[1:i+1], z[1:i+1], z[n+1:])
	z = z[:n+1]
	z[n]++ // Add -1.
	subVW(z[i+1:], z[i+1:], c)
	// Add 1.
	if z[n] == 1 {
		z[n] = 0
	} else {
		addVW(z, z, 1)
	}
	z.norm()
	return z
}
