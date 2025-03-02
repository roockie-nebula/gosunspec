package model125

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block125 - Pricing - Pricing Signal

const (
	ModelID          = 125
	ModelLabel       = "Pricing"
	ModelDescription = "Pricing Signal  "
)

const (
	ModEna  = "ModEna"
	Pad     = "Pad"
	RmpTms  = "RmpTms"
	RvtTms  = "RvtTms"
	Sig     = "Sig"
	SigType = "SigType"
	Sig_SF  = "Sig_SF"
	WinTms  = "WinTms"
)

type Block125 struct {
	ModEna  sunspec.Bitfield16  `sunspec:"offset=0,len=1,access=rw"`
	SigType sunspec.Enum16      `sunspec:"offset=1,len=1,access=rw"`
	Sig     int16               `sunspec:"offset=2,len=1,sf=Sig_SF,access=rw"`
	WinTms  uint16              `sunspec:"offset=3,len=1,access=rw"`
	RvtTms  uint16              `sunspec:"offset=4,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=5,len=1,access=rw"`
	Sig_SF  sunspec.ScaleFactor `sunspec:"offset=6,len=1,access=r"`
	Pad     sunspec.Pad         `sunspec:"offset=7,len=1,access=r"`
}

func (block *Block125) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "pricing",
		Length: 8,
		Blocks: []smdx.BlockElement{
			{
				Length: 8,
				Type:   "fixed",
				Points: []smdx.PointElement{
					{Id: ModEna, Offset: 0, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "Is price-based charge/discharge mode active?"},
					{Id: SigType, Offset: 1, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "SigType", Description: "Meaning of the pricing signal. When a Price schedule is used, type must match the schedule range variable description."},
					{Id: Sig, Offset: 2, Type: typelabel.Int16, ScaleFactor: "Sig_SF", Access: "rw", Length: 1, Mandatory: true, Label: "Sig", Description: "Utility/ESP specific pricing signal. Content depends on pricing signal type. When H/M/L type is specified. Low=0; Med=1; High=2."},
					{Id: WinTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "WinTms", Description: "Time window for charge/discharge pricing change."},
					{Id: RvtTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RvtTms", Description: "Timeout period for charge/discharge pricing change."},
					{Id: RmpTms, Offset: 5, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpTms", Description: "Ramp time for moving from current charge or discharge level to new level."},
					{Id: Sig_SF, Offset: 6, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Sig_SF", Description: "Pricing signal scale factor."},
					{Id: Pad, Offset: 7, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
