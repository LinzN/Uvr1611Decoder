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

const (
	TransferExpedited            = 0x2
	TransferSizeIndicated        = 0x1
	TransferMaskCommandSpecifier = 0xE0
	TransferMaskSize             = 0xC
	TransferSegmentToggle        = 0x10
	TransferAbort                = 0x80
)
