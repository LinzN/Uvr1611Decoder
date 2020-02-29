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

// The Handler interfaces defines a method to receive a frame.
type Handler interface {
	Handle(frame Frame)
}

// HandlerFunc defines the function type to handle a frame.
type HandlerFunc func(frame Frame)

type handler struct {
	fn HandlerFunc
}

// NewHandler returns a new handler which calls fn when a frame is received.
func NewHandler(fn HandlerFunc) Handler {
	return &handler{fn}
}

func (h *handler) Handle(frame Frame) {
	h.fn(frame)
}
