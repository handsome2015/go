// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is tested when running "go test -run Manual"
// without source arguments. Use for one-off debugging.

package p

type T1 interface{ M() }

func F1() T1

var _ = F1().(*X1 /* ERROR undeclared name: X1 */)

func _() {
	switch F1().(type) {
	case *X1 /* ERROR undeclared name: X1 */ :
	}
}

type T2 interface{ M() }

func F2() T2

var _ = F2(). /* ERROR impossible type assertion: F2\(\).\(\*X2\)\n\t\*X2 does not implement T2 \(missing M method\) */ (*X2)

type X2 struct{}

func _() {
	switch F2().(type) {
	case * /* ERROR impossible type switch case: \*X2\n\tF2\(\) \(value of type T2\) cannot have dynamic type \*X2 \(missing M method\) */ X2:
	}
}
