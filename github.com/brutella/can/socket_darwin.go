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

import "syscall"

// NewSockaddr returns a socket address based on the protocol and interface index.
// TODO(brutella) This method has no implementation.
func NewSockaddr(proto uint16, Ifindex int) syscall.Sockaddr {
	return &syscall.SockaddrUnix{}
}
