/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type FileProperties struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// File creation time in epoch milliseconds
	CreatedEpochMs int64 `json:"created_epoch_ms"`

	// File modification time in epoch milliseconds
	ModifiedEpochMs int64 `json:"modified_epoch_ms"`

	// File name
	Name string `json:"name"`

	// Size of the file in bytes
	Size int64 `json:"size"`
}
