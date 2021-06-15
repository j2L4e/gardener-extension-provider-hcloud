// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package apis is the main package for HCloud specific APIs
package apis

import (
	"strings"
)

// GetRegionFromZone returns the region for a given zone string
func GetRegionFromZone(zone string) string {
	zoneData := strings.SplitN(zone, "-", 2)
	return zoneData[0]
}

// FindRegionSpecForGardenerRegion finds a RegionSpec by name in the cloud profile config
func FindRegionSpecForGardenerRegion(name string, cloudProfileConfig *CloudProfileConfig) *RegionSpec {
	for _, r := range cloudProfileConfig.Regions {
		if r.Name == name {
			return &r
		}
	}
	return nil
}
