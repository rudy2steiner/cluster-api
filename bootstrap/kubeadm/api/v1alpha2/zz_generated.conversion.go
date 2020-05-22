// +build !ignore_autogenerated

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

package v1alpha2

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha3 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	v1beta1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*File)(nil), (*v1alpha3.File)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_File_To_v1alpha3_File(a.(*File), b.(*v1alpha3.File), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmConfig)(nil), (*v1alpha3.KubeadmConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfig_To_v1alpha3_KubeadmConfig(a.(*KubeadmConfig), b.(*v1alpha3.KubeadmConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.KubeadmConfig)(nil), (*KubeadmConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfig_To_v1alpha2_KubeadmConfig(a.(*v1alpha3.KubeadmConfig), b.(*KubeadmConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmConfigList)(nil), (*v1alpha3.KubeadmConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigList_To_v1alpha3_KubeadmConfigList(a.(*KubeadmConfigList), b.(*v1alpha3.KubeadmConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.KubeadmConfigList)(nil), (*KubeadmConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigList_To_v1alpha2_KubeadmConfigList(a.(*v1alpha3.KubeadmConfigList), b.(*KubeadmConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmConfigTemplate)(nil), (*v1alpha3.KubeadmConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate(a.(*KubeadmConfigTemplate), b.(*v1alpha3.KubeadmConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.KubeadmConfigTemplate)(nil), (*KubeadmConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigTemplate_To_v1alpha2_KubeadmConfigTemplate(a.(*v1alpha3.KubeadmConfigTemplate), b.(*KubeadmConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmConfigTemplateList)(nil), (*v1alpha3.KubeadmConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigTemplateList_To_v1alpha3_KubeadmConfigTemplateList(a.(*KubeadmConfigTemplateList), b.(*v1alpha3.KubeadmConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.KubeadmConfigTemplateList)(nil), (*KubeadmConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigTemplateList_To_v1alpha2_KubeadmConfigTemplateList(a.(*v1alpha3.KubeadmConfigTemplateList), b.(*KubeadmConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmConfigTemplateResource)(nil), (*v1alpha3.KubeadmConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigTemplateResource_To_v1alpha3_KubeadmConfigTemplateResource(a.(*KubeadmConfigTemplateResource), b.(*v1alpha3.KubeadmConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.KubeadmConfigTemplateResource)(nil), (*KubeadmConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigTemplateResource_To_v1alpha2_KubeadmConfigTemplateResource(a.(*v1alpha3.KubeadmConfigTemplateResource), b.(*KubeadmConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmConfigTemplateSpec)(nil), (*v1alpha3.KubeadmConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigTemplateSpec_To_v1alpha3_KubeadmConfigTemplateSpec(a.(*KubeadmConfigTemplateSpec), b.(*v1alpha3.KubeadmConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.KubeadmConfigTemplateSpec)(nil), (*KubeadmConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigTemplateSpec_To_v1alpha2_KubeadmConfigTemplateSpec(a.(*v1alpha3.KubeadmConfigTemplateSpec), b.(*KubeadmConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NTP)(nil), (*v1alpha3.NTP)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_NTP_To_v1alpha3_NTP(a.(*NTP), b.(*v1alpha3.NTP), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.NTP)(nil), (*NTP)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_NTP_To_v1alpha2_NTP(a.(*v1alpha3.NTP), b.(*NTP), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*User)(nil), (*v1alpha3.User)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_User_To_v1alpha3_User(a.(*User), b.(*v1alpha3.User), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.User)(nil), (*User)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_User_To_v1alpha2_User(a.(*v1alpha3.User), b.(*User), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*KubeadmConfigSpec)(nil), (*v1alpha3.KubeadmConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigSpec_To_v1alpha3_KubeadmConfigSpec(a.(*KubeadmConfigSpec), b.(*v1alpha3.KubeadmConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*KubeadmConfigStatus)(nil), (*v1alpha3.KubeadmConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeadmConfigStatus_To_v1alpha3_KubeadmConfigStatus(a.(*KubeadmConfigStatus), b.(*v1alpha3.KubeadmConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.File)(nil), (*File)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_File_To_v1alpha2_File(a.(*v1alpha3.File), b.(*File), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.KubeadmConfigSpec)(nil), (*KubeadmConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigSpec_To_v1alpha2_KubeadmConfigSpec(a.(*v1alpha3.KubeadmConfigSpec), b.(*KubeadmConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.KubeadmConfigStatus)(nil), (*KubeadmConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubeadmConfigStatus_To_v1alpha2_KubeadmConfigStatus(a.(*v1alpha3.KubeadmConfigStatus), b.(*KubeadmConfigStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_File_To_v1alpha3_File(in *File, out *v1alpha3.File, s conversion.Scope) error {
	out.Path = in.Path
	out.Owner = in.Owner
	out.Permissions = in.Permissions
	out.Encoding = v1alpha3.Encoding(in.Encoding)
	out.Content = in.Content
	return nil
}

// Convert_v1alpha2_File_To_v1alpha3_File is an autogenerated conversion function.
func Convert_v1alpha2_File_To_v1alpha3_File(in *File, out *v1alpha3.File, s conversion.Scope) error {
	return autoConvert_v1alpha2_File_To_v1alpha3_File(in, out, s)
}

func autoConvert_v1alpha3_File_To_v1alpha2_File(in *v1alpha3.File, out *File, s conversion.Scope) error {
	out.Path = in.Path
	out.Owner = in.Owner
	out.Permissions = in.Permissions
	out.Encoding = Encoding(in.Encoding)
	out.Content = in.Content
	// WARNING: in.ContentFrom requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_KubeadmConfig_To_v1alpha3_KubeadmConfig(in *KubeadmConfig, out *v1alpha3.KubeadmConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_KubeadmConfigSpec_To_v1alpha3_KubeadmConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeadmConfigStatus_To_v1alpha3_KubeadmConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_KubeadmConfig_To_v1alpha3_KubeadmConfig is an autogenerated conversion function.
func Convert_v1alpha2_KubeadmConfig_To_v1alpha3_KubeadmConfig(in *KubeadmConfig, out *v1alpha3.KubeadmConfig, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeadmConfig_To_v1alpha3_KubeadmConfig(in, out, s)
}

func autoConvert_v1alpha3_KubeadmConfig_To_v1alpha2_KubeadmConfig(in *v1alpha3.KubeadmConfig, out *KubeadmConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_KubeadmConfigSpec_To_v1alpha2_KubeadmConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_KubeadmConfigStatus_To_v1alpha2_KubeadmConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_KubeadmConfig_To_v1alpha2_KubeadmConfig is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmConfig_To_v1alpha2_KubeadmConfig(in *v1alpha3.KubeadmConfig, out *KubeadmConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmConfig_To_v1alpha2_KubeadmConfig(in, out, s)
}

func autoConvert_v1alpha2_KubeadmConfigList_To_v1alpha3_KubeadmConfigList(in *KubeadmConfigList, out *v1alpha3.KubeadmConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.KubeadmConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_KubeadmConfig_To_v1alpha3_KubeadmConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_KubeadmConfigList_To_v1alpha3_KubeadmConfigList is an autogenerated conversion function.
func Convert_v1alpha2_KubeadmConfigList_To_v1alpha3_KubeadmConfigList(in *KubeadmConfigList, out *v1alpha3.KubeadmConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeadmConfigList_To_v1alpha3_KubeadmConfigList(in, out, s)
}

func autoConvert_v1alpha3_KubeadmConfigList_To_v1alpha2_KubeadmConfigList(in *v1alpha3.KubeadmConfigList, out *KubeadmConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_KubeadmConfig_To_v1alpha2_KubeadmConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_KubeadmConfigList_To_v1alpha2_KubeadmConfigList is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmConfigList_To_v1alpha2_KubeadmConfigList(in *v1alpha3.KubeadmConfigList, out *KubeadmConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmConfigList_To_v1alpha2_KubeadmConfigList(in, out, s)
}

func autoConvert_v1alpha2_KubeadmConfigSpec_To_v1alpha3_KubeadmConfigSpec(in *KubeadmConfigSpec, out *v1alpha3.KubeadmConfigSpec, s conversion.Scope) error {
	out.ClusterConfiguration = (*v1beta1.ClusterConfiguration)(unsafe.Pointer(in.ClusterConfiguration))
	out.InitConfiguration = (*v1beta1.InitConfiguration)(unsafe.Pointer(in.InitConfiguration))
	out.JoinConfiguration = (*v1beta1.JoinConfiguration)(unsafe.Pointer(in.JoinConfiguration))
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]v1alpha3.File, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_File_To_v1alpha3_File(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Files = nil
	}
	out.PreKubeadmCommands = *(*[]string)(unsafe.Pointer(&in.PreKubeadmCommands))
	out.PostKubeadmCommands = *(*[]string)(unsafe.Pointer(&in.PostKubeadmCommands))
	out.Users = *(*[]v1alpha3.User)(unsafe.Pointer(&in.Users))
	out.NTP = (*v1alpha3.NTP)(unsafe.Pointer(in.NTP))
	out.Format = v1alpha3.Format(in.Format)
	return nil
}

func autoConvert_v1alpha3_KubeadmConfigSpec_To_v1alpha2_KubeadmConfigSpec(in *v1alpha3.KubeadmConfigSpec, out *KubeadmConfigSpec, s conversion.Scope) error {
	out.ClusterConfiguration = (*v1beta1.ClusterConfiguration)(unsafe.Pointer(in.ClusterConfiguration))
	out.InitConfiguration = (*v1beta1.InitConfiguration)(unsafe.Pointer(in.InitConfiguration))
	out.JoinConfiguration = (*v1beta1.JoinConfiguration)(unsafe.Pointer(in.JoinConfiguration))
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_File_To_v1alpha2_File(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Files = nil
	}
	out.PreKubeadmCommands = *(*[]string)(unsafe.Pointer(&in.PreKubeadmCommands))
	out.PostKubeadmCommands = *(*[]string)(unsafe.Pointer(&in.PostKubeadmCommands))
	out.Users = *(*[]User)(unsafe.Pointer(&in.Users))
	out.NTP = (*NTP)(unsafe.Pointer(in.NTP))
	out.Format = Format(in.Format)
	// WARNING: in.Verbosity requires manual conversion: does not exist in peer-type
	// WARNING: in.UseExperimentalRetryJoin requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_KubeadmConfigStatus_To_v1alpha3_KubeadmConfigStatus(in *KubeadmConfigStatus, out *v1alpha3.KubeadmConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.BootstrapData = *(*[]byte)(unsafe.Pointer(&in.BootstrapData))
	// WARNING: in.ErrorReason requires manual conversion: does not exist in peer-type
	// WARNING: in.ErrorMessage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_KubeadmConfigStatus_To_v1alpha2_KubeadmConfigStatus(in *v1alpha3.KubeadmConfigStatus, out *KubeadmConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.DataSecretName requires manual conversion: does not exist in peer-type
	out.BootstrapData = *(*[]byte)(unsafe.Pointer(&in.BootstrapData))
	// WARNING: in.FailureReason requires manual conversion: does not exist in peer-type
	// WARNING: in.FailureMessage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate(in *KubeadmConfigTemplate, out *v1alpha3.KubeadmConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_KubeadmConfigTemplateSpec_To_v1alpha3_KubeadmConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate is an autogenerated conversion function.
func Convert_v1alpha2_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate(in *KubeadmConfigTemplate, out *v1alpha3.KubeadmConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate(in, out, s)
}

func autoConvert_v1alpha3_KubeadmConfigTemplate_To_v1alpha2_KubeadmConfigTemplate(in *v1alpha3.KubeadmConfigTemplate, out *KubeadmConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_KubeadmConfigTemplateSpec_To_v1alpha2_KubeadmConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_KubeadmConfigTemplate_To_v1alpha2_KubeadmConfigTemplate is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmConfigTemplate_To_v1alpha2_KubeadmConfigTemplate(in *v1alpha3.KubeadmConfigTemplate, out *KubeadmConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmConfigTemplate_To_v1alpha2_KubeadmConfigTemplate(in, out, s)
}

func autoConvert_v1alpha2_KubeadmConfigTemplateList_To_v1alpha3_KubeadmConfigTemplateList(in *KubeadmConfigTemplateList, out *v1alpha3.KubeadmConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.KubeadmConfigTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_KubeadmConfigTemplateList_To_v1alpha3_KubeadmConfigTemplateList is an autogenerated conversion function.
func Convert_v1alpha2_KubeadmConfigTemplateList_To_v1alpha3_KubeadmConfigTemplateList(in *KubeadmConfigTemplateList, out *v1alpha3.KubeadmConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeadmConfigTemplateList_To_v1alpha3_KubeadmConfigTemplateList(in, out, s)
}

func autoConvert_v1alpha3_KubeadmConfigTemplateList_To_v1alpha2_KubeadmConfigTemplateList(in *v1alpha3.KubeadmConfigTemplateList, out *KubeadmConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmConfigTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_KubeadmConfigTemplate_To_v1alpha2_KubeadmConfigTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_KubeadmConfigTemplateList_To_v1alpha2_KubeadmConfigTemplateList is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmConfigTemplateList_To_v1alpha2_KubeadmConfigTemplateList(in *v1alpha3.KubeadmConfigTemplateList, out *KubeadmConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmConfigTemplateList_To_v1alpha2_KubeadmConfigTemplateList(in, out, s)
}

func autoConvert_v1alpha2_KubeadmConfigTemplateResource_To_v1alpha3_KubeadmConfigTemplateResource(in *KubeadmConfigTemplateResource, out *v1alpha3.KubeadmConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha2_KubeadmConfigSpec_To_v1alpha3_KubeadmConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_KubeadmConfigTemplateResource_To_v1alpha3_KubeadmConfigTemplateResource is an autogenerated conversion function.
func Convert_v1alpha2_KubeadmConfigTemplateResource_To_v1alpha3_KubeadmConfigTemplateResource(in *KubeadmConfigTemplateResource, out *v1alpha3.KubeadmConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeadmConfigTemplateResource_To_v1alpha3_KubeadmConfigTemplateResource(in, out, s)
}

func autoConvert_v1alpha3_KubeadmConfigTemplateResource_To_v1alpha2_KubeadmConfigTemplateResource(in *v1alpha3.KubeadmConfigTemplateResource, out *KubeadmConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha3_KubeadmConfigSpec_To_v1alpha2_KubeadmConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_KubeadmConfigTemplateResource_To_v1alpha2_KubeadmConfigTemplateResource is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmConfigTemplateResource_To_v1alpha2_KubeadmConfigTemplateResource(in *v1alpha3.KubeadmConfigTemplateResource, out *KubeadmConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmConfigTemplateResource_To_v1alpha2_KubeadmConfigTemplateResource(in, out, s)
}

func autoConvert_v1alpha2_KubeadmConfigTemplateSpec_To_v1alpha3_KubeadmConfigTemplateSpec(in *KubeadmConfigTemplateSpec, out *v1alpha3.KubeadmConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha2_KubeadmConfigTemplateResource_To_v1alpha3_KubeadmConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_KubeadmConfigTemplateSpec_To_v1alpha3_KubeadmConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha2_KubeadmConfigTemplateSpec_To_v1alpha3_KubeadmConfigTemplateSpec(in *KubeadmConfigTemplateSpec, out *v1alpha3.KubeadmConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeadmConfigTemplateSpec_To_v1alpha3_KubeadmConfigTemplateSpec(in, out, s)
}

func autoConvert_v1alpha3_KubeadmConfigTemplateSpec_To_v1alpha2_KubeadmConfigTemplateSpec(in *v1alpha3.KubeadmConfigTemplateSpec, out *KubeadmConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha3_KubeadmConfigTemplateResource_To_v1alpha2_KubeadmConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_KubeadmConfigTemplateSpec_To_v1alpha2_KubeadmConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha3_KubeadmConfigTemplateSpec_To_v1alpha2_KubeadmConfigTemplateSpec(in *v1alpha3.KubeadmConfigTemplateSpec, out *KubeadmConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubeadmConfigTemplateSpec_To_v1alpha2_KubeadmConfigTemplateSpec(in, out, s)
}

func autoConvert_v1alpha2_NTP_To_v1alpha3_NTP(in *NTP, out *v1alpha3.NTP, s conversion.Scope) error {
	out.Servers = *(*[]string)(unsafe.Pointer(&in.Servers))
	out.Enabled = (*bool)(unsafe.Pointer(in.Enabled))
	return nil
}

// Convert_v1alpha2_NTP_To_v1alpha3_NTP is an autogenerated conversion function.
func Convert_v1alpha2_NTP_To_v1alpha3_NTP(in *NTP, out *v1alpha3.NTP, s conversion.Scope) error {
	return autoConvert_v1alpha2_NTP_To_v1alpha3_NTP(in, out, s)
}

func autoConvert_v1alpha3_NTP_To_v1alpha2_NTP(in *v1alpha3.NTP, out *NTP, s conversion.Scope) error {
	out.Servers = *(*[]string)(unsafe.Pointer(&in.Servers))
	out.Enabled = (*bool)(unsafe.Pointer(in.Enabled))
	return nil
}

// Convert_v1alpha3_NTP_To_v1alpha2_NTP is an autogenerated conversion function.
func Convert_v1alpha3_NTP_To_v1alpha2_NTP(in *v1alpha3.NTP, out *NTP, s conversion.Scope) error {
	return autoConvert_v1alpha3_NTP_To_v1alpha2_NTP(in, out, s)
}

func autoConvert_v1alpha2_User_To_v1alpha3_User(in *User, out *v1alpha3.User, s conversion.Scope) error {
	out.Name = in.Name
	out.Gecos = (*string)(unsafe.Pointer(in.Gecos))
	out.Groups = (*string)(unsafe.Pointer(in.Groups))
	out.HomeDir = (*string)(unsafe.Pointer(in.HomeDir))
	out.Inactive = (*bool)(unsafe.Pointer(in.Inactive))
	out.Shell = (*string)(unsafe.Pointer(in.Shell))
	out.Passwd = (*string)(unsafe.Pointer(in.Passwd))
	out.PrimaryGroup = (*string)(unsafe.Pointer(in.PrimaryGroup))
	out.LockPassword = (*bool)(unsafe.Pointer(in.LockPassword))
	out.Sudo = (*string)(unsafe.Pointer(in.Sudo))
	out.SSHAuthorizedKeys = *(*[]string)(unsafe.Pointer(&in.SSHAuthorizedKeys))
	return nil
}

// Convert_v1alpha2_User_To_v1alpha3_User is an autogenerated conversion function.
func Convert_v1alpha2_User_To_v1alpha3_User(in *User, out *v1alpha3.User, s conversion.Scope) error {
	return autoConvert_v1alpha2_User_To_v1alpha3_User(in, out, s)
}

func autoConvert_v1alpha3_User_To_v1alpha2_User(in *v1alpha3.User, out *User, s conversion.Scope) error {
	out.Name = in.Name
	out.Gecos = (*string)(unsafe.Pointer(in.Gecos))
	out.Groups = (*string)(unsafe.Pointer(in.Groups))
	out.HomeDir = (*string)(unsafe.Pointer(in.HomeDir))
	out.Inactive = (*bool)(unsafe.Pointer(in.Inactive))
	out.Shell = (*string)(unsafe.Pointer(in.Shell))
	out.Passwd = (*string)(unsafe.Pointer(in.Passwd))
	out.PrimaryGroup = (*string)(unsafe.Pointer(in.PrimaryGroup))
	out.LockPassword = (*bool)(unsafe.Pointer(in.LockPassword))
	out.Sudo = (*string)(unsafe.Pointer(in.Sudo))
	out.SSHAuthorizedKeys = *(*[]string)(unsafe.Pointer(&in.SSHAuthorizedKeys))
	return nil
}

// Convert_v1alpha3_User_To_v1alpha2_User is an autogenerated conversion function.
func Convert_v1alpha3_User_To_v1alpha2_User(in *v1alpha3.User, out *User, s conversion.Scope) error {
	return autoConvert_v1alpha3_User_To_v1alpha2_User(in, out, s)
}
