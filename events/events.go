// Package events provides methods and structs for creating event-driven systems
package events

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dispatcher is event dispatcher
type Dispatcher struct {
	handlers map[string]handlers
	mx       sync.RWMutex
}

// Handler is a function that handles an event
type Handler func(payload any)

// ////////////////////////////////////////////////////////////////////////////////// //

// handlers is a slice with handlers
type handlers []Handler

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilDispatcher is returned when dispatcher is nil
	ErrNilDispatcher = errors.New("dispatcher is nil")

	// ErrEmptyName is returned when event name is empty
	ErrEmptyName = errors.New("event name is empty")

	// ErrNilHandler is returned when handler is nil
	ErrNilHandler = errors.New("handler must not be nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewDispatcher creates new event dispatcher instance
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[string]handlers),
		mx:       sync.RWMutex{},
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// AddHandler registers handler for specified event
func (d *Dispatcher) AddHandler(ev string, handler Handler) error {
	err := d.validateArguments(ev, handler, true)

	if err != nil {
		return err
	}

	d.mx.Lock()
	defer d.mx.Unlock()

	if d.getHandlerIndex(ev, handler) != -1 {
		return fmt.Errorf("handler already registered for given event (%s)", ev)
	}

	d.handlers[ev] = append(d.handlers[ev], handler)

	return nil
}

// RemoveHandler removes handler for specified event
func (d *Dispatcher) RemoveHandler(ev string, handler Handler) error {
	err := d.validateArguments(ev, handler, true)

	if err != nil {
		return err
	}

	d.mx.Lock()
	defer d.mx.Unlock()

	i := d.getHandlerIndex(ev, handler)

	if i == -1 {
		return fmt.Errorf("handler is not registered for given event (%s)", ev)
	}

	copy(d.handlers[ev][i:], d.handlers[ev][i+1:])
	d.handlers[ev][len(d.handlers[ev])-1] = nil
	d.handlers[ev] = d.handlers[ev][:len(d.handlers[ev])-1]

	return nil
}

// HasHandler returns true if there is a handler for given event
func (d *Dispatcher) HasHandler(ev string, handler Handler) bool {
	err := d.validateArguments(ev, handler, true)

	if err != nil {
		return false
	}

	d.mx.RLock()
	defer d.mx.RUnlock()

	return d.getHandlerIndex(ev, handler) != -1
}

// Dispatch dispatches event with given payload
func (d *Dispatcher) Dispatch(ev string, payload any) error {
	err := d.validateArguments(ev, nil, false)

	if err != nil {
		return err
	}

	d.mx.RLock()
	defer d.mx.RUnlock()

	if d.handlers[ev] == nil {
		return fmt.Errorf("no handlers for event %q", ev)
	}

	for _, h := range d.handlers[ev] {
		go h(payload)
	}

	return nil
}

// DispatchAndWait dispatches event with given payload and waits
// until all handlers will be executed
func (d *Dispatcher) DispatchAndWait(ev string, payload any) error {
	err := d.validateArguments(ev, nil, false)

	if err != nil {
		return err
	}

	d.mx.RLock()
	defer d.mx.RUnlock()

	if d.handlers[ev] == nil {
		return fmt.Errorf("no handlers for event %q", ev)
	}

	var wg sync.WaitGroup

	for _, h := range d.handlers[ev] {
		wg.Add(1)

		go func(h Handler) {
			defer wg.Done()
			defer func() { recover() }()
			h(payload)
		}(h)
	}

	wg.Wait()

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getHandlerIndex returns the index of handler in d.handlers[ev].
// Caller must hold d.mx (at minimum for reading) before calling this method.
func (d *Dispatcher) getHandlerIndex(ev string, handler Handler) int {
	hp := reflect.ValueOf(handler).Pointer()

	for i, h := range d.handlers[ev] {
		if reflect.ValueOf(h).Pointer() == hp {
			return i
		}
	}

	return -1
}

// validateArguments checks that the event name are non-empty, and optionally that
// the handler is non-nil
func (d *Dispatcher) validateArguments(ev string, handler Handler, isHandlerRequired bool) error {
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

// ////////////////////////////////////////////////////////////////////////////////// //
