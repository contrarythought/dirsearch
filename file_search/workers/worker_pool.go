package workers

import (
	"strings"
	"sync"
)

const (
	MAX_WORKERS = 3
	ROOT_PATH   = "C:"
)

type DirMap struct {
	DirMap map[string]bool
	Mu     sync.Mutex
}

func NewDirMap() *DirMap {
	return &DirMap{
		DirMap: make(map[string]bool),
	}
}

func (d *DirMap) Append(dir string) {
	d.Mu.Lock()
	defer d.Mu.Unlock()
	if d.DirMap[dir] == false {
		d.DirMap[dir] = true
	}
}

// TODO
func (d *DirMap) SearchDirRecur(startPath, target string) {
	if strings.Compare(ROOT_PATH, startPath) != 0 {

	}
}
