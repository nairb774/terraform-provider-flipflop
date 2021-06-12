package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFlipFlop(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flipflop" "ff" {
						value = "bar"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"flipflop.ff", "value", "bar"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "a", "bar"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "b", "bar"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "baz"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"flipflop.ff", "value", "baz"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "a", "bar"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "b", "baz"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "baz"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"flipflop.ff", "value", "baz"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "a", "bar"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "b", "baz"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "foo"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"flipflop.ff", "value", "foo"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "index", "0"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "a", "foo"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "b", "baz"),
				),
			},
			{
				Config: `
					resource "flipflop" "ff" {
						value = "baz"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"flipflop.ff", "value", "baz"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "index", "1"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "a", "foo"),
					resource.TestCheckResourceAttr(
						"flipflop.ff", "b", "baz"),
				),
			},
		},
	})
}
