package il2cpp

//#include "wrapper/Method.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type Method struct {
	handle C.IppMethod
}

func (m *Method) GetName() string {
	if m.handle == nil {
		return ""
	}

	return C.GoString(C.ippGetMethodName(m.handle))
}

func (m *Method) IsNull() bool {
	return m.handle == nil
}

func (m *Method) GetReturnType() Type {
	return Type{
		handle: C.ippGetMethodReturnType(m.handle),
	}
}

func (m *Method) GetParams() []Type {
	paramCount := C.ippGetMethodParamCount(m.handle)
	params := make([]Type, paramCount)
	for i := 0; i < int(paramCount); i++ {
		params[i] = Type{
			handle: C.ippGetMethodParam(m.handle, C.uint(i)),
		}
	}

	return params
}

func (m *Method) Pointer() uintptr {
	return uintptr(C.ippGetMethodPtr(m.handle))
}

func (m *Method) Invoke(args ...uintptr) (Object, error) {
	var argss unsafe.Pointer
	if len(args) > 0 {
		argss = unsafe.Pointer(&args[0])
	}
	res := C.ippInvokeMethod(m.handle, nil, argss)
	if res == nil {
		return Object{}, fmt.Errorf("failed to invoke method")
	}

	return Object{
		handle: res,
	}, nil
}

func (m *Method) InvokeObject(obj *Object) (*Object, error) {
	res := C.ippInvokeMethod(m.handle, obj.handle, nil)

	return &Object{
		handle: res,
	}, nil
}
