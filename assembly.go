package il2cpp

//#include "wrapper/Assembly.h"
import "C"

type Assembly uintptr

func (a Assembly) handle() C.IppAssembly {
	return C.IppAssembly(a)
}

func (a Assembly) GetName() string {
	return C.GoString(C.ippGetAssemblyName(a.handle()))
}

func (a Assembly) GetImage() Image {
	handle := C.ippGetAssemblyImage(a.handle())

	return Image{
		handle: handle,
	}
}
