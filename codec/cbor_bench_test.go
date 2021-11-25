// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

package codec

import "testing"

func benchmarkCBORGroup(t *testing.B) {
	benchmarkDivider()
	t.Run("Benchmark__Cbor_______Encode", Benchmark__Cbor_______Encode)
	t.Run("Benchmark__Fxcbor_____Encode", Benchmark__Fxcbor_____Encode)
	t.Run("Benchmark__Gcbor______Encode", Benchmark__Gcbor______Encode)

	benchmarkDivider()
	t.Run("Benchmark__Cbor_______Decode", Benchmark__Cbor_______Decode)
	t.Run("Benchmark__Fxcbor_____Decode", Benchmark__Fxcbor_____Decode)
	t.Run("Benchmark__Gcbor______Decode", Benchmark__Gcbor______Decode)
}

func BenchmarkCBORSuite(t *testing.B) { benchmarkSuite(t, benchmarkCBORGroup) }
