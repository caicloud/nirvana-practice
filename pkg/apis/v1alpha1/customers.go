package v1alpha1

import (
	"time"

	meta "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
)

type Customer struct {
	meta.Metadata `json:",inline"`
	Spec          *CustomerSpec   `json:"spec,omitempty"`
	Status        *CustomerStatus `json:"status,omitempty"`
}

type CustomerSpec struct {
	Sex string `json:"sex,omitempty"`
}

type CustomerStatus struct {
	LatestLoginTimestamp *time.Time `json:"latestLoginTimestamp,omitempty"`
}

type CustomersList struct {
	Total int         `json:"total"`
	Items []*Customer `json:"items,omitempty"`
}
