// +build !ignore_autogenerated

// Code generated by collections-gen. DO NOT EDIT.

package mapelement

import (
	"io"
	"reflect"
	"sort"
)

// WriterMap is a collection of map[string]io.Writer values.
type WriterMap struct {
	items []map[string]io.Writer
}

// NewWriterMap creates a new collection from a slice of map[string]io.Writer.
func NewWriterMap(items []map[string]io.Writer) *WriterMap {
	return &WriterMap{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *WriterMap) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of map[string]io.Writer values used by the
// collection.
func (c *WriterMap) Items() []map[string]io.Writer {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *WriterMap) EachIndex(fn func(map[string]io.Writer, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *WriterMap) Each(fn func(map[string]io.Writer)) {
	c.EachIndex(func(item map[string]io.Writer, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *WriterMap) IndexOf(el map[string]io.Writer) int {
	for i, item := range c.items {
		if reflect.DeepEqual(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *WriterMap) First() map[string]io.Writer {
	return c.Nth(0)
}

// FirstN returns the first n map[string]io.Writer items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *WriterMap) FirstN(n int) []map[string]io.Writer {
	if n > c.Len() {
		return c.Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *WriterMap) Last() map[string]io.Writer {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n map[string]io.Writer items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *WriterMap) LastN(n int) []map[string]io.Writer {
	if c.Len()-n < 0 {
		return c.Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *WriterMap) Get(pos int) map[string]io.Writer {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *WriterMap) Nth(pos int) map[string]io.Writer {
	return c.items[pos]
}

// Len returns the length of the underlying map[string]io.Writer slice.
func (c *WriterMap) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying map[string]io.Writer slice.
func (c *WriterMap) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *WriterMap) Append(items ...map[string]io.Writer) *WriterMap {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *WriterMap) Prepend(items ...map[string]io.Writer) *WriterMap {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying map[string]io.Writer slice.
func (c *WriterMap) Copy() *WriterMap {
	s := make([]map[string]io.Writer, c.Len(), c.Len())
	copy(s, c.items)

	return NewWriterMap(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *WriterMap) Filter(fn func(map[string]io.Writer) bool) *WriterMap {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue map[string]io.Writer

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *WriterMap) Collect(fn func(map[string]io.Writer) bool) *WriterMap {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *WriterMap) Reject(fn func(map[string]io.Writer) bool) *WriterMap {
	return c.Filter(func(v map[string]io.Writer) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *WriterMap) Partition(fn func(map[string]io.Writer) bool) (*WriterMap, *WriterMap) {
	lhs := make([]map[string]io.Writer, 0, c.Len())
	rhs := make([]map[string]io.Writer, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewWriterMap(lhs), NewWriterMap(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *WriterMap) Map(fn func(map[string]io.Writer) map[string]io.Writer) *WriterMap {
	return c.MapIndex(func(item map[string]io.Writer, _ int) map[string]io.Writer {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *WriterMap) MapIndex(fn func(map[string]io.Writer, int) map[string]io.Writer) *WriterMap {
	for i, item := range c.items {
		c.items[i] = fn(item, i)

	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero map[string]io.Writer value on the first invocation.
func (c *WriterMap) Reduce(fn func(reducer map[string]io.Writer, item map[string]io.Writer) map[string]io.Writer) map[string]io.Writer {
	var reducer map[string]io.Writer

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// map[string]io.Writer value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *WriterMap) Find(fn func(map[string]io.Writer) bool) map[string]io.Writer {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// map[string]io.Writer value. The second return value denotes whether the condition
// matched any item or not.
func (c *WriterMap) FindOk(fn func(map[string]io.Writer) bool) (map[string]io.Writer, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue map[string]io.Writer
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *WriterMap) Any(fn func(map[string]io.Writer) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *WriterMap) All(fn func(map[string]io.Writer) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *WriterMap) Contains(el map[string]io.Writer) bool {
	for _, item := range c.items {
		if reflect.DeepEqual(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *WriterMap) Sort(fn func(map[string]io.Writer, map[string]io.Writer) bool) *WriterMap {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *WriterMap) IsSorted(fn func(map[string]io.Writer, map[string]io.Writer) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *WriterMap) lessFunc(fn func(map[string]io.Writer, map[string]io.Writer) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse reverses the order of the collection items in place and returns c.
func (c *WriterMap) Reverse() *WriterMap {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
func (c *WriterMap) Remove(pos int) *WriterMap {
	c.items = append(c.items[:pos], c.items[pos+1:]...)
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *WriterMap) RemoveItem(item map[string]io.Writer) *WriterMap {
	for i, el := range c.items {
		if reflect.DeepEqual(el, item) {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
func (c *WriterMap) InsertItem(item map[string]io.Writer, pos int) *WriterMap {
	var zeroValue map[string]io.Writer
	c.items = append(c.items, zeroValue)
	copy(c.items[pos+1:], c.items[pos:])
	c.items[pos] = item
	return c
}

// Cut returns a copy of the underlying map[string]io.Writer slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *WriterMap) Cut(i, j int) []map[string]io.Writer {
	s := make([]map[string]io.Writer, 0, c.Len())
	s = append(s, c.items[:i]...)
	return append(s, c.items[j:]...)
}

// Slice returns the map[string]io.Writer items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *WriterMap) Slice(i, j int) []map[string]io.Writer {
	return c.items[i:j]
}

// ImmutableWriterMap is an immutable collection of map[string]io.Writer values.
type ImmutableWriterMap struct {
	items []map[string]io.Writer
}

// NewImmutableWriterMap creates a new immutable collection from a slice of map[string]io.Writer.
func NewImmutableWriterMap(items []map[string]io.Writer) *ImmutableWriterMap {
	return &ImmutableWriterMap{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ImmutableWriterMap) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of map[string]io.Writer values used by the
// collection.
func (c *ImmutableWriterMap) Items() []map[string]io.Writer {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ImmutableWriterMap) EachIndex(fn func(map[string]io.Writer, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ImmutableWriterMap) Each(fn func(map[string]io.Writer)) {
	c.EachIndex(func(item map[string]io.Writer, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ImmutableWriterMap) IndexOf(el map[string]io.Writer) int {
	for i, item := range c.items {
		if reflect.DeepEqual(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ImmutableWriterMap) First() map[string]io.Writer {
	return c.Nth(0)
}

// FirstN returns the first n map[string]io.Writer items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ImmutableWriterMap) FirstN(n int) []map[string]io.Writer {
	if n > c.Len() {
		return c.Copy().Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ImmutableWriterMap) Last() map[string]io.Writer {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n map[string]io.Writer items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ImmutableWriterMap) LastN(n int) []map[string]io.Writer {
	if c.Len()-n < 0 {
		return c.Copy().Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ImmutableWriterMap) Get(pos int) map[string]io.Writer {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ImmutableWriterMap) Nth(pos int) map[string]io.Writer {
	return c.items[pos]
}

// Len returns the length of the underlying map[string]io.Writer slice.
func (c *ImmutableWriterMap) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying map[string]io.Writer slice.
func (c *ImmutableWriterMap) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableWriterMap) Append(items ...map[string]io.Writer) *ImmutableWriterMap {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableWriterMap) Prepend(items ...map[string]io.Writer) *ImmutableWriterMap {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying map[string]io.Writer slice.
func (c *ImmutableWriterMap) Copy() *ImmutableWriterMap {
	s := make([]map[string]io.Writer, c.Len(), c.Len())
	copy(s, c.items)

	return NewImmutableWriterMap(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableWriterMap) Filter(fn func(map[string]io.Writer) bool) *ImmutableWriterMap {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue map[string]io.Writer

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableWriterMap) Collect(fn func(map[string]io.Writer) bool) *ImmutableWriterMap {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *ImmutableWriterMap) Reject(fn func(map[string]io.Writer) bool) *ImmutableWriterMap {
	return c.Filter(func(v map[string]io.Writer) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ImmutableWriterMap) Partition(fn func(map[string]io.Writer) bool) (*ImmutableWriterMap, *ImmutableWriterMap) {
	lhs := make([]map[string]io.Writer, 0, c.Len())
	rhs := make([]map[string]io.Writer, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewImmutableWriterMap(lhs), NewImmutableWriterMap(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableWriterMap) Map(fn func(map[string]io.Writer) map[string]io.Writer) *ImmutableWriterMap {
	return c.MapIndex(func(item map[string]io.Writer, _ int) map[string]io.Writer {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableWriterMap) MapIndex(fn func(map[string]io.Writer, int) map[string]io.Writer) *ImmutableWriterMap {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item, i)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero map[string]io.Writer value on the first invocation.
func (c *ImmutableWriterMap) Reduce(fn func(reducer map[string]io.Writer, item map[string]io.Writer) map[string]io.Writer) map[string]io.Writer {
	var reducer map[string]io.Writer

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// map[string]io.Writer value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *ImmutableWriterMap) Find(fn func(map[string]io.Writer) bool) map[string]io.Writer {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// map[string]io.Writer value. The second return value denotes whether the condition
// matched any item or not.
func (c *ImmutableWriterMap) FindOk(fn func(map[string]io.Writer) bool) (map[string]io.Writer, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue map[string]io.Writer
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ImmutableWriterMap) Any(fn func(map[string]io.Writer) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ImmutableWriterMap) All(fn func(map[string]io.Writer) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ImmutableWriterMap) Contains(el map[string]io.Writer) bool {
	for _, item := range c.items {
		if reflect.DeepEqual(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
// The result will be a copy of c which is sorted, the original collection is
// not altered.
func (c *ImmutableWriterMap) Sort(fn func(map[string]io.Writer, map[string]io.Writer) bool) *ImmutableWriterMap {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ImmutableWriterMap) IsSorted(fn func(map[string]io.Writer, map[string]io.Writer) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ImmutableWriterMap) lessFunc(fn func(map[string]io.Writer, map[string]io.Writer) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *ImmutableWriterMap) Reverse() *ImmutableWriterMap {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableWriterMap) Remove(pos int) *ImmutableWriterMap {
	d := c.Copy()
	d.items = append(d.items[:pos], d.items[pos+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *ImmutableWriterMap) RemoveItem(item map[string]io.Writer) *ImmutableWriterMap {
	d := c.Copy()

	for i, el := range d.items {
		if reflect.DeepEqual(el, item) {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableWriterMap) InsertItem(item map[string]io.Writer, pos int) *ImmutableWriterMap {
	var zeroValue map[string]io.Writer
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[pos+1:], d.items[pos:])
	d.items[pos] = item
	return d
}

// Cut returns a copy of the underlying map[string]io.Writer slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ImmutableWriterMap) Cut(i, j int) []map[string]io.Writer {
	d := c.Copy()
	return append(d.items[:i], d.items[j:]...)
}

// Slice returns the map[string]io.Writer items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ImmutableWriterMap) Slice(i, j int) []map[string]io.Writer {
	return c.Copy().items[i:j]
}
