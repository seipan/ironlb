package allocator

import (
	"net"
	"sync"
)

type IPAllocator struct {
	pool []net.IP
	mu   sync.Mutex
}

func NewIPAllocator(cidr string) (*IPAllocator, error) {
	return &IPAllocator{}, nil
}

func (a *IPAllocator) Allocate() (net.IP, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	return nil, nil
}

func (a *IPAllocator) Release(ip net.IP) {
	a.mu.Lock()
	defer a.mu.Unlock()
}
