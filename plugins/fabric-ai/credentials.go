package fabric-ai

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func Credentials() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.Credentials,
		DocsURL:       sdk.URL("https://fabric-ai.com/docs/credentials"), // TODO: Replace with actual URL
		ManagementURL: sdk.URL("https://console.fabric-ai.com/user/security/tokens"), // TODO: Replace with actual URL
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Credentials,
				MarkdownDescription: "Credentials used to authenticate to Fabric AI CLI.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 17,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Symbols:   true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			TryFabricAICLIConfigFile(),
		)}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"FABRIC-AI_CREDENTIALS": fieldname.Credentials, // TODO: Check if this is correct
}

// TODO: Check if the platform stores the Credentials in a local config file, and if so,
// implement the function below to add support for importing it.
func TryFabricAICLIConfigFile() sdk.Importer {
	return importer.TryFile("~/path/to/config/file.yml", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		// var config Config
		// if err := contents.ToYAML(&config); err != nil {
		// 	out.AddError(err)
		// 	return
		// }

		// if config.Credentials == "" {
		// 	return
		// }

		// out.AddCandidate(sdk.ImportCandidate{
		// 	Fields: map[sdk.FieldName]string{
		// 		fieldname.Credentials: config.Credentials,
		// 	},
		// })
	})
}

// TODO: Implement the config file schema
// type Config struct {
//	Credentials string
// }
