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

package can

import (
	"log"
	"reflect"
	"testing"
)

func TestFrame(t *testing.T) {
	f := Frame{
		ID:     0x701,
		Length: 1,
		Flags:  0,
		Res0:   0,
		Res1:   0,
		Data:   [8]uint8{0x05},
	}
	b, err := Marshal(f)

	if err != nil {
		t.Fatal(err)
	}

	if is, want := len(b), 16; is != want {
		log.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := b, []uint8{0x7, 0x1, 0x1, 0x0, 0x0, 0x0, 0x05, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}; reflect.DeepEqual(is, want) {
		log.Fatalf("is=%v want=%v", is, want)
	}

	if err := Unmarshal(b, &f); err != nil {
		t.Fatal(err)
	}

	if is, want := f.ID, uint32(0x701); is != want {
		log.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := f.Data, []uint8{0x05, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}; reflect.DeepEqual(is, want) {
		log.Fatalf("is=%v want=%v", is, want)
	}
}
