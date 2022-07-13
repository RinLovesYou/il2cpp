#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern const char* ippGetClassName(IppClass);
extern const char* ippGetClassNamespace(IppClass);
extern int ippGetClassFlags(IppClass);
extern IppClass ippGetClassParent(IppClass);
extern IppBool ippIsClassInterface(IppClass);
extern IppMethod ippGetClassMethods(IppClass, size_t*);
extern IppField ippGetClassFields(IppClass, size_t*);
extern IppProperty ippGetClassProperties(IppClass handle, size_t* iter);

#ifdef __cplusplus
}
#endif