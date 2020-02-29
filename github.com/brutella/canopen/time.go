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
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

// RefDate is the reference date of CAN timestamp messages.
var RefDate = time.Date(
	1984,     // year
	1,        // month
	1,        // day
	0,        // hour
	0,        // min
	0,        // sec
	0,        // nsec
	time.UTC, // location
)

// Timestamp returns the time encoded in the frame.
func (frm Frame) Timestamp() (*time.Time, error) {
	if t := frm.MessageType(); t != MessageTypeTimestamp {
		return nil, fmt.Errorf("Invalid message type % X", t)
	}

	if n := len(frm.Data); n != 8 {
		return nil, fmt.Errorf("Invalid data length %d", n)
	}

	var msec int32
	var day int16

	msecBuf := bytes.NewBuffer(frm.Data[:6])
	dayBuf := bytes.NewBuffer(frm.Data[6:8])

	binary.Read(msecBuf, binary.LittleEndian, &msec)
	binary.Read(dayBuf, binary.LittleEndian, &dayBuf)

	daySec := int64(day) * 24 * 60 * 60
	t := time.Unix(RefDate.Unix()+daySec, int64(msec)*1000)

	return &t, nil
}
