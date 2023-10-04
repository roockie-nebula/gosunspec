package model8

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block8 - Get Device Security Certificate - Security model for PKI

const (
	ModelID          = 8
	ModelLabel       = "Get Device Security Certificate"
	ModelDescription = "Security model for PKI"
)

const (
	Cert = "Cert"
	Fmt  = "Fmt"
	N    = "N"
)

type Block8Repeat struct {
	Cert uint16 `sunspec:"offset=0,access=r"`
}

type Block8 struct {
	Fmt sunspec.Enum16 `sunspec:"offset=0,access=r"`
	N   uint16         `sunspec:"offset=1,access=r"`

	Repeats []Block8Repeat
}

func (block *Block8) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 3,
		Blocks: []smdx.BlockElement{
			{
				Length: 2,
				Points: []smdx.PointElement{
					{Id: Fmt, Offset: 0, Type: typelabel.Enum16, Access: "r", Mandatory: true, Label: "Format", Description: "X.509 format of the certificate. DER or PEM."},
					{Id: N, Offset: 1, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "N", Description: "Number of registers to follow for the certificate"},
				},
			},
			{
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: Cert, Offset: 0, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "Cert", Description: "X.509 Certificate of the device"},
				},
			},
		}})
}
