package huang

import "unsafe"

func buf2String(buf []byte) string {
	return unsafe.String(&buf[0], len(buf))
}
