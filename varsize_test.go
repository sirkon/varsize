package varsize_test

import (
	"encoding/binary"
	"fmt"
	"github.com/sirkon/varsize"
	"math"
	"testing"
)

func ExampleInt() {
	fmt.Println(varsize.Int(63), varsize.Int(64))
	// Output:
	// 1 2
}

func TestInt(t *testing.T) {
	tests := []any{
		math.MaxInt,
		int8(math.MaxInt8),
		int8(math.MinInt8),
		int16(math.MaxInt16),
		int16(math.MinInt16),
		int32(math.MaxInt32),
		int32(math.MinInt32),
		int64(math.MaxInt64),
		int64(math.MaxInt32),
		int64(math.MinInt64),
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, i)
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T %d", tt, tt), func(t *testing.T) {
			var got int
			var want int
			buf := make([]byte, 32)
			switch v := tt.(type) {
			case int:
				got = varsize.Int(v)
				want = binary.PutVarint(buf, int64(v))
			case int8:
				got = varsize.Int(v)
				want = binary.PutVarint(buf, int64(v))
			case int16:
				got = varsize.Int(v)
				want = binary.PutVarint(buf, int64(v))
			case int32:
				got = varsize.Int(v)
				want = binary.PutVarint(buf, int64(v))
			case int64:
				got = varsize.Int(v)
				want = binary.PutVarint(buf, int64(v))
			default:
				t.Errorf("unsupported type %T", tt)
			}
			if got != want {
				t.Errorf("Int() = %v, want %v", got, want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	tests := []any{
		uint(math.MaxUint),
		uint8(math.MaxUint8),
		uint16(math.MaxUint16),
		uint32(math.MaxUint32),
		uint64(math.MaxUint64),
		uint64(math.MaxUint32),
	}
	for i := 0; i < 1000; i++ {
		tests = append(tests, uint(i))
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T %d", tt, tt), func(t *testing.T) {
			var got int
			var want int
			buf := make([]byte, 32)
			switch v := tt.(type) {
			case uint:
				got = varsize.Uint(v)
				want = binary.PutUvarint(buf, uint64(v))
			case uint8:
				got = varsize.Uint(v)
				want = binary.PutUvarint(buf, uint64(v))
			case uint16:
				got = varsize.Uint(v)
				want = binary.PutUvarint(buf, uint64(v))
			case uint32:
				got = varsize.Uint(v)
				want = binary.PutUvarint(buf, uint64(v))
			case uint64:
				got = varsize.Uint(v)
				want = binary.PutUvarint(buf, uint64(v))
			default:
				t.Errorf("unsupported type %T", tt)
			}
			if got != want {
				t.Errorf("Uint() = %v, want %v", got, want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	const N = 1024
	buf := make([]byte, N)
	for i := 0; i < N; i++ {
		t.Run(fmt.Sprintf("len %d", i), func(t *testing.T) {
			got := varsize.Len(buf[:i])
			var b [32]byte
			want := binary.PutUvarint(b[:], uint64(len(buf[:i])))
			if got != want {
				t.Errorf("Len() = %v, want %v", got, want)
			}
		})
	}
}
