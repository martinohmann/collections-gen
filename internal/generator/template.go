package generator

var collectionCode = `
// {{.type|public}} is a{{if .immutable}}n immutable{{end}} collection of {{.elemtype|raw}} values.
type {{.type|public}} struct {
	items []{{.elemtype|raw}}
}

// New{{.type|public}} creates a new{{if .immutable}} immutable{{end}} collection from a slice of {{.elemtype|raw}}.
func New{{.type|public}}(items []{{.elemtype|raw}}) *{{.type|public}} {
	return &{{.type|public}}{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *{{.type|public}}) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of {{.elemtype|raw}} values used by the
// collection.
func (c *{{.type|public}}) Items() []{{.elemtype|raw}} {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *{{.type|public}}) EachIndex(fn func({{.elemtype|raw}}, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *{{.type|public}}) Each(fn func({{.elemtype|raw}})) {
	c.EachIndex(func(item {{.elemtype|raw}}, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *{{.type|public}}) IndexOf(el {{.elemtype|raw}}) int {
	for i, item := range c.items {
{{ if .equalityFunc -}}
		if {{.equalityFunc}}(item, el) {
{{ else -}}
		if item == el {
{{ end -}}
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *{{.type|public}}) First() {{.elemtype|raw}} {
	return c.Nth(0)
}

// FirstN returns the first n {{.elemtype|raw}} items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *{{.type|public}}) FirstN(n int) []{{.elemtype|raw}} {
	if n > c.Len() {
{{ if .immutable -}}
        return c.Copy().Items()
{{- else -}}
		return c.Items()
{{- end }}
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *{{.type|public}}) Last() {{.elemtype|raw}} {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n {{.elemtype|raw}} items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *{{.type|public}}) LastN(n int) []{{.elemtype|raw}} {
	if c.Len()-n < 0 {
{{ if .immutable -}}
        return c.Copy().Items()
{{- else -}}
		return c.Items()
{{- end }}
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *{{.type|public}}) Get(pos int) {{.elemtype|raw}} {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *{{.type|public}}) Nth(pos int) {{.elemtype|raw}} {
	return c.items[pos]
}

// Len returns the length of the underlying {{.elemtype|raw}} slice.
func (c *{{.type|public}}) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying {{.elemtype|raw}} slice.
func (c *{{.type|public}}) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.{{if .immutable}} The
// original collection will not be modified.{{end}}
func (c *{{.type|public}}) Append(items ...{{.elemtype|raw}}) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
{{ else -}}
	c.items = append(c.items, items...)
	return c
{{ end -}}
}

// Prepend prepends items and returns the collection.{{if .immutable}} The
// original collection will not be modified.{{end}}
func (c *{{.type|public}}) Prepend(items ...{{.elemtype|raw}}) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
{{ else -}}
	c.items = append(items, c.items...)
	return c
{{ end -}}
}

// Copy creates a copy of the collection and the underlying {{.elemtype|raw}} slice.
func (c *{{.type|public}}) Copy() *{{.type|public}} {
	s := make([]{{.elemtype|raw}}, c.Len(), c.Len())
	copy(s, c.items)

	return New{{.type|public}}(s)
}

{{ if .immutable -}}
// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
{{- else -}}
// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
{{- end }}
func (c *{{.type|public}}) Filter(fn func({{.elemtype|raw}}) bool) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue {{.elemtype|raw}}

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
{{ else -}}
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue {{.elemtype|raw}}

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
{{ end -}}
}

{{ if .immutable -}}
// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
{{- else -}}
// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
{{- end }}
func (c *{{.type|public}}) Collect(fn func({{.elemtype|raw}}) bool) *{{.type|public}} {
	return c.Filter(fn)
}

{{ if .immutable -}}
// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
{{- else -}}
// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
{{- end }}
func (c *{{.type|public}}) Reject(fn func({{.elemtype|raw}}) bool) *{{.type|public}} {
	return c.Filter(func(v {{.elemtype|raw}}) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *{{.type|public}}) Partition(fn func({{.elemtype|raw}}) bool) (*{{.type|public}}, *{{.type|public}}) {
	lhs := make([]{{.elemtype|raw}}, 0, c.Len())
	rhs := make([]{{.elemtype|raw}}, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return New{{.type|public}}(lhs), New{{.type|public}}(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.{{if .immutable}} The result is a new collection. The original
// collection is not modified.{{end}}
func (c *{{.type|public}}) Map(fn func({{.elemtype|raw}}) {{.elemtype|raw}}) *{{.type|public}} {
	return c.MapIndex(func(item {{.elemtype|raw}}, _ int) {{.elemtype|raw}} {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn.{{if .immutable}} The result is a new collection. The original
// collection is not modified.{{end}}
func (c *{{.type|public}}) MapIndex(fn func({{.elemtype|raw}}, int) {{.elemtype|raw}}) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item, i)

	}

	return d
{{ else -}}
	for i, item := range c.items {
		c.items[i] = fn(item, i)

	}

	return c
{{ end -}}
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero {{.elemtype|raw}} value on the first invocation.
func (c *{{.type|public}}) Reduce(fn func(reducer {{.elemtype|raw}}, item {{.elemtype|raw}}) {{.elemtype|raw}}) {{.elemtype|raw}} {
	var reducer {{.elemtype|raw}}

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// {{.elemtype|raw}} value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *{{.type|public}}) Find(fn func({{.elemtype|raw}}) bool) {{.elemtype|raw}} {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// {{.elemtype|raw}} value. The second return value denotes whether the condition
// matched any item or not.
func (c *{{.type|public}}) FindOk(fn func({{.elemtype|raw}}) bool) ({{.elemtype|raw}}, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue {{.elemtype|raw}}
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *{{.type|public}}) Any(fn func({{.elemtype|raw}}) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *{{.type|public}}) All(fn func({{.elemtype|raw}}) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *{{.type|public}}) Contains(el {{.elemtype|raw}}) bool {
	for _, item := range c.items {
{{ if .equalityFunc -}}
		if {{.equalityFunc}}(item, el) {
{{ else -}}
		if item == el {
{{ end -}}
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
{{- if .immutable }}
// The result will be a copy of c which is sorted, the original collection is
// not altered.
{{- end }}
func (c *{{.type|public}}) Sort(fn func({{.elemtype|raw}}, {{.elemtype|raw}}) bool) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
{{ else -}}
	sort.Slice(c.items, c.lessFunc(fn))
	return c
{{ end -}}
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *{{.type|public}}) IsSorted(fn func({{.elemtype|raw}}, {{.elemtype|raw}}) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *{{.type|public}}) lessFunc(fn func({{.elemtype|raw}}, {{.elemtype|raw}}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

{{ if .immutable -}}
// Reverse copies the collection and returns it with the order of all items
// reversed.
{{- else -}}
// Reverse reverses the order of the collection items in place and returns c.
{{- end}}
func (c *{{.type|public}}) Reverse() *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
{{ else -}}
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
{{ end -}}
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
{{- if .immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.type|public}}) Remove(pos int) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()
	d.items = append(d.items[:pos], d.items[pos+1:]...)
	return d
{{ else -}}
	c.items = append(c.items[:pos], c.items[pos+1:]...)
	return c
{{ end -}}
}

// RemoveItem removes all instances of item from the collection and returns it.
{{- if .immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.type|public}}) RemoveItem(item {{.elemtype|raw}}) *{{.type|public}} {
{{ if .immutable -}}
	d := c.Copy()

	for i, el := range d.items {
{{ if .equalityFunc -}}
		if {{.equalityFunc}}(el, item) {
{{ else -}}
		if el == item {
{{ end -}}
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
{{ else -}}
	for i, el := range c.items {
{{ if .equalityFunc -}}
		if {{.equalityFunc}}(el, item) {
{{ else -}}
		if el == item {
{{ end -}}
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
{{ end -}}
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
{{- if .immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.type|public}}) InsertItem(item {{.elemtype|raw}}, pos int) *{{.type|public}} {
    var zeroValue {{.elemtype|raw}}
{{ if .immutable -}}
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[pos+1:], d.items[pos:])
	d.items[pos] = item
	return d
{{ else -}}
	c.items = append(c.items, zeroValue)
	copy(c.items[pos+1:], c.items[pos:])
	c.items[pos] = item
	return c
{{ end -}}
}

// Cut returns a copy of the underlying {{.elemtype|raw}} slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *{{.type|public}}) Cut(i, j int) []{{.elemtype|raw}} {
{{ if .immutable -}}
	d := c.Copy()
	return append(d.items[:i], d.items[j:]...)
{{ else -}}
	s := make([]{{.elemtype|raw}}, 0, c.Len())
	s = append(s, c.items[:i]...)
	return append(s, c.items[j:]...)
{{ end -}}
}

// Slice returns the {{.elemtype|raw}} items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *{{.type|public}}) Slice(i, j int) []{{.elemtype|raw}} {
{{ if .immutable -}}
	return c.Copy().items[i:j]
{{ else -}}
	return c.items[i:j]
{{ end -}}
}
`
