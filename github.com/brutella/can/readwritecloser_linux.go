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

package can

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"syscall"
)

func NewReadWriteCloserForInterface(i *net.Interface) (ReadWriteCloser, error) {
	s, _ := syscall.Socket(syscall.AF_CAN, syscall.SOCK_RAW, unix.CAN_RAW)
	addr := &unix.SockaddrCAN{Ifindex: i.Index}
	if err := unix.Bind(s, addr); err != nil {
		return nil, err
	}

	f := os.NewFile(uintptr(s), fmt.Sprintf("fd %d", s))

	return &readWriteCloser{f}, nil
}
