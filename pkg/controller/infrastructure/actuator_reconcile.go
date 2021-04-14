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
	"fmt"

	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/controller/infrastructure/ensurer"
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud"
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis"
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis/helper"
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis/transcoder"
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

type preparedReconcile struct {
	cloudProfileConfig *apis.CloudProfileConfig
	infraConfig        *apis.InfrastructureConfig
	region             *apis.RegionSpec
	token              string
}

func (a *actuator) prepareReconcile(ctx context.Context, infra *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) (*preparedReconcile, error) {
	cloudProfileConfig, err := transcoder.DecodeCloudProfileConfigFromControllerCluster(cluster)
	if err != nil {
		return nil, err
	}

	infraConfig, err := transcoder.DecodeInfrastructureConfigFromInfrastructure(infra)
	if err != nil {
		return nil, err
	}

	region := helper.FindRegion(infra.Spec.Region, cloudProfileConfig)
	if region == nil {
		return nil, fmt.Errorf("Region %q not found in cloud profile", infra.Spec.Region)
	}

	secret, err := extensionscontroller.GetSecretByReference(ctx, a.Client(), &infra.Spec.SecretRef)
	if err != nil {
		return nil, err
	}

	credentials, err := hcloud.ExtractCredentials(secret)
	if err != nil {
		return nil, err
	}

	token := credentials.HcloudCCM().HcloudToken

	prepared := &preparedReconcile{
		cloudProfileConfig: cloudProfileConfig,
		infraConfig:        infraConfig,
		region:             region,
		token:              token,
	}

	return prepared, nil
}

func (a *actuator) reconcile(ctx context.Context, infra *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	prepared, err := a.prepareReconcile(ctx, infra, cluster)
	if err != nil {
		return err
	}

	client := apis.GetClientForToken(string(prepared.token))

	err = ensurer.EnsureSSHPublicKey(ctx, client, infra.Spec.SSHPublicKey)
	if err != nil {
		return err
	}

	err = ensurer.EnsureNetworks(ctx, client, infra.Namespace, prepared.infraConfig.Networks)
	if err != nil {
		return err
	}

	return a.updateProviderStatus(ctx, infra)
}
