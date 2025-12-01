package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccResourceFlipFlopTri(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v1"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v2"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v3"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v4"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v5"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v5"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v5"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlopTri_EmptyValue(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = ""
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", ""),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "something"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "something"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "something"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = ""
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", ""),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "something"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlopTri_MultipleNoOpUpdates(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "initial"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlopTri_RapidFlipping(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v1"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v2"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v3"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v4"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v5"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v5"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v3"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v5"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v6"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v6"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v4"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v6"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v5"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "v7"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "v7"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "v7"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "v6"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "v5"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlopTri_KnownAtPlanTime(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "initial"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("value"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("top_index"), knownvalue.Int64Exact(0)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("middle_index"), knownvalue.Int64Exact(1)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("bottom_index"), knownvalue.Int64Exact(2)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("a"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("b"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("c"), knownvalue.StringExact("initial")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "changed"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("value"), knownvalue.StringExact("changed")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("top_index"), knownvalue.Int64Exact(2)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("middle_index"), knownvalue.Int64Exact(0)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("bottom_index"), knownvalue.Int64Exact(1)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("a"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("b"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("c"), knownvalue.StringExact("changed")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "changed"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "changed"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "third"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("value"), knownvalue.StringExact("third")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("top_index"), knownvalue.Int64Exact(1)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("middle_index"), knownvalue.Int64Exact(2)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("bottom_index"), knownvalue.Int64Exact(0)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("a"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("b"), knownvalue.StringExact("third")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("c"), knownvalue.StringExact("changed")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "third"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "third"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "changed"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "fourth"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("value"), knownvalue.StringExact("fourth")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("top_index"), knownvalue.Int64Exact(0)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("middle_index"), knownvalue.Int64Exact(1)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("bottom_index"), knownvalue.Int64Exact(2)),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("a"), knownvalue.StringExact("fourth")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("b"), knownvalue.StringExact("third")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("c"), knownvalue.StringExact("changed")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "fourth"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "fourth"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "third"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "changed"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlopTri_UnknownAtPlanTime(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "top_index", "0"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "middle_index", "1"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "bottom_index", "2"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "b", "initial"),
					resource.TestCheckResourceAttr("flipflop_tri.ff", "c", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop_tri" "ff" {
						value = timestamp()
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("flipflop_tri.ff", tfjsonpath.New("value")),
						plancheck.ExpectUnknownValue("flipflop_tri.ff", tfjsonpath.New("top_index")),
						plancheck.ExpectUnknownValue("flipflop_tri.ff", tfjsonpath.New("middle_index")),
						plancheck.ExpectUnknownValue("flipflop_tri.ff", tfjsonpath.New("bottom_index")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("a"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("b"), knownvalue.StringExact("initial")),
						plancheck.ExpectUnknownValue("flipflop_tri.ff", tfjsonpath.New("c")),
						plancheck.ExpectKnownValue("flipflop_tri.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop_tri.ff", "id", "ready"),
					resource.TestCheckResourceAttrSet("flipflop_tri.ff", "value"),
					resource.TestCheckResourceAttrSet("flipflop_tri.ff", "a"),
					resource.TestCheckResourceAttrSet("flipflop_tri.ff", "b"),
					resource.TestCheckResourceAttrSet("flipflop_tri.ff", "c"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
