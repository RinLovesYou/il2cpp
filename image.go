package il2cpp

//#include "il2cpp_structs.h"
import "C"

type Il2CppImage struct {
	image *C.Il2CppImage

	Name      string
	NameNoExt string
	Classes   []*Il2CppClass
}

func newImage(image *C.Il2CppImage) (*Il2CppImage, error) {
	if image == nil {
		return nil, errNil
	}

	var err error

	img := &Il2CppImage{image: image}
	img.Name, img.NameNoExt, err = img.getName()
	if err != nil {
		return nil, err
	}
	img.Classes, err = img.getClasses()

	return img, err
}

func (i *Il2CppImage) GetClass(name string) (*Il2CppClass, error) {
	return i.GetClassWhere(func(c *Il2CppClass) bool {
		return c.Name == name
	})
}

func (i *Il2CppImage) GetClassWhere(fn func(*Il2CppClass) bool) (*Il2CppClass, error) {
	for _, class := range i.Classes {
		if class == nil {
			continue
		}

		if class.class == nil {
			continue
		}

		if fn(class) {
			return class, nil
		}
	}

	return nil, errNotFound
}

func (i *Il2CppImage) getName() (name, nameNoExt string, err error) {
	if i.image == nil || i.image.name == nil || i.image.nameNoExt == nil {
		return "", "", errNil
	}

	return C.GoString(i.image.name), C.GoString(i.image.nameNoExt), nil
}

func (i *Il2CppImage) getClasses() ([]*Il2CppClass, error) {
	if i.image == nil {
		return nil, errNil
	}

	classCount, err := imageGetClassCount(i)
	if err != nil {
		return nil, err
	}

	classes := make([]*Il2CppClass, classCount)
	for j := uint64(0); j < classCount; j++ {
		class, err := imageGetClass(i, j)
		if err != nil {
			return nil, err
		}
		classes[j] = class
	}

	return classes, nil
}
