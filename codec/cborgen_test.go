package codec

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ipfs/go-ipld-cbor/encoding"
	"github.com/polydawn/refmt/obj/atlas"
)

var (
	ErrTypeNotMatch = errors.New("model type does not match")
	emptyAtlas      = atlas.MustBuild()
	marshaler       = encoding.NewPooledMarshaller(emptyAtlas)
	unmarshaler     = encoding.NewPooledUnmarshaller(emptyAtlas)
)

func fnCborgenEncodeFn(ts interface{}, bsIn []byte) (bs []byte, err error) {
	bin, err := marshaler.Marshal(ts)
	if err != nil {
		panic(err)
	}
	return bin, nil
}

func fnCborgenDecodeFn(buf []byte, ts interface{}) error {
	r := bytes.NewReader(buf)
	if err := unmarshaler.Decode(r, ts); err != nil {
		estr := err.Error()
		if strings.Contains(estr, "wrong type") {
			return ErrTypeNotMatch
		} else if strings.Contains(estr, "EOF") {
			return err
		} else {
			panic("unexpected error in unmarshalling")
		}
	}
	return nil
}

func Benchmark__Cborgen____Encode(b *testing.B) {
	fnBenchmarkEncode(b, "cborgen", benchTs, fnCborgenEncodeFn)
}

func Benchmark__Cborgen____Decode(b *testing.B) {
	fnBenchmarkDecode(b, "cborgen", benchTs, fnCborgenEncodeFn, fnCborgenDecodeFn, fnBenchNewTs)
}
