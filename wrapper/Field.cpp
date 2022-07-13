#include "il2cpp.h"

#include "Field.h"

const char* ippGetFieldName(IppField handle) {
    FieldInfo *field_ = reinterpret_cast<FieldInfo *>(handle);
    return field_->name;
}

int ippGetFieldFlags(IppField handle) {
    FieldInfo *field_ = reinterpret_cast<FieldInfo *>(handle);
    return il2cpp_field_get_flags(field_);
}

IppObject ippGetFieldValueObject(IppField handle, IppObject obj) {
    FieldInfo *field_ = reinterpret_cast<FieldInfo *>(handle);
    return reinterpret_cast<IppObject>(il2cpp_field_get_value_object(field_, reinterpret_cast<Il2CppObject*>(obj)));
}