package model132

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block132 - Volt-Watt - Volt-Watt

const (
	ModelID          = 132
	ModelLabel       = "Volt-Watt"
	ModelDescription = "Volt-Watt "
)

const (
	ActCrv       = "ActCrv"
	ActPt        = "ActPt"
	CrvNam       = "CrvNam"
	DeptRef      = "DeptRef"
	DeptRef_SF   = "DeptRef_SF"
	ModEna       = "ModEna"
	NCrv         = "NCrv"
	NPt          = "NPt"
	ReadOnly     = "ReadOnly"
	RmpDecTmm    = "RmpDecTmm"
	RmpIncDec_SF = "RmpIncDec_SF"
	RmpIncTmm    = "RmpIncTmm"
	RmpPt1Tms    = "RmpPt1Tms"
	RmpTms       = "RmpTms"
	RvrtTms      = "RvrtTms"
	V1           = "V1"
	V10          = "V10"
	V11          = "V11"
	V12          = "V12"
	V13          = "V13"
	V14          = "V14"
	V15          = "V15"
	V16          = "V16"
	V17          = "V17"
	V18          = "V18"
	V19          = "V19"
	V2           = "V2"
	V20          = "V20"
	V3           = "V3"
	V4           = "V4"
	V5           = "V5"
	V6           = "V6"
	V7           = "V7"
	V8           = "V8"
	V9           = "V9"
	V_SF         = "V_SF"
	W1           = "W1"
	W10          = "W10"
	W11          = "W11"
	W12          = "W12"
	W13          = "W13"
	W14          = "W14"
	W15          = "W15"
	W16          = "W16"
	W17          = "W17"
	W18          = "W18"
	W19          = "W19"
	W2           = "W2"
	W20          = "W20"
	W3           = "W3"
	W4           = "W4"
	W5           = "W5"
	W6           = "W6"
	W7           = "W7"
	W8           = "W8"
	W9           = "W9"
	WinTms       = "WinTms"
)

type Block132Repeat struct {
	ActPt     uint16         `sunspec:"offset=0,len=1,access=rw"`
	DeptRef   sunspec.Enum16 `sunspec:"offset=1,len=1,access=rw"`
	V1        uint16         `sunspec:"offset=2,len=1,sf=V_SF,access=rw"`
	W1        int16          `sunspec:"offset=3,len=1,sf=DeptRef_SF,access=rw"`
	V2        uint16         `sunspec:"offset=4,len=1,sf=V_SF,access=rw"`
	W2        int16          `sunspec:"offset=5,len=1,sf=DeptRef_SF,access=rw"`
	V3        uint16         `sunspec:"offset=6,len=1,sf=V_SF,access=rw"`
	W3        int16          `sunspec:"offset=7,len=1,sf=DeptRef_SF,access=rw"`
	V4        uint16         `sunspec:"offset=8,len=1,sf=V_SF,access=rw"`
	W4        int16          `sunspec:"offset=9,len=1,sf=DeptRef_SF,access=rw"`
	V5        uint16         `sunspec:"offset=10,len=1,sf=V_SF,access=rw"`
	W5        int16          `sunspec:"offset=11,len=1,sf=DeptRef_SF,access=rw"`
	V6        uint16         `sunspec:"offset=12,len=1,sf=V_SF,access=rw"`
	W6        int16          `sunspec:"offset=13,len=1,sf=DeptRef_SF,access=rw"`
	V7        uint16         `sunspec:"offset=14,len=1,sf=V_SF,access=rw"`
	W7        int16          `sunspec:"offset=15,len=1,sf=DeptRef_SF,access=rw"`
	V8        uint16         `sunspec:"offset=16,len=1,sf=V_SF,access=rw"`
	W8        int16          `sunspec:"offset=17,len=1,sf=DeptRef_SF,access=rw"`
	V9        uint16         `sunspec:"offset=18,len=1,sf=V_SF,access=rw"`
	W9        int16          `sunspec:"offset=19,len=1,sf=DeptRef_SF,access=rw"`
	V10       uint16         `sunspec:"offset=20,len=1,sf=V_SF,access=rw"`
	W10       int16          `sunspec:"offset=21,len=1,sf=DeptRef_SF,access=rw"`
	V11       uint16         `sunspec:"offset=22,len=1,sf=V_SF,access=rw"`
	W11       int16          `sunspec:"offset=23,len=1,sf=DeptRef_SF,access=rw"`
	V12       uint16         `sunspec:"offset=24,len=1,sf=V_SF,access=rw"`
	W12       int16          `sunspec:"offset=25,len=1,sf=DeptRef_SF,access=rw"`
	V13       uint16         `sunspec:"offset=26,len=1,sf=V_SF,access=rw"`
	W13       int16          `sunspec:"offset=27,len=1,sf=DeptRef_SF,access=rw"`
	V14       uint16         `sunspec:"offset=28,len=1,sf=V_SF,access=rw"`
	W14       int16          `sunspec:"offset=29,len=1,sf=DeptRef_SF,access=rw"`
	V15       uint16         `sunspec:"offset=30,len=1,sf=V_SF,access=rw"`
	W15       int16          `sunspec:"offset=31,len=1,sf=DeptRef_SF,access=rw"`
	V16       uint16         `sunspec:"offset=32,len=1,sf=V_SF,access=rw"`
	W16       int16          `sunspec:"offset=33,len=1,sf=DeptRef_SF,access=rw"`
	V17       uint16         `sunspec:"offset=34,len=1,sf=V_SF,access=rw"`
	W17       int16          `sunspec:"offset=35,len=1,sf=DeptRef_SF,access=rw"`
	V18       uint16         `sunspec:"offset=36,len=1,sf=V_SF,access=rw"`
	W18       int16          `sunspec:"offset=37,len=1,sf=DeptRef_SF,access=rw"`
	V19       uint16         `sunspec:"offset=38,len=1,sf=V_SF,access=rw"`
	W19       int16          `sunspec:"offset=39,len=1,sf=DeptRef_SF,access=rw"`
	V20       uint16         `sunspec:"offset=40,len=1,sf=V_SF,access=rw"`
	W20       int16          `sunspec:"offset=41,len=1,sf=DeptRef_SF,access=rw"`
	CrvNam    string         `sunspec:"offset=42,len=8,access=rw"`
	RmpPt1Tms uint16         `sunspec:"offset=50,len=1,access=rw"`
	RmpDecTmm uint16         `sunspec:"offset=51,len=1,sf=RmpIncDec_SF,access=rw"`
	RmpIncTmm uint16         `sunspec:"offset=52,len=1,sf=RmpIncDec_SF,access=rw"`
	ReadOnly  sunspec.Enum16 `sunspec:"offset=53,len=1,access=r"`
}

type Block132 struct {
	ActCrv       uint16              `sunspec:"offset=0,len=1,access=rw"`
	ModEna       sunspec.Bitfield16  `sunspec:"offset=1,len=1,access=rw"`
	WinTms       uint16              `sunspec:"offset=2,len=1,access=rw"`
	RvrtTms      uint16              `sunspec:"offset=3,len=1,access=rw"`
	RmpTms       uint16              `sunspec:"offset=4,len=1,access=rw"`
	NCrv         uint16              `sunspec:"offset=5,len=1,access=r"`
	NPt          uint16              `sunspec:"offset=6,len=1,access=r"`
	V_SF         sunspec.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	DeptRef_SF   sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	RmpIncDec_SF sunspec.ScaleFactor `sunspec:"offset=9,len=1,access=r"`

	Repeats []Block132Repeat
}

func (block *Block132) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "volt_watt",
		Length: 64,
		Blocks: []smdx.BlockElement{
			{
				Length: 10,
				Type:   "fixed",
				Points: []smdx.PointElement{
					{Id: ActCrv, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActCrv", Description: "Index of active curve. 0=no active curve."},
					{Id: ModEna, Offset: 1, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "Is Volt-Watt control active."},
					{Id: WinTms, Offset: 2, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "WinTms", Description: "Time window for volt-watt change."},
					{Id: RvrtTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RvrtTms", Description: "Timeout period for volt-watt curve selection."},
					{Id: RmpTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpTms", Description: "Ramp time for moving from current mode to new mode."},
					{Id: NCrv, Offset: 5, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NCrv", Description: "Number of curves supported (recommend min. 4)."},
					{Id: NPt, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NPt", Description: "Number of points in array (maximum 20)."},
					{Id: V_SF, Offset: 7, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "V_SF", Description: "Scale factor for percent VRef."},
					{Id: DeptRef_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "DeptRef_SF", Description: "Scale factor for % DeptRef."},
					{Id: RmpIncDec_SF, Offset: 9, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "RmpIncDec_SF", Description: "Scale factor for increment and decrement ramps."},
				},
			},
			{Name: "curve",
				Length: 54,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActPt", Description: "Number of active points in array."},
					{Id: DeptRef, Offset: 1, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "DeptRef", Description: "Defines the meaning of the Watts DeptRef.  1=% WMax 2=% WAvail"},
					{Id: V1, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Mandatory: true, Label: "V1", Description: "Point 1 Volts."},
					{Id: W1, Offset: 3, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Mandatory: true, Label: "W1", Description: "Point 1 Watts."},
					{Id: V2, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V2", Description: "Point 2 Volts."},
					{Id: W2, Offset: 5, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W2", Description: "Point 2 Watts."},
					{Id: V3, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V3", Description: "Point 3 Volts."},
					{Id: W3, Offset: 7, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W3", Description: "Point 3 Watts."},
					{Id: V4, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V4", Description: "Point 4 Volts."},
					{Id: W4, Offset: 9, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W4", Description: "Point 4 Watts."},
					{Id: V5, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V5", Description: "Point 5 Volts."},
					{Id: W5, Offset: 11, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W5", Description: "Point 5 Watts."},
					{Id: V6, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V6", Description: "Point 6 Volts."},
					{Id: W6, Offset: 13, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W6", Description: "Point 6 Watts."},
					{Id: V7, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V7", Description: "Point 7 Volts."},
					{Id: W7, Offset: 15, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W7", Description: "Point 7 Watts."},
					{Id: V8, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V8", Description: "Point 8 Volts."},
					{Id: W8, Offset: 17, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W8", Description: "Point 8 Watts."},
					{Id: V9, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V9", Description: "Point 9 Volts."},
					{Id: W9, Offset: 19, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W9", Description: "Point 9 Watts."},
					{Id: V10, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V10", Description: "Point 10 Volts."},
					{Id: W10, Offset: 21, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W10", Description: "Point 10 Watts."},
					{Id: V11, Offset: 22, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V11", Description: "Point 11 Volts."},
					{Id: W11, Offset: 23, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W11", Description: "Point 11 Watts."},
					{Id: V12, Offset: 24, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V12", Description: "Point 12 Volts."},
					{Id: W12, Offset: 25, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W12", Description: "Point 12 Watts."},
					{Id: V13, Offset: 26, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V13", Description: "Point 13 Volts."},
					{Id: W13, Offset: 27, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W13", Description: "Point 13 Watts."},
					{Id: V14, Offset: 28, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V14", Description: "Point 14 Volts."},
					{Id: W14, Offset: 29, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W14", Description: "Point 14 Watts."},
					{Id: V15, Offset: 30, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V15", Description: "Point 15 Volts."},
					{Id: W15, Offset: 31, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W15", Description: "Point 15 Watts."},
					{Id: V16, Offset: 32, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V16", Description: "Point 16 Volts."},
					{Id: W16, Offset: 33, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W16", Description: "Point 16 Watts."},
					{Id: V17, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V17", Description: "Point 17 Volts."},
					{Id: W17, Offset: 35, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W17", Description: "Point 17 Watts."},
					{Id: V18, Offset: 36, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V18", Description: "Point 18 Volts."},
					{Id: W18, Offset: 37, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W18", Description: "Point 18 Watts."},
					{Id: V19, Offset: 38, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V19", Description: "Point 19 Volts."},
					{Id: W19, Offset: 39, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W19", Description: "Point 19 Watts."},
					{Id: V20, Offset: 40, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V20", Description: "Point 20 Volts."},
					{Id: W20, Offset: 41, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "W20", Description: "Point 20 Watts."},
					{Id: CrvNam, Offset: 42, Type: typelabel.String, Access: "rw", Length: 8, Label: "CrvNam", Description: "Optional description for curve."},
					{Id: RmpPt1Tms, Offset: 50, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpPt1Tms", Description: "The time of the PT1 in seconds (time to accomplish a change of 95%)."},
					{Id: RmpDecTmm, Offset: 51, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% WMax/min", Access: "rw", Length: 1, Label: "RmpDecTmm", Description: "The maximum rate at which the watt value may be reduced in response to changes in the voltage value."},
					{Id: RmpIncTmm, Offset: 52, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% WMax/min", Access: "rw", Length: 1, Label: "RmpIncTmm", Description: "The maximum rate at which the watt value may be increased in response to changes in the voltage value."},
					{Id: ReadOnly, Offset: 53, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true, Label: "ReadOnly", Description: "Enumerated value indicates if curve is read-only or can be modified."},
				},
			},
		}})
}
