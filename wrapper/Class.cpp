#include "il2cpp.h"

#include "Class.h"
#include <stdio.h>

const char* ippGetClassName(IppClass handle) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return class_->name;
}

const char* ippGetClassNamespace(IppClass handle) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return class_->namespaze;
}

int ippGetClassFlags(IppClass handle) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return il2cpp_class_get_flags(class_);
}

IppClass ippGetClassParent(IppClass handle) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return reinterpret_cast<IppClass>(il2cpp_class_get_parent(class_));
}

IppBool ippIsClassInterface(IppClass handle) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return il2cpp_class_is_interface(class_) ? 1 : 0;
}

IppMethod ippGetClassMethods(IppClass handle, size_t* iter) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return reinterpret_cast<IppMethod>(il2cpp_class_get_methods(class_, reinterpret_cast<void*>(iter)));
}

IppField ippGetClassFields(IppClass handle, size_t* iter) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return reinterpret_cast<IppField>(il2cpp_class_get_fields(class_, reinterpret_cast<void*>(iter)));
}

IppProperty ippGetClassProperties(IppClass handle, size_t* iter) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    return reinterpret_cast<IppProperty>(il2cpp_class_get_properties(class_, reinterpret_cast<void*>(iter)));
}

IppObject ippGetClassTypeObject(IppClass handle) {
    Il2CppClass *class_ = reinterpret_cast<Il2CppClass *>(handle);
    Il2CppType* type = il2cpp_class_get_type(class_);
    return reinterpret_cast<IppObject>(il2cpp_type_get_object(type));
}