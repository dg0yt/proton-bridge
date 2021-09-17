// Copyright (c) 2021 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package pchan

import (
	"sort"
	"sync"
)

type PChan struct {
	lock        sync.Mutex
	items       []*Item
	ready, done chan struct{}
}

type Item struct {
	ch   *PChan
	val  interface{}
	prio int
	done chan struct{}
}

func (item *Item) Wait() {
	<-item.done
}

func (item *Item) GetPriority() int {
	item.ch.lock.Lock()
	defer item.ch.lock.Unlock()

	return item.prio
}

func (item *Item) SetPriority(priority int) {
	item.ch.lock.Lock()
	defer item.ch.lock.Unlock()

	item.prio = priority

	sort.Slice(item.ch.items, func(i, j int) bool {
		return item.ch.items[i].prio < item.ch.items[j].prio
	})
}

func New() *PChan {
	return &PChan{
		ready: make(chan struct{}),
		done:  make(chan struct{}),
	}
}

func (ch *PChan) Push(val interface{}, prio int) *Item {
	defer ch.notify()

	return ch.push(val, prio)
}

func (ch *PChan) Pop() (interface{}, int, bool) {
	select {
	case <-ch.ready:
		val, prio := ch.pop()
		return val, prio, true

	case <-ch.done:
		return nil, 0, false
	}
}

func (ch *PChan) Close() {
	select {
	case <-ch.done:
		return

	default:
		close(ch.done)
	}
}

func (ch *PChan) push(val interface{}, prio int) *Item {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	done := make(chan struct{})

	item := &Item{
		ch:   ch,
		val:  val,
		prio: prio,
		done: done,
	}

	ch.items = append(ch.items, item)

	return item
}

func (ch *PChan) pop() (interface{}, int) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	sort.Slice(ch.items, func(i, j int) bool {
		return ch.items[i].prio < ch.items[j].prio
	})

	var item *Item

	item, ch.items = ch.items[len(ch.items)-1], ch.items[:len(ch.items)-1]

	defer close(item.done)

	return item.val, item.prio
}

func (ch *PChan) notify() {
	go func() { ch.ready <- struct{}{} }()
}