package il2cpp

/*
#include <stdlib.h>
#include <stdbool.h>
#include "il2cpp_structs.h"
*/
import "C"
import (
	"errors"
	"syscall"
	"unsafe"
)

var errNil = errors.New("nil")
var errNotFound = errors.New("not found")

var (
	gameAssembly = syscall.NewLazyDLL("GameAssembly.dll")

	il2cpp_init                  = gameAssembly.NewProc("il2cpp_init")
	il2cpp_domain_get            = gameAssembly.NewProc("il2cpp_domain_get")
	il2cpp_domain_get_assemblies = gameAssembly.NewProc("il2cpp_domain_get_assemblies")

	il2cpp_thread_attach = gameAssembly.NewProc("il2cpp_thread_attach")
	il2cpp_thread_detach = gameAssembly.NewProc("il2cpp_thread_detach")

	il2cpp_format_exception = gameAssembly.NewProc("il2cpp_format_exception")

	il2cpp_resolve_icall = gameAssembly.NewProc("il2cpp_resolve_icall")

	il2cpp_array_length       = gameAssembly.NewProc("il2cpp_array_length")
	il2cpp_array_new          = gameAssembly.NewProc("il2cpp_array_new")
	il2cpp_array_element_size = gameAssembly.NewProc("il2cpp_array_element_size")
	il2cpp_array_get          = gameAssembly.NewProc("il2cpp_array_get")

	il2cpp_assembly_get_image = gameAssembly.NewProc("il2cpp_assembly_get_image")

	il2cpp_class_get_flags          = gameAssembly.NewProc("il2cpp_class_get_flags")
	il2cpp_class_from_il2cpp_type   = gameAssembly.NewProc("il2cpp_class_from_il2cpp_type")
	il2cpp_class_from_system_type   = gameAssembly.NewProc("il2cpp_class_from_system_type")
	il2cpp_class_get_events         = gameAssembly.NewProc("il2cpp_class_get_events")
	il2cpp_class_get_fields         = gameAssembly.NewProc("il2cpp_class_get_fields")
	il2cpp_class_get_methods        = gameAssembly.NewProc("il2cpp_class_get_methods")
	il2cpp_class_get_properties     = gameAssembly.NewProc("il2cpp_class_get_properties")
	il2cpp_class_get_nested_types   = gameAssembly.NewProc("il2cpp_class_get_nested_types")
	il2cpp_class_get_interfaces     = gameAssembly.NewProc("il2cpp_class_get_interfaces")
	il2cpp_class_get_parent         = gameAssembly.NewProc("il2cpp_class_get_parent")
	il2cpp_class_get_declaring_type = gameAssembly.NewProc("il2cpp_class_get_declaring_type")
	il2cpp_class_get_name           = gameAssembly.NewProc("il2cpp_class_get_name")
	il2cpp_class_get_namespace      = gameAssembly.NewProc("il2cpp_class_get_namespace")
	il2cpp_class_is_interface       = gameAssembly.NewProc("il2cpp_class_is_interface")
	il2cpp_class_from_type          = gameAssembly.NewProc("il2cpp_class_from_type")
	il2cpp_class_get_type           = gameAssembly.NewProc("il2cpp_class_get_type")
	il2cpp_class_get_type_token     = gameAssembly.NewProc("il2cpp_class_get_type_token")
	il2cpp_class_is_enum            = gameAssembly.NewProc("il2cpp_class_is_enum")

	il2cpp_field_get_flags        = gameAssembly.NewProc("il2cpp_field_get_flags")
	il2cpp_field_get_name         = gameAssembly.NewProc("il2cpp_field_get_name")
	il2cpp_field_get_type         = gameAssembly.NewProc("il2cpp_field_get_type")
	il2cpp_field_get_value_object = gameAssembly.NewProc("il2cpp_field_get_value_object")
	il2cpp_field_set_value        = gameAssembly.NewProc("il2cpp_field_set_value")
	il2cpp_field_static_set_value = gameAssembly.NewProc("il2cpp_field_static_set_value")

	il2cpp_gchandle_new        = gameAssembly.NewProc("il2cpp_gchandle_new")
	il2cpp_gchandle_free       = gameAssembly.NewProc("il2cpp_gchandle_free")
	il2cpp_gchandle_get_target = gameAssembly.NewProc("il2cpp_gchandle_get_target")

	il2cpp_method_get_return_type = gameAssembly.NewProc("il2cpp_method_get_return_type")
	il2cpp_method_get_name        = gameAssembly.NewProc("il2cpp_method_get_name")
	il2cpp_method_get_param_count = gameAssembly.NewProc("il2cpp_method_get_param_count")
	il2cpp_method_get_param       = gameAssembly.NewProc("il2cpp_method_get_param")
	il2cpp_method_get_flags       = gameAssembly.NewProc("il2cpp_method_get_flags")
	il2cpp_method_get_param_name  = gameAssembly.NewProc("il2cpp_method_get_param_name")

	il2cpp_property_get_flags      = gameAssembly.NewProc("il2cpp_property_get_flags")
	il2cpp_property_get_get_method = gameAssembly.NewProc("il2cpp_property_get_get_method")
	il2cpp_property_get_set_method = gameAssembly.NewProc("il2cpp_property_get_set_method")
	il2cpp_property_get_name       = gameAssembly.NewProc("il2cpp_property_get_name")

	il2cpp_object_get_class = gameAssembly.NewProc("il2cpp_object_get_class")
	il2cpp_object_get_size  = gameAssembly.NewProc("il2cpp_object_get_size")
	il2cpp_object_new       = gameAssembly.NewProc("il2cpp_object_new")
	il2cpp_object_unbox     = gameAssembly.NewProc("il2cpp_object_unbox")
	il2cpp_value_box        = gameAssembly.NewProc("il2cpp_value_box")

	il2cpp_runtime_invoke      = gameAssembly.NewProc("il2cpp_runtime_invoke")
	il2cpp_runtime_object_init = gameAssembly.NewProc("il2cpp_runtime_object_init")
	il2cpp_runtime_class_init  = gameAssembly.NewProc("il2cpp_runtime_class_init")

	il2cpp_string_length = gameAssembly.NewProc("il2cpp_string_length")
	il2cpp_string_chars  = gameAssembly.NewProc("il2cpp_string_chars")
	il2cpp_string_new    = gameAssembly.NewProc("il2cpp_string_new")

	il2cpp_type_get_object = gameAssembly.NewProc("il2cpp_type_get_object")
	il2cpp_type_get_name   = gameAssembly.NewProc("il2cpp_type_get_name")
	il2cpp_type_get_attrs  = gameAssembly.NewProc("il2cpp_type_get_attrs")
	il2cpp_type_equals     = gameAssembly.NewProc("il2cpp_type_equals")
	il2cpp_type_get_type   = gameAssembly.NewProc("il2cpp_type_get_type")

	il2cpp_image_get_name        = gameAssembly.NewProc("il2cpp_image_get_name")
	il2cpp_image_get_class       = gameAssembly.NewProc("il2cpp_image_get_class")
	il2cpp_image_get_class_count = gameAssembly.NewProc("il2cpp_image_get_class_count")
)

func Init(domainName string) (int, error) {
	cString := C.CString(domainName)
	defer C.free(unsafe.Pointer(cString))

	ret, _, err := il2cpp_init.Call(
		uintptr(unsafe.Pointer(cString)),
	)

	if err != syscall.Errno(0) {
		return 0, err
	}

	if ret <= 0 {
		return 0, errors.New("il2cpp_init failed")
	}

	return *(*int)(unsafe.Pointer(&ret)), nil
}

var domain *Il2CppDomain

func GetDomain() (*Il2CppDomain, error) {
	if domain != nil {
		return domain, nil
	}

	ret, _, err := il2cpp_domain_get.Call()
	if err != syscall.Errno(0) {
		return nil, err
	}

	if ret == 0 {
		return nil, errors.New("il2cpp_domain_get failed")
	}

	domain, err = newDomain((*C.Il2CppDomain)(unsafe.Pointer(ret)))

	return domain, err
}

func GetImage(name string) (*Il2CppImage, error) {
	domain, err := GetDomain()
	if err != nil {
		return nil, err
	}

	return domain.GetImage(name)
}

func threadAttach(d *Il2CppDomain) error {
	if d.domain == nil {
		return errNil
	}

	ret, _, err := il2cpp_thread_attach.Call(uintptr(unsafe.Pointer(d.domain)))
	if err != syscall.Errno(0) {
		return err
	}

	if ret == 0 {
		return errors.New("il2cpp_thread_attach failed")
	}

	return nil
}

func domainGetAssemblies(d *Il2CppDomain, assemblyCount *uint64) ([]*Il2CppAssembly, error) {
	if d.domain == nil {
		return nil, errNil
	}

	ret, _, err := il2cpp_domain_get_assemblies.Call(
		uintptr(unsafe.Pointer(d.domain)),
		uintptr(unsafe.Pointer(assemblyCount)),
	)

	if err != syscall.Errno(0) {
		return nil, err
	}

	if ret == 0 {
		return nil, errors.New("il2cpp_domain_get_assemblies failed")
	}

	assembliesPtr := (**C.Il2CppAssembly)(unsafe.Pointer(ret))

	if *assemblyCount == 0 || assembliesPtr == nil {
		return nil, errNil
	}

	assembliesC := (*[1 << 30]*C.Il2CppAssembly)(unsafe.Pointer(assembliesPtr))[:*assemblyCount:*assemblyCount]

	assemblies := make([]*Il2CppAssembly, *assemblyCount)
	for i, assembly := range assembliesC {
		assemblies[i], err = newAssembly(assembly)
		if err != nil {
			return nil, err
		}
	}

	return assemblies, nil
}

func imageGetClassCount(image *Il2CppImage) (uint64, error) {
	if image.image == nil {
		return 0, errNil
	}

	ret, _, err := il2cpp_image_get_class_count.Call(
		uintptr(unsafe.Pointer(image.image)),
	)
	if err != syscall.Errno(0) {
		return 0, err
	}

	if ret == 0 {
		return 0, errors.New("il2cpp_image_get_class_count failed")
	}

	return uint64(ret), nil
}

func imageGetClass(image *Il2CppImage, index uint64) (*Il2CppClass, error) {
	if image.image == nil {
		return nil, errNil
	}

	ret, _, err := il2cpp_image_get_class.Call(
		uintptr(unsafe.Pointer(image.image)),
		uintptr(index),
	)
	if err != syscall.Errno(0) {
		return nil, err
	}

	if ret == 0 {
		return nil, errors.New("il2cpp_image_get_class failed")
	}

	return newClass((*C.Il2CppClass)(unsafe.Pointer(ret)))
}

func gcHandleNew(obj *Il2CppObject, pinned bool) (uint32, error) {
	if obj.object == nil {
		return 0, errNil
	}

	pinnedPtr := uintptr(0)
	if pinned {
		pinnedPtr = 1
	}

	ret, _, err := il2cpp_gchandle_new.Call(
		uintptr(unsafe.Pointer(obj.object)),
		pinnedPtr,
	)
	if err != syscall.Errno(0) {
		return 0, err
	}

	if ret == 0 {
		return 0, errors.New("il2cpp_gc_handle_new failed")
	}

	return uint32(ret), nil
}

func gcHandleFree(handle uint32) error {
	_, _, err := il2cpp_gchandle_free.Call(uintptr(handle))
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

func classGetMethods(klass *Il2CppClass, iter *C.intptr_t) (*MethodInfo, error) {
	if klass.class == nil || iter == nil {
		return nil, errNil
	}

	ret, _, err := il2cpp_class_get_methods.Call(
		uintptr(unsafe.Pointer(klass.class)),
		uintptr(unsafe.Pointer(iter)),
	)
	if err != syscall.Errno(0) {
		return nil, err
	}

	if ret == 0 {
		return nil, errors.New("il2cpp_class_get_methods failed")
	}

	return newMethod((*C.MethodInfo)(unsafe.Pointer(ret)))
}

func typeGetName(t *Il2CppType) (string, error) {
	if t.xType == nil {
		return "", errNil
	}

	ret, _, err := il2cpp_type_get_name.Call(
		uintptr(unsafe.Pointer(t.xType)),
	)

	if err != syscall.Errno(0) {
		return "", err
	}

	if ret == 0 {
		return "", errors.New("il2cpp_type_get_name failed")
	}

	return C.GoString((*C.char)(unsafe.Pointer(ret))), nil
}

func paramGetTypeName(p *ParameterInfo) (string, error) {
	if p.parameter == nil {
		return "", errNil
	}

	ret, _, err := il2cpp_type_get_name.Call(
		uintptr(unsafe.Pointer(p.parameter)),
	)

	if err != syscall.Errno(0) {
		return "", err
	}

	if ret == 0 {
		return "", errors.New("il2cpp_type_get_name failed")
	}

	return C.GoString((*C.char)(unsafe.Pointer(ret))), nil
}

func methodGetParam(method *MethodInfo, index uint32) (*ParameterInfo, error) {
	if method.method == nil {
		return nil, errNil
	}

	ret, _, err := il2cpp_method_get_param.Call(
		uintptr(unsafe.Pointer(method.method)),
		uintptr(index),
	)

	if err != syscall.Errno(0) {
		return nil, err
	}

	if ret == 0 {
		return nil, errors.New("il2cpp_method_get_param failed")
	}

	name, err := methodGetParamName(method, index)
	if err != nil {
		return nil, err
	}

	return newParam((*C.ParameterInfo)(unsafe.Pointer(ret)), name)
}

func methodGetParamCount(method *MethodInfo) (uint32, error) {
	if method.method == nil {
		return 0, errNil
	}

	ret, _, err := il2cpp_method_get_param_count.Call(
		uintptr(unsafe.Pointer(method.method)),
	)

	if err != syscall.Errno(0) {
		return 0, err
	}

	return uint32(ret), nil
}

func methodGetParamName(method *MethodInfo, index uint32) (string, error) {
	if method.method == nil {
		return "", errNil
	}

	ret, _, err := il2cpp_method_get_param_name.Call(
		uintptr(unsafe.Pointer(method.method)),
		uintptr(index),
	)

	if err != syscall.Errno(0) {
		return "", err
	}

	if ret == 0 {
		return "", errors.New("il2cpp_method_get_param_name failed")
	}

	return C.GoString((*C.char)(unsafe.Pointer(ret))), nil
}
