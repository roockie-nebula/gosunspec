package model16

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block16 - Simple IP Network - Include this model for a simple IPv4 network stack

const (
	ModelID          = 16
	ModelLabel       = "Simple IP Network"
	ModelDescription = "Include this model for a simple IPv4 network stack"
)

const (
	Addr   = "Addr"
	Cfg    = "Cfg"
	Ctl    = "Ctl"
	DNS1   = "DNS1"
	DNS2   = "DNS2"
	Gw     = "Gw"
	LnkCtl = "LnkCtl"
	MAC    = "MAC"
	Msk    = "Msk"
	Nam    = "Nam"
	Pad    = "Pad"
)

type Block16 struct {
	Nam    string             `sunspec:"offset=0,len=4,access=rw"`
	Cfg    sunspec.Enum16     `sunspec:"offset=4"`
	Ctl    sunspec.Enum16     `sunspec:"offset=5,access=rw"`
	Addr   string             `sunspec:"offset=6,len=8,access=rw"`
	Msk    string             `sunspec:"offset=14,len=8,access=rw"`
	Gw     string             `sunspec:"offset=22,len=8,access=rw"`
	DNS1   string             `sunspec:"offset=30,len=8,access=rw"`
	DNS2   string             `sunspec:"offset=38,len=8,access=rw"`
	MAC    sunspec.Eui48      `sunspec:"offset=46,access=r"`
	LnkCtl sunspec.Bitfield16 `sunspec:"offset=50,access=rw"`
	Pad    sunspec.Pad        `sunspec:"offset=51"`
}

func (block *Block16) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 52,
		Blocks: []smdx.BlockElement{
			{
				Length: 52,
				Points: []smdx.PointElement{
					{Id: Nam, Offset: 0, Type: typelabel.String, Access: "rw", Length: 4, Label: "Name", Description: "Interface name.  (8 chars)"},
					{Id: Cfg, Offset: 4, Type: typelabel.Enum16, Mandatory: true, Label: "Config", Description: "Enumerated value.  Force IPv4 configuration method"},
					{Id: Ctl, Offset: 5, Type: typelabel.Enum16, Access: "rw", Mandatory: true, Label: "Control", Description: "Bitmask value Configure use of services"},
					{Id: Addr, Offset: 6, Type: typelabel.String, Access: "rw", Length: 8, Mandatory: true, Label: "Address", Description: "IP address"},
					{Id: Msk, Offset: 14, Type: typelabel.String, Access: "rw", Length: 8, Mandatory: true, Label: "Netmask", Description: "Netmask"},
					{Id: Gw, Offset: 22, Type: typelabel.String, Access: "rw", Length: 8, Label: "Gateway", Description: "Gateway IP address"},
					{Id: DNS1, Offset: 30, Type: typelabel.String, Access: "rw", Length: 8, Label: "DNS1", Description: "32 bit IP address of DNS server"},
					{Id: DNS2, Offset: 38, Type: typelabel.String, Access: "rw", Length: 8, Label: "DNS2", Description: "32 bit IP address of DNS server"},
					{Id: MAC, Offset: 46, Type: typelabel.Eui48, Access: "r", Label: "MAC", Description: "IEEE MAC address of this interface"},
					{Id: LnkCtl, Offset: 50, Type: typelabel.Bitfield16, Access: "rw", Label: "Link Control", Description: "Bitmask value.  Link control flags"},
					{Id: Pad, Offset: 51, Type: typelabel.Pad},
				},
			},
		}})
}
