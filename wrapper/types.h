#pragma once

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void *IppDomain;
typedef void *IppAssembly;
typedef void *IppImage;
typedef void* IppClass;
typedef void* IppMethod;
typedef void* IppType;
typedef void* IppObject;
typedef void* IppField;
typedef void* IppProperty;
typedef void* IppString;

typedef int IppBool;

typedef struct {
    float x;
    float y;
    float z;
} IppVector3;

#ifdef __cplusplus
}
#endif