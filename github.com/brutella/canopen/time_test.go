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
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 500 ms timestamp message
	b := []byte{
		0x00, 0x1, 0x00, 0x00,
		0x08,
		0x00,
		0x00,
		0x00,
		0xF4, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	frm := Frame{}

	Unmarshal(b, &frm)

	if is, want := frm.MessageType(), MessageTypeTimestamp; is != want {
		t.Fatalf("is=%X, want=%X", is, want)
	}

	if tm, err := frm.Timestamp(); err != nil {
		t.Fatal(err)
	} else if is, want := tm, time.Unix(RefDate.Unix(), 500*1000); is.Equal(want) == false {
		t.Fatalf("is=%v, want=%v", is, want)
	}
}
