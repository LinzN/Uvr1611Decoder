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

// A Request represents a CANopen request published on a CAN bus and received by another CANopen node.
type Request struct {
	// The Frame of the request
	Frame Frame

	// The ResponseID of the response frame
	ResponseID uint32
}

// NewRequest returns a request containing the frame to be sent
// and the expected response frame id.
func NewRequest(frm Frame, respID uint32) *Request {
	return &Request{
		Frame:      frm,
		ResponseID: respID,
	}
}
