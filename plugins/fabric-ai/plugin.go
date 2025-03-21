package fabric-ai

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "fabric-ai",
		Platform: schema.PlatformInfo{
			Name:     "Fabric AI CLI",
			Homepage: sdk.URL("https://fabric-ai.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			Credentials(),
		},
		Executables: []schema.Executable{
			FabricAICLICLI(),
		},
	}
}
