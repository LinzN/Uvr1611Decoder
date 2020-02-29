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

// splitN splits b into a list of n sized bytes
func splitN(b []byte, n int) [][]byte {
	if len(b) < n {
		return [][]byte{b}
	}

	var bs [][]byte
	var buf []byte
	for i := 0; i < len(b); i++ {
		if len(buf) == n {
			bs = append(bs, buf)
			buf = []byte{}
		}

		buf = append(buf, b[i])
	}

	if len(buf) > 0 {
		bs = append(bs, buf)
	}

	return bs
}
