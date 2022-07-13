package il2cpp

//#include "wrapper/Image.h"
import "C"
import "fmt"

type Image struct {
	handle C.IppImage
}

func (i *Image) GetName() string {
	return C.GoString(C.ippGetImageName(i.handle))
}

func (i *Image) GetNameExt() string {
	return C.GoString(C.ippGetImageNameWithExt(i.handle))
}

func (i *Image) GetClassCount() int {
	return int(C.ippGetImageClassCount(i.handle))
}

func (i *Image) GetClasses() []Class {
	count := i.GetClassCount()
	classes := make([]Class, count)
	for j := 0; j < count; j++ {
		classes[j] = i.getClass(j)
	}
	return classes
}

func (i *Image) getClass(index int) Class {
	return Class{
		handle: C.ippGetImageClass(i.handle, C.size_t(index)),
	}
}

func (i *Image) GetClass(name string) (Class, error) {
	return i.GetClassWhere(func(c Class) bool {
		return c.GetName() == name
	})
}

func (i *Image) GetClassWhere(predicate func(Class) bool) (Class, error) {
	classes := i.GetClasses()
	for _, c := range classes {
		if predicate(c) {
			return c, nil
		}
	}
	return Class{}, fmt.Errorf("class not found")
}
