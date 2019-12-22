// Copyright 2019 IBM Corp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainer":                  schema_pkg_apis_securecontainers_v1alpha1_SecureContainer(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImage":             schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImage(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfig":       schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfig(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigRef":    schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfigRef(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigSpec":   schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfigSpec(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigStatus": schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfigStatus(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageRef":          schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageRef(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageSpec":         schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageSpec(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageStatus":       schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageStatus(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerSpec":              schema_pkg_apis_securecontainers_v1alpha1_SecureContainerSpec(ref),
		"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerStatus":            schema_pkg_apis_securecontainers_v1alpha1_SecureContainerStatus(ref),
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainer is the Schema for the securecontainers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerStatus"),
						},
					},
					"object": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
				},
				Required: []string{"object"},
			},
		},
		Dependencies: []string{
			"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerSpec", "github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImage(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImage is the Schema for the securecontainerimages API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageSpec", "github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageConfig is the Schema for the securecontainerimageconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigSpec", "github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfigRef(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageConfigRef defines SecureContainerImage configuration",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageConfigSpec defines the desired state of SecureContainerImageConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"imageDir": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"runtimeClassName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"imageDir", "runtimeClassName"},
			},
		},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageConfigStatus defines the observed state of SecureContainerImageConfig",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageRef(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageRef defines the SecureContainerImage of SecureContainer",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageSpec defines the desired state of SecureContainerImage",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"vmImage": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecrets": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
					"SecureContainerImageConfigRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigRef"),
						},
					},
					"SecureContainerImageConfigSpec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigSpec"),
						},
					},
				},
				Required: []string{"vmImage", "SecureContainerImageConfigRef"},
			},
		},
		Dependencies: []string{
			"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigRef", "github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageConfigSpec", "k8s.io/api/core/v1.LocalObjectReference"},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerImageStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerImageStatus defines the observed state of SecureContainerImage",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerSpec defines the desired state of SecureContainer",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"SecureContainerImageRef": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageRef"),
						},
					},
				},
				Required: []string{"SecureContainerImageRef"},
			},
		},
		Dependencies: []string{
			"github.com/ibm/raksh/pkg/apis/securecontainers/v1alpha1.SecureContainerImageRef"},
	}
}

func schema_pkg_apis_securecontainers_v1alpha1_SecureContainerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecureContainerStatus defines the observed state of SecureContainer",
				Type:        []string{"object"},
			},
		},
	}
}