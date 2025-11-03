package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, _ client.Client, _ resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			// Configuration will be empty - Terraform provider will use environment variables
			// ELESTIO_EMAIL and ELESTIO_API_TOKEN from the pod environment
			Configuration: map[string]any{},
		}
		return ps, nil
	}
}
