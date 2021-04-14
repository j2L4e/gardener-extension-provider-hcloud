/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ensurer

import (
	"context"
	"fmt"
	"net"

	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

func EnsureNetworks(ctx context.Context, client *hcloud.Client, namespace string, networks *apis.Networks) error {
	if "" != networks.Workers {
		name := fmt.Sprintf("%s-workers", namespace)

		network, _, err := client.Network.GetByName(ctx, name)
		if err != nil {
			return err
		}
		if network == nil {
			_, ipRange, _ := net.ParseCIDR(networks.Workers)

			labels := map[string]string{ "hcloud.provider.extensions.gardener.cloud/role": "workers-network-v1" }

			opts := hcloud.NetworkCreateOpts{
				Name: name,
				IPRange: ipRange,
				Subnets: []hcloud.NetworkSubnet {
					hcloud.NetworkSubnet{
						Type: hcloud.NetworkSubnetTypeCloud,
						IPRange: ipRange,
						NetworkZone: hcloud.NetworkZoneEUCentral, // @TODO: only supported one at time of implementation
					}},
				Labels: labels,
			}

			_, _, err := client.Network.Create(ctx, opts)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
