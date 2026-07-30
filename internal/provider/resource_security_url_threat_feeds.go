// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityUrlThreatFeeds keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityUrlThreatFeeds{}

func newResourceSecurityUrlThreatFeeds() resource.Resource {
	return &resourceSecurityUrlThreatFeeds{
		resourceSecurityUrlThreatFeed: &resourceSecurityUrlThreatFeed{},
	}
}

type resourceSecurityUrlThreatFeeds struct {
	*resourceSecurityUrlThreatFeed
}

func (r *resourceSecurityUrlThreatFeeds) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_url_threat_feeds"
}

func (r *resourceSecurityUrlThreatFeeds) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityUrlThreatFeed.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_url_threat_feeds is deprecated. Please use fortisase_security_url_threat_feed instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityUrlThreatFeeds) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityUrlThreatFeed.Configure(ctx, req, resp)
	r.resourceSecurityUrlThreatFeed.resourceName = "fortisase_security_url_threat_feeds"
}
func (r *resourceSecurityUrlThreatFeeds) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
