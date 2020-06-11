package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WeeklySpec defines the desired state of Weekly
type WeeklySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Name string `json:"name"`
	Articles []ArticlesSpec `json:"articles"`
}

type ArticlesSpec struct {
	Title string `json:"title"`
	URL string `json:"url"`
}


// WeeklyStatus defines the observed state of Weekly
type WeeklyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Weekly is the Schema for the weeklies API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=weeklies,scope=Namespaced
type Weekly struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WeeklySpec   `json:"spec,omitempty"`
	Status WeeklyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WeeklyList contains a list of Weekly
type WeeklyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Weekly `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Weekly{}, &WeeklyList{})
}
