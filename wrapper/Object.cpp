#include "il2cpp.h"

#include "Object.h"
#include <string>
#include <Windows.h>

uint32_t ippNewGcHandle(IppObject handle, IppBool pinned)
{
   Il2CppObject *obj = reinterpret_cast<Il2CppObject *>(handle);
   return il2cpp_gchandle_new(obj, pinned == 1 ? true : false);
}

IppObject ippGetGcHandleTarget(uint32_t handle)
{
   return reinterpret_cast<IppObject>(il2cpp_gchandle_get_target(handle));
}
void ippFreeGcHandle(uint32_t handle)
{
   il2cpp_gchandle_free(handle);
}

const char *ippUnboxString(IppObject handle)
{
   Il2CppString *str = reinterpret_cast<Il2CppString *>(handle);

   int32_t length = il2cpp_string_length(str);

   if (length == 0)
      return "";

   uint16_t* chars = il2cpp_string_chars(str);
   
   return utf8_encode(chars, length);
}

const char* utf8_encode(uint16_t* wstr, int32_t length)
{
    int size_needed = WideCharToMultiByte(CP_UTF8, 0, (LPCWCH)&wstr[0], (int)length, NULL, 0, NULL, NULL);
    std::string strTo( size_needed, 0 );
    WideCharToMultiByte                  (CP_UTF8, 0, (LPCWCH)&wstr[0], (int)length, &strTo[0], size_needed, NULL, NULL);
    return strTo.c_str();
}