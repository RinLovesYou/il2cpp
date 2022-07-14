package il2cpp

//#include "wrapper/Object.h"
import "C"

type String struct {
	Handle   C.IppString
	gchandle uint
}

func NewString(stringe string) *String {
	str := C.ippNewString(C.CString(stringe))
	return &String{
		Handle:   str,
		gchandle: uint(C.ippNewGcHandle(C.IppObject(str), C.IppBool(0))),
	}
}

func (s *String) Free() {
	if s.Handle == nil {
		return
	}

	C.ippFreeGcHandle(C.uint(s.gchandle))
}
