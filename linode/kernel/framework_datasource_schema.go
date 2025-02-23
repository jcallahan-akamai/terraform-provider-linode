package kernel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var frameworkDatasourceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "The unique ID of this Kernel.",
			Required:    true,
		},
		"architecture": schema.StringAttribute{
			Description: "The architecture of this Kernel.",
			Computed:    true,
		},
		"deprecated": schema.BoolAttribute{
			Description: "Whether or not this Kernel is deprecated.",
			Computed:    true,
		},
		"kvm": schema.BoolAttribute{
			Description: "If this Kernel is suitable for KVM Linodes.",
			Computed:    true,
		},
		"label": schema.StringAttribute{
			Description: "The friendly name of this Kernel.",
			Computed:    true,
		},
		"pvops": schema.BoolAttribute{
			Description: "If this Kernel is suitable for paravirtualized operations.",
			Computed:    true,
		},
		"version": schema.StringAttribute{
			Description: "Linux Kernel version.",
			Computed:    true,
		},
		"xen": schema.BoolAttribute{
			Description: "If this Kernel is suitable for Xen Linodes.",
			Computed:    true,
		},
	},
}
