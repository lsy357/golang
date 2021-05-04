package singleton

import (
	"sync"
	"testing"
)

var (
	onceInit = &sync.Once{}

	expensiveInstance *expensive
)

// 用这个比较好
func init() {
	expensiveInstance = &expensive{}
	expensiveInstance.init()
}

func TestSingleton(t *testing.T) {

}

type expensive struct {
	ptr  uint64
	desc string
}

func (e *expensive) init() {
	// lots of operation
	e.desc = "init done"
}

func getSingletonWithOnceImpl() *expensive {
	if expensiveInstance == nil {
		onceInit.Do(func() {
			expensiveInstance = &expensive{}
			expensiveInstance.init()
		})
	}
	return expensiveInstance
}
