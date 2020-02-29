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

const (
	GoToOperational        uint8 = 0x1
	GoToStopped            uint8 = 0x2
	GoToPreOperation       uint8 = 0x80
	GoToResetNode          uint8 = 0x81
	GoToResetCommunication uint8 = 0x82
)

const (
	BootUp         uint8 = 0x0
	Stopped        uint8 = 0x04
	Operational    uint8 = 0x05
	PreOperational uint8 = 0x7f
)
