package workers

import "sync"

const (
	MAX_WORKERS = 3
)

type WorkerPool struct {
	NumWorkers int
	DirMap     *DirMap
	Mu         sync.Mutex
}

func NewWorkerPool(NumWorkers int) *WorkerPool {
	return &WorkerPool{
		NumWorkers: NumWorkers,
		DirMap:     NewDirMap(),
	}
}

type DirMap struct {
	DirMap map[string]bool
	Mu sync.Mutex
}

func NewDirMap() *DirMap {
	return &DirMap{
		DirMap: make(map[string]bool),
	}
}

func(d *DirMap) Append(dir string) {
	d.Mu.Lock()
	defer d.Mu.Unlock()
	if d.DirMap[dir] == false {
		d.DirMap[dir] = true
	}
}

func (wp *WorkerPool) StartWorkers(work_func func(string, string), startRoot, target string) {
	var wg sync.WaitGroup

	// run each worker
	for i := 0; i < wp.NumWorkers; i++ {
		wg.Add(1)
		go func() {
			work_func(startRoot, target)
			wg.Done()
		}()
	}

	wg.Wait()
}
