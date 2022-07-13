package il2cpp

//#include "wrapper/Type.h"
import "C"

type Type struct {
	handle C.IppType
}

func (t Type) GetName() string {
	if t.handle == nil {
		return ""
	}

	return C.GoString(C.ippGetTypeName(t.handle))
}
