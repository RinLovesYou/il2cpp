package il2cpp

//#include "wrapper/types.h"
import "C"
import "unsafe"

type Vector3 struct {
	v *C.IppVector3
}

func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{
		v: &C.IppVector3{x: C.float(x), y: C.float(y), z: C.float(z)},
	}
}

func Vector3FromPointer(ptr unsafe.Pointer) *Vector3 {
	return &Vector3{v: (*C.IppVector3)(ptr)}
}

func (v *Vector3) Handle() unsafe.Pointer {
	return unsafe.Pointer(v.v)
}

func (v *Vector3) X() float32 {
	return float32(v.v.x)
}

func (v *Vector3) Y() float32 {
	return float32(v.v.y)
}

func (v *Vector3) Z() float32 {
	return float32(v.v.z)
}

func (v *Vector3) Plus(other *Vector3) *Vector3 {
	return NewVector3(v.X()+other.X(), v.Y()+other.Y(), v.Z()+other.Z())
}

func (v *Vector3) Minus(other *Vector3) *Vector3 {
	return NewVector3(v.X()-other.X(), v.Y()-other.Y(), v.Z()-other.Z())
}

func (v *Vector3) Times(other float32) *Vector3 {
	return NewVector3(v.X()*other, v.Y()*other, v.Z()*other)
}

func (v *Vector3) Divide(other *Vector3) *Vector3 {
	return NewVector3(v.X()/other.X(), v.Y()/other.Y(), v.Z()/other.Z())
}
