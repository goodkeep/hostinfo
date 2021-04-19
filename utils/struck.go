package utils

type ALL struct {

	PulicIp string
	HostInfo Host
	NetInfo []Network
	CpuInfo []Cpu
	MemortInfo Memory
	DiskInfo []Disk

}
type AllResp struct {

	HostInfo string
	CpuInfo string
	MemortInfo string
	DiskInfo string
	NetInfo string
	PulicIp string

}
type Memory struct {
	Total string `json:"total"`
	Available string `json:"available"`
	Used string `json:"used"`
	UsedPercent string `json:"used_percent"`
	Free string `json:"free"`
	Swapcached string `json:"swapcached"`
	Swaptotal string `json:"swaptotal"`
	Swapfree string `json:"swapfree"`
}
type Cpu struct {
	Cpu int32 `json:"cpu"`
	Cores int32 `json:"cores"`
	ModelName string `json:"modelName"`
	Mhz float64 `json:"mhz"`
}
type Host struct {
	Hostname string `json:"hostname"`
	Os string `json:"os"`
	Platform string `json:"platform"`
	PlatformFamily string `json:"platformFamily"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion string `json:"kernelVersion"`
	KernelArch string `json:"kernelArch"`
	Hostid string `json:"hostid"`
}

type  Disk struct {
	Uuid string `json:"uuid"`
	Device string `json:"device"`
	Fstype string `json:"fstype"`
	Total string `json:"total"`
	Free string `json:"free"`
	Used string `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}
type IP string 
type Network struct {
	Name       string
	IP         string
	MACAddress string
}

type intfInfo struct {
	Name       string
	MacAddress string
	Ipv4       []string
}