package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceFlipFlop struct{}

type resourceFlipFlopModel struct {
	ID    types.String `tfsdk:"id"`
	Value types.String `tfsdk:"value"`
	A     types.String `tfsdk:"a"`
	B     types.String `tfsdk:"b"`
	Index types.Int64  `tfsdk:"index"`
}

var _ resource.Resource = (*resourceFlipFlop)(nil)
var _ resource.ResourceWithModifyPlan = (*resourceFlipFlop)(nil)

func (r *resourceFlipFlop) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceFlipFlopModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceFlipFlop) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFlipFlop) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (r *resourceFlipFlop) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No-op: this resource doesn't read from external state
}

func (r *resourceFlipFlop) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema.Description = "Resource that allows tracking prior values."
	resp.Schema.Attributes = map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "Resource identifier.",
			Computed:    true,
		},

		"value": schema.StringAttribute{
			Description: "The current value.",
			Required:    true,
		},

		"a": schema.StringAttribute{
			Description: "One of the prior recorded values.",
			Computed:    true,
		},

		"b": schema.StringAttribute{
			Description: "One of the prior recorded values.",
			Computed:    true,
		},

		"index": schema.Int64Attribute{
			Description: "Index of the currently active value. 0==a  1==b",
			Computed:    true,
		},
	}
}

func (r *resourceFlipFlop) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceFlipFlopModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceFlipFlop) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If this is a resource deletion (plan is null), do nothing
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan resourceFlipFlopModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if state is null (create operation)
	if req.State.Raw.IsNull() {
		// On create: set both a and b to value with index=0
		plan.ID = types.StringValue("ready")
		plan.A = plan.Value
		plan.B = plan.Value
		plan.Index = types.Int64Value(0)
	} else {
		// Get current state for update
		var state resourceFlipFlopModel
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Preserve ID
		plan.ID = state.ID

		// If value is unknown, we don't know if it will change or which slot to update
		if plan.Value.IsUnknown() {
			// The slot opposite to the current index is preserved (known)
			// The current index slot and index itself are unknown
			if state.Index.ValueInt64() == 0 {
				// Current is A, so A is known and B is unknown
				plan.A = state.A
				plan.B = types.StringUnknown()
			} else {
				// Current is B, so A is unknown and B is known
				plan.A = types.StringUnknown()
				plan.B = state.B
			}
			plan.Index = types.Int64Unknown()
		} else if !plan.Value.Equal(state.Value) {
			// On update when value changes: flip index and update corresponding field
			newIndex := int64(1) - state.Index.ValueInt64()
			plan.Index = types.Int64Value(newIndex)

			if newIndex == 0 {
				plan.A = plan.Value
				plan.B = state.B
			} else {
				plan.A = state.A
				plan.B = plan.Value
			}
		} else {
			// Value unchanged: preserve existing state
			plan.A = state.A
			plan.B = state.B
			plan.Index = state.Index
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func newResourceFlipFlop() resource.Resource {
	return &resourceFlipFlop{}
}
