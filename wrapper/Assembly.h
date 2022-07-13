#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IppImage ippGetAssemblyImage(IppAssembly);
extern const char* ippGetAssemblyName(IppAssembly);

#ifdef __cplusplus
}
#endif