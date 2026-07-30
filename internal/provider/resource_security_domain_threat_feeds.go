// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityDomainThreatFeeds keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityDomainThreatFeeds{}

func newResourceSecurityDomainThreatFeeds() resource.Resource {
	return &resourceSecurityDomainThreatFeeds{
		resourceSecurityDomainThreatFeed: &resourceSecurityDomainThreatFeed{},
	}
}

type resourceSecurityDomainThreatFeeds struct {
	*resourceSecurityDomainThreatFeed
}

func (r *resourceSecurityDomainThreatFeeds) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_domain_threat_feeds"
}

func (r *resourceSecurityDomainThreatFeeds) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityDomainThreatFeed.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_domain_threat_feeds is deprecated. Please use fortisase_security_domain_threat_feed instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityDomainThreatFeeds) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityDomainThreatFeed.Configure(ctx, req, resp)
	r.resourceSecurityDomainThreatFeed.resourceName = "fortisase_security_domain_threat_feeds"
}
func (r *resourceSecurityDomainThreatFeeds) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
