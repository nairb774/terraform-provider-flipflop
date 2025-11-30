package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	// schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

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
