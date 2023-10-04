package model17

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block17 - Serial Interface - Include this model for serial interface configuration support

const (
	ModelID          = 17
	ModelLabel       = "Serial Interface"
	ModelDescription = "Include this model for serial interface configuration support"
)

const (
	Bits = "Bits"
	Dup  = "Dup"
	Flw  = "Flw"
	Nam  = "Nam"
	Pcol = "Pcol"
	Pty  = "Pty"
	Rte  = "Rte"
	Typ  = "Typ"
)

type Block17 struct {
	Nam  string         `sunspec:"offset=0,len=4,access=rw"`
	Rte  uint32         `sunspec:"offset=4,access=rw"`
	Bits uint16         `sunspec:"offset=6,access=rw"`
	Pty  sunspec.Enum16 `sunspec:"offset=7,access=rw"`
	Dup  sunspec.Enum16 `sunspec:"offset=8,access=rw"`
	Flw  sunspec.Enum16 `sunspec:"offset=9,access=rw"`
	Typ  sunspec.Enum16 `sunspec:"offset=10"`
	Pcol sunspec.Enum16 `sunspec:"offset=11"`
}

func (block *Block17) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 12,
		Blocks: []smdx.BlockElement{
			{
				Length: 12,
				Points: []smdx.PointElement{
					{Id: Nam, Offset: 0, Type: typelabel.String, Access: "rw", Length: 4, Label: "Name", Description: "Interface name (8 chars)"},
					{Id: Rte, Offset: 4, Type: typelabel.Uint32, Units: "bps", Access: "rw", Mandatory: true, Label: "Rate", Description: "Interface baud rate in bits per second"},
					{Id: Bits, Offset: 6, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Bits", Description: "Number of data bits per character"},
					{Id: Pty, Offset: 7, Type: typelabel.Enum16, Access: "rw", Mandatory: true, Label: "Parity", Description: "Bitmask value.  Parity setting"},
					{Id: Dup, Offset: 8, Type: typelabel.Enum16, Access: "rw", Label: "Duplex", Description: "Enumerated value.  Duplex mode"},
					{Id: Flw, Offset: 9, Type: typelabel.Enum16, Access: "rw", Label: "Flow Control", Description: "Flow Control Method"},
					{Id: Typ, Offset: 10, Type: typelabel.Enum16, Label: "Interface Type", Description: "Enumerated value.  Interface type"},
					{Id: Pcol, Offset: 11, Type: typelabel.Enum16, Label: "Protocol", Description: "Enumerated value. Serial protocol selection"},
				},
			},
		}})
}
