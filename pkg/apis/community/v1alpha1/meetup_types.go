package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MeetupSpec defines the desired state of Meetup
type MeetupSpec struct {
	Name            string        `json:"name"`
	Date            string        `json:"date"`
	Time            string        `json:"time"`
	Place           string        `json:"place"`
	City            string        `json:"city"`
	Sponsor         string        `json:"sponsor"`
	RegistrationURL string        `json:"registration_url"`
	Image           string        `json:"image"`
	Tags            []string      `json:"tags"`
	Speakers        []SpeakerSpec `json:"speakers"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

type SpeakerSpec struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Company  string `json:"company"`
	Title    string `json:"title"`
}

// MeetupStatus defines the observed state of Meetup
type MeetupStatus struct {
	Send bool `json:"send"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Meetup is the Schema for the meetups API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=meetups,scope=Namespaced
type Meetup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MeetupSpec   `json:"spec,omitempty"`
	Status MeetupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeetupList contains a list of Meetup
type MeetupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Meetup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Meetup{}, &MeetupList{})
}
