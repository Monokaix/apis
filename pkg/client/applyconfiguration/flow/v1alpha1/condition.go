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
	v1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	batchv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/batch/v1alpha1"
)

// ConditionApplyConfiguration represents a declarative configuration of the Condition type for use
// with apply.
type ConditionApplyConfiguration struct {
	Phase           *v1alpha1.JobPhase                                   `json:"phase,omitempty"`
	CreateTimestamp *v1.Time                                             `json:"createTime,omitempty"`
	RunningDuration *v1.Duration                                         `json:"runningDuration,omitempty"`
	TaskStatusCount map[string]batchv1alpha1.TaskStateApplyConfiguration `json:"taskStatusCount,omitempty"`
}

// ConditionApplyConfiguration constructs a declarative configuration of the Condition type for use with
// apply.
func Condition() *ConditionApplyConfiguration {
	return &ConditionApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithPhase(value v1alpha1.JobPhase) *ConditionApplyConfiguration {
	b.Phase = &value
	return b
}

// WithCreateTimestamp sets the CreateTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreateTimestamp field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithCreateTimestamp(value v1.Time) *ConditionApplyConfiguration {
	b.CreateTimestamp = &value
	return b
}

// WithRunningDuration sets the RunningDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunningDuration field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithRunningDuration(value v1.Duration) *ConditionApplyConfiguration {
	b.RunningDuration = &value
	return b
}

// WithTaskStatusCount puts the entries into the TaskStatusCount field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the TaskStatusCount field,
// overwriting an existing map entries in TaskStatusCount field with the same key.
func (b *ConditionApplyConfiguration) WithTaskStatusCount(entries map[string]batchv1alpha1.TaskStateApplyConfiguration) *ConditionApplyConfiguration {
	if b.TaskStatusCount == nil && len(entries) > 0 {
		b.TaskStatusCount = make(map[string]batchv1alpha1.TaskStateApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.TaskStatusCount[k] = v
	}
	return b
}
