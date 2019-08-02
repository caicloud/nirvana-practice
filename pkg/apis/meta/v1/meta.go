package v1

import "time"

type Metadata struct {
	// UID is a ID-like unique identification suitable for traditional databases.
	UID string `json:"uid,omitempty"`

	// Name is a label-like unique identification suitable for Kubernetes
	Name string `json:"name,omitempty"`

	Alias             string            `json:"alias,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	Description       string            `json:"description,omitempty"`
	CreationTimestamp *time.Time        `json:"creationTimestamp,omitempty"`
}
