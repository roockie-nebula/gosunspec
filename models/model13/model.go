package model13

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block13 - IPv6 - Include to support an IPv6 protocol stack on this interface

const (
	ModelID          = 13
	ModelLabel       = "IPv6"
	ModelDescription = "Include to support an IPv6 protocol stack on this interface"
)

const (
	Addr    = "Addr"
	CIDR    = "CIDR"
	Cap     = "Cap"
	Cfg     = "Cfg"
	CfgSt   = "CfgSt"
	ChgSt   = "ChgSt"
	Ctl     = "Ctl"
	DNS1    = "DNS1"
	DNS2    = "DNS2"
	DomNam  = "DomNam"
	Gw      = "Gw"
	HostNam = "HostNam"
	NTP1    = "NTP1"
	NTP2    = "NTP2"
	Nam     = "Nam"
	Pad     = "Pad"
)

type Block13 struct {
	Nam     string             `sunspec:"offset=0,len=4,access=rw"`
	CfgSt   sunspec.Enum16     `sunspec:"offset=4"`
	ChgSt   sunspec.Bitfield16 `sunspec:"offset=5"`
	Cap     sunspec.Bitfield16 `sunspec:"offset=6"`
	Cfg     sunspec.Enum16     `sunspec:"offset=7,access=rw"`
	Ctl     sunspec.Enum16     `sunspec:"offset=8,access=rw"`
	Addr    string             `sunspec:"offset=9,len=20,access=rw"`
	CIDR    string             `sunspec:"offset=29,len=20,access=rw"`
	Gw      string             `sunspec:"offset=49,len=20,access=rw"`
	DNS1    string             `sunspec:"offset=69,len=20,access=rw"`
	DNS2    string             `sunspec:"offset=89,len=20,access=rw"`
	NTP1    string             `sunspec:"offset=109,len=20,access=rw"`
	NTP2    string             `sunspec:"offset=129,len=20,access=rw"`
	DomNam  string             `sunspec:"offset=149,len=12,access=rw"`
	HostNam string             `sunspec:"offset=161,len=12,access=rw"`
	Pad     sunspec.Pad        `sunspec:"offset=173"`
}

func (block *Block13) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 174,
		Blocks: []smdx.BlockElement{
			{
				Length: 174,
				Points: []smdx.PointElement{
					{Id: Nam, Offset: 0, Type: typelabel.String, Access: "rw", Length: 4, Label: "Name", Description: "Interface name"},
					{Id: CfgSt, Offset: 4, Type: typelabel.Enum16, Mandatory: true, Label: "Config Status", Description: "Enumerated value.  Configuration status"},
					{Id: ChgSt, Offset: 5, Type: typelabel.Bitfield16, Mandatory: true, Label: "Change Status", Description: "Bitmask value.  A configuration change is pending"},
					{Id: Cap, Offset: 6, Type: typelabel.Bitfield16, Mandatory: true, Label: "Config Capability", Description: "Bitmask value. Identify capable sources of configuration"},
					{Id: Cfg, Offset: 7, Type: typelabel.Enum16, Access: "rw", Mandatory: true, Label: "IPv6 Config", Description: "Enumerated value.  Configuration method used."},
					{Id: Ctl, Offset: 8, Type: typelabel.Enum16, Access: "rw", Mandatory: true, Label: "Control", Description: "Bitmask value.  Configure use of services"},
					{Id: Addr, Offset: 9, Type: typelabel.String, Access: "rw", Length: 20, Mandatory: true, Label: "IP", Description: "IPv6 numeric address as a dotted string xxxx.xxxx.xxxx.xxxx"},
					{Id: CIDR, Offset: 29, Type: typelabel.String, Access: "rw", Length: 20, Label: "CIDR", Description: "Classless Inter-Domain Routing Number"},
					{Id: Gw, Offset: 49, Type: typelabel.String, Access: "rw", Length: 20, Label: "Gateway", Description: "IPv6 numeric address as a dotted string xxxx.xxxx.xxxx.xxxx"},
					{Id: DNS1, Offset: 69, Type: typelabel.String, Access: "rw", Length: 20, Label: "DNS1", Description: "IPv6 numeric DNS address as a dotted string xxxx.xxxx.xxxx.xxxx"},
					{Id: DNS2, Offset: 89, Type: typelabel.String, Access: "rw", Length: 20, Label: "DNS2", Description: "IPv6 numeric DNS address as a dotted string xxxx.xxxx.xxxx.xxxx"},
					{Id: NTP1, Offset: 109, Type: typelabel.String, Access: "rw", Length: 20, Label: "NTP1", Description: "IPv6 numeric NTP address as a name or dotted string xxxx.xxxx.xxxx.xxxx"},
					{Id: NTP2, Offset: 129, Type: typelabel.String, Access: "rw", Length: 20, Label: "NTP2", Description: "IPv6 numeric NTP address as a name or dotted string xxxx.xxxx.xxxx.xxxx"},
					{Id: DomNam, Offset: 149, Type: typelabel.String, Access: "rw", Length: 12, Label: "Domain", Description: "Domain name (24 chars max)"},
					{Id: HostNam, Offset: 161, Type: typelabel.String, Access: "rw", Length: 12, Label: "Host Name", Description: "Host name (24 chars max)"},
					{Id: Pad, Offset: 173, Type: typelabel.Pad},
				},
			},
		}})
}
