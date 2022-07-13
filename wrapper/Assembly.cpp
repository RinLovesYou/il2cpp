#include "il2cpp.h"

#include "Assembly.h"

IppImage ippGetAssemblyImage(IppAssembly handle)
{
   Il2CppAssembly *assembly = reinterpret_cast<Il2CppAssembly *>(handle);
   return reinterpret_cast<IppImage>(assembly->image);
}

const char* ippGetAssemblyName(IppAssembly handle)
{
    Il2CppAssembly *assembly = reinterpret_cast<Il2CppAssembly *>(handle);
    return assembly->image->nameNoExt;
}