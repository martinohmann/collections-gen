// +build !ignore_autogenerated

// Code generated by collections-gen. DO NOT EDIT.

package customtype

import (
	"sort"
)

// ImmutableTypeCollection is an immutable collection of *Type values.
type ImmutableTypeCollection struct {
	items []*Type
}

// NewImmutableTypeCollection creates a new immutable collection from a slice of *Type.
func NewImmutableTypeCollection(items []*Type) *ImmutableTypeCollection {
	return &ImmutableTypeCollection{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ImmutableTypeCollection) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of *Type values used by the
// collection.
func (c *ImmutableTypeCollection) Items() []*Type {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ImmutableTypeCollection) EachIndex(fn func(*Type, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ImmutableTypeCollection) Each(fn func(*Type)) {
	c.EachIndex(func(item *Type, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ImmutableTypeCollection) IndexOf(el *Type) int {
	for i, item := range c.items {
		if Equal(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ImmutableTypeCollection) First() *Type {
	return c.Nth(0)
}

// FirstN returns the first n *Type items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ImmutableTypeCollection) FirstN(n int) []*Type {
	if n > c.Len() {
		return c.Copy().Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ImmutableTypeCollection) Last() *Type {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n *Type items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ImmutableTypeCollection) LastN(n int) []*Type {
	if c.Len()-n < 0 {
		return c.Copy().Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ImmutableTypeCollection) Get(pos int) *Type {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ImmutableTypeCollection) Nth(pos int) *Type {
	return c.items[pos]
}

// Len returns the length of the underlying *Type slice.
func (c *ImmutableTypeCollection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying *Type slice.
func (c *ImmutableTypeCollection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableTypeCollection) Append(items ...*Type) *ImmutableTypeCollection {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableTypeCollection) Prepend(items ...*Type) *ImmutableTypeCollection {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying *Type slice.
func (c *ImmutableTypeCollection) Copy() *ImmutableTypeCollection {
	s := make([]*Type, c.Len(), c.Len())
	copy(s, c.items)

	return NewImmutableTypeCollection(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableTypeCollection) Filter(fn func(*Type) bool) *ImmutableTypeCollection {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue *Type

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableTypeCollection) Collect(fn func(*Type) bool) *ImmutableTypeCollection {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *ImmutableTypeCollection) Reject(fn func(*Type) bool) *ImmutableTypeCollection {
	return c.Filter(func(v *Type) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ImmutableTypeCollection) Partition(fn func(*Type) bool) (*ImmutableTypeCollection, *ImmutableTypeCollection) {
	lhs := make([]*Type, 0, c.Len())
	rhs := make([]*Type, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewImmutableTypeCollection(lhs), NewImmutableTypeCollection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableTypeCollection) Map(fn func(*Type) *Type) *ImmutableTypeCollection {
	return c.MapIndex(func(item *Type, _ int) *Type {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableTypeCollection) MapIndex(fn func(*Type, int) *Type) *ImmutableTypeCollection {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item, i)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero *Type value on the first invocation.
func (c *ImmutableTypeCollection) Reduce(fn func(reducer *Type, item *Type) *Type) *Type {
	var reducer *Type

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// *Type value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *ImmutableTypeCollection) Find(fn func(*Type) bool) *Type {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// *Type value. The second return value denotes whether the condition
// matched any item or not.
func (c *ImmutableTypeCollection) FindOk(fn func(*Type) bool) (*Type, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue *Type
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ImmutableTypeCollection) Any(fn func(*Type) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ImmutableTypeCollection) All(fn func(*Type) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ImmutableTypeCollection) Contains(el *Type) bool {
	for _, item := range c.items {
		if Equal(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
// The result will be a copy of c which is sorted, the original collection is
// not altered.
func (c *ImmutableTypeCollection) Sort(fn func(*Type, *Type) bool) *ImmutableTypeCollection {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ImmutableTypeCollection) IsSorted(fn func(*Type, *Type) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ImmutableTypeCollection) lessFunc(fn func(*Type, *Type) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *ImmutableTypeCollection) Reverse() *ImmutableTypeCollection {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableTypeCollection) Remove(pos int) *ImmutableTypeCollection {
	d := c.Copy()
	d.items = append(d.items[:pos], d.items[pos+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *ImmutableTypeCollection) RemoveItem(item *Type) *ImmutableTypeCollection {
	d := c.Copy()

	for i, el := range d.items {
		if Equal(el, item) {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableTypeCollection) InsertItem(item *Type, pos int) *ImmutableTypeCollection {
	var zeroValue *Type
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[pos+1:], d.items[pos:])
	d.items[pos] = item
	return d
}

// Cut returns a copy of the underlying *Type slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ImmutableTypeCollection) Cut(i, j int) []*Type {
	d := c.Copy()
	return append(d.items[:i], d.items[j:]...)
}

// Slice returns the *Type items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ImmutableTypeCollection) Slice(i, j int) []*Type {
	return c.Copy().items[i:j]
}