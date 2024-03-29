// +build !ignore_autogenerated

// Code generated by collections-gen. DO NOT EDIT.

package underlying

import (
	"reflect"
	"sort"
)

// ReaderMapCollection is a collection of ReaderMap values.
type ReaderMapCollection struct {
	items []ReaderMap
}

// NewReaderMapCollection creates a new collection from a slice of ReaderMap.
func NewReaderMapCollection(items []ReaderMap) *ReaderMapCollection {
	return &ReaderMapCollection{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ReaderMapCollection) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of ReaderMap values used by the
// collection.
func (c *ReaderMapCollection) Items() []ReaderMap {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ReaderMapCollection) EachIndex(fn func(ReaderMap, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ReaderMapCollection) Each(fn func(ReaderMap)) {
	c.EachIndex(func(item ReaderMap, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ReaderMapCollection) IndexOf(el ReaderMap) int {
	for i, item := range c.items {
		if reflect.DeepEqual(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ReaderMapCollection) First() ReaderMap {
	return c.Nth(0)
}

// FirstN returns the first n ReaderMap items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ReaderMapCollection) FirstN(n int) []ReaderMap {
	if n > c.Len() {
		return c.Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ReaderMapCollection) Last() ReaderMap {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n ReaderMap items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ReaderMapCollection) LastN(n int) []ReaderMap {
	if c.Len()-n < 0 {
		return c.Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ReaderMapCollection) Get(pos int) ReaderMap {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ReaderMapCollection) Nth(pos int) ReaderMap {
	return c.items[pos]
}

// Len returns the length of the underlying ReaderMap slice.
func (c *ReaderMapCollection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying ReaderMap slice.
func (c *ReaderMapCollection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *ReaderMapCollection) Append(items ...ReaderMap) *ReaderMapCollection {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *ReaderMapCollection) Prepend(items ...ReaderMap) *ReaderMapCollection {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying ReaderMap slice.
func (c *ReaderMapCollection) Copy() *ReaderMapCollection {
	s := make([]ReaderMap, c.Len(), c.Len())
	copy(s, c.items)

	return NewReaderMapCollection(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *ReaderMapCollection) Filter(fn func(ReaderMap) bool) *ReaderMapCollection {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue ReaderMap

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *ReaderMapCollection) Collect(fn func(ReaderMap) bool) *ReaderMapCollection {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *ReaderMapCollection) Reject(fn func(ReaderMap) bool) *ReaderMapCollection {
	return c.Filter(func(v ReaderMap) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ReaderMapCollection) Partition(fn func(ReaderMap) bool) (*ReaderMapCollection, *ReaderMapCollection) {
	lhs := make([]ReaderMap, 0, c.Len())
	rhs := make([]ReaderMap, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewReaderMapCollection(lhs), NewReaderMapCollection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *ReaderMapCollection) Map(fn func(ReaderMap) ReaderMap) *ReaderMapCollection {
	return c.MapIndex(func(item ReaderMap, _ int) ReaderMap {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *ReaderMapCollection) MapIndex(fn func(ReaderMap, int) ReaderMap) *ReaderMapCollection {
	for i, item := range c.items {
		c.items[i] = fn(item, i)

	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero ReaderMap value on the first invocation.
func (c *ReaderMapCollection) Reduce(fn func(reducer ReaderMap, item ReaderMap) ReaderMap) ReaderMap {
	var reducer ReaderMap

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// ReaderMap value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *ReaderMapCollection) Find(fn func(ReaderMap) bool) ReaderMap {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// ReaderMap value. The second return value denotes whether the condition
// matched any item or not.
func (c *ReaderMapCollection) FindOk(fn func(ReaderMap) bool) (ReaderMap, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue ReaderMap
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ReaderMapCollection) Any(fn func(ReaderMap) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ReaderMapCollection) All(fn func(ReaderMap) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ReaderMapCollection) Contains(el ReaderMap) bool {
	for _, item := range c.items {
		if reflect.DeepEqual(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *ReaderMapCollection) Sort(fn func(ReaderMap, ReaderMap) bool) *ReaderMapCollection {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ReaderMapCollection) IsSorted(fn func(ReaderMap, ReaderMap) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ReaderMapCollection) lessFunc(fn func(ReaderMap, ReaderMap) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse reverses the order of the collection items in place and returns c.
func (c *ReaderMapCollection) Reverse() *ReaderMapCollection {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
func (c *ReaderMapCollection) Remove(pos int) *ReaderMapCollection {
	c.items = append(c.items[:pos], c.items[pos+1:]...)
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *ReaderMapCollection) RemoveItem(item ReaderMap) *ReaderMapCollection {
	for i, el := range c.items {
		if reflect.DeepEqual(el, item) {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
func (c *ReaderMapCollection) InsertItem(item ReaderMap, pos int) *ReaderMapCollection {
	var zeroValue ReaderMap
	c.items = append(c.items, zeroValue)
	copy(c.items[pos+1:], c.items[pos:])
	c.items[pos] = item
	return c
}

// Cut returns a copy of the underlying ReaderMap slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ReaderMapCollection) Cut(i, j int) []ReaderMap {
	s := make([]ReaderMap, 0, c.Len())
	s = append(s, c.items[:i]...)
	return append(s, c.items[j:]...)
}

// Slice returns the ReaderMap items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ReaderMapCollection) Slice(i, j int) []ReaderMap {
	return c.items[i:j]
}
