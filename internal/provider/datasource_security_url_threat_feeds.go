// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityUrlThreatFeeds keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityUrlThreatFeeds{}

func newDatasourceSecurityUrlThreatFeeds() datasource.DataSource {
	return &datasourceSecurityUrlThreatFeeds{
		datasourceSecurityUrlThreatFeed: &datasourceSecurityUrlThreatFeed{},
	}
}

type datasourceSecurityUrlThreatFeeds struct {
	*datasourceSecurityUrlThreatFeed
}

func (r *datasourceSecurityUrlThreatFeeds) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_url_threat_feeds"
}

func (r *datasourceSecurityUrlThreatFeeds) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityUrlThreatFeed.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_url_threat_feeds is deprecated. Please use fortisase_security_url_threat_feed instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityUrlThreatFeeds) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityUrlThreatFeed.Configure(ctx, req, resp)
	r.datasourceSecurityUrlThreatFeed.resourceName = "fortisase_security_url_threat_feeds"
}
