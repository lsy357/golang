package unsafe

import (
	"fmt"
	"hello/go_sdk/test"
	"reflect"
	"testing"
	"unsafe"
)

/**

指针变量、unsafe.Pointer、uintptr：
	1. * 类型: 普通指针类型，用于传递对象地址，不能进行指针运算
	2. unsafe.Pointer: 通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值（必须转换到某一类型的普通指针）
	3. uintptr: 用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象，即 uintptr 类型的目标会被回收

	常见的操作为 unsafe.Pointer 转为 unitptr 后进行指针运算，如增加偏移量，运算完成后再转为 unsafe.Pointer

*/

func TestUnsafePointer(t *testing.T) {
	v1 := uint(12)
	v2 := int(13)

	fmt.Println(reflect.TypeOf(v1))  // uint
	fmt.Println(reflect.TypeOf(v2))  // int
	fmt.Println(reflect.TypeOf(&v1)) // *uint
	fmt.Println(reflect.TypeOf(&v2)) // *int

	// 把int的指针转为unit 类型指针
	p := (*uint)(unsafe.Pointer(&v2))
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(*p)
}

// 结构体的第一个成员地址就是结构体的地址
func TestVisitPrivateField(t *testing.T) {
	v := new(test.Visible)
	t.Logf("init v's val: %s", v.ToString())

	// int64 内存偏移量：64位，8字节
	si64 := uintptr(unsafe.Sizeof(int64(0))) // 8
	t.Logf("size of int64: %d", int(si64))
	// 直接操控内存
	pf2 := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + si64))
	*pf2 = 123
	t.Logf("changed v's val: %s", v.ToString())
}
