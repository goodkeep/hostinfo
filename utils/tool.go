package utils

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

const  GB int=1024*1024*1024


//获取cpu信息
func GetCpuInfo()*[]Cpu  {
	var cpuinfo []Cpu
	cpuInfo,_:=cpu.Info()
	for _,ci:=range cpuInfo{
		var cpu Cpu
		a,_:=json.Marshal(ci)
		json.Unmarshal(a,&cpu)
		cpuinfo=append(cpuinfo,cpu)
	}
	return &cpuinfo
}
//获取主机信息
func GetHostInfo()*Host  {
var hostinfo Host
	hinfo,_:=host.Info()
a,_:=json.Marshal(hinfo)
json.Unmarshal(a,&hostinfo)

	return &hostinfo
}
//获取内存信息
func GetMemInfo() *Memory  {
	var memory Memory
	memInfo,_:=mem.VirtualMemory()
	memory.Total=InitString(memInfo.Total)+" GB"
	memory.Available=InitString(memInfo.Available)+" GB"
	memory.Used=InitString(memInfo.Used)+" GB"
	memory.UsedPercent=InitString(memInfo.UsedPercent)
	memory.Swapcached=InitString(memInfo.SwapCached)+" GB"
	memory.Swaptotal=InitString(memInfo.SwapTotal)+" GB"
	memory.Swapfree=InitString(memInfo.SwapFree)+" GB"
	return &memory
}
//获取硬盘信息
func GetDiskInfo()[]Disk  {
	var diskinfo []Disk
	parts,_:=disk.Partitions(false)


	for _,v:=range parts{
		var part Disk

		partinfo,_:=disk.Usage(v.Mountpoint)
		part.Uuid=GetUuid(v.Device)
		part.Used=InitString(int64(partinfo.Used))+" GB"
		part.Total=InitString(int64(partinfo.Total))+" GB"
		part.Fstype=v.Fstype
		part.Device=v.Device
		part.Free=InitString(partinfo.Free)+" GB"
		part.UsedPercent=partinfo.UsedPercent
		
		diskinfo=append(diskinfo,part)

	}
	return diskinfo
}
//将传入参数转为strin类型。而且转为GB
func InitString(a interface{})string  {
	var SB string
	switch a.(type) {
	case int:
		SB=strconv.Itoa(a.(int)/GB)
	case int32:
		SB=strconv.Itoa(int(a.(int32))/GB)
	case int64:
		SB=strconv.Itoa(int(a.(int64))/GB)
	case float32:
		SB=strconv.Itoa(int(a.(float32))/GB)
	case float64:
		SB=strconv.Itoa(int(a.(float64))/GB)
	case uint64:
		SB=strconv.Itoa(int(a.(uint64))/GB)
	case uint32:
		SB=strconv.Itoa(int(a.(uint32))/GB)
	}

	return SB
}
//将任意数据先转换为json。然后转为字符串
func AllJson(a interface{} )string  {
	i,_:=json.Marshal(a)
	return string(i)
}
//get网卡信息
func GetNetworkInfo() []Network {
	intf, err := net.Interfaces()
	if err != nil {
		log.Fatal("get network info failed: %v", err)
		return nil
	}
	var is = make([]intfInfo, len(intf))
	var q int
	var networkinfo []Network
	for i, v := range intf {
		ips, err := v.Addrs()
		if err != nil {
			log.Fatal("get network addr failed: %v", err)
			return nil
		}
		//此处过滤loopback（本地回环）和isatap（isatap隧道）
		var network Network
		if !strings.Contains(v.Name, "Loopback") && !strings.Contains(v.Name, "isatap") {

			is[i].Name = v.Name
			is[i].MacAddress = v.HardwareAddr.String()
			q+=1
			for _, ip := range ips {

				if strings.Contains(ip.String(), ".") {
					is[i].Ipv4 = append(is[i].Ipv4, ip.String())
				}
			}
			network.Name = is[i].Name
			network.MACAddress = is[i].MacAddress
			if len(is[i].Ipv4) > 0 {
				network.IP = is[i].Ipv4[0]
			}


			//fmt.Println("network:=", network)
		}
		networkinfo=append(networkinfo,network)
	}

	return networkinfo
}
//获取公网ip。如果有
func PublicIp() string {
	var ip string
	resp,err:=http.Get("http://ip.dhcp.cn/?ip")
	if err !=nil{
		return " "
	}
	defer resp.Body.Close()
	ipinfo,_:=ioutil.ReadAll(resp.Body)
	ip=string(ipinfo)
	return ip
}
//获取磁盘UUid，这里写的比较差劲
func GetUuid(dev string)string  {
	diskin,err:=exec.Command("blkid",dev).Output()
	if err !=nil{
		return " "
	}
	diskinfo:=string(diskin)
	var uuid string
	i :=0
	c:=0
	for k,v:=range diskinfo {
		if string(v)==" "&&string(diskinfo[k+1])=="U"{

			for j,q:=range diskinfo[k:] {
				if string(q)==`"`{
					i+=1
					if i==1{
						c=k+j+1
					}
					if i==2{
						uuid=string(diskinfo[c:k+j])
						break
					}else {
						continue
					}
				}

				continue

			}
		}

	}
	return uuid
}
