#include "il2cpp.h"

#include "Property.h"

const char* ippGetPropertyName(IppProperty handle) {
    PropertyInfo *property_ = reinterpret_cast<PropertyInfo *>(handle);
    return property_->name;
}

extern IppMethod ippGetPropertyGetter(IppProperty handle) {
    PropertyInfo *property_ = reinterpret_cast<PropertyInfo *>(handle);
    return reinterpret_cast<IppMethod>(property_->get);
}
extern IppMethod ippGetPropertySetter(IppProperty handle) {
    PropertyInfo *property_ = reinterpret_cast<PropertyInfo *>(handle);
    return reinterpret_cast<IppMethod>(il2cpp_property_get_set_method(property_));
}