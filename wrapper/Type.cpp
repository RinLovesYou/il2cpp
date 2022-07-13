#include "il2cpp.h"

#include "Type.h"

char* ippGetTypeName(IppType handle) {
    Il2CppType *type = reinterpret_cast<Il2CppType *>(handle);
    return il2cpp_type_get_name(type);
}