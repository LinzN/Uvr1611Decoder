/*
 * Copyright (C) 2020. Niklas Linz - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the LGPLv3 license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the LGPLv3 license with
 * this file. If not, please write to: niklas.linz@enigmar.de
 *
 */

package sdo

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	var b = []byte{1, 2, 3, 4, 5}

	splitted := splitN(b, 2)

	if x := len(splitted); x != 3 {
		t.Fatal(x)
	}

	if is, want := splitted[0], []byte{1, 2}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := splitted[1], []byte{3, 4}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := splitted[2], []byte{5}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
