// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model102

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block102 - Inverter (Split-Phase) - Include this model for split phase inverter monitoring

const (
	ModelID = 102
)

const (
	A       = "A"
	A_SF    = "A_SF"
	AphA    = "AphA"
	AphB    = "AphB"
	AphC    = "AphC"
	DCA     = "DCA"
	DCA_SF  = "DCA_SF"
	DCV     = "DCV"
	DCV_SF  = "DCV_SF"
	DCW     = "DCW"
	DCW_SF  = "DCW_SF"
	Evt1    = "Evt1"
	Evt2    = "Evt2"
	EvtVnd1 = "EvtVnd1"
	EvtVnd2 = "EvtVnd2"
	EvtVnd3 = "EvtVnd3"
	EvtVnd4 = "EvtVnd4"
	Hz      = "Hz"
	Hz_SF   = "Hz_SF"
	PF      = "PF"
	PF_SF   = "PF_SF"
	PPVphAB = "PPVphAB"
	PPVphBC = "PPVphBC"
	PPVphCA = "PPVphCA"
	PhVphA  = "PhVphA"
	PhVphB  = "PhVphB"
	PhVphC  = "PhVphC"
	St      = "St"
	StVnd   = "StVnd"
	TmpCab  = "TmpCab"
	TmpOt   = "TmpOt"
	TmpSnk  = "TmpSnk"
	TmpTrns = "TmpTrns"
	Tmp_SF  = "Tmp_SF"
	VA      = "VA"
	VA_SF   = "VA_SF"
	VAr     = "VAr"
	VAr_SF  = "VAr_SF"
	V_SF    = "V_SF"
	W       = "W"
	WH      = "WH"
	WH_SF   = "WH_SF"
	W_SF    = "W_SF"
)

type Block102 struct {
	A       uint16           `sunspec:"offset=0,sf=A_SF"`
	AphA    uint16           `sunspec:"offset=1,sf=A_SF"`
	AphB    uint16           `sunspec:"offset=2,sf=A_SF"`
	AphC    uint16           `sunspec:"offset=3,sf=A_SF"`
	A_SF    core.ScaleFactor `sunspec:"offset=4"`
	PPVphAB uint16           `sunspec:"offset=5,sf=V_SF"`
	PPVphBC uint16           `sunspec:"offset=6,sf=V_SF"`
	PPVphCA uint16           `sunspec:"offset=7,sf=V_SF"`
	PhVphA  uint16           `sunspec:"offset=8,sf=V_SF"`
	PhVphB  uint16           `sunspec:"offset=9,sf=V_SF"`
	PhVphC  uint16           `sunspec:"offset=10,sf=V_SF"`
	V_SF    core.ScaleFactor `sunspec:"offset=11"`
	W       int16            `sunspec:"offset=12,sf=W_SF"`
	W_SF    core.ScaleFactor `sunspec:"offset=13"`
	Hz      uint16           `sunspec:"offset=14,sf=Hz_SF"`
	Hz_SF   core.ScaleFactor `sunspec:"offset=15"`
	VA      int16            `sunspec:"offset=16,sf=VA_SF"`
	VA_SF   core.ScaleFactor `sunspec:"offset=17"`
	VAr     int16            `sunspec:"offset=18,sf=VAr_SF"`
	VAr_SF  core.ScaleFactor `sunspec:"offset=19"`
	PF      int16            `sunspec:"offset=20,sf=PF_SF"`
	PF_SF   core.ScaleFactor `sunspec:"offset=21"`
	WH      core.Acc32       `sunspec:"offset=22,sf=WH_SF"`
	WH_SF   core.ScaleFactor `sunspec:"offset=24"`
	DCA     uint16           `sunspec:"offset=25,sf=DCA_SF"`
	DCA_SF  core.ScaleFactor `sunspec:"offset=26"`
	DCV     uint16           `sunspec:"offset=27,sf=DCV_SF"`
	DCV_SF  core.ScaleFactor `sunspec:"offset=28"`
	DCW     int16            `sunspec:"offset=29,sf=DCW_SF"`
	DCW_SF  core.ScaleFactor `sunspec:"offset=30"`
	TmpCab  int16            `sunspec:"offset=31,sf=Tmp_SF"`
	TmpSnk  int16            `sunspec:"offset=32,sf=Tmp_SF"`
	TmpTrns int16            `sunspec:"offset=33,sf=Tmp_SF"`
	TmpOt   int16            `sunspec:"offset=34,sf=Tmp_SF"`
	Tmp_SF  core.ScaleFactor `sunspec:"offset=35"`
	St      core.Enum16      `sunspec:"offset=36"`
	StVnd   core.Enum16      `sunspec:"offset=37"`
	Evt1    core.Bitfield32  `sunspec:"offset=38"`
	Evt2    core.Bitfield32  `sunspec:"offset=40"`
	EvtVnd1 core.Bitfield32  `sunspec:"offset=42"`
	EvtVnd2 core.Bitfield32  `sunspec:"offset=44"`
	EvtVnd3 core.Bitfield32  `sunspec:"offset=46"`
	EvtVnd4 core.Bitfield32  `sunspec:"offset=48"`
}

func (self *Block102) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "inverter",
		Length: 50,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 50,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: A, Offset: 0, Type: typelabel.Uint16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AphA, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AphB, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AphC, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "A_SF", Units: "A"},
					smdx.PointElement{Id: A_SF, Offset: 4, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: PPVphAB, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PPVphBC, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PPVphCA, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PhVphA, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: PhVphB, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: PhVphC, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: V_SF, Offset: 11, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: W, Offset: 12, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: W_SF, Offset: 13, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: Hz, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Mandatory: true},
					smdx.PointElement{Id: Hz_SF, Offset: 15, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: VA, Offset: 16, Type: typelabel.Int16, ScaleFactor: "VA_SF", Units: "VA"},
					smdx.PointElement{Id: VA_SF, Offset: 17, Type: typelabel.Sunssf},
					smdx.PointElement{Id: VAr, Offset: 18, Type: typelabel.Int16, ScaleFactor: "VAr_SF", Units: "var"},
					smdx.PointElement{Id: VAr_SF, Offset: 19, Type: typelabel.Sunssf},
					smdx.PointElement{Id: PF, Offset: 20, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "Pct"},
					smdx.PointElement{Id: PF_SF, Offset: 21, Type: typelabel.Sunssf},
					smdx.PointElement{Id: WH, Offset: 22, Type: typelabel.Acc32, ScaleFactor: "WH_SF", Units: "Wh", Mandatory: true},
					smdx.PointElement{Id: WH_SF, Offset: 24, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: DCA, Offset: 25, Type: typelabel.Uint16, ScaleFactor: "DCA_SF", Units: "A"},
					smdx.PointElement{Id: DCA_SF, Offset: 26, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCV, Offset: 27, Type: typelabel.Uint16, ScaleFactor: "DCV_SF", Units: "V"},
					smdx.PointElement{Id: DCV_SF, Offset: 28, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCW, Offset: 29, Type: typelabel.Int16, ScaleFactor: "DCW_SF", Units: "W"},
					smdx.PointElement{Id: DCW_SF, Offset: 30, Type: typelabel.Sunssf},
					smdx.PointElement{Id: TmpCab, Offset: 31, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: TmpSnk, Offset: 32, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C"},
					smdx.PointElement{Id: TmpTrns, Offset: 33, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C"},
					smdx.PointElement{Id: TmpOt, Offset: 34, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C"},
					smdx.PointElement{Id: Tmp_SF, Offset: 35, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: St, Offset: 36, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: StVnd, Offset: 37, Type: typelabel.Enum16},
					smdx.PointElement{Id: Evt1, Offset: 38, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: Evt2, Offset: 40, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: EvtVnd1, Offset: 42, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: EvtVnd2, Offset: 44, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: EvtVnd3, Offset: 46, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: EvtVnd4, Offset: 48, Type: typelabel.Bitfield32},
				},
			},
		}})
}
