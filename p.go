package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {

	/*mask := "255.255.255.0"
	mk := strings.Split(mask, ".")
	a, _ := strconv.Atoi(mk[0])

	//如何让a 转成byte类型？

	ones, bits := net.IPv4Mask(byte(a), 255, 255, 248).Size()

	fmt.Println(ones, bits)*/
	bb()

}

func aa() {
	_, ipNet, err := net.ParseCIDR("10.95.134.192/29")
	if err != nil {
		fmt.Println(err)
	}
	val := make([]byte, len(ipNet.Mask))
	copy(val, ipNet.Mask)

	var s []string
	for _, i := range val[:] {
		s = append(s, strconv.Itoa(int(i)))
	}
	fmt.Printf("%s/%s\n", ipNet.IP, strings.Join(s, "."))

}

type IpServiceStruct struct {
	IpAdress     string `json:"ip_adress"`     //ip地址
	SubnetMask   string `json:"subnet_mask"`   //子网掩码
	Gateway      string `json:"gateway"`       //网关
	PreferredDNS string `json:"preferred_dns"` //首选dns
	StandbyDNS   string `json:"standby_dns"`   //备用dns
}

func bb() {

	uplog := IpServiceStruct{}
	/*err := json.Unmarshal(msg, &uplog)
	if err != nil {
		return err
	}*/
	uplog.IpAdress = "10.0.102.124"
	uplog.SubnetMask = "255.255.255.0"
	mask := uplog.SubnetMask
	mk := strings.Split(mask, ".")
	mk0, _ := strconv.Atoi(mk[0])
	mk1, _ := strconv.Atoi(mk[1])
	mk2, _ := strconv.Atoi(mk[2])
	mk3, _ := strconv.Atoi(mk[3])
	ones, bits := net.IPv4Mask(byte(mk0), byte(mk1), byte(mk2), byte(mk3)).Size()
	fmt.Println(ones, bits)
	smark := strconv.Itoa(ones)
	//组装数据 ip/mask
	uplog.IpAdress = uplog.IpAdress + "/" + smark
	//gateway
	uplog.Gateway = "10.0.102.1"

	fmt.Println(uplog)
}
