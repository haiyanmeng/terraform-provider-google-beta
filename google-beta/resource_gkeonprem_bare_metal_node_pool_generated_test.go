// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccGkeonpremBareMetalNodePool_gkeonpremBareMetalNodePoolBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalNodePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalNodePool_gkeonpremBareMetalNodePoolBasicExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_bare_metal_node_pool.nodepool-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "bare_metal_cluster", "location"},
			},
		},
	})
}

func testAccGkeonpremBareMetalNodePool_gkeonpremBareMetalNodePoolBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_bare_metal_cluster" "default-basic" {
  provider = google-beta
  name = "default-basic"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  bare_metal_version = "1.12.3"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
          labels = {}
          node_ip = "10.200.0.9"
        }
      }
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.13"
      ingress_vip = "10.200.0.14"
    }
    metal_lb_config {
      address_pools {
        pool = "pool1"
        addresses = [
          "10.200.0.14/32",
          "10.200.0.15/32",
          "10.200.0.16/32",
          "10.200.0.17/32",
          "10.200.0.18/32",
          "fd00:1::f/128",
          "fd00:1::10/128",
          "fd00:1::11/128",
          "fd00:1::12/128"
        ]
      }
    }
  }
  storage {
    lvp_share_config {
      lvp_config {
        path = "/mnt/localpv-share"
        storage_class = "local-shared"
      }
      shared_path_pv_count = 5
    }
    lvp_node_mounts_config {
      path = "/mnt/localpv-disk"
      storage_class = "local-disks"
    }
  }
  security_config {
    authorization {
      admin_users {
        username = "admin@hashicorptest.com"
      }
    }
  }
}

resource "google_gkeonprem_bare_metal_node_pool" "nodepool-basic" {
  provider = google-beta
  name =  "np-nodepool%{random_suffix}"
  bare_metal_cluster =  google_gkeonprem_bare_metal_cluster.default-basic.name
  location = "us-west1"
  node_pool_config {
    operating_system = "LINUX"
    node_configs {
      node_ip = "10.200.0.11"
    }
  }
}
`, context)
}

func TestAccGkeonpremBareMetalNodePool_gkeonpremBareMetalNodePoolFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalNodePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalNodePool_gkeonpremBareMetalNodePoolFullExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_bare_metal_node_pool.nodepool-full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "bare_metal_cluster", "location"},
			},
		},
	})
}

func testAccGkeonpremBareMetalNodePool_gkeonpremBareMetalNodePoolFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_bare_metal_cluster" "default-full" {
  provider = google-beta
  name = "default-full"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  bare_metal_version = "1.12.3"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
          labels = {}
          node_ip = "10.200.0.9"
        }
      }
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.13"
      ingress_vip = "10.200.0.14"
    }
    metal_lb_config {
      address_pools {
        pool = "pool1"
        addresses = [
          "10.200.0.14/32",
          "10.200.0.15/32",
          "10.200.0.16/32",
          "10.200.0.17/32",
          "10.200.0.18/32",
          "fd00:1::f/128",
          "fd00:1::10/128",
          "fd00:1::11/128",
          "fd00:1::12/128"
        ]
      }
    }
  }
  storage {
    lvp_share_config {
      lvp_config {
        path = "/mnt/localpv-share"
        storage_class = "local-shared"
      }
      shared_path_pv_count = 5
    }
    lvp_node_mounts_config {
      path = "/mnt/localpv-disk"
      storage_class = "local-disks"
    }
  }
  security_config {
    authorization {
      admin_users {
        username = "admin@hashicorptest.com"
      }
    }
  }
}

resource "google_gkeonprem_bare_metal_node_pool" "nodepool-full" {
  provider = google-beta
  name =  "np-nodepool%{random_suffix}"
  display_name = "test-name"
  bare_metal_cluster =  google_gkeonprem_bare_metal_cluster.default-full.name
  location = "us-west1"
  annotations = {}
  node_pool_config {
    operating_system = "LINUX"
    labels = {}
    node_configs {
      node_ip = "10.200.0.11"
      labels = {}
    }
    taints {
      key = "test-key"
      value = "test-value"
      effect = "NO_EXECUTE"
    }
  }
}
`, context)
}

func testAccCheckGkeonpremBareMetalNodePoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gkeonprem_bare_metal_node_pool" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("GkeonpremBareMetalNodePool still exists at %s", url)
			}
		}

		return nil
	}
}
