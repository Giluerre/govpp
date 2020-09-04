// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01
// source: .vppapi/core/l2tp.api.json

// Package l2tp contains generated bindings for API file l2tp.api.
//
// Contents:
//   7 aliases
//  11 enums
//   6 structs
//   1 union
//  10 messages
//
package l2tp

import (
	"fmt"
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	"net"
	"strconv"
	"strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "l2tp"
	APIVersion = "2.0.0"
	VersionCrc = 0x1ecf7730
)

// AddressFamily defines enum 'address_family'.
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var (
	AddressFamily_name = map[uint32]string{
		0: "ADDRESS_IP4",
		1: "ADDRESS_IP6",
	}
	AddressFamily_value = map[string]uint32{
		"ADDRESS_IP4": 0,
		"ADDRESS_IP6": 1,
	}
)

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint32(x)]
	if ok {
		return s
	}
	return "AddressFamily(" + strconv.Itoa(int(x)) + ")"
}

// IfStatusFlags defines enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var (
	IfStatusFlags_name = map[uint32]string{
		1: "IF_STATUS_API_FLAG_ADMIN_UP",
		2: "IF_STATUS_API_FLAG_LINK_UP",
	}
	IfStatusFlags_value = map[string]uint32{
		"IF_STATUS_API_FLAG_ADMIN_UP": 1,
		"IF_STATUS_API_FLAG_LINK_UP":  2,
	}
)

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IfStatusFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IfStatusFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IfType defines enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var (
	IfType_name = map[uint32]string{
		1: "IF_API_TYPE_HARDWARE",
		2: "IF_API_TYPE_SUB",
		3: "IF_API_TYPE_P2P",
		4: "IF_API_TYPE_PIPE",
	}
	IfType_value = map[string]uint32{
		"IF_API_TYPE_HARDWARE": 1,
		"IF_API_TYPE_SUB":      2,
		"IF_API_TYPE_P2P":      3,
		"IF_API_TYPE_PIPE":     4,
	}
)

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return "IfType(" + strconv.Itoa(int(x)) + ")"
}

// IPDscp defines enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var (
	IPDscp_name = map[uint8]string{
		0:  "IP_API_DSCP_CS0",
		8:  "IP_API_DSCP_CS1",
		10: "IP_API_DSCP_AF11",
		12: "IP_API_DSCP_AF12",
		14: "IP_API_DSCP_AF13",
		16: "IP_API_DSCP_CS2",
		18: "IP_API_DSCP_AF21",
		20: "IP_API_DSCP_AF22",
		22: "IP_API_DSCP_AF23",
		24: "IP_API_DSCP_CS3",
		26: "IP_API_DSCP_AF31",
		28: "IP_API_DSCP_AF32",
		30: "IP_API_DSCP_AF33",
		32: "IP_API_DSCP_CS4",
		34: "IP_API_DSCP_AF41",
		36: "IP_API_DSCP_AF42",
		38: "IP_API_DSCP_AF43",
		40: "IP_API_DSCP_CS5",
		46: "IP_API_DSCP_EF",
		48: "IP_API_DSCP_CS6",
		50: "IP_API_DSCP_CS7",
	}
	IPDscp_value = map[string]uint8{
		"IP_API_DSCP_CS0":  0,
		"IP_API_DSCP_CS1":  8,
		"IP_API_DSCP_AF11": 10,
		"IP_API_DSCP_AF12": 12,
		"IP_API_DSCP_AF13": 14,
		"IP_API_DSCP_CS2":  16,
		"IP_API_DSCP_AF21": 18,
		"IP_API_DSCP_AF22": 20,
		"IP_API_DSCP_AF23": 22,
		"IP_API_DSCP_CS3":  24,
		"IP_API_DSCP_AF31": 26,
		"IP_API_DSCP_AF32": 28,
		"IP_API_DSCP_AF33": 30,
		"IP_API_DSCP_CS4":  32,
		"IP_API_DSCP_AF41": 34,
		"IP_API_DSCP_AF42": 36,
		"IP_API_DSCP_AF43": 38,
		"IP_API_DSCP_CS5":  40,
		"IP_API_DSCP_EF":   46,
		"IP_API_DSCP_CS6":  48,
		"IP_API_DSCP_CS7":  50,
	}
)

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return "IPDscp(" + strconv.Itoa(int(x)) + ")"
}

// IPEcn defines enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var (
	IPEcn_name = map[uint8]string{
		0: "IP_API_ECN_NONE",
		1: "IP_API_ECN_ECT0",
		2: "IP_API_ECN_ECT1",
		3: "IP_API_ECN_CE",
	}
	IPEcn_value = map[string]uint8{
		"IP_API_ECN_NONE": 0,
		"IP_API_ECN_ECT0": 1,
		"IP_API_ECN_ECT1": 2,
		"IP_API_ECN_CE":   3,
	}
)

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return "IPEcn(" + strconv.Itoa(int(x)) + ")"
}

// IPProto defines enum 'ip_proto'.
type IPProto uint32

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_AH       IPProto = 50
	IP_API_PROTO_ESP      IPProto = 51
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var (
	IPProto_name = map[uint32]string{
		0:   "IP_API_PROTO_HOPOPT",
		1:   "IP_API_PROTO_ICMP",
		2:   "IP_API_PROTO_IGMP",
		6:   "IP_API_PROTO_TCP",
		17:  "IP_API_PROTO_UDP",
		47:  "IP_API_PROTO_GRE",
		50:  "IP_API_PROTO_AH",
		51:  "IP_API_PROTO_ESP",
		88:  "IP_API_PROTO_EIGRP",
		89:  "IP_API_PROTO_OSPF",
		132: "IP_API_PROTO_SCTP",
		255: "IP_API_PROTO_RESERVED",
	}
	IPProto_value = map[string]uint32{
		"IP_API_PROTO_HOPOPT":   0,
		"IP_API_PROTO_ICMP":     1,
		"IP_API_PROTO_IGMP":     2,
		"IP_API_PROTO_TCP":      6,
		"IP_API_PROTO_UDP":      17,
		"IP_API_PROTO_GRE":      47,
		"IP_API_PROTO_AH":       50,
		"IP_API_PROTO_ESP":      51,
		"IP_API_PROTO_EIGRP":    88,
		"IP_API_PROTO_OSPF":     89,
		"IP_API_PROTO_SCTP":     132,
		"IP_API_PROTO_RESERVED": 255,
	}
)

func (x IPProto) String() string {
	s, ok := IPProto_name[uint32(x)]
	if ok {
		return s
	}
	return "IPProto(" + strconv.Itoa(int(x)) + ")"
}

// L2tLookupKey defines enum 'l2t_lookup_key'.
type L2tLookupKey uint8

const (
	L2T_LOOKUP_KEY_API_SRC_ADDR   L2tLookupKey = 0
	L2T_LOOKUP_KEY_API_DST_ADDR   L2tLookupKey = 1
	L2T_LOOKUP_KEY_API_SESSION_ID L2tLookupKey = 2
)

var (
	L2tLookupKey_name = map[uint8]string{
		0: "L2T_LOOKUP_KEY_API_SRC_ADDR",
		1: "L2T_LOOKUP_KEY_API_DST_ADDR",
		2: "L2T_LOOKUP_KEY_API_SESSION_ID",
	}
	L2tLookupKey_value = map[string]uint8{
		"L2T_LOOKUP_KEY_API_SRC_ADDR":   0,
		"L2T_LOOKUP_KEY_API_DST_ADDR":   1,
		"L2T_LOOKUP_KEY_API_SESSION_ID": 2,
	}
)

func (x L2tLookupKey) String() string {
	s, ok := L2tLookupKey_name[uint8(x)]
	if ok {
		return s
	}
	return "L2tLookupKey(" + strconv.Itoa(int(x)) + ")"
}

// LinkDuplex defines enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var (
	LinkDuplex_name = map[uint32]string{
		0: "LINK_DUPLEX_API_UNKNOWN",
		1: "LINK_DUPLEX_API_HALF",
		2: "LINK_DUPLEX_API_FULL",
	}
	LinkDuplex_value = map[string]uint32{
		"LINK_DUPLEX_API_UNKNOWN": 0,
		"LINK_DUPLEX_API_HALF":    1,
		"LINK_DUPLEX_API_FULL":    2,
	}
)

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return "LinkDuplex(" + strconv.Itoa(int(x)) + ")"
}

// MtuProto defines enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var (
	MtuProto_name = map[uint32]string{
		1: "MTU_PROTO_API_L3",
		2: "MTU_PROTO_API_IP4",
		3: "MTU_PROTO_API_IP6",
		4: "MTU_PROTO_API_MPLS",
		5: "MTU_PROTO_API_N",
	}
	MtuProto_value = map[string]uint32{
		"MTU_PROTO_API_L3":   1,
		"MTU_PROTO_API_IP4":  2,
		"MTU_PROTO_API_IP6":  3,
		"MTU_PROTO_API_MPLS": 4,
		"MTU_PROTO_API_N":    5,
	}
)

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return "MtuProto(" + strconv.Itoa(int(x)) + ")"
}

// RxMode defines enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var (
	RxMode_name = map[uint32]string{
		0: "RX_MODE_API_UNKNOWN",
		1: "RX_MODE_API_POLLING",
		2: "RX_MODE_API_INTERRUPT",
		3: "RX_MODE_API_ADAPTIVE",
		4: "RX_MODE_API_DEFAULT",
	}
	RxMode_value = map[string]uint32{
		"RX_MODE_API_UNKNOWN":   0,
		"RX_MODE_API_POLLING":   1,
		"RX_MODE_API_INTERRUPT": 2,
		"RX_MODE_API_ADAPTIVE":  3,
		"RX_MODE_API_DEFAULT":   4,
	}
)

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RxMode(" + strconv.Itoa(int(x)) + ")"
}

// SubIfFlags defines enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var (
	SubIfFlags_name = map[uint32]string{
		1:   "SUB_IF_API_FLAG_NO_TAGS",
		2:   "SUB_IF_API_FLAG_ONE_TAG",
		4:   "SUB_IF_API_FLAG_TWO_TAGS",
		8:   "SUB_IF_API_FLAG_DOT1AD",
		16:  "SUB_IF_API_FLAG_EXACT_MATCH",
		32:  "SUB_IF_API_FLAG_DEFAULT",
		64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
		128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
		254: "SUB_IF_API_FLAG_MASK_VNET",
		256: "SUB_IF_API_FLAG_DOT1AH",
	}
	SubIfFlags_value = map[string]uint32{
		"SUB_IF_API_FLAG_NO_TAGS":           1,
		"SUB_IF_API_FLAG_ONE_TAG":           2,
		"SUB_IF_API_FLAG_TWO_TAGS":          4,
		"SUB_IF_API_FLAG_DOT1AD":            8,
		"SUB_IF_API_FLAG_EXACT_MATCH":       16,
		"SUB_IF_API_FLAG_DEFAULT":           32,
		"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
		"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
		"SUB_IF_API_FLAG_MASK_VNET":         254,
		"SUB_IF_API_FLAG_DOT1AH":            256,
	}
)

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := SubIfFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "SubIfFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// AddressWithPrefix defines alias 'address_with_prefix'.
type AddressWithPrefix Prefix

func ParseAddressWithPrefix(s string) (AddressWithPrefix, error) {
	prefix, err := ParsePrefix(s)
	if err != nil {
		return AddressWithPrefix{}, err
	}
	return AddressWithPrefix(prefix), nil
}
func (x AddressWithPrefix) String() string {
	return Prefix(x).String()
}
func (x *AddressWithPrefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *AddressWithPrefix) UnmarshalText(text []byte) error {
	prefix, err := ParseAddressWithPrefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

// IP4Address defines alias 'ip4_address'.
type IP4Address [4]uint8

func ParseIP4Address(s string) (IP4Address, error) {
	ip := net.ParseIP(s).To4()
	if ip == nil {
		return IP4Address{}, fmt.Errorf("invalid IP address: %s", s)
	}
	var ipaddr IP4Address
	copy(ipaddr[:], ip.To4())
	return ipaddr, nil
}

func (x IP4Address) ToIP() net.IP {
	return net.IP(x[:]).To4()
}
func (x IP4Address) String() string {
	return x.ToIP().String()
}
func (x *IP4Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *IP4Address) UnmarshalText(text []byte) error {
	ipaddr, err := ParseIP4Address(string(text))
	if err != nil {
		return err
	}
	*x = ipaddr
	return nil
}

// IP4AddressWithPrefix defines alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address defines alias 'ip6_address'.
type IP6Address [16]uint8

func ParseIP6Address(s string) (IP6Address, error) {
	ip := net.ParseIP(s).To16()
	if ip == nil {
		return IP6Address{}, fmt.Errorf("invalid IP address: %s", s)
	}
	var ipaddr IP6Address
	copy(ipaddr[:], ip.To16())
	return ipaddr, nil
}

func (x IP6Address) ToIP() net.IP {
	return net.IP(x[:]).To16()
}
func (x IP6Address) String() string {
	return x.ToIP().String()
}
func (x *IP6Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *IP6Address) UnmarshalText(text []byte) error {
	ipaddr, err := ParseIP6Address(string(text))
	if err != nil {
		return err
	}
	*x = ipaddr
	return nil
}

// IP6AddressWithPrefix defines alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// MacAddress defines alias 'mac_address'.
type MacAddress [6]uint8

func ParseMacAddress(s string) (MacAddress, error) {
	var macaddr MacAddress
	mac, err := net.ParseMAC(s)
	if err != nil {
		return macaddr, err
	}
	copy(macaddr[:], mac[:])
	return macaddr, nil
}
func (x MacAddress) ToMAC() net.HardwareAddr {
	return net.HardwareAddr(x[:])
}
func (x MacAddress) String() string {
	return x.ToMAC().String()
}
func (x *MacAddress) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *MacAddress) UnmarshalText(text []byte) error {
	mac, err := ParseMacAddress(string(text))
	if err != nil {
		return err
	}
	*x = mac
	return nil
}

// Address defines type 'address'.
type Address struct {
	Af AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	Un AddressUnion  `binapi:"address_union,name=un" json:"un,omitempty"`
}

func ParseAddress(s string) (Address, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return Address{}, fmt.Errorf("invalid address: %s", s)
	}
	var addr Address
	if ip.To4() == nil {
		addr.Af = ADDRESS_IP6
		var ip6 IP6Address
		copy(ip6[:], ip.To16())
		addr.Un.SetIP6(ip6)
	} else {
		addr.Af = ADDRESS_IP4
		var ip4 IP4Address
		copy(ip4[:], ip.To4())
		addr.Un.SetIP4(ip4)
	}
	return addr, nil
}
func (x Address) ToIP() net.IP {
	if x.Af == ADDRESS_IP6 {
		ip6 := x.Un.GetIP6()
		return net.IP(ip6[:]).To16()
	} else {
		ip4 := x.Un.GetIP4()
		return net.IP(ip4[:]).To4()
	}
}
func (x Address) String() string {
	return x.ToIP().String()
}
func (x *Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *Address) UnmarshalText(text []byte) error {
	addr, err := ParseAddress(string(text))
	if err != nil {
		return err
	}
	*x = addr
	return nil
}

// IP4Prefix defines type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address `binapi:"ip4_address,name=address" json:"address,omitempty"`
	Len     uint8      `binapi:"u8,name=len" json:"len,omitempty"`
}

func ParseIP4Prefix(s string) (prefix IP4Prefix, err error) {
	hasPrefix := strings.Contains(s, "/")
	if hasPrefix {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseIP4Address(ip.String())
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	} else {
		ip := net.ParseIP(s)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if ip.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseIP4Address(ip.String())
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	}
	return prefix, nil
}
func (x IP4Prefix) ToIPNet() *net.IPNet {
	mask := net.CIDRMask(int(x.Len), 32)
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}
func (x IP4Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}
func (x *IP4Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *IP4Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParseIP4Prefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// IP6Prefix defines type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address `binapi:"ip6_address,name=address" json:"address,omitempty"`
	Len     uint8      `binapi:"u8,name=len" json:"len,omitempty"`
}

func ParseIP6Prefix(s string) (prefix IP6Prefix, err error) {
	hasPrefix := strings.Contains(s, "/")
	if hasPrefix {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseIP6Address(ip.String())
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	} else {
		ip := net.ParseIP(s)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if ip.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseIP6Address(ip.String())
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	}
	return prefix, nil
}
func (x IP6Prefix) ToIPNet() *net.IPNet {
	mask := net.CIDRMask(int(x.Len), 128)
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}
func (x IP6Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}
func (x *IP6Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *IP6Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParseIP6Prefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// Mprefix defines type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	GrpAddressLength uint16        `binapi:"u16,name=grp_address_length" json:"grp_address_length,omitempty"`
	GrpAddress       AddressUnion  `binapi:"address_union,name=grp_address" json:"grp_address,omitempty"`
	SrcAddress       AddressUnion  `binapi:"address_union,name=src_address" json:"src_address,omitempty"`
}

// Prefix defines type 'prefix'.
type Prefix struct {
	Address Address `binapi:"address,name=address" json:"address,omitempty"`
	Len     uint8   `binapi:"u8,name=len" json:"len,omitempty"`
}

func ParsePrefix(ip string) (prefix Prefix, err error) {
	hasPrefix := strings.Contains(ip, "/")
	if hasPrefix {
		netIP, network, err := net.ParseCIDR(ip)
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseAddress(netIP.String())
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
	} else {
		netIP := net.ParseIP(ip)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if netIP.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseAddress(netIP.String())
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
	}
	return prefix, nil
}
func (x Prefix) ToIPNet() *net.IPNet {
	var mask net.IPMask
	if x.Address.Af == ADDRESS_IP4 {
		mask = net.CIDRMask(int(x.Len), 32)
	} else {
		mask = net.CIDRMask(int(x.Len), 128)
	}
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}
func (x Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}
func (x *Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParsePrefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// PrefixMatcher defines type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8 `binapi:"u8,name=le" json:"le,omitempty"`
	Ge uint8 `binapi:"u8,name=ge" json:"ge,omitempty"`
}

// AddressUnion defines union 'address_union'.
type AddressUnion struct {
	// IP4 *IP4Address
	// IP6 *IP6Address
	XXX_UnionData [16]byte
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 4)
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(4))
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 16)
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(16))
	return
}

// L2tpv3CreateTunnel defines message 'l2tpv3_create_tunnel'.
type L2tpv3CreateTunnel struct {
	ClientAddress     Address `binapi:"address,name=client_address" json:"client_address,omitempty"`
	OurAddress        Address `binapi:"address,name=our_address" json:"our_address,omitempty"`
	LocalSessionID    uint32  `binapi:"u32,name=local_session_id" json:"local_session_id,omitempty"`
	RemoteSessionID   uint32  `binapi:"u32,name=remote_session_id" json:"remote_session_id,omitempty"`
	LocalCookie       uint64  `binapi:"u64,name=local_cookie" json:"local_cookie,omitempty"`
	RemoteCookie      uint64  `binapi:"u64,name=remote_cookie" json:"remote_cookie,omitempty"`
	L2SublayerPresent bool    `binapi:"bool,name=l2_sublayer_present" json:"l2_sublayer_present,omitempty"`
	EncapVrfID        uint32  `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
}

func (m *L2tpv3CreateTunnel) Reset()               { *m = L2tpv3CreateTunnel{} }
func (*L2tpv3CreateTunnel) GetMessageName() string { return "l2tpv3_create_tunnel" }
func (*L2tpv3CreateTunnel) GetCrcString() string   { return "596892cb" }
func (*L2tpv3CreateTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3CreateTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.ClientAddress.Af
	size += 1 * 16 // m.ClientAddress.Un
	size += 4      // m.OurAddress.Af
	size += 1 * 16 // m.OurAddress.Un
	size += 4      // m.LocalSessionID
	size += 4      // m.RemoteSessionID
	size += 8      // m.LocalCookie
	size += 8      // m.RemoteCookie
	size += 1      // m.L2SublayerPresent
	size += 4      // m.EncapVrfID
	return size
}
func (m *L2tpv3CreateTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.ClientAddress.Af))
	buf.EncodeBytes(m.ClientAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.OurAddress.Af))
	buf.EncodeBytes(m.OurAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.LocalSessionID)
	buf.EncodeUint32(m.RemoteSessionID)
	buf.EncodeUint64(m.LocalCookie)
	buf.EncodeUint64(m.RemoteCookie)
	buf.EncodeBool(m.L2SublayerPresent)
	buf.EncodeUint32(m.EncapVrfID)
	return buf.Bytes(), nil
}
func (m *L2tpv3CreateTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ClientAddress.Af = AddressFamily(buf.DecodeUint32())
	copy(m.ClientAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.OurAddress.Af = AddressFamily(buf.DecodeUint32())
	copy(m.OurAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LocalSessionID = buf.DecodeUint32()
	m.RemoteSessionID = buf.DecodeUint32()
	m.LocalCookie = buf.DecodeUint64()
	m.RemoteCookie = buf.DecodeUint64()
	m.L2SublayerPresent = buf.DecodeBool()
	m.EncapVrfID = buf.DecodeUint32()
	return nil
}

// L2tpv3CreateTunnelReply defines message 'l2tpv3_create_tunnel_reply'.
type L2tpv3CreateTunnelReply struct {
	Retval    int32          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *L2tpv3CreateTunnelReply) Reset()               { *m = L2tpv3CreateTunnelReply{} }
func (*L2tpv3CreateTunnelReply) GetMessageName() string { return "l2tpv3_create_tunnel_reply" }
func (*L2tpv3CreateTunnelReply) GetCrcString() string   { return "5383d31f" }
func (*L2tpv3CreateTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3CreateTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *L2tpv3CreateTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *L2tpv3CreateTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// L2tpv3InterfaceEnableDisable defines message 'l2tpv3_interface_enable_disable'.
type L2tpv3InterfaceEnableDisable struct {
	EnableDisable bool           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex     InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *L2tpv3InterfaceEnableDisable) Reset() { *m = L2tpv3InterfaceEnableDisable{} }
func (*L2tpv3InterfaceEnableDisable) GetMessageName() string {
	return "l2tpv3_interface_enable_disable"
}
func (*L2tpv3InterfaceEnableDisable) GetCrcString() string { return "3865946c" }
func (*L2tpv3InterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3InterfaceEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex
	return size
}
func (m *L2tpv3InterfaceEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *L2tpv3InterfaceEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// L2tpv3InterfaceEnableDisableReply defines message 'l2tpv3_interface_enable_disable_reply'.
type L2tpv3InterfaceEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2tpv3InterfaceEnableDisableReply) Reset() { *m = L2tpv3InterfaceEnableDisableReply{} }
func (*L2tpv3InterfaceEnableDisableReply) GetMessageName() string {
	return "l2tpv3_interface_enable_disable_reply"
}
func (*L2tpv3InterfaceEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*L2tpv3InterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3InterfaceEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *L2tpv3InterfaceEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *L2tpv3InterfaceEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// L2tpv3SetLookupKey defines message 'l2tpv3_set_lookup_key'.
type L2tpv3SetLookupKey struct {
	Key L2tLookupKey `binapi:"l2t_lookup_key,name=key" json:"key,omitempty"`
}

func (m *L2tpv3SetLookupKey) Reset()               { *m = L2tpv3SetLookupKey{} }
func (*L2tpv3SetLookupKey) GetMessageName() string { return "l2tpv3_set_lookup_key" }
func (*L2tpv3SetLookupKey) GetCrcString() string   { return "c9892c86" }
func (*L2tpv3SetLookupKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3SetLookupKey) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Key
	return size
}
func (m *L2tpv3SetLookupKey) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Key))
	return buf.Bytes(), nil
}
func (m *L2tpv3SetLookupKey) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Key = L2tLookupKey(buf.DecodeUint8())
	return nil
}

// L2tpv3SetLookupKeyReply defines message 'l2tpv3_set_lookup_key_reply'.
type L2tpv3SetLookupKeyReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2tpv3SetLookupKeyReply) Reset()               { *m = L2tpv3SetLookupKeyReply{} }
func (*L2tpv3SetLookupKeyReply) GetMessageName() string { return "l2tpv3_set_lookup_key_reply" }
func (*L2tpv3SetLookupKeyReply) GetCrcString() string   { return "e8d4e804" }
func (*L2tpv3SetLookupKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3SetLookupKeyReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *L2tpv3SetLookupKeyReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *L2tpv3SetLookupKeyReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// L2tpv3SetTunnelCookies defines message 'l2tpv3_set_tunnel_cookies'.
type L2tpv3SetTunnelCookies struct {
	SwIfIndex       InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	NewLocalCookie  uint64         `binapi:"u64,name=new_local_cookie" json:"new_local_cookie,omitempty"`
	NewRemoteCookie uint64         `binapi:"u64,name=new_remote_cookie" json:"new_remote_cookie,omitempty"`
}

func (m *L2tpv3SetTunnelCookies) Reset()               { *m = L2tpv3SetTunnelCookies{} }
func (*L2tpv3SetTunnelCookies) GetMessageName() string { return "l2tpv3_set_tunnel_cookies" }
func (*L2tpv3SetTunnelCookies) GetCrcString() string   { return "b3f4faf7" }
func (*L2tpv3SetTunnelCookies) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3SetTunnelCookies) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 8 // m.NewLocalCookie
	size += 8 // m.NewRemoteCookie
	return size
}
func (m *L2tpv3SetTunnelCookies) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint64(m.NewLocalCookie)
	buf.EncodeUint64(m.NewRemoteCookie)
	return buf.Bytes(), nil
}
func (m *L2tpv3SetTunnelCookies) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.NewLocalCookie = buf.DecodeUint64()
	m.NewRemoteCookie = buf.DecodeUint64()
	return nil
}

// L2tpv3SetTunnelCookiesReply defines message 'l2tpv3_set_tunnel_cookies_reply'.
type L2tpv3SetTunnelCookiesReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2tpv3SetTunnelCookiesReply) Reset()               { *m = L2tpv3SetTunnelCookiesReply{} }
func (*L2tpv3SetTunnelCookiesReply) GetMessageName() string { return "l2tpv3_set_tunnel_cookies_reply" }
func (*L2tpv3SetTunnelCookiesReply) GetCrcString() string   { return "e8d4e804" }
func (*L2tpv3SetTunnelCookiesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3SetTunnelCookiesReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *L2tpv3SetTunnelCookiesReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *L2tpv3SetTunnelCookiesReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwIfL2tpv3TunnelDetails defines message 'sw_if_l2tpv3_tunnel_details'.
type SwIfL2tpv3TunnelDetails struct {
	SwIfIndex         InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InterfaceName     string         `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
	ClientAddress     Address        `binapi:"address,name=client_address" json:"client_address,omitempty"`
	OurAddress        Address        `binapi:"address,name=our_address" json:"our_address,omitempty"`
	LocalSessionID    uint32         `binapi:"u32,name=local_session_id" json:"local_session_id,omitempty"`
	RemoteSessionID   uint32         `binapi:"u32,name=remote_session_id" json:"remote_session_id,omitempty"`
	LocalCookie       []uint64       `binapi:"u64[2],name=local_cookie" json:"local_cookie,omitempty"`
	RemoteCookie      uint64         `binapi:"u64,name=remote_cookie" json:"remote_cookie,omitempty"`
	L2SublayerPresent bool           `binapi:"bool,name=l2_sublayer_present" json:"l2_sublayer_present,omitempty"`
}

func (m *SwIfL2tpv3TunnelDetails) Reset()               { *m = SwIfL2tpv3TunnelDetails{} }
func (*SwIfL2tpv3TunnelDetails) GetMessageName() string { return "sw_if_l2tpv3_tunnel_details" }
func (*SwIfL2tpv3TunnelDetails) GetCrcString() string   { return "1dab5c7e" }
func (*SwIfL2tpv3TunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwIfL2tpv3TunnelDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 64     // m.InterfaceName
	size += 4      // m.ClientAddress.Af
	size += 1 * 16 // m.ClientAddress.Un
	size += 4      // m.OurAddress.Af
	size += 1 * 16 // m.OurAddress.Un
	size += 4      // m.LocalSessionID
	size += 4      // m.RemoteSessionID
	size += 8 * 2  // m.LocalCookie
	size += 8      // m.RemoteCookie
	size += 1      // m.L2SublayerPresent
	return size
}
func (m *SwIfL2tpv3TunnelDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.InterfaceName, 64)
	buf.EncodeUint32(uint32(m.ClientAddress.Af))
	buf.EncodeBytes(m.ClientAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.OurAddress.Af))
	buf.EncodeBytes(m.OurAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.LocalSessionID)
	buf.EncodeUint32(m.RemoteSessionID)
	for i := 0; i < 2; i++ {
		var x uint64
		if i < len(m.LocalCookie) {
			x = uint64(m.LocalCookie[i])
		}
		buf.EncodeUint64(x)
	}
	buf.EncodeUint64(m.RemoteCookie)
	buf.EncodeBool(m.L2SublayerPresent)
	return buf.Bytes(), nil
}
func (m *SwIfL2tpv3TunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.InterfaceName = buf.DecodeString(64)
	m.ClientAddress.Af = AddressFamily(buf.DecodeUint32())
	copy(m.ClientAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.OurAddress.Af = AddressFamily(buf.DecodeUint32())
	copy(m.OurAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LocalSessionID = buf.DecodeUint32()
	m.RemoteSessionID = buf.DecodeUint32()
	m.LocalCookie = make([]uint64, 2)
	for i := 0; i < len(m.LocalCookie); i++ {
		m.LocalCookie[i] = buf.DecodeUint64()
	}
	m.RemoteCookie = buf.DecodeUint64()
	m.L2SublayerPresent = buf.DecodeBool()
	return nil
}

// SwIfL2tpv3TunnelDump defines message 'sw_if_l2tpv3_tunnel_dump'.
type SwIfL2tpv3TunnelDump struct{}

func (m *SwIfL2tpv3TunnelDump) Reset()               { *m = SwIfL2tpv3TunnelDump{} }
func (*SwIfL2tpv3TunnelDump) GetMessageName() string { return "sw_if_l2tpv3_tunnel_dump" }
func (*SwIfL2tpv3TunnelDump) GetCrcString() string   { return "51077d14" }
func (*SwIfL2tpv3TunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwIfL2tpv3TunnelDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SwIfL2tpv3TunnelDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SwIfL2tpv3TunnelDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_l2tp_binapi_init() }
func file_l2tp_binapi_init() {
	api.RegisterMessage((*L2tpv3CreateTunnel)(nil), "l2tpv3_create_tunnel_596892cb")
	api.RegisterMessage((*L2tpv3CreateTunnelReply)(nil), "l2tpv3_create_tunnel_reply_5383d31f")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisable)(nil), "l2tpv3_interface_enable_disable_3865946c")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisableReply)(nil), "l2tpv3_interface_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*L2tpv3SetLookupKey)(nil), "l2tpv3_set_lookup_key_c9892c86")
	api.RegisterMessage((*L2tpv3SetLookupKeyReply)(nil), "l2tpv3_set_lookup_key_reply_e8d4e804")
	api.RegisterMessage((*L2tpv3SetTunnelCookies)(nil), "l2tpv3_set_tunnel_cookies_b3f4faf7")
	api.RegisterMessage((*L2tpv3SetTunnelCookiesReply)(nil), "l2tpv3_set_tunnel_cookies_reply_e8d4e804")
	api.RegisterMessage((*SwIfL2tpv3TunnelDetails)(nil), "sw_if_l2tpv3_tunnel_details_1dab5c7e")
	api.RegisterMessage((*SwIfL2tpv3TunnelDump)(nil), "sw_if_l2tpv3_tunnel_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*L2tpv3CreateTunnel)(nil),
		(*L2tpv3CreateTunnelReply)(nil),
		(*L2tpv3InterfaceEnableDisable)(nil),
		(*L2tpv3InterfaceEnableDisableReply)(nil),
		(*L2tpv3SetLookupKey)(nil),
		(*L2tpv3SetLookupKeyReply)(nil),
		(*L2tpv3SetTunnelCookies)(nil),
		(*L2tpv3SetTunnelCookiesReply)(nil),
		(*SwIfL2tpv3TunnelDetails)(nil),
		(*SwIfL2tpv3TunnelDump)(nil),
	}
}