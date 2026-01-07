// Package events provides methods and structs for creating event-driven systems
package events

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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
type Handler func(payload any)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilDispatcher is returned when dispatcher is nil
	ErrNilDispatcher = fmt.Errorf("Dispatcher is nil")

	// ErrEmptyName is returned when event name is empty
	ErrEmptyName     = fmt.Errorf("Event name is empty")

	// ErrNilHandler is returned when handler is nil
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

// AddHandler registers handler for specified event
func (d *Dispatcher) AddHandler(ev string, handler Handler) error {
	err := validateArguments(d, ev, handler, true)

	if err != nil {
		return err
	}

	d.mx.Lock()

	if d.getHandlerIndex(ev, handler) != -1 {
		d.mx.Unlock()
		return fmt.Errorf("Handler already registered for given event (%s)", ev)
	}

	d.handlers[ev] = append(d.handlers[ev], handler)

	d.mx.Unlock()
	return nil
}

// RemoveHandler removes handler for specified event
func (d *Dispatcher) RemoveHandler(ev string, handler Handler) error {
	err := validateArguments(d, ev, handler, true)

	if err != nil {
		return err
	}

	d.mx.Lock()

	i := d.getHandlerIndex(ev, handler)

	if i == -1 {
		d.mx.Unlock()
		return fmt.Errorf("Handler is not registered for given event (%s)", ev)
	}

	copy(d.handlers[ev][i:], d.handlers[ev][i+1:])
	d.handlers[ev][len(d.handlers[ev])-1] = nil
	d.handlers[ev] = d.handlers[ev][:len(d.handlers[ev])-1]

	d.mx.Unlock()
	return nil
}

// HasHandler returns true if there is a handler for given event
func (d *Dispatcher) HasHandler(ev string, handler Handler) bool {
	err := validateArguments(d, ev, handler, true)

	if err != nil {
		return false
	}

	d.mx.RLock()
	defer d.mx.RUnlock()

	return d.getHandlerIndex(ev, handler) != -1
}

// Dispatch dispatches event with given payload
func (d *Dispatcher) Dispatch(ev string, payload any) error {
	err := validateArguments(d, ev, nil, false)

	if err != nil {
		return err
	}

	d.mx.RLock()

	if d.handlers[ev] == nil {
		d.mx.RUnlock()
		return fmt.Errorf("No handlers for event %q", ev)
	}

	d.mx.RUnlock()
	d.mx.RLock()

	for _, h := range d.handlers[ev] {
		go h(payload)
	}

	d.mx.RUnlock()
	return nil
}

// DispatchAndWait dispatches event with given payload and waits
// until all handlers will be executed
func (d *Dispatcher) DispatchAndWait(ev string, payload any) error {
	err := validateArguments(d, ev, nil, false)

	if err != nil {
		return err
	}

	d.mx.RLock()

	if d.handlers[ev] == nil {
		d.mx.RUnlock()
		return fmt.Errorf("No handlers for event %q", ev)
	}

	d.mx.RUnlock()
	d.mx.RLock()

	cur, tot := 0, len(d.handlers[ev])
	ch := make(chan bool, tot)

	for _, h := range d.handlers[ev] {
		go execWrapper(ch, h, payload)
	}

	d.mx.RUnlock()

	for range ch {
		cur++
		if cur == tot {
			break
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getHandlerIndex returns index of handler in the slice
func (d *Dispatcher) getHandlerIndex(ev string, handler Handler) int {
	hp := reflect.ValueOf(handler).Pointer()

	for i, h := range d.handlers[ev] {
		if reflect.ValueOf(h).Pointer() == hp {
			return i
		}
	}

	return -1
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validateArguments validates arguments
func validateArguments(d *Dispatcher, ev string, handler Handler, isHandlerRequired bool) error {
	if d == nil || d.handlers == nil {
		return ErrNilDispatcher
	}

	if ev == "" {
		return ErrEmptyName
	}

	if isHandlerRequired && handler == nil {
		return ErrNilHandler
	}

	return nil
}

// execWrapper exec wrapper runs given handler function and sends true to the channel
// when function is successfully executed
func execWrapper(ch chan bool, handler Handler, payload any) {
	handler(payload)
	ch <- true
}
