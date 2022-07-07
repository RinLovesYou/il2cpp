package il2cpp

//#include "il2cpp_structs.h"
import "C"
import "unsafe"

type Il2CppClass struct {
	class   *C.Il2CppClass
	methods []*MethodInfo

	Name string
}

func newClass(c *C.Il2CppClass) (*Il2CppClass, error) {
	if c == nil {
		return nil, errNil
	}

	var err error

	class := &Il2CppClass{class: c}
	class.Name, err = class.getName()

	return class, err
}

func (c *Il2CppClass) getName() (string, error) {
	if c.class == nil || c.class.name == nil {
		return "", errNil
	}

	return C.GoString(c.class.name), nil
}

func (c *Il2CppClass) GetMethods() ([]*MethodInfo, error) {
	if c.methods != nil {
		return c.methods, nil
	}

	if c.class == nil {
		return nil, errNil
	}

	c.methods = make([]*MethodInfo, 0)

	methodCount := c.class.method_count
	if methodCount == 0 || c.class.methods == nil {
		return nil, errNil
	}
	methodsC := (*[1 << 30]*C.MethodInfo)(unsafe.Pointer(c.class.methods))[:methodCount:methodCount]

	for _, method := range methodsC {
		if method == nil {
			continue
		}

		methodInfo, err := newMethod(method)
		if err != nil {
			return nil, err
		}

		c.methods = append(c.methods, methodInfo)
	}

	return c.methods, nil
}

func (c *Il2CppClass) HasMethod(name string) bool {
	_, err := c.GetMethod(name)
	return err == nil
}

func (c *Il2CppClass) GetMethod(name string) (*MethodInfo, error) {
	return c.GetMethodWhere(func(method *MethodInfo) bool {
		return method.Name == name
	})
}

func (c *Il2CppClass) GetMethodWhere(fn func(*MethodInfo) bool) (*MethodInfo, error) {
	methods, err := c.GetMethods()
	if err != nil {
		return nil, err
	}

	for _, method := range methods {
		if method != nil && method.method != nil {
			if fn(method) {
				return method, nil
			}
		}
	}

	return nil, errNotFound
}
