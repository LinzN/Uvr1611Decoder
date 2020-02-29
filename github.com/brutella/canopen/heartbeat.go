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

// NewHeartbeatFrame returns a new CANopen heartbeat frame containing the node id and state.
func NewHeartbeatFrame(id uint8, state uint8) Frame {
	return Frame{
		CobID: uint16(MessageTypeHeartbeat) + uint16(id),
		Data:  []byte{byte(state)},
	}
}
