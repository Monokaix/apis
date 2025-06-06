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
	batchv1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// DependsOnApplyConfiguration represents a declarative configuration of the DependsOn type for use
// with apply.
type DependsOnApplyConfiguration struct {
	Name      []string                 `json:"name,omitempty"`
	Iteration *batchv1alpha1.Iteration `json:"iteration,omitempty"`
}

// DependsOnApplyConfiguration constructs a declarative configuration of the DependsOn type for use with
// apply.
func DependsOn() *DependsOnApplyConfiguration {
	return &DependsOnApplyConfiguration{}
}

// WithName adds the given value to the Name field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Name field.
func (b *DependsOnApplyConfiguration) WithName(values ...string) *DependsOnApplyConfiguration {
	for i := range values {
		b.Name = append(b.Name, values[i])
	}
	return b
}

// WithIteration sets the Iteration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Iteration field is set to the value of the last call.
func (b *DependsOnApplyConfiguration) WithIteration(value batchv1alpha1.Iteration) *DependsOnApplyConfiguration {
	b.Iteration = &value
	return b
}
