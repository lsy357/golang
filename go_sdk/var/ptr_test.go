package _var

import (
	"testing"
)

/**

1. map/slice 都是指针类型变量，chan 也是指针变量: TestChanVarType

*/

func TestChanVarType(t *testing.T) {
	strChan := make(chan string, 1)
	t.Logf("chan val: %v, assert ptr: %p", strChan, strChan)
}
