package test

import "fmt"

type Visible struct {
	privateField1 int64
	privateField2 int32
	PublicField1  string
}

func (v *Visible) ToString() string {
	return fmt.Sprintf("private_field1: %d, private_field2: %d, pubilc_field1: %s",
		v.privateField1, v.privateField2, v.PublicField1)
}
