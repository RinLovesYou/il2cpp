#include "il2cpp.h"

#include "Domain.h"

IppDomain ippGetDomain()
{
    return reinterpret_cast<IppDomain>(il2cpp_domain_get());
}

const char* ippGetDomainName(IppDomain handle)
{
    Il2CppDomain* domain = reinterpret_cast<Il2CppDomain*>(handle);
    return domain->friendly_name;
}

const IppAssembly* ippGetDomainAssemblies(IppDomain handle, size_t* size) {
    Il2CppDomain* domain = reinterpret_cast<Il2CppDomain*>(handle);
    return reinterpret_cast<const IppAssembly*>(il2cpp_domain_get_assemblies(domain, size));
}

void* ippAttachThread(IppDomain handle) {
    Il2CppDomain* domain = reinterpret_cast<Il2CppDomain*>(handle);
    return reinterpret_cast<void*>(il2cpp_thread_attach(domain));
}