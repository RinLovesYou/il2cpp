#include "il2cpp.h"

#include "Image.h"

const char *ippGetImageName(IppImage image)
{
   Il2CppImage *il2cppImage = reinterpret_cast<Il2CppImage *>(image);
   return il2cppImage->nameNoExt;
}

const char *ippGetImageNameWithExt(IppImage image)
{
   Il2CppImage *il2cppImage = reinterpret_cast<Il2CppImage *>(image);
   return il2cppImage->name;
}

size_t ippGetImageClassCount(IppImage image) {
    Il2CppImage *il2cppImage = reinterpret_cast<Il2CppImage *>(image);
    return il2cpp_image_get_class_count(il2cppImage);
}

IppClass ippGetImageClass(IppImage image, size_t index) {
    Il2CppImage *il2cppImage = reinterpret_cast<Il2CppImage *>(image);
    return reinterpret_cast<IppClass>(il2cpp_image_get_class(il2cppImage, index));
}