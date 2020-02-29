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

// Index represents the 2-byte index in an object index.
type Index struct {
	B0 byte
	B1 byte
}

// ObjectIndex represents the index of an object.
type ObjectIndex struct {
	Index    Index
	SubIndex byte
}

// NewObjectIndex returns an object index from a 2-byte index and 1-byte sub index.
func NewObjectIndex(index uint16, subIndex uint8) ObjectIndex {
	return ObjectIndex{
		Index: Index{
			B0: byte(index & 0xFF),
			B1: byte(index >> 8),
		},
		SubIndex: subIndex,
	}
}
