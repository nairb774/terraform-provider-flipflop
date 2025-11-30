package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

func TestAccResourceFlipFlop_Validation(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = ""
					}
				`,
				ExpectError: regexp.MustCompile("Attribute value string length must be at least 1"),
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
