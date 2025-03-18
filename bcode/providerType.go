package bcode

type providerType struct {
	Openstack  int
	Kubernetes int
	NCP        int
	Netapp     int
	ALL        int
}

var ProviderType = providerType{
	Openstack:  1,
	Kubernetes: 2,
	NCP:        3,
	Netapp:     4,
	ALL:        100,
}
