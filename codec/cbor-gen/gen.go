package main

import (
	codec "go-codec-bench"

	cbg "github.com/daotl/cbor-gen"
)

func main() {
	if err := cbg.WriteTupleEncodersToFile(
		"values_gen_cbor.go", "codec", true, nil,
		codec.StringUint64T{},
		codec.AnonInTestStruc{},
		codec.TestSimpleFields{},
		codec.TestStrucCommon{},
		codec.TestStruc{},
	); err != nil {
		panic(err)
	}
}
