package il2cpp

//#include "il2cpp_structs.h"
import "C"

type Il2CppObject struct {
	object *C.Il2CppObject
	handle uint32
}

func newObject(object *C.Il2CppObject) (*Il2CppObject, error) {
	if object == nil {
		return nil, errNil
	}

	var err error

	obj := &Il2CppObject{object: object}
	obj.handle, err = gcHandleNew(obj, false)

	return obj, err
}

func (o *Il2CppObject) Free() error {
	return gcHandleFree(o.handle)
}
