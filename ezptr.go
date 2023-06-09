package ezpwngo

import (
	"fmt"
	"strconv"
)

type Ptr64 int64

func NewPtr64FromString(s string) Ptr64 {
	v, _ := strconv.ParseInt(s, 16, 64)
	return Ptr64(v)
}

func (p Ptr64) ToBytes() []byte {
	// convert into bytes with little endian
	r := make([]byte, 8)
	for i := 0; i < 8; i++ {
		r[i] = byte(p >> (i * 8) & 0xFF)
	}
	return r
}
func (p Ptr64) ToString() string {
	return fmt.Sprintf("%x", p)
}
