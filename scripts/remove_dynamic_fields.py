#!/usr/bin/env python3
"""
Remove dynamic type fields from Elestio provider schema.

The local_field and local_field_sensitive attributes use cty.DynamicPseudoType
which Upjet v2 cannot convert to Kubernetes CRD schema. These fields are
optional metadata storage and not needed for core service provisioning.
"""
import json
import sys

def remove_dynamic_fields(schema):
    """Remove local_field and local_field_sensitive from all resources."""
    modified = False

    if "provider_schemas" not in schema:
        return schema, modified

    for provider_name, provider_data in schema["provider_schemas"].items():
        if "resource_schemas" not in provider_data:
            continue

        for resource_name, resource_data in provider_data["resource_schemas"].items():
            if "block" not in resource_data:
                continue

            block = resource_data["block"]
            if "attributes" not in block:
                continue

            # Remove dynamic type fields
            removed = []
            for field in ["local_field", "local_field_sensitive"]:
                if field in block["attributes"]:
                    del block["attributes"][field]
                    removed.append(field)
                    modified = True

            if removed:
                print(f"Removed {', '.join(removed)} from {resource_name}", file=sys.stderr)

    return schema, modified

def main():
    # Read schema
    with open("config/schema.json", "r") as f:
        schema = json.load(f)

    # Remove dynamic fields
    schema, modified = remove_dynamic_fields(schema)

    if not modified:
        print("No dynamic fields found to remove", file=sys.stderr)
        return 0

    # Write back as single-line JSON (matching original format)
    with open("config/schema.json", "w") as f:
        json.dump(schema, f, separators=(',', ':'))

    print("Successfully removed dynamic type fields from schema.json", file=sys.stderr)
    return 0

if __name__ == "__main__":
    sys.exit(main())
