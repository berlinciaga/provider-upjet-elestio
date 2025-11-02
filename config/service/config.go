// ABOUTME: Configuration for Elestio Service resource
// ABOUTME: Defines CRD generation settings and resource behavior
package service

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for elestio_service resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elestio_service", func(r *config.Resource) {
		// Set the API group short name (will be service.elestio.crossplane.io)
		r.ShortGroup = "service"

		// Set the Kind name for the CRD
		r.Kind = "Service"

		// Skip fields that Upjet cannot convert due to type issues
		// These fields have nested_type without a top-level type which causes conversion errors
		// Workaround: We'll handle these fields separately or use alternative API fields
		fieldsToSkip := []string{"admin", "database_admin", "ssh_keys", "ssh_public_keys"}
		for _, field := range fieldsToSkip {
			delete(r.TerraformResource.Schema, field)
		}

		// Note: We'll add sensitive fields and references in later steps
	})
}
