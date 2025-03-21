package fabric-ai

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestCredentialsProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, Credentials().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.Credentials: "S&lt;vQ-&amp;AsB.EXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"FABRIC-AI_CREDENTIALS": "S&lt;vQ-&amp;AsB.EXAMPLE",
				},
			},
		},
	})
}

func TestCredentialsImporter(t *testing.T) {
	plugintest.TestImporter(t, Credentials().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"FABRIC-AI_CREDENTIALS": "S&lt;vQ-&amp;AsB.EXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Credentials: "S&lt;vQ-&amp;AsB.EXAMPLE",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in fabric-ai/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "S&lt;vQ-&amp;AsB.EXAMPLE",
			// 		},
			// 	},
			},
		},
	})
}
