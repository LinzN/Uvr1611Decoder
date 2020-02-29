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

package main

import (
	"flag"
	"fmt"
	"github.com/brutella/can"
	"github.com/brutella/uvr"
	"log"
)

func readOutlet(outlet uvr.Outlet, client *uvr.Client) (descr string, mode string, val string) {
	if value, err := client.Read(outlet.Description); err == nil {
		descr = value.(string)
	}

	if value, err := client.Read(outlet.Mode); err == nil {
		mode = value.(string)
	}

	if value, err := client.Read(outlet.State); err == nil {
		val = value.(string)
	}

	return
}

func readInlet(inlet uvr.Inlet, client *uvr.Client) (descr string, state string, val float32) {
	if value, err := client.Read(inlet.Description); err == nil {
		descr = value.(string)
	}

	if value, err := client.Read(inlet.State); err == nil {
		state = value.(string)
	}

	if value, err := client.Read(inlet.Value); err == nil {
		if float, ok := value.(float32); ok == true {
			val = float
		}
	}

	return
}

func readNotify(notify uvr.Notify, client *uvr.Client) (Notify string, Status string) {
	if value, err := client.Read(notify.Notify); err == nil {
		Notify = value.(string)
	}
	if value, err := client.Read(notify.Status); err == nil {
		Status = value.(string)
	}
	return
}

func readOutlets(client *uvr.Client) {
	outlets := []uvr.Outlet{
		uvr.NewOutlet(0x1),
		uvr.NewOutlet(0x2),
		uvr.NewOutlet(0x3),
		uvr.NewOutlet(0x4),
		uvr.NewOutlet(0x5),
		uvr.NewOutlet(0x6),
		uvr.NewOutlet(0x7),
		uvr.NewOutlet(0x8),
		uvr.NewOutlet(0x9),
		uvr.NewOutlet(0xa),
		uvr.NewOutlet(0xb),
		uvr.NewOutlet(0xc),
		uvr.NewOutlet(0xd),
	}

	fmt.Printf("\"outlets\": [")
	for index, outlet := range outlets {
		descr, mode, val := readOutlet(outlet, client)
		fmt.Printf("{\"index\":%d, \"description\":\"%s\", \"mode\":\"%s\", \"value\":\"%s\"}", index+1, descr, mode, val)
		if index < 12 {
			fmt.Printf(",\n")
		} else {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("]")
}

func readInlets(client *uvr.Client) {
	inlets := []uvr.Inlet{
		uvr.NewInlet(0x1),
		uvr.NewInlet(0x2),
		uvr.NewInlet(0x3),
		uvr.NewInlet(0x4),
		uvr.NewInlet(0x5),
		uvr.NewInlet(0x6),
		uvr.NewInlet(0x7),
		uvr.NewInlet(0x8),
		uvr.NewInlet(0x9),
		uvr.NewInlet(0xa),
		uvr.NewInlet(0xb),
		uvr.NewInlet(0xc),
		uvr.NewInlet(0xd),
		uvr.NewInlet(0xe),
		uvr.NewInlet(0xf),
		uvr.NewInlet(0x10),
	}
	fmt.Printf("\"inlets\": [")
	for index, inlet := range inlets {
		descr, state, val := readInlet(inlet, client)
		fmt.Printf("{\"index\":%d, \"description\":\"%s\", \"state\":\"%s\", \"value\":%f}", index+1, descr, state, val)
		if index < 15 {
			fmt.Printf(",\n")
		} else {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("]")
}

func readNotifies(client *uvr.Client) {
	notifies := []uvr.Notify{
		uvr.NewNotify(0x1),
		uvr.NewNotify(0x2),
		uvr.NewNotify(0x3),
		uvr.NewNotify(0x4),
		uvr.NewNotify(0x5),
	}
	fmt.Printf("\"notifies\": [")
	for index, notify := range notifies {
		Notify, Status := readNotify(notify, client)
		fmt.Printf("{\"index\":%d, \"notify\":\"%s\", \"state\":\"%s\"}", index+1, Notify, Status)
		if index < 4 {
			fmt.Printf(",\n")
		} else {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("]")
}

func HandleCANopen(frame can.Frame) {
	log.Printf("%X % X\n", frame.ID, frame.Data)
}

func main() {
	var (
		clientId = flag.Int("client_id", 16, "id of the client; range from [1...254]")
		serverId = flag.Int("server_id", 1, "id of the server to which the client connects to: range from [1...254]")
		iface    = flag.String("if", "can0", "name of the can network interface")
	)

	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	bus, err := can.NewBusForInterfaceWithName(*iface)

	if err != nil {
		log.Fatal(err)
	}
	// bus.SubscribeFunc(HandleCANopen)
	go bus.ConnectAndPublish()

	nodeID := uint8(*clientId)
	uvrID := uint8(*serverId)

	c := uvr.NewClient(nodeID, bus)
	c.Connect(uvrID)
	fmt.Printf("{")
	readInlets(c)
	fmt.Printf(",\n")
	readOutlets(c)
	fmt.Printf(",\n")
	readNotifies(c)
	fmt.Printf("}\n")
	c.Disconnect(uvrID)
	bus.Disconnect()
}
