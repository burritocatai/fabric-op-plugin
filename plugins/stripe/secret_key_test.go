package stripe

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestSecretKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, SecretKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"STRIPE_API_KEY":    "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
				"STRIPE_SECRET_KEY": "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[string]string{
						fieldname.Key: "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
					},
				},
				{
					Fields: map[string]string{
						fieldname.Key: "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
					},
				},
			},
		},
		"Stripe config file": {
			Files: map[string]string{
				"~/.config/stripe/config.toml": plugintest.LoadFixture(t, "config.toml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					NameHint: "default – test",
					Fields: map[string]string{
						fieldname.Key:  "sk_uKVoEC2LqU1aXSbKM1ptxFB3QxWiSTMTnbr0CGvkEBMFOs2vetsHc148WMhtrVRAAsP4fcRd35Fz7ykqbhLoa04ZoA7AcRKvUEXAMPLE",
						fieldname.Mode: ModeTest,
					},
				},
				{
					NameHint: "default",
					Fields: map[string]string{
						fieldname.Key:  "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
						fieldname.Mode: ModeLive,
					},
				},
			},
		},
	})
}

func TestSecretKeyProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, SecretKey().Provisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[string]string{
				fieldname.Key: "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"STRIPE_API_KEY": "sk_TEm8TYekzqaEKmSIDRb4PXJQAoq94iL6PZx4C1RQlr1Ls5kn67RVRJjhBfmejEX8OS4T7GxCWBnqBuIG20SzdEwopINxyEL05EXAMPLE",
				},
			},
		},
	})
}
