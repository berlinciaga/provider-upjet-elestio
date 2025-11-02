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

		// Note: We'll add sensitive fields and references in later steps
	})
}
