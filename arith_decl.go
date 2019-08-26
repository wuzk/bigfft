// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bigfft

import . "math/big"

// implemented in arith_$GOARCH.s
//func addVV(z, x, y []Word) (c Word)
//func subVV(z, x, y []Word) (c Word)
//func addVW(z, x []Word, y Word) (c Word)
//func subVW(z, x []Word, y Word) (c Word)
//func shlVU(z, x []Word, s uint) (c Word)
//func mulAddVWW(z, x []Word, y, r Word) (c Word)
//func addMulVVW(z, x []Word, y Word) (c Word)
// Trampolines to math/big assembly implementations.

#include "textflag.h"

// func addVV(z, x, y []Word) (c Word)
TEXT ·addVV(SB),NOSPLIT,$0
	B	math∕big·addVV(SB)

// func subVV(z, x, y []Word) (c Word)
TEXT ·subVV(SB),NOSPLIT,$0
	B	math∕big·subVV(SB)

// func addVW(z, x []Word, y Word) (c Word)
TEXT ·addVW(SB),NOSPLIT,$0
	B	math∕big·addVW(SB)

// func subVW(z, x []Word, y Word) (c Word)
TEXT ·subVW(SB),NOSPLIT,$0
	B	math∕big·subVW(SB)

// func shlVU(z, x []Word, s uint) (c Word)
TEXT ·shlVU(SB),NOSPLIT,$0
	B	math∕big·shlVU(SB)

// func shrVU(z, x []Word, s uint) (c Word)
TEXT ·shrVU(SB),NOSPLIT,$0
	B	math∕big·shrVU(SB)

// func mulAddVWW(z, x []Word, y, r Word) (c Word)
TEXT ·mulAddVWW(SB),NOSPLIT,$0
	B	math∕big·mulAddVWW(SB)

// func addMulVVW(z, x []Word, y Word) (c Word)
TEXT ·addMulVVW(SB),NOSPLIT,$0
	B	math∕big·addMulVVW(SB)
