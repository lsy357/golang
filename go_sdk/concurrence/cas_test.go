package main

import (
	"sync/atomic"
	"testing"
	"unsafe"

	"github.com/sirupsen/logrus"
)

type PointerContainer struct {
	ptr unsafe.Pointer // 指向 desc
}

type SyncTestStruct struct {
	desc string
}

func TestAtomicUpdate(t *testing.T) {
	raw, target := &SyncTestStruct{"raw"}, &SyncTestStruct{"updated"}
	pc := &PointerContainer{unsafe.Pointer(raw)}
	logrus.Infof("update raw addr: %p, val: %v", pc.ptr, *(*SyncTestStruct)(pc.ptr))
	pc.AtomicUpdate(target)
	logrus.Infof("update result addr: %p, val: %v", pc.ptr, *(*SyncTestStruct)(pc.ptr))
}

// cas 更新指定对象：更新后持有的变量名不变（PointerContainer 封装），内存变化（PointerContainer.ptr 变化）
func (pc *PointerContainer) AtomicUpdate(target *SyncTestStruct) {
	logrus.Infof("before update addr: %p, val: %v, expected addr: %p, val: %v",
		pc.ptr, *(*SyncTestStruct)(pc.ptr), target, *target)
	for i := 0; i < 100; i++ {
		p := atomic.LoadPointer(&pc.ptr)
		if atomic.CompareAndSwapPointer(&pc.ptr, p, unsafe.Pointer(target)) {
			logrus.Infof("update success, addr: %p, val: %v, try count: %v", pc.ptr, *(*SyncTestStruct)(pc.ptr), i)
			return
		}
	}
}
