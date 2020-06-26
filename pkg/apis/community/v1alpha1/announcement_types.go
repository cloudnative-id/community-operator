package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnnouncementSpec defines the desired state of Announcement
type AnnouncementSpec struct {
	Title     string   `json:"title"`
	Message   string   `json:"message"`
	Community string   `json:"community"`
	Image     string   `json:"image"`
	Tags      []string `json:"tags"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// AnnouncementStatus defines the observed state of Announcement
type AnnouncementStatus struct {
	Send bool `json:"send"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Announcement is the Schema for the announcements API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=announcements,scope=Namespaced
type Announcement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnnouncementSpec   `json:"spec,omitempty"`
	Status AnnouncementStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnnouncementList contains a list of Announcement
type AnnouncementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Announcement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Announcement{}, &AnnouncementList{})
}
