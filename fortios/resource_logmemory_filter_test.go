// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSLogMemoryFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogMemoryFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogMemoryFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogMemoryFilterExists("fortios_logmemory_filter.trname"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "anomaly", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "dns", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "filter_type", "include"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "forward_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "gtp", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "local_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "multicast_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "severity", "information"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "sniffer_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "ssh", "enable"),
					resource.TestCheckResourceAttr("fortios_logmemory_filter.trname", "voip", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogMemoryFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogMemoryFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogMemoryFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogMemoryFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogMemoryFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogMemoryFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogMemoryFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logmemory_filter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogMemoryFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogMemoryFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogMemoryFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logmemory_filter" "trname" {
  anomaly           = "enable"
  dns               = "enable"
  filter_type       = "include"
  forward_traffic   = "enable"
  gtp               = "enable"
  local_traffic     = "enable"
  multicast_traffic = "enable"
  severity          = "information"
  sniffer_traffic   = "enable"
  ssh               = "enable"
  voip              = "enable"
}
`)
}
