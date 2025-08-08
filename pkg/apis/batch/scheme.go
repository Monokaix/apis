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

package batch

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	v1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	v1alpha2 "volcano.sh/apis/pkg/apis/batch/v1alpha2"
)

var (
	// Scheme is the runtime scheme for the batch API group
	Scheme = runtime.NewScheme()
	// SchemeGroupVersion is the group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
)

func init() {
	// Add both versions to the scheme
	utilruntime.Must(v1alpha1.AddToScheme(Scheme))
	utilruntime.Must(v1alpha2.AddToScheme(Scheme))
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
