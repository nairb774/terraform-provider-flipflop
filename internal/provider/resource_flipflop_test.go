package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccResourceFlipFlop(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = "bar"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "bar"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "bar"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "bar"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "baz"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "baz"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "bar"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "baz"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "baz"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "baz"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "bar"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "baz"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "foo"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "foo"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "foo"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "baz"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "baz"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "baz"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "foo"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "baz"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlop_EmptyValue(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = ""
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", ""),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", ""),
					resource.TestCheckResourceAttr("flipflop.ff", "b", ""),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "something"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "something"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", ""),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "something"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = ""
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", ""),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", ""),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "something"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlop_MultipleNoOpUpdates(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "initial"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "initial"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlop_RapidFlipping(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = "v1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "v1"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "v1"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "v2"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "v2"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "v1"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "v3"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "v3"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "v3"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "v2"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "v4"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "v4"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "v3"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "v4"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "v5"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "value", "v5"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "v5"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "v4"),
				),
			},
		},
	})
}

func TestAccResourceFlipFlop_KnownAtPlanTime(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = "initial"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("value"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("index"), knownvalue.Int64Exact(0)),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("a"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("b"), knownvalue.StringExact("initial")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "initial"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "changed"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("value"), knownvalue.StringExact("changed")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("index"), knownvalue.Int64Exact(1)),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("a"), knownvalue.StringExact("initial")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("b"), knownvalue.StringExact("changed")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "changed"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "initial"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "changed"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "third"
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("id"), knownvalue.StringExact("ready")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("value"), knownvalue.StringExact("third")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("index"), knownvalue.Int64Exact(0)),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("a"), knownvalue.StringExact("third")),
						plancheck.ExpectKnownValue("flipflop.ff", tfjsonpath.New("b"), knownvalue.StringExact("changed")),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("flipflop.ff", "id", "ready"),
					resource.TestCheckResourceAttr("flipflop.ff", "value", "third"),
					resource.TestCheckResourceAttr("flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr("flipflop.ff", "a", "third"),
					resource.TestCheckResourceAttr("flipflop.ff", "b", "changed"),
				),
			},
		},
	})
}
