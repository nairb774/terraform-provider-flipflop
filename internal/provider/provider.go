package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type flipflopProvider struct {
	version string
}

var _ provider.Provider = (*flipflopProvider)(nil)

func (p *flipflopProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *flipflopProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *flipflopProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "flipflop"
	resp.Version = p.version
}

func (p *flipflopProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newResourceFlipFlop,
		newResourceFlipFlopTri,
	}
}

func (p *flipflopProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &flipflopProvider{version: version}
	}
}
