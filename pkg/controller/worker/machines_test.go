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

// Package worker provides the HCloud implementation for worker machines
package worker

import (
	"github.com/23technologies/gardener-extension-provider-hcloud/pkg/hcloud/apis/mock"
	"github.com/gardener/gardener/extensions/pkg/controller/common"
	"github.com/gardener/gardener/extensions/pkg/controller/worker/genericactuator"
	mcmv1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Machines", func() {
	var mockTestEnv    mock.MockTestEnv
	var workerDelegate genericactuator.WorkerDelegate

	var _ = BeforeSuite(func() {
		mockTestEnv = mock.NewMockTestEnv()

		newWorkerDelegate, err := NewWorkerDelegate(common.NewClientContext(nil, nil, nil), nil, "", nil, nil)
		Expect(err).NotTo(HaveOccurred())

		workerDelegate = newWorkerDelegate
	})

	var _ = AfterSuite(func() {
		mockTestEnv.Teardown()
	})

	Describe("#MachineClass", func() {
		It("should return the correct kind of the machine class", func() {
			Expect(workerDelegate.MachineClass()).To(Equal(&mcmv1alpha1.MachineClass{}))
		})
	})

	Describe("#MachineClassKind", func() {
		It("should return the correct kind of the machine class", func() {
			Expect(workerDelegate.MachineClassKind()).To(Equal("MachineClass"))
		})
	})
})