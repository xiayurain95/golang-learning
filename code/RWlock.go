package code

import "sync"

type RWLock struct {
	RDLock, WRLock sync.Mutex
	nRead          int
}

func (l *RWLock) RLock() {
	l.RDLock.Lock()
	l.nRead++
	if l.nRead == 1 {
		l.WRLock.Lock()
	}
	l.RDLock.Unlock()
}
func (l *RWLock) RUnlock() {
	l.RDLock.Lock()
	if l.nRead == 1 {
		l.WRLock.Unlock()
	}
	l.nRead--
	l.RDLock.Unlock()
}
func (l *RWLock) WLock() {
	l.WRLock.Lock()
}
func (l *RWLock) WUnlock() {
	l.WRLock.Lock()
}
