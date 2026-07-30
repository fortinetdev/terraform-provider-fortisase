// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityIpThreatFeeds keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityIpThreatFeeds{}

func newResourceSecurityIpThreatFeeds() resource.Resource {
	return &resourceSecurityIpThreatFeeds{
		resourceSecurityIpThreatFeed: &resourceSecurityIpThreatFeed{},
	}
}

type resourceSecurityIpThreatFeeds struct {
	*resourceSecurityIpThreatFeed
}

func (r *resourceSecurityIpThreatFeeds) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ip_threat_feeds"
}

func (r *resourceSecurityIpThreatFeeds) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityIpThreatFeed.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_ip_threat_feeds is deprecated. Please use fortisase_security_ip_threat_feed instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityIpThreatFeeds) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityIpThreatFeed.Configure(ctx, req, resp)
	r.resourceSecurityIpThreatFeed.resourceName = "fortisase_security_ip_threat_feeds"
}
func (r *resourceSecurityIpThreatFeeds) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
