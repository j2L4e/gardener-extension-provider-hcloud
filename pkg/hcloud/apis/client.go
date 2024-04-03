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

// Package apis is the main package for HCloud specific APIs
package apis

import (
	"os"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

var singletons = make(map[string]*hcloud.Client)

// GetClientForToken returns an underlying HCloud client for the given token.
//
// PARAMETERS
// token string Token to look up client instance for
func GetClientForToken(token string) *hcloud.Client {
	client, ok := singletons[token]

	if !ok {
		opts := []hcloud.ClientOption{
			hcloud.WithToken(token),
			hcloud.WithApplication("gardener-extension-provider-hcloud", "v0.0.0"),
		}
		if endpoint := os.Getenv("HCLOUD_ENDPOINT"); endpoint != "" {
			opts = append(opts, hcloud.WithEndpoint(endpoint))
		}
		client = hcloud.NewClient(opts...)
	}

	return client
}

// SetClientForToken sets a preconfigured HCloud client for the given token.
//
// PARAMETERS
// token  string         Token to look up client instance for
// client *hcloud.Client Preconfigured HCloud client
func SetClientForToken(token string, client *hcloud.Client) {
	if client == nil {
		delete(singletons, token)
	} else {
		singletons[token] = client
	}
}
