package model807

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block807 - Flow Battery String Model -

const (
	ModelID          = 807
	ModelLabel       = "Flow Battery String Model"
	ModelDescription = ""
)

const (
	CellVAvg        = "CellVAvg"
	CellVMax        = "CellVMax"
	CellVMaxMod     = "CellVMaxMod"
	CellVMaxStk     = "CellVMaxStk"
	CellVMin        = "CellVMin"
	CellVMinMod     = "CellVMinMod"
	CellVMinStk     = "CellVMinStk"
	CellV_SF        = "CellV_SF"
	Evt1            = "Evt1"
	Evt2            = "Evt2"
	EvtVnd1         = "EvtVnd1"
	EvtVnd2         = "EvtVnd2"
	Idx             = "Idx"
	ModAnoTmp       = "ModAnoTmp"
	ModCatTmp       = "ModCatTmp"
	ModCellVAvg     = "ModCellVAvg"
	ModCellVMax     = "ModCellVMax"
	ModCellVMaxCell = "ModCellVMaxCell"
	ModCellVMin     = "ModCellVMin"
	ModCellVMinCell = "ModCellVMinCell"
	ModConFail      = "ModConFail"
	ModConSt        = "ModConSt"
	ModDisRsn       = "ModDisRsn"
	ModEvt1         = "ModEvt1"
	ModEvt2         = "ModEvt2"
	ModIdx          = "ModIdx"
	ModNStk         = "ModNStk"
	ModOCV          = "ModOCV"
	ModSetCon       = "ModSetCon"
	ModSetEna       = "ModSetEna"
	ModSoC          = "ModSoC"
	ModSt           = "ModSt"
	ModV            = "ModV"
	ModVAvg         = "ModVAvg"
	ModVMax         = "ModVMax"
	ModVMaxMod      = "ModVMaxMod"
	ModVMin         = "ModVMin"
	ModVMinMod      = "ModVMinMod"
	ModV_SF         = "ModV_SF"
	NMod            = "NMod"
	NModCon         = "NModCon"
	OCV_SF          = "OCV_SF"
	Pad1            = "Pad1"
	SoC_SF          = "SoC_SF"
	TmpAvg          = "TmpAvg"
	TmpMax          = "TmpMax"
	TmpMaxMod       = "TmpMaxMod"
	TmpMin          = "TmpMin"
	TmpMinMod       = "TmpMinMod"
	Tmp_SF          = "Tmp_SF"
)

type Block807Repeat struct {
	ModIdx          uint16             `sunspec:"offset=0"`
	ModNStk         uint16             `sunspec:"offset=1"`
	ModSt           sunspec.Bitfield32 `sunspec:"offset=2"`
	ModSoC          uint16             `sunspec:"offset=4,sf=SoC_SF"`
	ModOCV          uint16             `sunspec:"offset=5,sf=OCV_SF"`
	ModV            uint16             `sunspec:"offset=6,sf=ModV_SF"`
	ModCellVMax     uint16             `sunspec:"offset=7,sf=CellV_SF"`
	ModCellVMaxCell uint16             `sunspec:"offset=8"`
	ModCellVMin     uint16             `sunspec:"offset=9,sf=CellV_SF"`
	ModCellVMinCell uint16             `sunspec:"offset=10"`
	ModCellVAvg     uint16             `sunspec:"offset=11,sf=CellV_SF"`
	ModAnoTmp       uint16             `sunspec:"offset=12,sf=Tmp_SF"`
	ModCatTmp       uint16             `sunspec:"offset=13,sf=Tmp_SF"`
	ModConSt        sunspec.Bitfield32 `sunspec:"offset=14"`
	ModEvt1         sunspec.Bitfield32 `sunspec:"offset=16"`
	ModEvt2         sunspec.Bitfield32 `sunspec:"offset=18"`
	ModConFail      sunspec.Enum16     `sunspec:"offset=20"`
	ModSetEna       sunspec.Enum16     `sunspec:"offset=21,access=rw"`
	ModSetCon       sunspec.Enum16     `sunspec:"offset=22,access=rw"`
	ModDisRsn       sunspec.Enum16     `sunspec:"offset=23"`
}

type Block807 struct {
	Idx         uint16              `sunspec:"offset=0"`
	NMod        uint16              `sunspec:"offset=1"`
	NModCon     uint16              `sunspec:"offset=2"`
	ModVMax     uint16              `sunspec:"offset=3,sf=ModV_SF"`
	ModVMaxMod  uint16              `sunspec:"offset=4"`
	ModVMin     uint16              `sunspec:"offset=5,sf=ModV_SF"`
	ModVMinMod  uint16              `sunspec:"offset=6"`
	ModVAvg     uint16              `sunspec:"offset=7,sf=ModV_SF"`
	CellVMax    uint16              `sunspec:"offset=8,sf=CellV_SF"`
	CellVMaxMod uint16              `sunspec:"offset=9"`
	CellVMaxStk uint16              `sunspec:"offset=10"`
	CellVMin    uint16              `sunspec:"offset=11,sf=CellV_SF"`
	CellVMinMod uint16              `sunspec:"offset=12"`
	CellVMinStk uint16              `sunspec:"offset=13"`
	CellVAvg    uint16              `sunspec:"offset=14,sf=CellV_SF"`
	TmpMax      int16               `sunspec:"offset=15,sf=Tmp_SF"`
	TmpMaxMod   uint16              `sunspec:"offset=16"`
	TmpMin      int16               `sunspec:"offset=17,sf=Tmp_SF"`
	TmpMinMod   uint16              `sunspec:"offset=18"`
	TmpAvg      int16               `sunspec:"offset=19,sf=Tmp_SF"`
	Evt1        sunspec.Bitfield32  `sunspec:"offset=20"`
	Evt2        sunspec.Bitfield32  `sunspec:"offset=22"`
	EvtVnd1     sunspec.Bitfield32  `sunspec:"offset=24"`
	EvtVnd2     sunspec.Bitfield32  `sunspec:"offset=26"`
	ModV_SF     sunspec.ScaleFactor `sunspec:"offset=28"`
	CellV_SF    sunspec.ScaleFactor `sunspec:"offset=29"`
	Tmp_SF      sunspec.ScaleFactor `sunspec:"offset=30"`
	SoC_SF      sunspec.ScaleFactor `sunspec:"offset=31"`
	OCV_SF      sunspec.ScaleFactor `sunspec:"offset=32"`
	Pad1        sunspec.Pad         `sunspec:"offset=33"`

	Repeats []Block807Repeat
}

func (block *Block807) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "flow_battery_string",
		Length: 58,
		Blocks: []smdx.BlockElement{
			{
				Length: 34,
				Points: []smdx.PointElement{
					{Id: Idx, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "String Index", Description: "Index of the string within the bank."},
					{Id: NMod, Offset: 1, Type: typelabel.Uint16, Mandatory: true, Label: "Module Count", Description: "Number of modules in this string."},
					{Id: NModCon, Offset: 2, Type: typelabel.Uint16, Mandatory: true, Label: "Connected Module Count", Description: "Number of electrically connected modules in this string."},
					{Id: ModVMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "ModV_SF", Units: "V", Mandatory: true, Label: "Max Module Voltage", Description: "Maximum voltage for all modules in the string."},
					{Id: ModVMaxMod, Offset: 4, Type: typelabel.Uint16, Label: "Max Module Voltage Module", Description: "Module with the maximum voltage."},
					{Id: ModVMin, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "ModV_SF", Units: "V", Mandatory: true, Label: "Min Module Voltage", Description: "Minimum voltage for all modules in the string."},
					{Id: ModVMinMod, Offset: 6, Type: typelabel.Uint16, Label: "Min Module Voltage Module", Description: "Module with the minimum voltage."},
					{Id: ModVAvg, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "ModV_SF", Units: "V", Mandatory: true, Label: "Average Module Voltage", Description: "Average voltage for all modules in the string."},
					{Id: CellVMax, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Max Cell Voltage", Description: "Maximum voltage for all cells in the string."},
					{Id: CellVMaxMod, Offset: 9, Type: typelabel.Uint16, Label: "Max Cell Voltage Module", Description: "Module containing the cell with the maximum voltage."},
					{Id: CellVMaxStk, Offset: 10, Type: typelabel.Uint16, Label: "Max Cell Voltage Stack", Description: "Stack containing the cell with the maximum voltage."},
					{Id: CellVMin, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Min Cell Voltage", Description: "Minimum voltage for all cells in the string."},
					{Id: CellVMinMod, Offset: 12, Type: typelabel.Uint16, Label: "Min Cell Voltage Module", Description: "Module containing the cell with the minimum voltage."},
					{Id: CellVMinStk, Offset: 13, Type: typelabel.Uint16, Label: "Min Cell Voltage Stack", Description: "Stack containing the cell with the minimum voltage."},
					{Id: CellVAvg, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Average Cell Voltage", Description: "Average voltage for all cells in the string."},
					{Id: TmpMax, Offset: 15, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Max Temperature", Description: "Maximum electrolyte temperature for all modules in the string."},
					{Id: TmpMaxMod, Offset: 16, Type: typelabel.Uint16, Label: "Max Temperature Module", Description: "Module with the maximum temperature."},
					{Id: TmpMin, Offset: 17, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Min Temperature", Description: "Minimum electrolyte temperature for all modules in the string."},
					{Id: TmpMinMod, Offset: 18, Type: typelabel.Uint16, Label: "Min Temperature Module", Description: "Module with the minimum temperature."},
					{Id: TmpAvg, Offset: 19, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Average Temperature", Description: "Average electrolyte temperature for all modules in the string."},
					{Id: Evt1, Offset: 20, Type: typelabel.Bitfield32, Mandatory: true, Label: "String Event 1", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: Evt2, Offset: 22, Type: typelabel.Bitfield32, Mandatory: true, Label: "String Event 2", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: EvtVnd1, Offset: 24, Type: typelabel.Bitfield32, Mandatory: true, Label: "Vendor Event Bitfield 1", Description: "Vendor defined events."},
					{Id: EvtVnd2, Offset: 26, Type: typelabel.Bitfield32, Mandatory: true, Label: "Vendor Event Bitfield 2", Description: "Vendor defined events."},
					{Id: ModV_SF, Offset: 28, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: CellV_SF, Offset: 29, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: Tmp_SF, Offset: 30, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: SoC_SF, Offset: 31, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: OCV_SF, Offset: 32, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: Pad1, Offset: 33, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
				},
			},
			{Name: "module",
				Length: 24,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ModIdx, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Module Index", Description: "Index of the module within the string."},
					{Id: ModNStk, Offset: 1, Type: typelabel.Uint16, Mandatory: true, Label: "Stack Count", Description: "Number of stacks in this module."},
					{Id: ModSt, Offset: 2, Type: typelabel.Bitfield32, Mandatory: true, Label: "Module Status", Description: "Current status of the module."},
					{Id: ModSoC, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%", Mandatory: true, Label: "Module State of Charge", Description: "State of charge for this module."},
					{Id: ModOCV, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "OCV_SF", Units: "V", Mandatory: true, Label: "Open Circuit Voltage", Description: "Open circuit voltage for this module."},
					{Id: ModV, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "ModV_SF", Units: "V", Mandatory: true, Label: "External Voltage", Description: "External voltage fo this module."},
					{Id: ModCellVMax, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Maximum Cell Voltage", Description: "Maximum voltage for all cells in this module."},
					{Id: ModCellVMaxCell, Offset: 8, Type: typelabel.Uint16, Label: "Max Cell Voltage Cell", Description: "Cell with the maximum cell voltage."},
					{Id: ModCellVMin, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Minimum Cell Voltage", Description: "Minimum voltage for all cells in this module."},
					{Id: ModCellVMinCell, Offset: 10, Type: typelabel.Uint16, Label: "Min Cell Voltage Cell", Description: "Cell with the minimum cell voltage."},
					{Id: ModCellVAvg, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Average Cell Voltage", Description: "Average voltage for all cells in this module."},
					{Id: ModAnoTmp, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "Tmp_SF", Units: "C", Label: "Anolyte Temperature", Description: ""},
					{Id: ModCatTmp, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "Tmp_SF", Units: "C", Label: "Catholyte Temperature", Description: ""},
					{Id: ModConSt, Offset: 14, Type: typelabel.Bitfield32, Label: "Contactor Status", Description: ""},
					{Id: ModEvt1, Offset: 16, Type: typelabel.Bitfield32, Mandatory: true, Label: "Module Event 1", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: ModEvt2, Offset: 18, Type: typelabel.Bitfield32, Mandatory: true, Label: "Module Event 2", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: ModConFail, Offset: 20, Type: typelabel.Enum16, Label: "Connection Failure Reason", Description: ""},
					{Id: ModSetEna, Offset: 21, Type: typelabel.Enum16, Access: "rw", Label: "Enable/Disable Module", Description: "Enables and disables the module."},
					{Id: ModSetCon, Offset: 22, Type: typelabel.Enum16, Access: "rw", Label: "Connect/Disconnect Module ", Description: "Connects and disconnects the module."},
					{Id: ModDisRsn, Offset: 23, Type: typelabel.Enum16, Label: "Disabled Reason", Description: "Reason why the module is currently disabled."},
				},
			},
		}})
}
