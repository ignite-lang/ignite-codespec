// Package ledger takes, pays, ships and cancels customer orders. It keeps its orders as values in a
// map and its events in a slice the caller drains after each call: no channel, no goroutine and no
// pointer shared between owners.
package ledger

import (
	"errors"
	"fmt"
)

type OrderID string

type Status string

const (
	Pending   Status = "pending"
	Paid      Status = "paid"
	Shipped   Status = "shipped"
	Cancelled Status = "cancelled"
)

type Order struct {
	ID       OrderID
	Customer string
	Total    int64
	Status   Status
	Tracking string
	Gift     bool
}

var ErrInvalidTotal = errors.New("total must be positive")

type ErrWrongState struct {
	Status Status
}

func (e ErrWrongState) Error() string { return fmt.Sprintf("order is %s", e.Status) }

type EventKind string

const (
	OrderPlaced    EventKind = "OrderPlaced"
	OrderPaid      EventKind = "OrderPaid"
	OrderShipped   EventKind = "OrderShipped"
	OrderCancelled EventKind = "OrderCancelled"
)

// Event is one change of the ledger. Kind names it; the fields it does not use are empty.
type Event struct {
	Kind     EventKind
	ID       OrderID
	Customer string
	Total    int64
	Tracking string
}

type Ledger struct {
	orders map[OrderID]Order
	events []Event
	seq    int
}

func New() Ledger {
	return Ledger{orders: map[OrderID]Order{}, events: []Event{}}
}

func (l *Ledger) Place(customer string, total int64, gift bool) (OrderID, error) {
	if total <= 0 {
		return "", ErrInvalidTotal
	}
	l.seq++
	id := OrderID(fmt.Sprintf("o-%d", l.seq))
	l.orders[id] = Order{ID: id, Customer: customer, Total: total, Status: Pending, Gift: gift}
	l.events = append(l.events, Event{Kind: OrderPlaced, ID: id, Customer: customer, Total: total})
	return id, nil
}

func (l *Ledger) Pay(id OrderID) error {
	o, ok := l.orders[id]
	if !ok {
		return fmt.Errorf("no order %s", id)
	}
	if o.Status != Pending {
		return ErrWrongState{Status: o.Status}
	}
	o.Status = Paid
	l.orders[id] = o
	l.events = append(l.events, Event{Kind: OrderPaid, ID: id})
	return nil
}

func (l *Ledger) Ship(id OrderID, tracking string) error {
	o, ok := l.orders[id]
	if !ok {
		return fmt.Errorf("no order %s", id)
	}
	if o.Status != Paid {
		return ErrWrongState{Status: o.Status}
	}
	o.Status = Shipped
	o.Tracking = tracking
	l.orders[id] = o
	l.events = append(l.events, Event{Kind: OrderShipped, ID: id, Tracking: tracking})
	return nil
}

func (l *Ledger) Cancel(id OrderID) error {
	o, ok := l.orders[id]
	if !ok {
		return fmt.Errorf("no order %s", id)
	}
	if o.Status != Pending && o.Status != Paid {
		return ErrWrongState{Status: o.Status}
	}
	o.Status = Cancelled
	l.orders[id] = o
	l.events = append(l.events, Event{Kind: OrderCancelled, ID: id})
	return nil
}

func (l *Ledger) Get(id OrderID) (Order, bool) {
	o, ok := l.orders[id]
	if !ok {
		return Order{}, false
	}
	return o, true
}

// PendingOrders lists the pending orders in map order, which is random; the caller sorts them.
func (l *Ledger) PendingOrders() []Order {
	out := []Order{}
	for _, o := range l.orders {
		if o.Status == Pending {
			out = append(out, o)
		}
	}
	return out
}

// Drain returns the events since the last Drain, oldest first, and forgets them.
func (l *Ledger) Drain() []Event {
	out := l.events
	l.events = []Event{}
	return out
}
