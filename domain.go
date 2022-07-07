package il2cpp

//#include "il2cpp_structs.h"
import "C"

type Il2CppDomain struct {
	domain *C.Il2CppDomain

	Name   string
	Images []*Il2CppImage
}

func newDomain(d *C.Il2CppDomain) (*Il2CppDomain, error) {
	if d == nil {
		return nil, errNil
	}

	var err error

	domain := &Il2CppDomain{domain: d}
	domain.Name, err = domain.getName()
	if err != nil {
		return nil, err
	}
	domain.Images, err = domain.getImages()

	return domain, err
}

func (d *Il2CppDomain) ThreadAttach() error {
	return threadAttach(d)
}

func (d *Il2CppDomain) GetImage(name string) (*Il2CppImage, error) {
	return d.GetImageWhere(func(i *Il2CppImage) bool {
		return i.NameNoExt == name
	})
}

func (d *Il2CppDomain) GetImageWhere(fn func(i *Il2CppImage) bool) (*Il2CppImage, error) {
	for _, image := range d.Images {
		if image != nil && image.image != nil {
			if fn(image) {
				return image, nil
			}
		}
	}

	return nil, errNotFound
}

func (d *Il2CppDomain) getName() (string, error) {
	if d.domain == nil || d.domain.friendly_name == nil {
		return "", errNil
	}

	return C.GoString(d.domain.friendly_name), nil
}

func (d *Il2CppDomain) getImages() ([]*Il2CppImage, error) {
	d.ThreadAttach()

	assemblyCount := uint64(0)
	assemblies, err := domainGetAssemblies(d, &assemblyCount)
	if err != nil {
		return nil, err
	}

	images := make([]*Il2CppImage, assemblyCount)
	for i, assembly := range assemblies {
		if assembly == nil || assembly.assembly.image == nil {
			continue
		}

		images[i], err = newImage(assembly.assembly.image)
		if err != nil {
			return nil, err
		}
	}

	return images, nil
}
