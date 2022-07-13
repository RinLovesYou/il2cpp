#include "il2cpp.h"

#include "Method.h"

const char *ippGetMethodName(IppMethod handle)
{
   MethodInfo *method = reinterpret_cast<MethodInfo *>(handle);
   return method->name;
}

uint32_t ippGetMethodParamCount(IppMethod handle)
{
   MethodInfo *method = reinterpret_cast<MethodInfo *>(handle);
   return il2cpp_method_get_param_count(method);
}

IppType ippGetMethodParam(IppMethod handle, uint32_t index)
{
   MethodInfo *method = reinterpret_cast<MethodInfo *>(handle);
   return reinterpret_cast<IppType>(il2cpp_method_get_param(method, index));
}

IppType ippGetMethodReturnType(IppMethod handle)
{
   MethodInfo *method = reinterpret_cast<MethodInfo *>(handle);
   return reinterpret_cast<IppType>(il2cpp_method_get_return_type(method));
}

IppObject ippInvokeMethod(IppMethod handle, IppObject obj, void *params)
{
   Il2CppException *exception = NULL;
   MethodInfo *method = reinterpret_cast<MethodInfo *>(handle);

    void **parameters = reinterpret_cast<void **>(params);

   return reinterpret_cast<IppObject>(il2cpp_runtime_invoke(method, (void*)obj, parameters, nullptr));
}

void *ippGetMethodPtr(IppMethod handle)
{
   MethodInfo *method = reinterpret_cast<MethodInfo *>(handle);
   return reinterpret_cast<void *>(method->methodPointer);
}