//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	shootdnsrewriting "github.com/gardener/gardener/plugin/pkg/shoot/dnsrewriting/apis/shootdnsrewriting"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*shootdnsrewriting.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_shootdnsrewriting_Configuration(a.(*Configuration), b.(*shootdnsrewriting.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*shootdnsrewriting.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_shootdnsrewriting_Configuration_To_v1alpha1_Configuration(a.(*shootdnsrewriting.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Configuration_To_shootdnsrewriting_Configuration(in *Configuration, out *shootdnsrewriting.Configuration, s conversion.Scope) error {
	out.CommonSuffixes = *(*[]string)(unsafe.Pointer(&in.CommonSuffixes))
	return nil
}

// Convert_v1alpha1_Configuration_To_shootdnsrewriting_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_shootdnsrewriting_Configuration(in *Configuration, out *shootdnsrewriting.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_shootdnsrewriting_Configuration(in, out, s)
}

func autoConvert_shootdnsrewriting_Configuration_To_v1alpha1_Configuration(in *shootdnsrewriting.Configuration, out *Configuration, s conversion.Scope) error {
	out.CommonSuffixes = *(*[]string)(unsafe.Pointer(&in.CommonSuffixes))
	return nil
}

// Convert_shootdnsrewriting_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_shootdnsrewriting_Configuration_To_v1alpha1_Configuration(in *shootdnsrewriting.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_shootdnsrewriting_Configuration_To_v1alpha1_Configuration(in, out, s)
}
