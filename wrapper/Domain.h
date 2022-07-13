#pragma once

#include "types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IppDomain ippGetDomain();
extern const char* ippGetDomainName(IppDomain handle);
extern const IppAssembly* ippGetDomainAssemblies(IppDomain handle, size_t* size);
extern void* ippAttachThread(IppDomain handle);

#ifdef __cplusplus
}
#endif