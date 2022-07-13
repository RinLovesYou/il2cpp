#pragma once

#include "structs.h"

#define WIN32_LEAN_AND_MEAN
#include <Windows.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef Il2CppDomain *(*Il2CppDomainGet)();
typedef Il2CppAssembly **(*Il2CppDomainGetAssembies)(Il2CppDomain *, size_t *);
typedef size_t (*Il2CppImageGetClassCount)(Il2CppImage*);
typedef Il2CppClass *(*Il2CppImageGetClass)(Il2CppImage*, size_t);
typedef int (*Il2CppClassGetFlags)(Il2CppClass*);
typedef Il2CppClass *(*Il2CppClassGetParent)(Il2CppClass*);
typedef bool (*Il2CppClassIsInterface)(Il2CppClass*);
typedef MethodInfo *(*Il2CppClassGetMethods)(Il2CppClass*, void*);
typedef uint32_t (*Il2CppMethodGetParamCount)(MethodInfo*);
typedef Il2CppType *(*Il2CppMethodGetParam)(MethodInfo*, uint32_t);
typedef Il2CppType *(*Il2CppMethodGetReturnType)(MethodInfo*);
typedef char* (*Il2CppTypeGetName)(Il2CppType*);
typedef void* (*Il2CppThreadAttach)(Il2CppDomain*);
typedef uint32_t (*Il2CppGcHandleNew)(Il2CppObject*, bool);
typedef Il2CppObject* (*Il2CppGcHandleGetTarget)(uint32_t);
typedef void (*Il2CppGcHandleFree)(uint32_t);
typedef FieldInfo* (*Il2CppClassGetFields)(Il2CppClass*, void*);
typedef int (*Il2CppFieldGetFlags)(FieldInfo*);
typedef Il2CppObject* (*Il2CppFieldGetValueObject)(FieldInfo*, Il2CppObject*);
typedef int32_t (*Il2CppStringLength)(Il2CppString*);
typedef Il2CppChar* (*Il2CppStringGetChars)(Il2CppString*);
typedef PropertyInfo* (*Il2CppClassGetProperties)(Il2CppClass*, void*);
typedef MethodInfo* (*Il2CppPropertyGetGetMethod)(PropertyInfo*);
typedef MethodInfo* (*Il2CppPropertyGetSetMethod)(PropertyInfo*);
typedef Il2CppObject* (*Il2CppRuntimeInvoke)(MethodInfo*, void*, void**, Il2CppException **exc);

const HMODULE GameAssembly = LoadLibraryA("GameAssembly.dll");

const Il2CppDomainGet il2cpp_domain_get = (Il2CppDomainGet)GetProcAddress(GameAssembly, "il2cpp_domain_get");
const Il2CppDomainGetAssembies il2cpp_domain_get_assemblies = (Il2CppDomainGetAssembies)GetProcAddress(GameAssembly, "il2cpp_domain_get_assemblies");
const Il2CppImageGetClassCount il2cpp_image_get_class_count = (Il2CppImageGetClassCount)GetProcAddress(GameAssembly, "il2cpp_image_get_class_count");
const Il2CppImageGetClass il2cpp_image_get_class = (Il2CppImageGetClass)GetProcAddress(GameAssembly, "il2cpp_image_get_class");
const Il2CppClassGetFlags il2cpp_class_get_flags = (Il2CppClassGetFlags)GetProcAddress(GameAssembly, "il2cpp_class_get_flags");
const Il2CppClassGetParent il2cpp_class_get_parent = (Il2CppClassGetParent)GetProcAddress(GameAssembly, "il2cpp_class_get_parent");
const Il2CppClassIsInterface il2cpp_class_is_interface = (Il2CppClassIsInterface)GetProcAddress(GameAssembly, "il2cpp_class_is_interface");
const Il2CppClassGetMethods il2cpp_class_get_methods = (Il2CppClassGetMethods)GetProcAddress(GameAssembly, "il2cpp_class_get_methods");
const Il2CppMethodGetParamCount il2cpp_method_get_param_count = (Il2CppMethodGetParamCount)GetProcAddress(GameAssembly, "il2cpp_method_get_param_count");
const Il2CppMethodGetParam il2cpp_method_get_param = (Il2CppMethodGetParam)GetProcAddress(GameAssembly, "il2cpp_method_get_param");
const Il2CppMethodGetReturnType il2cpp_method_get_return_type = (Il2CppMethodGetReturnType)GetProcAddress(GameAssembly, "il2cpp_method_get_return_type");
const Il2CppTypeGetName il2cpp_type_get_name = (Il2CppTypeGetName)GetProcAddress(GameAssembly, "il2cpp_type_get_name");
const Il2CppThreadAttach il2cpp_thread_attach = (Il2CppThreadAttach)GetProcAddress(GameAssembly, "il2cpp_thread_attach");
const Il2CppGcHandleNew il2cpp_gchandle_new = (Il2CppGcHandleNew)GetProcAddress(GameAssembly, "il2cpp_gchandle_new");
const Il2CppGcHandleGetTarget il2cpp_gchandle_get_target = (Il2CppGcHandleGetTarget)GetProcAddress(GameAssembly, "il2cpp_gchandle_get_target");
const Il2CppGcHandleFree il2cpp_gchandle_free = (Il2CppGcHandleFree)GetProcAddress(GameAssembly, "il2cpp_gchandle_free");
const Il2CppClassGetFields il2cpp_class_get_fields = (Il2CppClassGetFields)GetProcAddress(GameAssembly, "il2cpp_class_get_fields");
const Il2CppFieldGetFlags il2cpp_field_get_flags = (Il2CppFieldGetFlags)GetProcAddress(GameAssembly, "il2cpp_field_get_flags");
const Il2CppFieldGetValueObject il2cpp_field_get_value_object = (Il2CppFieldGetValueObject)GetProcAddress(GameAssembly, "il2cpp_field_get_value_object");
const Il2CppStringLength il2cpp_string_length = (Il2CppStringLength)GetProcAddress(GameAssembly, "il2cpp_string_length");
const Il2CppStringGetChars il2cpp_string_chars = (Il2CppStringGetChars)GetProcAddress(GameAssembly, "il2cpp_string_chars");
const Il2CppClassGetProperties il2cpp_class_get_properties = (Il2CppClassGetProperties)GetProcAddress(GameAssembly, "il2cpp_class_get_properties");
const Il2CppPropertyGetGetMethod il2cpp_property_get_get_method = (Il2CppPropertyGetGetMethod)GetProcAddress(GameAssembly, "il2cpp_property_get_get_method");
const Il2CppPropertyGetSetMethod il2cpp_property_get_set_method = (Il2CppPropertyGetSetMethod)GetProcAddress(GameAssembly, "il2cpp_property_get_set_method");
const Il2CppRuntimeInvoke il2cpp_runtime_invoke = (Il2CppRuntimeInvoke)GetProcAddress(GameAssembly, "il2cpp_runtime_invoke");

#ifdef __cplusplus
}
#endif