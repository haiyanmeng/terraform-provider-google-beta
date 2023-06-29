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

func TestAccSQLDatabase_sqlDatabaseBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSQLDatabaseDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSQLDatabase_sqlDatabaseBasicExample(context),
			},
			{
				ResourceName:      "google_sql_database.database",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSQLDatabase_sqlDatabaseBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sql_database" "database" {
  name     = "tf-test-my-database%{random_suffix}"
  instance = google_sql_database_instance.instance.name
}

# See versions at https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance#database_version
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-my-database-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "MYSQL_8_0"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = "%{deletion_protection}"
}
`, context)
}

func TestAccSQLDatabase_sqlDatabaseDeletionPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSQLDatabaseDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSQLDatabase_sqlDatabaseDeletionPolicyExample(context),
			},
			{
				ResourceName:            "google_sql_database.database_deletion_policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy"},
			},
		},
	})
}

func testAccSQLDatabase_sqlDatabaseDeletionPolicyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sql_database" "database_deletion_policy" {
  name     = "tf-test-my-database%{random_suffix}"
  instance = google_sql_database_instance.instance.name
  deletion_policy = "ABANDON"
}

# See versions at https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance#database_version
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-my-database-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "POSTGRES_14"
  settings {
    tier = "db-g1-small"
  }

  deletion_protection  = "%{deletion_protection}"
}
`, context)
}

func testAccCheckSQLDatabaseDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_sql_database" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SQLBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
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
				return fmt.Errorf("SQLDatabase still exists at %s", url)
			}
		}

		return nil
	}
}
