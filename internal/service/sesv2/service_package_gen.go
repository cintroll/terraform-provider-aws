// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sesv2

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
			Factory:  DataSourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
		},
		{
			Factory:  ResourceConfigurationSetEventDestination,
			TypeName: "aws_sesv2_configuration_set_event_destination",
		},
		{
			Factory:  ResourceDedicatedIPAssignment,
			TypeName: "aws_sesv2_dedicated_ip_assignment",
		},
		{
			Factory:  ResourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
		},
		{
			Factory:  ResourceEmailIdentity,
			TypeName: "aws_sesv2_email_identity",
		},
		{
			Factory:  ResourceEmailIdentityFeedbackAttributes,
			TypeName: "aws_sesv2_email_identity_feedback_attributes",
		},
		{
			Factory:  ResourceEmailIdentityMailFromAttributes,
			TypeName: "aws_sesv2_email_identity_mail_from_attributes",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SESV2
}

var ServicePackage = &servicePackage{}
