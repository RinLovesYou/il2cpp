package il2cpp

//#include "wrapper/Domain.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type Domain struct {
	handle C.IppDomain

	name       string
	assemblies []Assembly
	images     []Image
}

func GetDomain() *Domain {
	return &Domain{
		handle: C.ippGetDomain(),
	}
}

func (d *Domain) Name() string {
	if d.name != "" {
		return d.name
	}
	d.name = C.GoString(C.ippGetDomainName(d.handle))

	return d.name
}

func (d *Domain) AttachThread() {
	C.ippAttachThread(d.handle)
}

func (d *Domain) GetAssemblies() []Assembly {
	var size C.size_t
	asms := C.ippGetDomainAssemblies(d.handle, &size)

	assemblies := (*[1 << 30]C.IppAssembly)(unsafe.Pointer(asms))[:size:size]

	asms_go := make([]Assembly, size)
	for i, asm := range assemblies {
		asms_go[i] = Assembly(asm)
	}

	return asms_go
}

func (d *Domain) GetImages() []Image {
	assemblies := d.GetAssemblies()
	size := len(assemblies)

	images := make([]Image, size)
	for i, asm := range assemblies {
		images[i] = asm.GetImage()
	}

	return images
}

func (d *Domain) GetImage(name string) (Image, error) {
	return d.GetImageWhere(func(image Image) bool {
		return image.GetName() == name
	})
}

func (d *Domain) GetImageWhere(predicate func(Image) bool) (Image, error) {
	for _, image := range d.GetImages() {
		if predicate(image) {
			return image, nil
		}
	}

	return Image{}, fmt.Errorf("image not found")
}
