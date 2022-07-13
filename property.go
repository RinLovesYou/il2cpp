package il2cpp

//#include "wrapper/Property.h"
import "C"

type Property struct {
	handle C.IppProperty
}

func (p *Property) GetName() string {
	if p.handle == nil {
		return ""
	}
	return C.GoString(C.ippGetPropertyName(p.handle))
}

func (p *Property) GetGet() *Method {
	if p.handle == nil {
		return &Method{}
	}
	return &Method{handle: C.ippGetPropertyGetter(p.handle)}
}

func (p *Property) GetSet() *Method {
	if p.handle == nil {
		return &Method{}
	}
	return &Method{handle: C.ippGetPropertySetter(p.handle)}
}
