package il2cpp

//#include "il2cpp_structs.h"
//#include <string.h>
import "C"

type ParameterInfo struct {
	parameter *C.ParameterInfo

	Name     string
	TypeName string
}

func newParam(p *C.ParameterInfo, name string) (*ParameterInfo, error) {
	if p == nil || p.parameter_type == nil {
		return nil, errNil
	}

	var err error

	param := &ParameterInfo{parameter: p}
	param.Name = name

	param.TypeName, err = paramGetTypeName(param)

	return param, err
}
