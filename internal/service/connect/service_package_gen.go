// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package connect

import (
	"context"

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
			Factory:  DataSourceBotAssociation,
			TypeName: "aws_connect_bot_association",
		},
		{
			Factory:  DataSourceContactFlow,
			TypeName: "aws_connect_contact_flow",
		},
		{
			Factory:  DataSourceContactFlowModule,
			TypeName: "aws_connect_contact_flow_module",
		},
		{
			Factory:  DataSourceHoursOfOperation,
			TypeName: "aws_connect_hours_of_operation",
		},
		{
			Factory:  DataSourceInstance,
			TypeName: "aws_connect_instance",
		},
		{
			Factory:  DataSourceInstanceStorageConfig,
			TypeName: "aws_connect_instance_storage_config",
		},
		{
			Factory:  DataSourceLambdaFunctionAssociation,
			TypeName: "aws_connect_lambda_function_association",
		},
		{
			Factory:  DataSourcePrompt,
			TypeName: "aws_connect_prompt",
		},
		{
			Factory:  DataSourceQueue,
			TypeName: "aws_connect_queue",
		},
		{
			Factory:  DataSourceQuickConnect,
			TypeName: "aws_connect_quick_connect",
		},
		{
			Factory:  DataSourceRoutingProfile,
			TypeName: "aws_connect_routing_profile",
		},
		{
			Factory:  DataSourceSecurityProfile,
			TypeName: "aws_connect_security_profile",
		},
		{
			Factory:  DataSourceUser,
			TypeName: "aws_connect_user",
		},
		{
			Factory:  DataSourceUserHierarchyGroup,
			TypeName: "aws_connect_user_hierarchy_group",
		},
		{
			Factory:  DataSourceUserHierarchyStructure,
			TypeName: "aws_connect_user_hierarchy_structure",
		},
		{
			Factory:  DataSourceVocabulary,
			TypeName: "aws_connect_vocabulary",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBotAssociation,
			TypeName: "aws_connect_bot_association",
		},
		{
			Factory:  ResourceContactFlow,
			TypeName: "aws_connect_contact_flow",
			Name:     "Contact Flow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceContactFlowModule,
			TypeName: "aws_connect_contact_flow_module",
			Name:     "Contact Flow Module",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceHoursOfOperation,
			TypeName: "aws_connect_hours_of_operation",
			Name:     "Hours Of Operation",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_connect_instance",
		},
		{
			Factory:  ResourceInstanceStorageConfig,
			TypeName: "aws_connect_instance_storage_config",
		},
		{
			Factory:  ResourceLambdaFunctionAssociation,
			TypeName: "aws_connect_lambda_function_association",
		},
		{
			Factory:  ResourcePhoneNumber,
			TypeName: "aws_connect_phone_number",
			Name:     "Phone Number",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceQueue,
			TypeName: "aws_connect_queue",
			Name:     "Queue",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceQuickConnect,
			TypeName: "aws_connect_quick_connect",
			Name:     "Quick Connect",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRoutingProfile,
			TypeName: "aws_connect_routing_profile",
			Name:     "Routing Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSecurityProfile,
			TypeName: "aws_connect_security_profile",
			Name:     "Security Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_connect_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUserHierarchyGroup,
			TypeName: "aws_connect_user_hierarchy_group",
			Name:     "User Hierarchy Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUserHierarchyStructure,
			TypeName: "aws_connect_user_hierarchy_structure",
		},
		{
			Factory:  ResourceVocabulary,
			TypeName: "aws_connect_vocabulary",
			Name:     "Vocabulary",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Connect
}

var ServicePackage = &servicePackage{}
