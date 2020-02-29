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

import (
	"github.com/brutella/can"
	"time"
)

// A Client handles message communication by sending a request
// and waiting for the response.
type Client struct {
	Bus     *can.Bus
	Timeout time.Duration
}

// Do sends a request and waits for a response.
// If the response frame doesn't arrive on time, an error is returned.
func (c *Client) Do(req *Request) (*Response, error) {
	if err := c.Bus.Publish(req.Frame.CANFrame()); err != nil {
		return nil, err
	}

	resp := <-can.Wait(c.Bus, req.ResponseID, c.Timeout)

	return &Response{CANopenFrame(resp.Frame), req}, resp.Err
}
