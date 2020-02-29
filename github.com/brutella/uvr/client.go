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
	"time"
)

type Client struct {
	id        uint8
	bus       *can.Bus
	heartbeat chan<- struct{}
}

func NewClient(id uint8, bus *can.Bus) *Client {
	return &Client{id, bus, nil}
}

func (c *Client) Connect(id uint8) error {
	c.heartbeat = canopen.ProduceHeartbeat(c.id, canopen.Operational, c.bus, time.Second*10)
	return Connect(id, c.id, c.bus)
}

func (c *Client) Disconnect(id uint8) error {
	c.heartbeat <- struct{}{}
	return Disconnect(id, c.id, c.bus)
}

func (c *Client) Read(i canopen.ObjectIndex) (interface{}, error) {
	return ReadFromIndex(i, c.id, c.bus)
}

func (c *Client) Write(b []byte, i canopen.ObjectIndex) error {
	return WriteToIndex(i, b, c.id, c.bus)
}
