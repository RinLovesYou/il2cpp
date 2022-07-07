package il2cpp

//#include "il2cpp_structs.h"
import "C"
import "unsafe"

type MethodInfo struct {
	method *C.MethodInfo

	Name       string
	Parameters []*ParameterInfo
}

func newMethod(m *C.MethodInfo) (*MethodInfo, error) {
	if m == nil {
		return nil, errNil
	}

	var err error

	method := &MethodInfo{method: m}
	method.Name, err = method.getName()
	if err != nil {
		return nil, err
	}
	method.Parameters = make([]*ParameterInfo, 0)

	paramCount, err := methodGetParamCount(method)
	if err != nil {
		return nil, err
	}
	for i := uint32(0); i < paramCount; i++ {
		param, err := methodGetParam(method, uint32(i))
		if err != nil {
			continue
		}

		method.Parameters = append(method.Parameters, param)
	}

	return method, nil
}

func (m *MethodInfo) getName() (string, error) {
	if m.method == nil || m.method.name == nil {
		return "", errNil
	}

	return C.GoString(m.method.name), nil
}

func (m *MethodInfo) GetPointer() (uintptr, error) {
	if m.method == nil || m.method.methodPointer == nil {
		return 0, errNil
	}

	return uintptr(unsafe.Pointer(m.method.methodPointer)), nil
}
