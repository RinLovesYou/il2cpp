#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern const char* ippGetMethodName(IppMethod);
extern uint32_t ippGetMethodParamCount(IppMethod);
extern IppType ippGetMethodParam(IppMethod, uint32_t);
extern IppType ippGetMethodReturnType(IppMethod);
extern IppObject ippInvokeMethod(IppMethod, IppObject, void*);
extern void* ippGetMethodPtr(IppMethod);

#ifdef __cplusplus
}
#endif