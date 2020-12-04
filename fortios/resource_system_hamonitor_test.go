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

func TestAccFortiOSSystemHaMonitor_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemHaMonitor_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemHaMonitorConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemHaMonitorExists("fortios_system_hamonitor.trname"),
					resource.TestCheckResourceAttr("fortios_system_hamonitor.trname", "monitor_vlan", "disable"),
					resource.TestCheckResourceAttr("fortios_system_hamonitor.trname", "vlan_hb_interval", "5"),
					resource.TestCheckResourceAttr("fortios_system_hamonitor.trname", "vlan_hb_lost_threshold", "3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemHaMonitorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemHaMonitor: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemHaMonitor is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemHaMonitor(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemHaMonitor: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemHaMonitor: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemHaMonitorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_hamonitor" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemHaMonitor(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemHaMonitor %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemHaMonitorConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_hamonitor" "trname" {
  monitor_vlan           = "disable"
  vlan_hb_interval       = 5
  vlan_hb_lost_threshold = 3
}
`)
}
