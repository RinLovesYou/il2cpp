# A wrapper around il2cpp

## Quickstart
First of all, quite importantly, this wrapper makes some assumptions about your environment. It assumes that the current golang binary is a c-shared library<br>
injected into a Unity process built using il2cpp. It also assumes that the initialization process has gone far enough for all assemblies to be loaded.<br>

As a quick example, here we'll get an image, and print out every Class, its methods, their parameters, and their types

```go
package main

import (
	"GoMod/il2cpp"
	"fmt"
)

func init() {
	asm, err := il2cpp.GetImage("Assembly-CSharp")
	if err != nil {
		fmt.Printf("Failed to find Assembly-CSharp assembly: %s\n", err.Error())
		return
	}

	for _, class := range asm.Classes {
		fmt.Println("class:", class.Name)
		fmt.Println()

		methods, err := class.GetMethods()
		if err != nil {
			fmt.Printf("Failed to get methods for class %s: %s\n", class.Name, err.Error())
			return
		}

		for _, method := range methods {
			fmt.Println("method:", method.Name)
			fmt.Println()

			for _, param := range method.Parameters {
				fmt.Println("param:", param.Name, "type:", param.TypeName)
				fmt.Println()
			}
		}
	}
}

func main() {

}

```

You can compile this using `go build --buildmode=c-shared -o GoMod.dll .`
