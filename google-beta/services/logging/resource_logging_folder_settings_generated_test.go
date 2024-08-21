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

package logging_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccLoggingFolderSettings_loggingFolderSettingsAllExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"key_name":      acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingFolderSettings_loggingFolderSettingsAllExample(context),
			},
			{
				ResourceName:            "google_logging_folder_settings.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"folder"},
			},
		},
	})
}

func testAccLoggingFolderSettings_loggingFolderSettingsAllExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_logging_folder_settings" "example" {
  disable_default_sink = true
  folder               = google_folder.my_folder.folder_id
  kms_key_name         = "%{key_name}"
  storage_location     = "us-central1"
  depends_on           = [ google_kms_crypto_key_iam_member.iam ]
}

resource "google_folder" "my_folder" {
  display_name = "tf-test-folder-name%{random_suffix}"
  parent       = "organizations/%{org_id}"
  deletion_protection = false
}

data "google_logging_folder_settings" "settings" {
  folder = google_folder.my_folder.folder_id
}

resource "google_kms_crypto_key_iam_member" "iam" {
  crypto_key_id = "%{key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${data.google_logging_folder_settings.settings.kms_service_account_id}"
}
`, context)
}
