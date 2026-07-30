// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDomainThreatFeeds keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDomainThreatFeeds{}

func newDatasourceSecurityDomainThreatFeeds() datasource.DataSource {
	return &datasourceSecurityDomainThreatFeeds{
		datasourceSecurityDomainThreatFeed: &datasourceSecurityDomainThreatFeed{},
	}
}

type datasourceSecurityDomainThreatFeeds struct {
	*datasourceSecurityDomainThreatFeed
}

func (r *datasourceSecurityDomainThreatFeeds) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_domain_threat_feeds"
}

func (r *datasourceSecurityDomainThreatFeeds) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDomainThreatFeed.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_domain_threat_feeds is deprecated. Please use fortisase_security_domain_threat_feed instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDomainThreatFeeds) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDomainThreatFeed.Configure(ctx, req, resp)
	r.datasourceSecurityDomainThreatFeed.resourceName = "fortisase_security_domain_threat_feeds"
}
