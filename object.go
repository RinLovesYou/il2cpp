package il2cpp

//#include "wrapper/Object.h"
import "C"
import "unsafe"

type Object struct {
	Handle   C.IppObject
	gchandle uint
}

func NewObject(handle unsafe.Pointer) *Object {
	if handle == nil {
		return nil
	}

	return &Object{
		Handle:   C.IppObject(handle),
		gchandle: uint(C.ippNewGcHandle(C.IppObject(handle), C.IppBool(0))),
	}
}

func (o *Object) getHandle() C.IppObject {
	return C.ippGetGcHandleTarget(C.uint(o.gchandle))
}

func (o *Object) IsNull() bool {
	return o.Handle == nil
}

func (o *Object) Free() {
	if o.Handle == nil {
		return
	}

	C.ippFreeGcHandle(C.uint(o.gchandle))
}

func (o *Object) Unbox() unsafe.Pointer {
	if o.Handle == nil {
		return nil
	}

	return C.ippUnboxObject(o.Handle)
}

func (o *Object) UnboxString() string {
	if o.Handle == nil {
		return ""
	}

	return C.GoString(C.ippUnboxString(o.Handle))
}
