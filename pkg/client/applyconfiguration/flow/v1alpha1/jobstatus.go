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
	batchv1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// JobStatusApplyConfiguration represents a declarative configuration of the JobStatus type for use
// with apply.
type JobStatusApplyConfiguration struct {
	Name             *string                               `json:"name,omitempty"`
	State            *batchv1alpha1.JobPhase               `json:"state,omitempty"`
	StartTimestamp   *v1.Time                              `json:"startTimestamp,omitempty"`
	EndTimestamp     *v1.Time                              `json:"endTimestamp,omitempty"`
	RestartCount     *int32                                `json:"restartCount,omitempty"`
	RunningHistories []JobRunningHistoryApplyConfiguration `json:"runningHistories,omitempty"`
}

// JobStatusApplyConfiguration constructs a declarative configuration of the JobStatus type for use with
// apply.
func JobStatus() *JobStatusApplyConfiguration {
	return &JobStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithName(value string) *JobStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithState(value batchv1alpha1.JobPhase) *JobStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithStartTimestamp sets the StartTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTimestamp field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithStartTimestamp(value v1.Time) *JobStatusApplyConfiguration {
	b.StartTimestamp = &value
	return b
}

// WithEndTimestamp sets the EndTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndTimestamp field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithEndTimestamp(value v1.Time) *JobStatusApplyConfiguration {
	b.EndTimestamp = &value
	return b
}

// WithRestartCount sets the RestartCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartCount field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithRestartCount(value int32) *JobStatusApplyConfiguration {
	b.RestartCount = &value
	return b
}

// WithRunningHistories adds the given value to the RunningHistories field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RunningHistories field.
func (b *JobStatusApplyConfiguration) WithRunningHistories(values ...*JobRunningHistoryApplyConfiguration) *JobStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRunningHistories")
		}
		b.RunningHistories = append(b.RunningHistories, *values[i])
	}
	return b
}
