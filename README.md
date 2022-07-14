# A wrapper around il2cpp

[Support Server](https://discord.gg/zgzkyGvTS8) (discord)

## Quickstart
First of all, quite importantly, this wrapper makes some assumptions about your environment. It assumes that the current golang binary is a c-shared library<br>
injected into a Unity process built using il2cpp. It also assumes that the initialization process has gone far enough for all assemblies to be loaded.<br>

This wrapper updates as my GoMod project evolves, which you can get in the support server above.

Here is a simple example fetching some metadata about images/classes/methods/properties/fields

```go
package main

import (
	"github.com/RinLovesYou/il2cpp"
	"fmt"
)

func init() {
	domain := il2cpp.GetDomain()
	domain.AttachThread()

	for _, image := range domain.GetImages() {
		utils.Log("Image: %s", image.GetName())

		for _, class := range image.GetClasses() {
			utils.Log("Class: %s", class.GetName())

			for _, method := range class.GetMethods() {
				utils.Log("Method: %s", method.GetName())
				utils.Log("ReturnType: %s", method.GetReturnType().GetName())
				for _, param := range method.GetParams() {
					utils.Log("Param: %s", param.GetName())
				}
			}

			for _, property := range class.GetProperties() {
				utils.Log("Property: %s", property.GetName())
				utils.Log("ReturnType: %s", property.GetGet().GetReturnType().GetName())
			}

			for _, field := range class.GetFields() {
				utils.Log("Field: %s", field.GetName())
			}

		}
	}
}

func main() {

}

```

You can compile this using `go build --buildmode=c-shared -o GoMod.dll .`
