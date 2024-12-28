package huang

import (
	"bufio"
	"crypto/rand"
	"io"
	"sync"
)

type entropyPool struct {
	mu     sync.Mutex
	buffer io.Reader
}

func (p *entropyPool) Read(bs []byte) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.buffer.Read(bs)
}

func SetEntropyPool(ep io.Reader) {
	src = ep
}

var src io.Reader

func init() {
	SetEntropyPool(&entropyPool{buffer: bufio.NewReader(rand.Reader)})
}
