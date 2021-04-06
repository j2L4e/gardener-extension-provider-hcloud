/*
 * Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package controlplane

import (
	"context"
	"errors"

	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis"
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis/transcoder"
)

func decodeCloudProfileConfigFromGardenContext(ctx context.Context, webhookcontext webhookcontext.GardenContext) (*apis.CloudProfileConfig, error) {
	cluster, err := webhookcontext.GetCluster(ctx)
	if err != nil {
		return nil, err
	}

	cloudProfileConfig, err := transcoder.DecodeConfigFromCloudProfile(cluster.CloudProfile)
	if err != nil {
		return nil, errors.Wrapf(err, "decoding cloudprofileconfig failed")
	}

	return cloudProfileConfig, nil
}