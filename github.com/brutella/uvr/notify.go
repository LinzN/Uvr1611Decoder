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

type Notify struct {
	Notify canopen.ObjectIndex
	Status canopen.ObjectIndex
}

func NewNotify(subIndex uint8) Notify {
	return Notify{
		Notify: canopen.NewObjectIndex(0x2418, subIndex),
		Status: canopen.NewObjectIndex(0x2426, subIndex),
	}
}
