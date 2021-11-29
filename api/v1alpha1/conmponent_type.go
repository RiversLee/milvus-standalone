package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)


type Component struct {
	// +kubebuilder:validation:Optional
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	ImagePullPolicy *corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// +kubebuilder:validation:Optional
	ImagePullSecrets []corev1.PullPolicy `json:"imagePullSecrets,omitempty"`

	// +kubebuilder:validation:Optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}