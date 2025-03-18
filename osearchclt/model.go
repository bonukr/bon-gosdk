package osearchclt

type indexName struct {
	Rack      string
	Server    string
	Storage   string
	Network   string
	Openstack string
}

var IndexName = indexName{
	Rack:      "oke-metric-ao-rack",
	Server:    "oke-metric-ao-server",
	Storage:   "oke-metric-ao-storage",
	Network:   "oke-metric-ao-network",
	Openstack: "oke-metric-ao-openstack",
}

type RackPowerInfo struct {
	ID     string `json:"id"`
	UUID   string `json:"uuid"`
	Status string `json:"status"`
}

type RackInfo struct {
	CenterID string `json:"center_id"`
	ID       string `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`

	Power []RackPowerInfo `json:"power"`
}

type ServerNicInfo struct {
	NicID      string `json:"nic_id"`
	NicType    string `json:"nic_type"`
	IPAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
}

type ServerDetailInfo struct {
	Temperature        string `json:"temperature"`
	TemperatureStatus  string `json:"temperature_status"`
	FanSpeed           string `json:"fan_speed"`
	FanSpeedStatus     string `json:"fan_speed_status"`
	PowerSensorStatus  string `json:"power_sensor_status"`
	SashStatus         string `json:"sash_status"`
	PowerConsumption   string `json:"power_consumption"`
	VMCnt              string `json:"vm_cnt"`
	HypervisorUUID     string `json:"hypervisor_uuid"`
	HypervisorHostName string `json:"hypervisor_host_name"`

	Nic []ServerNicInfo `json:"nic"`
}

type ServerInfo struct {
	CenterID string `json:"center_id"`
	RackID   string `json:"rack_id"`
	ID       string `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Status   string `json:"status"`

	Detail ServerDetailInfo `json:"detail"`
}

type StorageVolumnInfo struct {
	ID        string `json:"id"`
	AllocSize string `json:"alloc_size"`
	UsedSize  string `json:"used_size"`
}

type StorageDetailInfo struct {
	Temperature   string `json:"temperature"`
	IPAddress     string `json:"ip_address"`
	AvailableSize string `json:"available_size"`
	AllocSize     string `json:"alloc_size"`

	Volumn []StorageVolumnInfo `json:"volumn"`
}

type StorageInfo struct {
	CenterID string `json:"center_id"`
	RackID   string `json:"rack_id"`
	ID       string `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Status   string `json:"status"`

	Detail StorageDetailInfo `json:"detail"`
}

type NetworkPortInfo struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	LinkServerID    string `json:"link_server_id"`
	LinkServerNicID string `json:"link_server_nic_id"`
}

type NetworkDetailInfo struct {
	Temperature   string `json:"temperature"`
	IPAddress     string `json:"ip_address"`
	MacAddress    string `json:"mac_address"`
	OobIPAddress  string `json:"oob_ip_address"`
	OobMacAddress string `json:"oob_mac_address"`
	TotalPortCnt  string `json:"total_port_cnt"`
	ActivePortCnt string `json:"active_port_cnt"`

	Port []NetworkPortInfo `json:"port"`
}

type NetworkInfo struct {
	CenterID string `json:"center_id"`
	RackID   string `json:"rack_id"`
	ID       string `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Status   string `json:"status"`

	Detail NetworkDetailInfo `json:"detail"`
}

type Msg struct {
	CollectTime   string `json:"collect_time"`
	CollectedTime string `json:"collected_time"`
	CollectorName string `json:"collector_name"`
	ResourceType  string `json:"resource_type"`
	ResourceUUID  string `json:"resource_uuid"`

	Rack    RackInfo    `json:"rack"`
	Server  ServerInfo  `json:"server"`
	Storage StorageInfo `json:"storage"`
	Network NetworkInfo `json:"network"`
}
