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

func TestAccNetworkServicesMesh_networkServicesMeshBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesMeshDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesMesh_networkServicesMeshBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_mesh.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesMesh_networkServicesMeshBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_services_mesh" "default" {
  provider    = google-beta
  name        = "tf-test-my-mesh%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  interception_port = 443
}
`, context)
}

func TestAccNetworkServicesMesh_networkServicesMeshNoPortExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesMeshDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesMesh_networkServicesMeshNoPortExample(context),
			},
			{
				ResourceName:            "google_network_services_mesh.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesMesh_networkServicesMeshNoPortExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_services_mesh" "default" {
  provider    = google-beta
  name        = "tf-test-my-mesh-noport%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
}
`, context)
}

func testAccCheckNetworkServicesMeshDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_mesh" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/meshes/{{name}}")
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
				return fmt.Errorf("NetworkServicesMesh still exists at %s", url)
			}
		}

		return nil
	}
}
