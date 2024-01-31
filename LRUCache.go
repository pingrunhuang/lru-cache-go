package lrucachego

import (
	"errors"
	"fmt"
)

type Item struct {
	key      any
	value    any
	expires  int16
	priority int16
}

type Cache struct {
	maxsize int16
	cache   map[any]Item
	_time   int16
}

func (self Cache) time() int16 {
	self._time += 1
	return self._time
}

func (self Cache) get(key any) (*Item, error) {
	if self.cache == nil {
		return nil, errors.New("empty cache")
	}

	item, ok := self.cache[key]
	if ok {
		return &item, nil
	} else {
		return nil, errors.New(fmt.Sprintf("no item with key=%v", key))
	}
}

func (self Cache) evict(int16) {
	if len(self.cache) != 0 {
		for k := range self.cache {
			delete(self.cache, k)
			break
		}
	}
}

func (self Cache) set(key any, value any, maxage int16, priority int16) bool {
	// remove
	now := self.time()
	_, ok := self.cache[key]
	switch {
	case ok:
		delete(self.cache, key)
	case len(self.cache) > int(self.maxsize):
		self.evict(now)
	}
	expires := now + maxage
	newItem := &Item{key, value, expires, priority}
	self.cache[key] = *newItem
	return true
}
