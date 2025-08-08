/*
Copyright 2024 The Volcano Authors.

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

// +k8s:conversion-gen=volcano.sh/apis/pkg/apis/batch
// +k8s:conversion-gen:explicit-conversion=true

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"volcano.sh/apis/pkg/apis/batch/v1alpha2"
)

const (
	// TaskSizeAnnotation is the annotation key for task size.
	TaskSizeAnnotation = "batch.volcano.sh/task-size"
)

func Convert_v1alpha1_JobSpec_To_v1alpha2_JobSpec(in *JobSpec, out *v1alpha2.JobSpec, s conversion.Scope) error {
	// Copy common fields
	out.SchedulerName = in.SchedulerName
	out.MinAvailable = in.MinAvailable

	// Convert Volumes field by field
	out.Volumes = make([]v1alpha2.VolumeSpec, len(in.Volumes))
	for i := range in.Volumes {
		out.Volumes[i].MountPath = in.Volumes[i].MountPath
		out.Volumes[i].VolumeClaimName = in.Volumes[i].VolumeClaimName
		out.Volumes[i].VolumeClaim = in.Volumes[i].VolumeClaim
	}

	// Convert Policies field by field
	out.Policies = make([]v1alpha2.LifecyclePolicy, len(in.Policies))
	for i := range in.Policies {
		out.Policies[i].Action = in.Policies[i].Action
		out.Policies[i].Event = in.Policies[i].Event
		out.Policies[i].Events = in.Policies[i].Events
		out.Policies[i].ExitCode = in.Policies[i].ExitCode
		out.Policies[i].Timeout = in.Policies[i].Timeout
	}

	out.Plugins = in.Plugins
	out.RunningEstimate = in.RunningEstimate
	out.Queue = in.Queue
	out.MaxRetry = in.MaxRetry
	out.TTLSecondsAfterFinished = in.TTLSecondsAfterFinished
	out.PriorityClassName = in.PriorityClassName
	out.MinSuccess = in.MinSuccess

	// Convert NetworkTopology field by field
	if in.NetworkTopology != nil {
		out.NetworkTopology = &v1alpha2.NetworkTopologySpec{
			Mode:               v1alpha2.NetworkTopologyMode(in.NetworkTopology.Mode),
			HighestTierAllowed: in.NetworkTopology.HighestTierAllowed,
		}
	}

	// Convert Tasks field
	out.Tasks = make([]v1alpha2.TaskSpecWithSize, len(in.Tasks))
	for i := range in.Tasks {
		if err := Convert_v1alpha1_TaskSpec_To_v1alpha2_TaskSpecWithSize(&in.Tasks[i], &out.Tasks[i], s); err != nil {
			return err
		}
	}

	return nil
}

func Convert_v1alpha2_JobSpec_To_v1alpha1_JobSpec(in *v1alpha2.JobSpec, out *JobSpec, s conversion.Scope) error {
	// Copy common fields
	out.SchedulerName = in.SchedulerName
	out.MinAvailable = in.MinAvailable

	// Convert Volumes field by field
	out.Volumes = make([]VolumeSpec, len(in.Volumes))
	for i := range in.Volumes {
		out.Volumes[i].MountPath = in.Volumes[i].MountPath
		out.Volumes[i].VolumeClaimName = in.Volumes[i].VolumeClaimName
		out.Volumes[i].VolumeClaim = in.Volumes[i].VolumeClaim
	}

	// Convert Policies field by field
	out.Policies = make([]LifecyclePolicy, len(in.Policies))
	for i := range in.Policies {
		out.Policies[i].Action = in.Policies[i].Action
		out.Policies[i].Event = in.Policies[i].Event
		out.Policies[i].Events = in.Policies[i].Events
		out.Policies[i].ExitCode = in.Policies[i].ExitCode
		out.Policies[i].Timeout = in.Policies[i].Timeout
	}

	out.Plugins = in.Plugins
	out.RunningEstimate = in.RunningEstimate
	out.Queue = in.Queue
	out.MaxRetry = in.MaxRetry
	out.TTLSecondsAfterFinished = in.TTLSecondsAfterFinished
	out.PriorityClassName = in.PriorityClassName
	out.MinSuccess = in.MinSuccess

	// Convert NetworkTopology field by field
	if in.NetworkTopology != nil {
		out.NetworkTopology = &NetworkTopologySpec{
			Mode:               NetworkTopologyMode(in.NetworkTopology.Mode),
			HighestTierAllowed: in.NetworkTopology.HighestTierAllowed,
		}
	}

	// Convert Tasks field
	out.Tasks = make([]TaskSpec, len(in.Tasks))
	for i := range in.Tasks {
		if err := Convert_v1alpha2_TaskSpecWithSize_To_v1alpha1_TaskSpec(&in.Tasks[i], &out.Tasks[i], s); err != nil {
			return err
		}
	}

	return nil
}

func Convert_v1alpha1_TaskSpec_To_v1alpha2_TaskSpecWithSize(in *TaskSpec, out *v1alpha2.TaskSpecWithSize, _ conversion.Scope) error {
	// Copy all TaskSpec fields directly since TaskSpec is inlined in TaskSpecWithSize
	out.TaskSpec.Name = in.Name
	out.TaskSpec.Replicas = in.Replicas
	out.TaskSpec.MinAvailable = in.MinAvailable
	out.TaskSpec.Template = in.Template

	// Convert Policies field by field
	out.TaskSpec.Policies = make([]v1alpha2.LifecyclePolicy, len(in.Policies))
	for i := range in.Policies {
		out.TaskSpec.Policies[i].Action = in.Policies[i].Action
		out.TaskSpec.Policies[i].Event = in.Policies[i].Event
		out.TaskSpec.Policies[i].Events = in.Policies[i].Events
		out.TaskSpec.Policies[i].ExitCode = in.Policies[i].ExitCode
		out.TaskSpec.Policies[i].Timeout = in.Policies[i].Timeout
	}

	out.TaskSpec.TopologyPolicy = v1alpha2.NumaPolicy(in.TopologyPolicy)
	out.TaskSpec.MaxRetry = in.MaxRetry

	// Convert DependsOn field by field
	if in.DependsOn != nil {
		out.TaskSpec.DependsOn = &v1alpha2.DependsOn{
			Name:      in.DependsOn.Name,
			Iteration: v1alpha2.Iteration(in.DependsOn.Iteration),
		}
	}

	// v1alpha1 doesn't have Size concept, set to 1 as default
	out.Size = 1
	return nil
}

func Convert_v1alpha2_TaskSpecWithSize_To_v1alpha1_TaskSpec(in *v1alpha2.TaskSpecWithSize, out *TaskSpec, _ conversion.Scope) error {
	// Copy all TaskSpec fields directly from the inlined TaskSpec
	out.Name = in.TaskSpec.Name
	out.Replicas = in.TaskSpec.Replicas
	out.MinAvailable = in.TaskSpec.MinAvailable
	out.Template = in.TaskSpec.Template

	// Convert Policies field by field
	out.Policies = make([]LifecyclePolicy, len(in.TaskSpec.Policies))
	for i := range in.TaskSpec.Policies {
		out.Policies[i].Action = in.TaskSpec.Policies[i].Action
		out.Policies[i].Event = in.TaskSpec.Policies[i].Event
		out.Policies[i].Events = in.TaskSpec.Policies[i].Events
		out.Policies[i].ExitCode = in.TaskSpec.Policies[i].ExitCode
		out.Policies[i].Timeout = in.TaskSpec.Policies[i].Timeout
	}

	out.TopologyPolicy = NumaPolicy(in.TaskSpec.TopologyPolicy)
	out.MaxRetry = in.TaskSpec.MaxRetry

	// Convert DependsOn field by field
	if in.TaskSpec.DependsOn != nil {
		out.DependsOn = &DependsOn{
			Name:      in.TaskSpec.DependsOn.Name,
			Iteration: Iteration(in.TaskSpec.DependsOn.Iteration),
		}
	}

	// v1alpha1 doesn't have Size field, so we simply ignore the Size value
	// The Size information is lost during this conversion, which is expected
	return nil
}
