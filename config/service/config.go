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

		// Configure sensitive fields to be stored in connection secrets
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			// Add IPv4 address if available
			if ipv4, ok := attr["ipv4"].(string); ok && ipv4 != "" {
				conn["ipv4"] = []byte(ipv4)
			}

			// Add IPv6 address if available
			if ipv6, ok := attr["ipv6"].(string); ok && ipv6 != "" {
				conn["ipv6"] = []byte(ipv6)
			}

			// Add CNAME if available
			if cname, ok := attr["cname"].(string); ok && cname != "" {
				conn["cname"] = []byte(cname)
				conn["endpoint"] = []byte(cname) // Also expose as generic endpoint
			}

			// Add admin email if available
			if adminEmail, ok := attr["admin_email"].(string); ok && adminEmail != "" {
				conn["admin_email"] = []byte(adminEmail)
			}

			// Add admin user if available
			if adminUser, ok := attr["admin_user"].(string); ok && adminUser != "" {
				conn["admin_user"] = []byte(adminUser)
			}

			return conn, nil
		}
	})
}
