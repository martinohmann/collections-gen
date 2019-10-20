package underlying

import "io"

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/underlying -p github.com/martinohmann/collections-gen/examples/underlying

// +collections-gen=true
//
// Generate ReaderMapCollection using ReaderMap as element type.
//
// +collections-gen:options=out-name=reader_map
//
// Generate ReaderMapUnderlyingCollection using map[string]io.Reader as element
// type.
//
// +collections-gen:options=underlying,name=ReaderMapUnderlyingCollection,out-name=reader_map_underlying

type ReaderMap map[string]io.Reader

// +collections-gen=true
//
// Generate WriterMapCollection using map[string]io.Writer as element type.
//
// +collections-gen:options=name=WriterMapCollection,out-name=writer_map

type wm = map[string]io.Writer
