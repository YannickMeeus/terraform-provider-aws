// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ssm_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDocument,
			TypeName: "aws_ssm_document",
		},
		{
			Factory:  DataSourceInstances,
			TypeName: "aws_ssm_instances",
		},
		{
			Factory:  DataSourceMaintenanceWindows,
			TypeName: "aws_ssm_maintenance_windows",
		},
		{
			Factory:  DataSourceParameter,
			TypeName: "aws_ssm_parameter",
		},
		{
			Factory:  DataSourceParametersByPath,
			TypeName: "aws_ssm_parameters_by_path",
		},
		{
			Factory:  DataSourcePatchBaseline,
			TypeName: "aws_ssm_patch_baseline",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceActivation,
			TypeName: "aws_ssm_activation",
			Name:     "Activation",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceAssociation,
			TypeName: "aws_ssm_association",
		},
		{
			Factory:  resourceDefaultPatchBaseline,
			TypeName: "aws_ssm_default_patch_baseline",
			Name:     "Default Patch Baseline",
		},
		{
			Factory:  ResourceDocument,
			TypeName: "aws_ssm_document",
			Name:     "Document",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Document",
			},
		},
		{
			Factory:  ResourceMaintenanceWindow,
			TypeName: "aws_ssm_maintenance_window",
			Name:     "Maintenance Window",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "MaintenanceWindow",
			},
		},
		{
			Factory:  ResourceMaintenanceWindowTarget,
			TypeName: "aws_ssm_maintenance_window_target",
		},
		{
			Factory:  ResourceMaintenanceWindowTask,
			TypeName: "aws_ssm_maintenance_window_task",
		},
		{
			Factory:  ResourceParameter,
			TypeName: "aws_ssm_parameter",
			Name:     "Parameter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Parameter",
			},
		},
		{
			Factory:  resourcePatchBaseline,
			TypeName: "aws_ssm_patch_baseline",
			Name:     "Patch Baseline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "PatchBaseline",
			},
		},
		{
			Factory:  ResourcePatchGroup,
			TypeName: "aws_ssm_patch_group",
		},
		{
			Factory:  ResourceResourceDataSync,
			TypeName: "aws_ssm_resource_data_sync",
		},
		{
			Factory:  ResourceServiceSetting,
			TypeName: "aws_ssm_service_setting",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSM
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ssm_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ssm_sdkv2.NewFromConfig(cfg, func(o *ssm_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
