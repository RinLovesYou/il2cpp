package il2cpp

//#include "wrapper/Field.h"
import "C"

type Field struct {
	handle C.IppField
}

func (f *Field) GetName() string {
	if f.handle == nil {
		return ""
	}
	return C.GoString(C.ippGetFieldName(f.handle))
}

func (f *Field) GetFlags() BindingFlags {
	if f.handle == nil {
		return 0
	}
	return BindingFlags(C.ippGetFieldFlags(f.handle))
}

func (f *Field) HasFlag(flag BindingFlags) bool {
	return f.GetFlags()&flag != 0
}

func (f *Field) GetValue() *Object {
	return f.GetValueObject(&Object{})
}

func (f *Field) GetValueObject(o *Object) *Object {
	var return_val C.IppObject

	if f.HasFlag(FIELD_STATIC) {
		return_val = C.ippGetFieldValueObject(f.handle, nil)
	} else {
		return_val = C.ippGetFieldValueObject(f.handle, o.handle)
	}

	return &Object{handle: return_val}
}
