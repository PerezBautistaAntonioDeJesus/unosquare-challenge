package models

import "time"

type Service struct {
	Metadata Metadata
	Spec     Spec
}

type Metadata struct {
	Name              string
	Namespase         string
	Uid               string
	ResourceVersion   string
	CreationTimestamp time.Time
}

type Spec struct {
	ClusterIP             string
	ClusterIPs            []string
	Type                  string
	SessionAffinity       string
	IpFamilies            []string
	InternalTrafficPolicy string
}
