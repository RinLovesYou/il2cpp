package il2cpp

//#include "wrapper/Class.h"
import "C"
import "fmt"

type Class struct {
	handle C.IppClass
}

func (c *Class) GetName() string {
	if c.handle == nil {
		return ""
	}
	return C.GoString(C.ippGetClassName(c.handle))
}

func (c *Class) GetNamespace() string {
	if c.handle == nil {
		return ""
	}
	return C.GoString(C.ippGetClassNamespace(c.handle))
}

func (c *Class) TypeObject() *Object {
	return &Object{
		Handle: C.ippGetClassTypeObject(c.handle),
	}
}

func (c *Class) GetFlags() BindingFlags {
	if c.handle == nil {
		return 0
	}
	return BindingFlags(C.ippGetClassFlags(c.handle))
}

func (c *Class) GetParent() Class {
	return Class{
		handle: C.ippGetClassParent(c.handle),
	}
}

func (c *Class) IsInterface() bool {
	return C.ippIsClassInterface(c.handle) == 1
}

func (c *Class) GetMethods() []Method {
	methods := make([]Method, 0)
	methodIter := C.size_t(0)
	method := C.ippGetClassMethods(c.handle, &methodIter)
	for method != nil {
		methods = append(methods, Method{handle: method})
		method = C.ippGetClassMethods(c.handle, &methodIter)
	}

	return methods
}

func (c *Class) GetFields() []Field {
	fields := make([]Field, 0)
	fieldIter := C.size_t(0)
	field := C.ippGetClassFields(c.handle, &fieldIter)
	for field != nil {
		fields = append(fields, Field{handle: field})
		field = C.ippGetClassFields(c.handle, &fieldIter)
	}

	return fields
}

func (c Class) GetProperties() []Property {
	properties := make([]Property, 0)
	propertyIter := C.size_t(0)
	property := C.ippGetClassProperties(c.handle, &propertyIter)
	for property != nil {
		properties = append(properties, Property{handle: property})
		property = C.ippGetClassProperties(c.handle, &propertyIter)
	}

	return properties
}

func (c *Class) GetMethod(name string) (Method, error) {
	return c.GetMethodWhere(func(m Method) bool {
		return m.GetName() == name
	})
}

func (c *Class) GetField(name string) (Field, error) {
	return c.GetFieldWhere(func(f Field) bool {
		return f.GetName() == name
	})
}

func (c *Class) GetProperty(name string) (Property, error) {
	return c.GetPropertyWhere(func(p Property) bool {
		return p.GetName() == name
	})
}

func (c *Class) GetMethodWhere(predicate func(Method) bool) (Method, error) {
	methods := c.GetMethods()
	for _, m := range methods {
		if predicate(m) {
			return m, nil
		}
	}
	return Method{}, fmt.Errorf("method not found")
}

func (c *Class) GetFieldWhere(predicate func(Field) bool) (Field, error) {
	fields := c.GetFields()
	for _, f := range fields {
		if predicate(f) {
			return f, nil
		}
	}
	return Field{}, fmt.Errorf("field not found")
}

func (c *Class) GetPropertyWhere(predicate func(Property) bool) (Property, error) {
	properties := c.GetProperties()
	for _, p := range properties {
		if predicate(p) {
			return p, nil
		}
	}
	return Property{}, fmt.Errorf("property not found")
}

func (c *Class) GetMethodsWhere(predicate func(Method) bool) []Method {
	methods := c.GetMethods()
	filtered := make([]Method, 0)
	for _, m := range methods {
		if predicate(m) {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

func (c *Class) GetFieldsWhere(predicate func(Field) bool) []Field {
	fields := c.GetFields()
	filtered := make([]Field, 0)
	for _, f := range fields {
		if predicate(f) {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

func (c *Class) GetPropertiesWhere(predicate func(Property) bool) []Property {
	properties := c.GetProperties()
	filtered := make([]Property, 0)
	for _, p := range properties {
		if predicate(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (c *Class) IsNull() bool {
	return c.handle == nil
}

func (c *Class) HasMethod(name string) bool {
	_, err := c.GetMethod(name)
	return err == nil
}

func (c *Class) HasField(name string) bool {
	_, err := c.GetField(name)
	return err == nil
}

func (c *Class) HasProperty(name string) bool {
	_, err := c.GetProperty(name)
	return err == nil
}
