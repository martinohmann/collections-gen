package mapelement

import "io"

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/map-element -p github.com/martinohmann/collections-gen/examples/map-element

// +collections-gen=true
// +collections-gen:options=underlying,out-name=reader_map
// +collections-gen:options=underlying,immutable,out-name=reader_map

type ReaderMap map[string]io.Reader

// +collections-gen=true
// +collections-gen:options=name=WriterMap,out-name=writer_map
// +collections-gen:options=immutable,name=ImmutableWriterMap,out-name=writer_map

type wm = map[string]io.Writer
