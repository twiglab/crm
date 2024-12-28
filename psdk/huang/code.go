package huang

import (
	"io"
	"time"
)

var base62chars = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

var weights = []int{
	3, 4, 15, 16, 5, 6, 13, 14, 9, 10, 11, 12, 1, 2, 7, 8,
	17, 24, 19, 22, 21, 18, 23, 20,
	30, 29, 27, 25, 28, 26, 32, 31,
	35, 33, 36, 34,
}

const (
	yyMMddhhmmss   = "060102150405"
	yyyyMMddhhmmss = "20060102150405"
)

var defCode24 = Code24By('H')
var defCode32 = Code32By('H')

func Code24() string {
	return defCode24()
}

func Code32() string {
	return defCode32()
}

func huang(size int, f string, typ byte) func() string {
	return func() string {
		buf := make([]byte, size)
		nonce62(buf)
		now := time.Now()
		now.AppendFormat(buf[:0], f)
		buf[size-2] = typ
		buf[size-1] = sum(buf)

		return buf2String(buf)
	}
}

func Code24By(typ byte) func() string {
	return huang(24, yyMMddhhmmss, typ)
}

func Code32By(typ byte) func() string {
	return huang(32, yyyyMMddhhmmss, typ)
}

func Code36By(typ byte) func() string {
	return huang(36, yyyyMMddhhmmss, typ)
}

func nonce62(buf []byte) {
	io.ReadFull(src, buf)
	for i, b := range buf {
		buf[i] = base62chars[b%62]
	}
}

func ord(b byte) (n int) {
	if b >= '0' && b <= '9' {
		n = int(b - '0')
	}
	if b >= 'A' && b <= 'Z' {
		n = int(b - 'A')
	}
	if b >= 'a' && b <= 'z' {
		n = int(b - 'a')
	}
	return
}

func sum(bs []byte) byte {
	s := 0
	for i := 0; i < len(bs)-2; i++ {
		n := ord(bs[i])
		s += n * weights[i]
	}
	return base62chars[s%62]
}

func Check(code string) bool {
	s := 0
	for i := 0; i < len(code)-2; i++ {
		n := ord(code[i])
		s += n * weights[i]
	}
	sum := base62chars[s%62]

	return sum == code[len(code)-1]
}
