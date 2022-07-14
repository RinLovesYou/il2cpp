#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern uint32_t ippNewGcHandle(IppObject, IppBool);
extern IppObject ippGetGcHandleTarget(uint32_t);
extern void ippFreeGcHandle(uint32_t);
extern const char* ippUnboxString(IppObject);
extern void* ippUnboxObject(IppObject);
extern IppString ippNewString(const char*);

const char* utf8_encode(uint16_t*, int32_t);


#ifdef __cplusplus
}
#endif