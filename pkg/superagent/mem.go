package superagent

import (
	"reflect"
	"unsafe"
)

//go:linkname memcpy reflect.typedmemmove
func memcpy(t reflect.Type, dst, src unsafe.Pointer)
