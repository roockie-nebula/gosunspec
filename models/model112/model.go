package model112

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block112 - Inverter (Split Phase) FLOAT - Include this model for split phase inverter monitoring using float values

const (
	ModelID          = 112
	ModelLabel       = "Inverter (Split Phase) FLOAT"
	ModelDescription = "Include this model for split phase inverter monitoring using float values"
)

const (
	A       = "A"
	AphA    = "AphA"
	AphB    = "AphB"
	AphC    = "AphC"
	DCA     = "DCA"
	DCV     = "DCV"
	DCW     = "DCW"
	Evt1    = "Evt1"
	Evt2    = "Evt2"
	EvtVnd1 = "EvtVnd1"
	EvtVnd2 = "EvtVnd2"
	EvtVnd3 = "EvtVnd3"
	EvtVnd4 = "EvtVnd4"
	Hz      = "Hz"
	PF      = "PF"
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
	VA      = "VA"
	VAr     = "VAr"
	W       = "W"
	WH      = "WH"
)

type Block112 struct {
	A       float32            `sunspec:"offset=0"`
	AphA    float32            `sunspec:"offset=2"`
	AphB    float32            `sunspec:"offset=4"`
	AphC    float32            `sunspec:"offset=6"`
	PPVphAB float32            `sunspec:"offset=8"`
	PPVphBC float32            `sunspec:"offset=10"`
	PPVphCA float32            `sunspec:"offset=12"`
	PhVphA  float32            `sunspec:"offset=14"`
	PhVphB  float32            `sunspec:"offset=16"`
	PhVphC  float32            `sunspec:"offset=18"`
	W       float32            `sunspec:"offset=20"`
	Hz      float32            `sunspec:"offset=22"`
	VA      float32            `sunspec:"offset=24"`
	VAr     float32            `sunspec:"offset=26"`
	PF      float32            `sunspec:"offset=28"`
	WH      float32            `sunspec:"offset=30"`
	DCA     float32            `sunspec:"offset=32"`
	DCV     float32            `sunspec:"offset=34"`
	DCW     float32            `sunspec:"offset=36"`
	TmpCab  float32            `sunspec:"offset=38"`
	TmpSnk  float32            `sunspec:"offset=40"`
	TmpTrns float32            `sunspec:"offset=42"`
	TmpOt   float32            `sunspec:"offset=44"`
	St      sunspec.Enum16     `sunspec:"offset=46"`
	StVnd   sunspec.Enum16     `sunspec:"offset=47"`
	Evt1    sunspec.Bitfield32 `sunspec:"offset=48"`
	Evt2    sunspec.Bitfield32 `sunspec:"offset=50"`
	EvtVnd1 sunspec.Bitfield32 `sunspec:"offset=52"`
	EvtVnd2 sunspec.Bitfield32 `sunspec:"offset=54"`
	EvtVnd3 sunspec.Bitfield32 `sunspec:"offset=56"`
	EvtVnd4 sunspec.Bitfield32 `sunspec:"offset=58"`
}

func (block *Block112) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "inverter",
		Length: 60,
		Blocks: []smdx.BlockElement{
			{
				Length: 60,
				Points: []smdx.PointElement{
					{Id: A, Offset: 0, Type: typelabel.Float32, Units: "A", Mandatory: true, Label: "Amps", Description: "AC Current"},
					{Id: AphA, Offset: 2, Type: typelabel.Float32, Units: "A", Mandatory: true, Label: "Amps PhaseA", Description: "Phase A Current"},
					{Id: AphB, Offset: 4, Type: typelabel.Float32, Units: "A", Mandatory: true, Label: "Amps PhaseB", Description: "Phase B Current"},
					{Id: AphC, Offset: 6, Type: typelabel.Float32, Units: "A", Label: "Amps PhaseC", Description: "Phase C Current"},
					{Id: PPVphAB, Offset: 8, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage AB", Description: "Phase Voltage AB"},
					{Id: PPVphBC, Offset: 10, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage BC", Description: "Phase Voltage BC"},
					{Id: PPVphCA, Offset: 12, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage CA", Description: "Phase Voltage CA"},
					{Id: PhVphA, Offset: 14, Type: typelabel.Float32, Units: "V", Mandatory: true, Label: "Phase Voltage AN", Description: "Phase Voltage AN"},
					{Id: PhVphB, Offset: 16, Type: typelabel.Float32, Units: "V", Mandatory: true, Label: "Phase Voltage BN", Description: "Phase Voltage BN"},
					{Id: PhVphC, Offset: 18, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage CN", Description: "Phase Voltage CN"},
					{Id: W, Offset: 20, Type: typelabel.Float32, Units: "W", Mandatory: true, Label: "Watts", Description: "AC Power"},
					{Id: Hz, Offset: 22, Type: typelabel.Float32, Units: "Hz", Mandatory: true, Label: "Hz", Description: "Line Frequency"},
					{Id: VA, Offset: 24, Type: typelabel.Float32, Units: "VA", Label: "VA", Description: "AC Apparent Power"},
					{Id: VAr, Offset: 26, Type: typelabel.Float32, Units: "var", Label: "VAr", Description: "AC Reactive Power"},
					{Id: PF, Offset: 28, Type: typelabel.Float32, Units: "Pct", Label: "PF", Description: "AC Power Factor"},
					{Id: WH, Offset: 30, Type: typelabel.Float32, Units: "Wh", Mandatory: true, Label: "WattHours", Description: "AC Energy"},
					{Id: DCA, Offset: 32, Type: typelabel.Float32, Units: "A", Label: "DC Amps", Description: "DC Current"},
					{Id: DCV, Offset: 34, Type: typelabel.Float32, Units: "V", Label: "DC Voltage", Description: "DC Voltage"},
					{Id: DCW, Offset: 36, Type: typelabel.Float32, Units: "W", Label: "DC Watts", Description: "DC Power"},
					{Id: TmpCab, Offset: 38, Type: typelabel.Float32, Units: "C", Mandatory: true, Label: "Cabinet Temperature", Description: "Cabinet Temperature"},
					{Id: TmpSnk, Offset: 40, Type: typelabel.Float32, Units: "C", Label: "Heat Sink Temperature", Description: "Heat Sink Temperature"},
					{Id: TmpTrns, Offset: 42, Type: typelabel.Float32, Units: "C", Label: "Transformer Temperature", Description: "Transformer Temperature"},
					{Id: TmpOt, Offset: 44, Type: typelabel.Float32, Units: "C", Label: "Other Temperature", Description: "Other Temperature"},
					{Id: St, Offset: 46, Type: typelabel.Enum16, Mandatory: true, Label: "Operating State", Description: "Enumerated value.  Operating state"},
					{Id: StVnd, Offset: 47, Type: typelabel.Enum16, Label: "Vendor Operating State", Description: "Vendor specific operating state code"},
					{Id: Evt1, Offset: 48, Type: typelabel.Bitfield32, Mandatory: true, Label: "Event1", Description: "Bitmask value. Event fields"},
					{Id: Evt2, Offset: 50, Type: typelabel.Bitfield32, Mandatory: true, Label: "Event Bitfield 2", Description: "Reserved for future use"},
					{Id: EvtVnd1, Offset: 52, Type: typelabel.Bitfield32, Label: "Vendor Event Bitfield 1", Description: "Vendor defined events"},
					{Id: EvtVnd2, Offset: 54, Type: typelabel.Bitfield32, Label: "Vendor Event Bitfield 2", Description: "Vendor defined events"},
					{Id: EvtVnd3, Offset: 56, Type: typelabel.Bitfield32, Label: "Vendor Event Bitfield 3", Description: "Vendor defined events"},
					{Id: EvtVnd4, Offset: 58, Type: typelabel.Bitfield32, Label: "Vendor Event Bitfield 4", Description: "Vendor defined events"},
				},
			},
		}})
}
