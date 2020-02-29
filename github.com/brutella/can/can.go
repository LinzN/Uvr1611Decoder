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

// Package can provides an implemention of a CAN bus to send and receive CAN frames.
package can

const (
	// MaskIDSff is used to extract the valid 11-bit CAN identifier bits from the frame ID of a standard frame format.
	MaskIDSff = 0x000007FF
	// MaskIDEff is used to extract the valid 29-bit CAN identifier bits from the frame ID of an extended frame format.
	MaskIDEff = 0x1FFFFFFF
	// MaskErr is used to extract the the error flag (0 = data frame, 1 = error message) from the frame ID.
	MaskErr = 0x20000000
	// MaskRtr is used to extract the rtr flag (1 = rtr frame) from the frame ID
	MaskRtr = 0x40000000
	// MaskEff is used to extract the eff flag (0 = standard frame, 1 = extended frame) from the frame ID
	MaskEff = 0x80000000
)

const (
	// MaxFrameDataLength defines the max length of a CAN data frame defined in ISO 11898-1.
	MaxFrameDataLength = 8
	// MaxExtFrameDataLength defines the max length of an CAN extended data frame defined in ISO ISO 11898-7.
	MaxExtFrameDataLength = 64
)
