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

type filter struct {
	id      uint32
	handler Handler
}

func newFilter(id uint32, handler Handler) Handler {
	return &filter{
		id:      id,
		handler: handler,
	}
}

func (f *filter) Handle(frame Frame) {
	if frame.ID == f.id {
		f.handler.Handle(frame)
	}
}
