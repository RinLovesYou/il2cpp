package il2cpp

//#include "il2cpp_structs.h"
import "C"

type Il2CppType struct {
	xType *C.Il2CppType

	Name string
}

func newType(xType *C.Il2CppType) (*Il2CppType, error) {
	if xType == nil {
		return nil, errNil
	}

	var err error

	t := &Il2CppType{xType: xType}
	t.Name, err = t.getName()

	return t, err
}

func (t *Il2CppType) getName() (string, error) {
	return typeGetName(t)
}
