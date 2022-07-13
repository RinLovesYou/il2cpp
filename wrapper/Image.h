#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern const char* ippGetImageName(IppImage image);
extern const char *ippGetImageNameWithExt(IppImage image);
extern size_t ippGetImageClassCount(IppImage image);
extern IppClass ippGetImageClass(IppImage image, size_t index);

#ifdef __cplusplus
}
#endif