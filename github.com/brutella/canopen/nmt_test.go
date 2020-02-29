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

package canopen

import (
	"reflect"
	"testing"
)

func TestHeartbeatFrame(t *testing.T) {
	frm := NewHeartbeatFrame(0x1, Operational)
	b, err := Marshal(frm)

	if err != nil {
		t.Fatal(err)
	}

	if is, want := b, []byte{0x01, 0x07, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("\n% X !=\n% X\n", is, want)
	}
}

func TestHeartbeat(t *testing.T) {
	frm := NewHeartbeatFrame(0x1, Operational)

	if is, want := frm.MessageType(), MessageTypeHeartbeat; is != want {
		t.Fatalf("is=%X, want=%X", is, want)
	}
}
