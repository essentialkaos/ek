// Package events provides methods and structs for creating event-driven systems
package events

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"reflect"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dispatcher is event dispatcher
type Dispatcher struct {
	handlers map[string]Handlers
	mx       *sync.RWMutex
}

// Handlers is a slice with handlers
type Handlers []Handler

// Handler is a function that handles an event
type Handler func(payload interface{})

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNilDispatcher = fmt.Errorf("Dispatcher wasn't created properly")
	ErrEmptyType     = fmt.Errorf("Event type is empty")
	ErrNilHandler    = fmt.Errorf("Handler must not be nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewDispatcher creates new event dispatcher instance
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[string]Handlers),
		mx:       &sync.RWMutex{},
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// AddHandler registers handler for events with given event type
func (d *Dispatcher) AddHandler(typ string, handler Handler) error {
	err := validateArguments(d, typ, handler, true)

	if err != nil {
		return err
	}

	d.mx.Lock()

	if d.getHandlerIndex(typ, handler) != -1 {
		d.mx.Unlock()
		return fmt.Errorf("Handler already registered for given event type (%s)", typ)
	}

	d.handlers[typ] = append(d.handlers[typ], handler)

	d.mx.Unlock()
	return nil
}

// RemoveHandler removes handler for given event type
func (d *Dispatcher) RemoveHandler(typ string, handler Handler) error {
	err := validateArguments(d, typ, handler, true)

	if err != nil {
		return err
	}

	d.mx.Lock()

	i := d.getHandlerIndex(typ, handler)

	if i == -1 {
		d.mx.Unlock()
		return fmt.Errorf("Handler is not registered for given event type (%s)", typ)
	}

	copy(d.handlers[typ][i:], d.handlers[typ][i+1:])
	d.handlers[typ][len(d.handlers[typ])-1] = nil
	d.handlers[typ] = d.handlers[typ][:len(d.handlers[typ])-1]

	d.mx.Unlock()
	return nil
}

// HasHandler returns true if given handler is registered for given event type
func (d *Dispatcher) HasHandler(typ string, handler Handler) bool {
	err := validateArguments(d, typ, handler, true)

	if err != nil {
		return false
	}

	d.mx.RLock()
	defer d.mx.RUnlock()

	return d.getHandlerIndex(typ, handler) != -1
}

// Dispatch dispatches event with given type and payload
func (d *Dispatcher) Dispatch(typ string, payload interface{}) error {
	err := validateArguments(d, typ, nil, false)

	if err != nil {
		return err
	}

	d.mx.RLock()

	if d.handlers[typ] == nil {
		d.mx.RUnlock()
		return fmt.Errorf("No handlers for event %q", typ)
	}

	d.mx.RUnlock()
	d.mx.RLock()

	for _, h := range d.handlers[typ] {
		go h(payload)
	}

	d.mx.RUnlock()
	return nil
}

// DispatchAndWait dispatches event with given type and payload and waits
// until all handlers will be executed
func (d *Dispatcher) DispatchAndWait(typ string, payload interface{}) error {
	err := validateArguments(d, typ, nil, false)

	if err != nil {
		return err
	}

	d.mx.RLock()

	if d.handlers[typ] == nil {
		d.mx.RUnlock()
		return fmt.Errorf("No handlers for event %q", typ)
	}

	d.mx.RUnlock()
	d.mx.RLock()

	cur, tot := 0, len(d.handlers[typ])
	ch := make(chan bool, tot)

	for _, h := range d.handlers[typ] {
		go execWrapper(ch, h, payload)
	}

	d.mx.RUnlock()

MAIN:
	for {
		select {
		case <-ch:
			cur++
			if cur == tot {
				break MAIN
			}
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getHandlerIndex returns index of handler in the slice
func (d *Dispatcher) getHandlerIndex(typ string, handler Handler) int {
	hp := reflect.ValueOf(handler).Pointer()

	for i, h := range d.handlers[typ] {
		if reflect.ValueOf(h).Pointer() == hp {
			return i
		}
	}

	return -1
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validateArguments validates arguments
func validateArguments(d *Dispatcher, typ string, handler Handler, isHandlerRequired bool) error {
	if d == nil || d.handlers == nil {
		return ErrNilDispatcher
	}

	if typ == "" {
		return ErrEmptyType
	}

	if isHandlerRequired && handler == nil {
		return ErrNilHandler
	}

	return nil
}

// execWrapper exec wrapper runs given handler function and sends true to the channel
// when function is successfully executed
func execWrapper(ch chan bool, handler Handler, payload interface{}) {
	handler(payload)
	ch <- true
}
