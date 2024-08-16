/*
Copyright The Volcano Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "volcano.sh/apis/pkg/apis/nodeinfo/v1alpha1"
)

// NumatopoSpecApplyConfiguration represents an declarative configuration of the NumatopoSpec type for use
// with apply.
type NumatopoSpecApplyConfiguration struct {
	Policies    map[v1alpha1.PolicyName]string            `json:"policies,omitempty"`
	ResReserved map[string]string                         `json:"resReserved,omitempty"`
	NumaResMap  map[string]ResourceInfoApplyConfiguration `json:"numares,omitempty"`
	CPUDetail   map[string]CPUInfoApplyConfiguration      `json:"cpuDetail,omitempty"`
}

// NumatopoSpecApplyConfiguration constructs an declarative configuration of the NumatopoSpec type for use with
// apply.
func NumatopoSpec() *NumatopoSpecApplyConfiguration {
	return &NumatopoSpecApplyConfiguration{}
}

// WithPolicies puts the entries into the Policies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Policies field,
// overwriting an existing map entries in Policies field with the same key.
func (b *NumatopoSpecApplyConfiguration) WithPolicies(entries map[v1alpha1.PolicyName]string) *NumatopoSpecApplyConfiguration {
	if b.Policies == nil && len(entries) > 0 {
		b.Policies = make(map[v1alpha1.PolicyName]string, len(entries))
	}
	for k, v := range entries {
		b.Policies[k] = v
	}
	return b
}

// WithResReserved puts the entries into the ResReserved field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ResReserved field,
// overwriting an existing map entries in ResReserved field with the same key.
func (b *NumatopoSpecApplyConfiguration) WithResReserved(entries map[string]string) *NumatopoSpecApplyConfiguration {
	if b.ResReserved == nil && len(entries) > 0 {
		b.ResReserved = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ResReserved[k] = v
	}
	return b
}

// WithNumaResMap puts the entries into the NumaResMap field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NumaResMap field,
// overwriting an existing map entries in NumaResMap field with the same key.
func (b *NumatopoSpecApplyConfiguration) WithNumaResMap(entries map[string]ResourceInfoApplyConfiguration) *NumatopoSpecApplyConfiguration {
	if b.NumaResMap == nil && len(entries) > 0 {
		b.NumaResMap = make(map[string]ResourceInfoApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.NumaResMap[k] = v
	}
	return b
}

// WithCPUDetail puts the entries into the CPUDetail field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the CPUDetail field,
// overwriting an existing map entries in CPUDetail field with the same key.
func (b *NumatopoSpecApplyConfiguration) WithCPUDetail(entries map[string]CPUInfoApplyConfiguration) *NumatopoSpecApplyConfiguration {
	if b.CPUDetail == nil && len(entries) > 0 {
		b.CPUDetail = make(map[string]CPUInfoApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.CPUDetail[k] = v
	}
	return b
}