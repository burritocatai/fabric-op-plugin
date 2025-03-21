package fabric-ai

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func FabricAICLICLI() schema.Executable {
	return schema.Executable{
		Name:      "Fabric AI CLI CLI", // TODO: Check if this is correct
		Runs:      []string{"fabric"},
		DocsURL:   sdk.URL("https://fabric-ai.com/docs/cli"), // TODO: Replace with actual URL
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.Credentials,
			},
		},
	}
}
