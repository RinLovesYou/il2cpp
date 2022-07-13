package il2cpp

// #cgo CPPFLAGS: -I./wrapper
// #cgo CXXFLAGS: -std=c++11
// #cgo CXXFLAGS: -Wno-subobject-linkage
import "C"

type MethodFlags int32
type BindingFlags int32
type MethodImplFlags int32

const (
	FIELD_STATIC BindingFlags = 16
)
