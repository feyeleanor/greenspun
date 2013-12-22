package greenspun

import "sync"

type Mutex struct {
	sync.Mutex
}


func (m Mutex) CriticalSection(f func()) {
	m.Lock()
	f()
	m.Unlock()
}
