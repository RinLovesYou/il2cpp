package il2cpp

/*
#include <stdlib.h>
#include "il2cpp_structs.h"
*/
import "C"
import (
	"errors"
	"syscall"
	"unsafe"
)

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

	return *(*int)(unsafe.Pointer(ret)), nil
}

func GetDomain() (*Il2CppDomain, error) {
	ret, _, err := il2cpp_domain_get.Call()
	if err != syscall.Errno(0) {
		return nil, err
	}

	if ret == 0 {
		return nil, errors.New("il2cpp_domain_get failed")
	}

	domain := (*C.Il2CppDomain)(unsafe.Pointer(ret))

	return &Il2CppDomain{
		domain,
	}, nil
}
