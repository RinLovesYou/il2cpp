package il2cpp

//#include "wrapper/Object.h"
import "C"
import "unsafe"

type Object struct {
	handle   C.IppObject
	gchandle uint
}

func NewObject(handle unsafe.Pointer) *Object {
	if handle == nil {
		return nil
	}

	return &Object{
		handle: C.IppObject(handle),
		//gchandle: uint(C.ippNewGcHandle(C.IppObject(handle), C.IppBool(0))),
	}
}

func (o *Object) getHandle() C.IppObject {
	return C.ippGetGcHandleTarget(C.uint(o.gchandle))
}

func (o *Object) IsNull() bool {
	return o.handle == nil
}

func (o *Object) Free() {
	if o.handle == nil {
		return
	}

	C.ippFreeGcHandle(C.uint(o.gchandle))
}

func (o *Object) UnboxString() string {
	if o.handle == nil {
		return ""
	}

	return C.GoString(C.ippUnboxString(o.handle))
}
