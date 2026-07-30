// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityIpThreatFeeds keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityIpThreatFeeds{}

func newDatasourceSecurityIpThreatFeeds() datasource.DataSource {
	return &datasourceSecurityIpThreatFeeds{
		datasourceSecurityIpThreatFeed: &datasourceSecurityIpThreatFeed{},
	}
}

type datasourceSecurityIpThreatFeeds struct {
	*datasourceSecurityIpThreatFeed
}

func (r *datasourceSecurityIpThreatFeeds) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ip_threat_feeds"
}

func (r *datasourceSecurityIpThreatFeeds) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityIpThreatFeed.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_ip_threat_feeds is deprecated. Please use fortisase_security_ip_threat_feed instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityIpThreatFeeds) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityIpThreatFeed.Configure(ctx, req, resp)
	r.datasourceSecurityIpThreatFeed.resourceName = "fortisase_security_ip_threat_feeds"
}
