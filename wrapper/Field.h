#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern const char* ippGetFieldName(IppField);
extern int ippGetFieldFlags(IppField);
extern IppObject ippGetFieldValueObject(IppField, IppObject);

#ifdef __cplusplus
}
#endif