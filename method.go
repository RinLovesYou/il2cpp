package il2cpp

//#include "il2cpp_structs.h"
import "C"

type MethodInfo struct {
	method *C.MethodInfo

	Name string
}

func newMethod(m *C.MethodInfo) (*MethodInfo, error) {
	if m == nil {
		return nil, errNil
	}

	var err error

	method := &MethodInfo{method: m}
	method.Name, err = method.getName()

	return method, err
}

func (m *MethodInfo) getName() (string, error) {
	if m.method == nil || m.method.name == nil {
		return "", errNil
	}

	return C.GoString(m.method.name), nil
}
