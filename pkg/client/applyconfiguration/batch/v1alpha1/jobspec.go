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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JobSpecApplyConfiguration represents a declarative configuration of the JobSpec type for use
// with apply.
type JobSpecApplyConfiguration struct {
	SchedulerName           *string                                `json:"schedulerName,omitempty"`
	MinAvailable            *int32                                 `json:"minAvailable,omitempty"`
	Volumes                 []VolumeSpecApplyConfiguration         `json:"volumes,omitempty"`
	Tasks                   []TaskSpecApplyConfiguration           `json:"tasks,omitempty"`
	Policies                []LifecyclePolicyApplyConfiguration    `json:"policies,omitempty"`
	Plugins                 map[string][]string                    `json:"plugins,omitempty"`
	RunningEstimate         *v1.Duration                           `json:"runningEstimate,omitempty"`
	Queue                   *string                                `json:"queue,omitempty"`
	MaxRetry                *int32                                 `json:"maxRetry,omitempty"`
	TTLSecondsAfterFinished *int32                                 `json:"ttlSecondsAfterFinished,omitempty"`
	PriorityClassName       *string                                `json:"priorityClassName,omitempty"`
	MinSuccess              *int32                                 `json:"minSuccess,omitempty"`
	NetworkTopology         *NetworkTopologySpecApplyConfiguration `json:"networkTopology,omitempty"`
}

// JobSpecApplyConfiguration constructs a declarative configuration of the JobSpec type for use with
// apply.
func JobSpec() *JobSpecApplyConfiguration {
	return &JobSpecApplyConfiguration{}
}

// WithSchedulerName sets the SchedulerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerName field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithSchedulerName(value string) *JobSpecApplyConfiguration {
	b.SchedulerName = &value
	return b
}

// WithMinAvailable sets the MinAvailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinAvailable field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithMinAvailable(value int32) *JobSpecApplyConfiguration {
	b.MinAvailable = &value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *JobSpecApplyConfiguration) WithVolumes(values ...*VolumeSpecApplyConfiguration) *JobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumes")
		}
		b.Volumes = append(b.Volumes, *values[i])
	}
	return b
}

// WithTasks adds the given value to the Tasks field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tasks field.
func (b *JobSpecApplyConfiguration) WithTasks(values ...*TaskSpecApplyConfiguration) *JobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTasks")
		}
		b.Tasks = append(b.Tasks, *values[i])
	}
	return b
}

// WithPolicies adds the given value to the Policies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Policies field.
func (b *JobSpecApplyConfiguration) WithPolicies(values ...*LifecyclePolicyApplyConfiguration) *JobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPolicies")
		}
		b.Policies = append(b.Policies, *values[i])
	}
	return b
}

// WithPlugins puts the entries into the Plugins field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Plugins field,
// overwriting an existing map entries in Plugins field with the same key.
func (b *JobSpecApplyConfiguration) WithPlugins(entries map[string][]string) *JobSpecApplyConfiguration {
	if b.Plugins == nil && len(entries) > 0 {
		b.Plugins = make(map[string][]string, len(entries))
	}
	for k, v := range entries {
		b.Plugins[k] = v
	}
	return b
}

// WithRunningEstimate sets the RunningEstimate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunningEstimate field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithRunningEstimate(value v1.Duration) *JobSpecApplyConfiguration {
	b.RunningEstimate = &value
	return b
}

// WithQueue sets the Queue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Queue field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithQueue(value string) *JobSpecApplyConfiguration {
	b.Queue = &value
	return b
}

// WithMaxRetry sets the MaxRetry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxRetry field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithMaxRetry(value int32) *JobSpecApplyConfiguration {
	b.MaxRetry = &value
	return b
}

// WithTTLSecondsAfterFinished sets the TTLSecondsAfterFinished field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TTLSecondsAfterFinished field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithTTLSecondsAfterFinished(value int32) *JobSpecApplyConfiguration {
	b.TTLSecondsAfterFinished = &value
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithPriorityClassName(value string) *JobSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithMinSuccess sets the MinSuccess field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinSuccess field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithMinSuccess(value int32) *JobSpecApplyConfiguration {
	b.MinSuccess = &value
	return b
}

// WithNetworkTopology sets the NetworkTopology field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkTopology field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithNetworkTopology(value *NetworkTopologySpecApplyConfiguration) *JobSpecApplyConfiguration {
	b.NetworkTopology = value
	return b
}
