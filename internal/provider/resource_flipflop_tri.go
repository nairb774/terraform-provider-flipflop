package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceFlipFlopTri struct{}

type resourceFlipFlopTriModel struct {
	ID          types.String `tfsdk:"id"`
	Value       types.String `tfsdk:"value"`
	A           types.String `tfsdk:"a"`
	B           types.String `tfsdk:"b"`
	C           types.String `tfsdk:"c"`
	TopIndex    types.Int64  `tfsdk:"top_index"`
	MiddleIndex types.Int64  `tfsdk:"middle_index"`
	BottomIndex types.Int64  `tfsdk:"bottom_index"`
}

var _ resource.Resource = (*resourceFlipFlopTri)(nil)
var _ resource.ResourceWithModifyPlan = (*resourceFlipFlopTri)(nil)

func (r *resourceFlipFlopTri) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceFlipFlopTriModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceFlipFlopTri) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFlipFlopTri) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tri"
}

func (r *resourceFlipFlopTri) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No-op: this resource doesn't read from external state
}

func (r *resourceFlipFlopTri) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema.Description = "Resource that allows tracking prior values with three states."
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

		"c": schema.StringAttribute{
			Description: "One of the prior recorded values.",
			Computed:    true,
		},

		"top_index": schema.Int64Attribute{
			Description: "Index of the most recent value. Points to a (0), b (1), or c (2).",
			Computed:    true,
		},

		"middle_index": schema.Int64Attribute{
			Description: "Index of the second most recent value. Points to a (0), b (1), or c (2).",
			Computed:    true,
		},

		"bottom_index": schema.Int64Attribute{
			Description: "Index of the oldest value. Points to a (0), b (1), or c (2).",
			Computed:    true,
		},
	}
}

func (r *resourceFlipFlopTri) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceFlipFlopTriModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceFlipFlopTri) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If this is a resource deletion (plan is null), do nothing
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan resourceFlipFlopTriModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if state is null (create operation)
	if req.State.Raw.IsNull() {
		// On create: set all slots to value with top_index=0, middle_index=1, bottom_index=2
		plan.ID = types.StringValue("ready")
		plan.A = plan.Value
		plan.B = plan.Value
		plan.C = plan.Value
		plan.TopIndex = types.Int64Value(0)
		plan.MiddleIndex = types.Int64Value(1)
		plan.BottomIndex = types.Int64Value(2)
	} else {
		// Get current state for update
		var state resourceFlipFlopTriModel
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Preserve ID
		plan.ID = state.ID

		if !plan.Value.Equal(state.Value) {
			// On update when value changes: rotate indices
			// New top is what was bottom, new middle is what was top, new bottom is what was middle
			plan.TopIndex = state.BottomIndex
			plan.MiddleIndex = state.TopIndex
			plan.BottomIndex = state.MiddleIndex

			// Update the slot at the new top index with the new value
			// Keep other slots unchanged
			plan.A = state.A
			plan.B = state.B
			plan.C = state.C

			switch plan.TopIndex.ValueInt64() {
			case 0:
				plan.A = plan.Value
			case 1:
				plan.B = plan.Value
			case 2:
				plan.C = plan.Value
			}
		} else {
			// Value unchanged: preserve existing state
			plan.A = state.A
			plan.B = state.B
			plan.C = state.C
			plan.TopIndex = state.TopIndex
			plan.MiddleIndex = state.MiddleIndex
			plan.BottomIndex = state.BottomIndex
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func newResourceFlipFlopTri() resource.Resource {
	return &resourceFlipFlopTri{}
}
