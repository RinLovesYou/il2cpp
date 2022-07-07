package il2cpp

//#include "il2cpp_structs.h"
import "C"

type Il2CppAssembly struct {
	assembly *C.Il2CppAssembly
}

func newAssembly(assembly *C.Il2CppAssembly) (*Il2CppAssembly, error) {
	if assembly == nil {
		return nil, errNil
	}

	return &Il2CppAssembly{assembly: assembly}, nil
}

func (a *Il2CppAssembly) Name() string {
	if a.assembly == nil {
		return ""
	}

	return C.GoString(a.assembly.aname.name)
}

func (a *Il2CppAssembly) Image() (*Il2CppImage, error) {
	if a.assembly == nil || a.assembly.image == nil {
		return nil, errNil
	}

	return newImage(a.assembly.image)
}
