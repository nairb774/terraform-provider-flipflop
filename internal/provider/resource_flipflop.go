package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFlipFlop() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource that allows tracking prior values.",

		CreateContext: resourceFlipFlopCreate,
		ReadContext:   resourceFlipFlopRead,
		UpdateContext: resourceFlipFlopUpdate,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{
			"value": {
				// This description is used by the documentation generator and the language server.
				Description:  "The current value.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"a": {
				Description: "One of the prior recorded values.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"b": {
				Description: "One of the prior recorded values.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"index": {
				Description: "Index of the currently active value. 0==a  1==b",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},

		CustomizeDiff: customdiff.All(
			customdiff.IfValueChange("value",
				func(_ context.Context, old, new, _ interface{}) bool { return old != new },
				func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
					value := diff.Get("value").(string)
					if diff.Id() == "" {
						if err := diff.SetNew("a", value); err != nil {
							return err
						}
						if err := diff.SetNew("b", value); err != nil {
							return err
						}

						return diff.SetNew("index", 0)
					}

					index := diff.Get("index").(int)
					// Update index:
					index = 1 - index
					if err := diff.SetNew("index", index); err != nil {
						return err
					}

					var field string
					if index == 0 {
						field = "a"
					} else {
						field = "b"
					}

					if diff.NewValueKnown("value") {
						return diff.SetNew(field, value)
					}

					return diff.SetNewComputed(field)
				}),
		),
	}
}

func resourceFlipFlopCreate(ctx context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("ready")
	return nil
}

func resourceFlipFlopRead(ctx context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourceFlipFlopUpdate(ctx context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}
