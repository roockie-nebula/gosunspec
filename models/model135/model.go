package model135

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block135 - LFRT - Low Frequency Ride-through

const (
	ModelID          = 135
	ModelLabel       = "LFRT"
	ModelDescription = "Low Frequency Ride-through"
)

const (
	ActCrv   = "ActCrv"
	ActPt    = "ActPt"
	CrvNam   = "CrvNam"
	Hz1      = "Hz1"
	Hz10     = "Hz10"
	Hz11     = "Hz11"
	Hz12     = "Hz12"
	Hz13     = "Hz13"
	Hz14     = "Hz14"
	Hz15     = "Hz15"
	Hz16     = "Hz16"
	Hz17     = "Hz17"
	Hz18     = "Hz18"
	Hz19     = "Hz19"
	Hz2      = "Hz2"
	Hz20     = "Hz20"
	Hz3      = "Hz3"
	Hz4      = "Hz4"
	Hz5      = "Hz5"
	Hz6      = "Hz6"
	Hz7      = "Hz7"
	Hz8      = "Hz8"
	Hz9      = "Hz9"
	Hz_SF    = "Hz_SF"
	ModEna   = "ModEna"
	NCrv     = "NCrv"
	NPt      = "NPt"
	Pad      = "Pad"
	ReadOnly = "ReadOnly"
	RmpTms   = "RmpTms"
	RvrtTms  = "RvrtTms"
	Tms1     = "Tms1"
	Tms10    = "Tms10"
	Tms11    = "Tms11"
	Tms12    = "Tms12"
	Tms13    = "Tms13"
	Tms14    = "Tms14"
	Tms15    = "Tms15"
	Tms16    = "Tms16"
	Tms17    = "Tms17"
	Tms18    = "Tms18"
	Tms19    = "Tms19"
	Tms2     = "Tms2"
	Tms20    = "Tms20"
	Tms3     = "Tms3"
	Tms4     = "Tms4"
	Tms5     = "Tms5"
	Tms6     = "Tms6"
	Tms7     = "Tms7"
	Tms8     = "Tms8"
	Tms9     = "Tms9"
	Tms_SF   = "Tms_SF"
	WinTms   = "WinTms"
)

type Block135Repeat struct {
	ActPt    uint16         `sunspec:"offset=0,len=1,access=rw"`
	Tms1     uint16         `sunspec:"offset=1,len=1,sf=Tms_SF,access=rw"`
	Hz1      uint16         `sunspec:"offset=2,len=1,sf=Hz_SF,access=rw"`
	Tms2     uint16         `sunspec:"offset=3,len=1,sf=Tms_SF,access=rw"`
	Hz2      uint16         `sunspec:"offset=4,len=1,sf=Hz_SF,access=rw"`
	Tms3     uint16         `sunspec:"offset=5,len=1,sf=Tms_SF,access=rw"`
	Hz3      uint16         `sunspec:"offset=6,len=1,sf=Hz_SF,access=rw"`
	Tms4     uint16         `sunspec:"offset=7,len=1,sf=Tms_SF,access=rw"`
	Hz4      uint16         `sunspec:"offset=8,len=1,sf=Hz_SF,access=rw"`
	Tms5     uint16         `sunspec:"offset=9,len=1,sf=Tms_SF,access=rw"`
	Hz5      uint16         `sunspec:"offset=10,len=1,sf=Hz_SF,access=rw"`
	Tms6     uint16         `sunspec:"offset=11,len=1,sf=Tms_SF,access=rw"`
	Hz6      uint16         `sunspec:"offset=12,len=1,sf=Hz_SF,access=rw"`
	Tms7     uint16         `sunspec:"offset=13,len=1,sf=Tms_SF,access=rw"`
	Hz7      uint16         `sunspec:"offset=14,len=1,sf=Hz_SF,access=rw"`
	Tms8     uint16         `sunspec:"offset=15,len=1,sf=Tms_SF,access=rw"`
	Hz8      uint16         `sunspec:"offset=16,len=1,sf=Hz_SF,access=rw"`
	Tms9     uint16         `sunspec:"offset=17,len=1,sf=Tms_SF,access=rw"`
	Hz9      uint16         `sunspec:"offset=18,len=1,sf=Hz_SF,access=rw"`
	Tms10    uint16         `sunspec:"offset=19,len=1,sf=Tms_SF,access=rw"`
	Hz10     uint16         `sunspec:"offset=20,len=1,sf=Hz_SF,access=rw"`
	Tms11    uint16         `sunspec:"offset=21,len=1,sf=Tms_SF,access=rw"`
	Hz11     uint16         `sunspec:"offset=22,len=1,sf=Hz_SF,access=rw"`
	Tms12    uint16         `sunspec:"offset=23,len=1,sf=Tms_SF,access=rw"`
	Hz12     uint16         `sunspec:"offset=24,len=1,sf=Hz_SF,access=rw"`
	Tms13    uint16         `sunspec:"offset=25,len=1,sf=Tms_SF,access=rw"`
	Hz13     uint16         `sunspec:"offset=26,len=1,sf=Hz_SF,access=rw"`
	Tms14    uint16         `sunspec:"offset=27,len=1,sf=Tms_SF,access=rw"`
	Hz14     uint16         `sunspec:"offset=28,len=1,sf=Hz_SF,access=rw"`
	Tms15    uint16         `sunspec:"offset=29,len=1,sf=Tms_SF,access=rw"`
	Hz15     uint16         `sunspec:"offset=30,len=1,sf=Hz_SF,access=rw"`
	Tms16    uint16         `sunspec:"offset=31,len=1,sf=Tms_SF,access=rw"`
	Hz16     uint16         `sunspec:"offset=32,len=1,sf=Hz_SF,access=rw"`
	Tms17    uint16         `sunspec:"offset=33,len=1,sf=Tms_SF,access=rw"`
	Hz17     uint16         `sunspec:"offset=34,len=1,sf=Hz_SF,access=rw"`
	Tms18    uint16         `sunspec:"offset=35,len=1,sf=Tms_SF,access=rw"`
	Hz18     uint16         `sunspec:"offset=36,len=1,sf=Hz_SF,access=rw"`
	Tms19    uint16         `sunspec:"offset=37,len=1,sf=Tms_SF,access=rw"`
	Hz19     uint16         `sunspec:"offset=38,len=1,sf=Hz_SF,access=rw"`
	Tms20    uint16         `sunspec:"offset=39,len=1,sf=Tms_SF,access=rw"`
	Hz20     uint16         `sunspec:"offset=40,len=1,sf=Hz_SF,access=rw"`
	CrvNam   string         `sunspec:"offset=41,len=8,access=rw"`
	ReadOnly sunspec.Enum16 `sunspec:"offset=49,len=1,access=r"`
}

type Block135 struct {
	ActCrv  uint16              `sunspec:"offset=0,len=1,access=rw"`
	ModEna  sunspec.Bitfield16  `sunspec:"offset=1,len=1,access=rw"`
	WinTms  uint16              `sunspec:"offset=2,len=1,access=rw"`
	RvrtTms uint16              `sunspec:"offset=3,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=4,len=1,access=rw"`
	NCrv    uint16              `sunspec:"offset=5,len=1,access=r"`
	NPt     uint16              `sunspec:"offset=6,len=1,access=r"`
	Tms_SF  sunspec.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	Hz_SF   sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	Pad     sunspec.Pad         `sunspec:"offset=9,len=1,access=r"`

	Repeats []Block135Repeat
}

func (block *Block135) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lfrt",
		Length: 60,
		Blocks: []smdx.BlockElement{
			{
				Length: 10,
				Type:   "fixed",
				Points: []smdx.PointElement{
					{Id: ActCrv, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActCrv", Description: "Index of active curve. 0=no active curve."},
					{Id: ModEna, Offset: 1, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "LHzRT control mode. Enable active curve.  Bitfield value."},
					{Id: WinTms, Offset: 2, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "WinTms", Description: "Time window for LFRT change."},
					{Id: RvrtTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RvrtTms", Description: "Timeout period for LFRT curve selection."},
					{Id: RmpTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpTms", Description: "Ramp time for moving from current mode to new mode."},
					{Id: NCrv, Offset: 5, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NCrv", Description: "Number of curves supported (recommend 4)."},
					{Id: NPt, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NPt", Description: "Number of curve points supported (maximum of 20)."},
					{Id: Tms_SF, Offset: 7, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Tms_SF", Description: "Scale factor for duration."},
					{Id: Hz_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Hz_SF", Description: "Scale factor for frequency."},
					{Id: Pad, Offset: 9, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
			{Name: "curve",
				Length: 50,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActPt", Description: "Number of active points in array."},
					{Id: Tms1, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Mandatory: true, Label: "Tms1", Description: "Point 1 must disconnect duration."},
					{Id: Hz1, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Mandatory: true, Label: "Hz1", Description: "Point 1 must disconnect frequency."},
					{Id: Tms2, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms2", Description: "Point 2 must disconnect duration."},
					{Id: Hz2, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz2", Description: "Point 2 must disconnect frequency."},
					{Id: Tms3, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms3", Description: "Point 3 must disconnect duration."},
					{Id: Hz3, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz3", Description: "Point 3 must disconnect frequency."},
					{Id: Tms4, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms4", Description: "Point 4 must disconnect duration."},
					{Id: Hz4, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz4", Description: "Point 4 must disconnect frequency."},
					{Id: Tms5, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms5", Description: "Point 5 must disconnect duration."},
					{Id: Hz5, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz5", Description: "Point 5 must disconnect frequency."},
					{Id: Tms6, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms6", Description: "Point 6 must disconnect duration."},
					{Id: Hz6, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz6", Description: "Point 6 must disconnect frequency."},
					{Id: Tms7, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms7", Description: "Point 7 must disconnect duration."},
					{Id: Hz7, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz7", Description: "Point 7 must disconnect frequency."},
					{Id: Tms8, Offset: 15, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms8", Description: "Point 8 must disconnect duration."},
					{Id: Hz8, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz8", Description: "Point 8 must disconnect frequency."},
					{Id: Tms9, Offset: 17, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms9", Description: "Point 9 must disconnect duration."},
					{Id: Hz9, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz9", Description: "Point 9 must disconnect frequency."},
					{Id: Tms10, Offset: 19, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms10", Description: "Point 10 must disconnect duration."},
					{Id: Hz10, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz10", Description: "Point 10 must disconnect frequency."},
					{Id: Tms11, Offset: 21, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms11", Description: "Point 11 must disconnect duration."},
					{Id: Hz11, Offset: 22, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz11", Description: "Point 11 must disconnect frequency."},
					{Id: Tms12, Offset: 23, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms12", Description: "Point 12 must disconnect duration."},
					{Id: Hz12, Offset: 24, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz12", Description: "Point 12 must disconnect frequency."},
					{Id: Tms13, Offset: 25, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms13", Description: "Point 13 must disconnect duration."},
					{Id: Hz13, Offset: 26, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz13", Description: "Point 13 must disconnect frequency."},
					{Id: Tms14, Offset: 27, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms14", Description: "Point 14 must disconnect duration."},
					{Id: Hz14, Offset: 28, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz14", Description: "Point 14 must disconnect frequency."},
					{Id: Tms15, Offset: 29, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms15", Description: "Point 15 must disconnect duration."},
					{Id: Hz15, Offset: 30, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz15", Description: "Point 15 must disconnect frequency."},
					{Id: Tms16, Offset: 31, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms16", Description: "Point 16 must disconnect duration."},
					{Id: Hz16, Offset: 32, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz16", Description: "Point 16 must disconnect frequency."},
					{Id: Tms17, Offset: 33, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms17", Description: "Point 17 must disconnect duration."},
					{Id: Hz17, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz17", Description: "Point 17 must disconnect frequency."},
					{Id: Tms18, Offset: 35, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms18", Description: "Point 18 must disconnect duration."},
					{Id: Hz18, Offset: 36, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz18", Description: "Point 18 must disconnect frequency."},
					{Id: Tms19, Offset: 37, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms19", Description: "Point 19 must disconnect duration."},
					{Id: Hz19, Offset: 38, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz19", Description: "Point 19 must disconnect frequency."},
					{Id: Tms20, Offset: 39, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms20", Description: "Point 20 must disconnect duration."},
					{Id: Hz20, Offset: 40, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "Hz20", Description: "Point 20 must disconnect frequency."},
					{Id: CrvNam, Offset: 41, Type: typelabel.String, Access: "rw", Length: 8, Label: "CrvNam", Description: "Optional description for curve."},
					{Id: ReadOnly, Offset: 49, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true, Label: "ReadOnly", Description: "Enumerated value indicates if curve is read-only or can be modified."},
				},
			},
		}})
}
