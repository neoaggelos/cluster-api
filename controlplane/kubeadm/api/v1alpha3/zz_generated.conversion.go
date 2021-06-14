// +build !ignore_autogenerated_kubeadm_controlplane_v1alpha3

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	clusterapiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	apiv1alpha3 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	v1alpha4 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha4"
	errors "sigs.k8s.io/cluster-api/errors"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlane)(nil), (*v1alpha4.KubeadmControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(a.(*KubeadmControlPlane), b.(*v1alpha4.KubeadmControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.KubeadmControlPlane)(nil), (*KubeadmControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlane_To_v1alpha3_KubeadmControlPlane(a.(*v1alpha4.KubeadmControlPlane), b.(*KubeadmControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneList)(nil), (*v1alpha4.KubeadmControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(a.(*KubeadmControlPlaneList), b.(*v1alpha4.KubeadmControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.KubeadmControlPlaneList)(nil), (*KubeadmControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneList_To_v1alpha3_KubeadmControlPlaneList(a.(*v1alpha4.KubeadmControlPlaneList), b.(*KubeadmControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneStatus)(nil), (*v1alpha4.KubeadmControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(a.(*KubeadmControlPlaneStatus), b.(*v1alpha4.KubeadmControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.KubeadmControlPlaneStatus)(nil), (*KubeadmControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1alpha3_KubeadmControlPlaneStatus(a.(*v1alpha4.KubeadmControlPlaneStatus), b.(*KubeadmControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*KubeadmControlPlaneSpec)(nil), (*v1alpha4.KubeadmControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmControlPlaneSpec_To_v1alpha4_KubeadmControlPlaneSpec(a.(*KubeadmControlPlaneSpec), b.(*v1alpha4.KubeadmControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha4.KubeadmControlPlaneSpec)(nil), (*KubeadmControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1alpha3_KubeadmControlPlaneSpec(a.(*v1alpha4.KubeadmControlPlaneSpec), b.(*KubeadmControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(in *KubeadmControlPlane, out *v1alpha4.KubeadmControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_KubeadmControlPlaneSpec_To_v1alpha4_KubeadmControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(in *KubeadmControlPlane, out *v1alpha4.KubeadmControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlane_To_v1alpha3_KubeadmControlPlane(in *v1alpha4.KubeadmControlPlane, out *KubeadmControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1alpha3_KubeadmControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1alpha3_KubeadmControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlane_To_v1alpha3_KubeadmControlPlane is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlane_To_v1alpha3_KubeadmControlPlane(in *v1alpha4.KubeadmControlPlane, out *KubeadmControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlane_To_v1alpha3_KubeadmControlPlane(in, out, s)
}

func autoConvert_v1alpha3_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(in *KubeadmControlPlaneList, out *v1alpha4.KubeadmControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha4.KubeadmControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(in *KubeadmControlPlaneList, out *v1alpha4.KubeadmControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlaneList_To_v1alpha3_KubeadmControlPlaneList(in *v1alpha4.KubeadmControlPlaneList, out *KubeadmControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_KubeadmControlPlane_To_v1alpha3_KubeadmControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneList_To_v1alpha3_KubeadmControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneList_To_v1alpha3_KubeadmControlPlaneList(in *v1alpha4.KubeadmControlPlaneList, out *KubeadmControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneList_To_v1alpha3_KubeadmControlPlaneList(in, out, s)
}

func autoConvert_v1alpha3_KubeadmControlPlaneSpec_To_v1alpha4_KubeadmControlPlaneSpec(in *KubeadmControlPlaneSpec, out *v1alpha4.KubeadmControlPlaneSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Version = in.Version
	// WARNING: in.InfrastructureTemplate requires manual conversion: does not exist in peer-type
	if err := apiv1alpha3.Convert_v1alpha3_KubeadmConfigSpec_To_v1alpha4_KubeadmConfigSpec(&in.KubeadmConfigSpec, &out.KubeadmConfigSpec, s); err != nil {
		return err
	}
	// WARNING: in.UpgradeAfter requires manual conversion: does not exist in peer-type
	// WARNING: in.NodeDrainTimeout requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha4_KubeadmControlPlaneSpec_To_v1alpha3_KubeadmControlPlaneSpec(in *v1alpha4.KubeadmControlPlaneSpec, out *KubeadmControlPlaneSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Version = in.Version
	// WARNING: in.MachineTemplate requires manual conversion: does not exist in peer-type
	if err := apiv1alpha3.Convert_v1alpha4_KubeadmConfigSpec_To_v1alpha3_KubeadmConfigSpec(&in.KubeadmConfigSpec, &out.KubeadmConfigSpec, s); err != nil {
		return err
	}
	// WARNING: in.RolloutAfter requires manual conversion: does not exist in peer-type
	// WARNING: in.RolloutStrategy requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(in *KubeadmControlPlaneStatus, out *v1alpha4.KubeadmControlPlaneStatus, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.ReadyReplicas = in.ReadyReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureReason = errors.KubeadmControlPlaneStatusError(in.FailureReason)
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			if err := clusterapiapiv1alpha3.Convert_v1alpha3_Condition_To_v1alpha4_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha3_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(in *KubeadmControlPlaneStatus, out *v1alpha4.KubeadmControlPlaneStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlaneStatus_To_v1alpha3_KubeadmControlPlaneStatus(in *v1alpha4.KubeadmControlPlaneStatus, out *KubeadmControlPlaneStatus, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.ReadyReplicas = in.ReadyReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureReason = errors.KubeadmControlPlaneStatusError(in.FailureReason)
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(clusterapiapiv1alpha3.Conditions, len(*in))
		for i := range *in {
			if err := clusterapiapiv1alpha3.Convert_v1alpha4_Condition_To_v1alpha3_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1alpha3_KubeadmControlPlaneStatus is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1alpha3_KubeadmControlPlaneStatus(in *v1alpha4.KubeadmControlPlaneStatus, out *KubeadmControlPlaneStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneStatus_To_v1alpha3_KubeadmControlPlaneStatus(in, out, s)
}
