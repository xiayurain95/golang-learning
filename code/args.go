package code

import (
	"fmt"
	"sync"
)

type info struct {
	mu   sync.Mutex
	rwMu sync.RWMutex
	cnt  int
}

func SyncUpdate() {
	infoIns := new(info)
	infoIns.mu.Lock()
	infoIns.cnt++
	infoIns.mu.Unlock()
	fmt.Printf("%v", infoIns.cnt)
}

func RWFunc() {
	infoIns := new(info)
	infoIns.rwMu.RLock()
	infoIns.cnt++
	infoIns.rwMu.RUnlock()
	fmt.Printf("%v", infoIns.cnt)
}
