#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern const char* ippGetPropertyName(IppProperty);
extern IppMethod ippGetPropertyGetter(IppProperty);
extern IppMethod ippGetPropertySetter(IppProperty);

#ifdef __cplusplus
}
#endif