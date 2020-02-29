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
	"github.com/brutella/canopen"
)

type Inlet struct {
	Description canopen.ObjectIndex
	Value       canopen.ObjectIndex
	State       canopen.ObjectIndex
}

func NewInlet(subIndex uint8) Inlet {
	return Inlet{
		Description: canopen.NewObjectIndex(0x2084, subIndex),
		Value:       canopen.NewObjectIndex(0x208d, subIndex),
		State:       canopen.NewObjectIndex(0x208e, subIndex),
	}
}
