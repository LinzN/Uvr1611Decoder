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

package uvr

import (
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"github.com/brutella/canopen/sdo"
)

func WriteToIndex(idx canopen.ObjectIndex, b []byte, nodeID uint8, bus *can.Bus) error {
	download := sdo.Download{
		ObjectIndex:   idx,
		RequestCobID:  uint16(SSDOClientToServer2) + uint16(nodeID),
		ResponseCobID: uint16(SSDOServerToClient2) + uint16(nodeID),
		Data:          b,
	}

	return download.Do(bus)
}
