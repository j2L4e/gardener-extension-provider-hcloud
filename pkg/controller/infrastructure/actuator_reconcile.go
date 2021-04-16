/*
 * Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 *
 */

package infrastructure

import (
	"context"

	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/controller/infrastructure/ensurer"
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis"
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

func (a *actuator) reconcile(ctx context.Context, infra *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	actuatorConfig, err := a.getActuatorConfig(ctx, infra, cluster)
	if err != nil {
		return err
	}

	client := apis.GetClientForToken(string(actuatorConfig.token))

	err = ensurer.EnsureSSHPublicKey(ctx, client, infra.Spec.SSHPublicKey)
	if err != nil {
		return err
	}

	err = ensurer.EnsureNetworks(ctx, client, infra.Namespace, actuatorConfig.infraConfig.Networks)
	if err != nil {
		return err
	}
/*
	opts := hcloudclient.ImageListOpts{
		IncludeDeprecated: true,
	}

	images, _, err := client.Image.List(ctx, opts)
	if err != nil {
		return err
	}
	panic(images)
*/
	return a.updateProviderStatus(ctx, infra)
}
