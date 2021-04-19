package controller

import (
	"fmt"
	"linux_version/utils"
)

func Getinfo()  {
//	var allResp utils.AllResp
	var all  utils.ALL

cpu:=utils.GetCpuInfo()

host:=utils.GetHostInfo()
mem:=utils.GetMemInfo()
disk:=utils.GetDiskInfo()
network:=utils.GetNetworkInfo()
all.PulicIp=utils.PublicIp()
all.NetInfo=network
all.CpuInfo=*cpu
all.DiskInfo=disk
all.HostInfo=*host
all.MemortInfo=*mem
fmt.Println(utils.AllJson(all))

//allResp.MemortInfo=utils.AllJson(utils.GetMemInfo())
//allResp.DiskInfo=utils.AllJson(utils.GetDiskInfo())
//allResp.HostInfo=utils.AllJson(utils.GetHostInfo())
//allResp.CpuInfo=utils.AllJson(utils.GetCpuInfo())
//allResp.NetInfo=utils.AllJson(utils.GetNetworkInfo())
//allResp.PulicIp=utils.PublicIp()
//fmt.Println(utils.AllJson(allResp))
}