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
	"fmt"
	"github.com/brutella/canopen"
	"strings"
)

type Outlet struct {
	Description canopen.ObjectIndex
	StartDelay  canopen.ObjectIndex
	RunOnTime   canopen.ObjectIndex
	Mode        canopen.ObjectIndex
	State       canopen.ObjectIndex
	SpeedStage  canopen.ObjectIndex
}

func NewOutlet(subIndex uint8) Outlet {
	return Outlet{
		Description: canopen.NewObjectIndex(0x20a5, subIndex),
		StartDelay:  canopen.NewObjectIndex(0x20a3, subIndex),
		RunOnTime:   canopen.NewObjectIndex(0x20a4, subIndex),
		Mode:        canopen.NewObjectIndex(0x20a1, subIndex),
		State:       canopen.NewObjectIndex(0x20aa, subIndex),
		SpeedStage:  canopen.NewObjectIndex(0x20ab, subIndex),
	}
}

func StringToBool(str string) (bool, error) {
	switch strings.TrimSpace(str) {
	case OutletStateOn:
		return true, nil
	case OutletStateOff:
		return false, nil
	default:
		break
	}

	return false, fmt.Errorf("Unknown string value %v (%X)", str, str)
}
