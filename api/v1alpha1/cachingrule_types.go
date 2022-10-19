package v1alpha1

import (
	gingersnapapi "github.com/gingersnap-project/operator/gen/gingersnap-api/config/cache/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CachingRuleStatus defines the observed state of CachingRule
type CachingRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CachingRule is the Schema for the cachingrules API
type CachingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   gingersnapapi.CachingRuleSpec `json:"spec,omitempty"`
	Status CachingRuleStatus             `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CachingRuleList contains a list of CachingRule
type CachingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CachingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CachingRule{}, &CachingRuleList{})
}
