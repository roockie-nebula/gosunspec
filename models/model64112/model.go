package model64112

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block64112 - OutBack FM Charge Controller -

const (
	ModelID          = 64112
	ModelLabel       = "OutBack FM Charge Controller"
	ModelDescription = ""
)

const (
	AH_SF                         = "AH_SF"
	CC_Config_AUX_Divert_Hyst_V   = "CC_Config_AUX_Divert_Hyst_V"
	CC_Config_AUX_Divert_Rel_V    = "CC_Config_AUX_Divert_Rel_V"
	CC_Config_AUX_Divert_dly_time = "CC_Config_AUX_Divert_dly_time"
	CC_Config_AUX_Divert_h_time   = "CC_Config_AUX_Divert_h_time"
	CC_Config_AUX_Error_batt_V    = "CC_Config_AUX_Error_batt_V"
	CC_Config_AUX_L_Batt_disc     = "CC_Config_AUX_L_Batt_disc"
	CC_Config_AUX_L_Batt_dly      = "CC_Config_AUX_L_Batt_dly"
	CC_Config_AUX_L_Batt_rcon     = "CC_Config_AUX_L_Batt_rcon"
	CC_Config_AUX_Nlite_Off_hist  = "CC_Config_AUX_Nlite_Off_hist"
	CC_Config_AUX_Nlite_On_hist   = "CC_Config_AUX_Nlite_On_hist"
	CC_Config_AUX_Nlite_On_tm     = "CC_Config_AUX_Nlite_On_tm"
	CC_Config_AUX_Nlite_ThrsV     = "CC_Config_AUX_Nlite_ThrsV"
	CC_Config_AUX_PV_trg_h_tm     = "CC_Config_AUX_PV_trg_h_tm"
	CC_Config_AUX_PV_triggerV     = "CC_Config_AUX_PV_triggerV"
	CC_Config_AUX_Vent_fan_V      = "CC_Config_AUX_Vent_fan_V"
	CC_Config_AUX_control         = "CC_Config_AUX_control"
	CC_Config_AUX_mode            = "CC_Config_AUX_mode"
	CC_Config_AUX_polarity        = "CC_Config_AUX_polarity"
	CC_Config_AUX_state           = "CC_Config_AUX_state"
	CC_Config_DataLog_Absorb_T    = "CC_Config_DataLog_Absorb_T"
	CC_Config_DataLog_Clear       = "CC_Config_DataLog_Clear"
	CC_Config_DataLog_Clr_Comp    = "CC_Config_DataLog_Clr_Comp"
	CC_Config_DataLog_Cur_Day_off = "CC_Config_DataLog_Cur_Day_off"
	CC_Config_DataLog_Daily_AH    = "CC_Config_DataLog_Daily_AH"
	CC_Config_DataLog_Daily_KWH   = "CC_Config_DataLog_Daily_KWH"
	CC_Config_DataLog_Day_offset  = "CC_Config_DataLog_Day_offset"
	CC_Config_DataLog_Float_T     = "CC_Config_DataLog_Float_T"
	CC_Config_DataLog_Max_Batt_V  = "CC_Config_DataLog_Max_Batt_V"
	CC_Config_DataLog_Max_Input_V = "CC_Config_DataLog_Max_Input_V"
	CC_Config_DataLog_Max_Out_A   = "CC_Config_DataLog_Max_Out_A"
	CC_Config_DataLog_Max_Out_W   = "CC_Config_DataLog_Max_Out_W"
	CC_Config_DataLog_Min_Batt_V  = "CC_Config_DataLog_Min_Batt_V"
	CC_Config_MPPT_mode           = "CC_Config_MPPT_mode"
	CC_Config_MajorFWRev          = "CC_Config_MajorFWRev"
	CC_Config_MidFWRev            = "CC_Config_MidFWRev"
	CC_Config_MinorFWRev          = "CC_Config_MinorFWRev"
	CC_Config_U_Pick_Duty_cyc     = "CC_Config_U_Pick_Duty_cyc"
	CC_Config_absorb_End_A        = "CC_Config_absorb_End_A"
	CC_Config_absorb_Hr           = "CC_Config_absorb_Hr"
	CC_Config_absorb_V            = "CC_Config_absorb_V"
	CC_Config_auto_equalize       = "CC_Config_auto_equalize"
	CC_Config_auto_restart        = "CC_Config_auto_restart"
	CC_Config_equalize_Hr         = "CC_Config_equalize_Hr"
	CC_Config_equalize_V          = "CC_Config_equalize_V"
	CC_Config_fault               = "CC_Config_fault"
	CC_Config_float_V             = "CC_Config_float_V"
	CC_Config_grid_tie            = "CC_Config_grid_tie"
	CC_Config_max_Chg_A           = "CC_Config_max_Chg_A"
	CC_Config_rebulk_V            = "CC_Config_rebulk_V"
	CC_Config_snooze_mode_A       = "CC_Config_snooze_mode_A"
	CC_Config_sweep_max           = "CC_Config_sweep_max"
	CC_Config_sweep_width         = "CC_Config_sweep_width"
	CC_Config_temp_comp           = "CC_Config_temp_comp"
	CC_Config_temp_comp_hlimt     = "CC_Config_temp_comp_hlimt"
	CC_Config_temp_comp_llimt     = "CC_Config_temp_comp_llimt"
	CC_Config_wakeup_VOC          = "CC_Config_wakeup_VOC"
	CC_Config_wakeup_interval     = "CC_Config_wakeup_interval"
	C_SF                          = "C_SF"
	H_SF                          = "H_SF"
	KWH_SF                        = "KWH_SF"
	P_SF                          = "P_SF"
	Port                          = "Port"
	V_SF                          = "V_SF"
)

type Block64112 struct {
	Port                          uint16              `sunspec:"offset=0"`
	V_SF                          sunspec.ScaleFactor `sunspec:"offset=1"`
	C_SF                          sunspec.ScaleFactor `sunspec:"offset=2"`
	H_SF                          sunspec.ScaleFactor `sunspec:"offset=3"`
	P_SF                          sunspec.ScaleFactor `sunspec:"offset=4"`
	AH_SF                         sunspec.ScaleFactor `sunspec:"offset=5"`
	KWH_SF                        sunspec.ScaleFactor `sunspec:"offset=6"`
	CC_Config_fault               sunspec.Bitfield16  `sunspec:"offset=7"`
	CC_Config_absorb_V            uint16              `sunspec:"offset=8,sf=V_SF"`
	CC_Config_absorb_Hr           uint16              `sunspec:"offset=9,sf=H_SF"`
	CC_Config_absorb_End_A        uint16              `sunspec:"offset=10,sf=V_SF"`
	CC_Config_rebulk_V            uint16              `sunspec:"offset=11,sf=V_SF"`
	CC_Config_float_V             uint16              `sunspec:"offset=12,sf=V_SF"`
	CC_Config_max_Chg_A           uint16              `sunspec:"offset=13,sf=V_SF"`
	CC_Config_equalize_V          uint16              `sunspec:"offset=14,sf=V_SF"`
	CC_Config_equalize_Hr         uint16              `sunspec:"offset=15"`
	CC_Config_auto_equalize       uint16              `sunspec:"offset=16"`
	CC_Config_MPPT_mode           sunspec.Enum16      `sunspec:"offset=17"`
	CC_Config_sweep_width         sunspec.Enum16      `sunspec:"offset=18"`
	CC_Config_sweep_max           sunspec.Enum16      `sunspec:"offset=19"`
	CC_Config_U_Pick_Duty_cyc     uint16              `sunspec:"offset=20,sf=V_SF"`
	CC_Config_grid_tie            sunspec.Enum16      `sunspec:"offset=21"`
	CC_Config_temp_comp           sunspec.Enum16      `sunspec:"offset=22"`
	CC_Config_temp_comp_llimt     uint16              `sunspec:"offset=23,sf=V_SF"`
	CC_Config_temp_comp_hlimt     uint16              `sunspec:"offset=24,sf=V_SF"`
	CC_Config_auto_restart        sunspec.Enum16      `sunspec:"offset=25"`
	CC_Config_wakeup_VOC          uint16              `sunspec:"offset=26,sf=V_SF"`
	CC_Config_snooze_mode_A       uint16              `sunspec:"offset=27,sf=V_SF"`
	CC_Config_wakeup_interval     uint16              `sunspec:"offset=28"`
	CC_Config_AUX_mode            sunspec.Enum16      `sunspec:"offset=29"`
	CC_Config_AUX_control         sunspec.Enum16      `sunspec:"offset=30"`
	CC_Config_AUX_state           sunspec.Enum16      `sunspec:"offset=31"`
	CC_Config_AUX_polarity        sunspec.Enum16      `sunspec:"offset=32"`
	CC_Config_AUX_L_Batt_disc     uint16              `sunspec:"offset=33,sf=V_SF"`
	CC_Config_AUX_L_Batt_rcon     uint16              `sunspec:"offset=34,sf=V_SF"`
	CC_Config_AUX_L_Batt_dly      uint16              `sunspec:"offset=35"`
	CC_Config_AUX_Vent_fan_V      uint16              `sunspec:"offset=36,sf=V_SF"`
	CC_Config_AUX_PV_triggerV     uint16              `sunspec:"offset=37,sf=V_SF"`
	CC_Config_AUX_PV_trg_h_tm     uint16              `sunspec:"offset=38"`
	CC_Config_AUX_Nlite_ThrsV     uint16              `sunspec:"offset=39,sf=V_SF"`
	CC_Config_AUX_Nlite_On_tm     uint16              `sunspec:"offset=40,sf=H_SF"`
	CC_Config_AUX_Nlite_On_hist   uint16              `sunspec:"offset=41"`
	CC_Config_AUX_Nlite_Off_hist  uint16              `sunspec:"offset=42"`
	CC_Config_AUX_Error_batt_V    uint16              `sunspec:"offset=43,sf=V_SF"`
	CC_Config_AUX_Divert_h_time   uint16              `sunspec:"offset=44,sf=V_SF"`
	CC_Config_AUX_Divert_dly_time uint16              `sunspec:"offset=45"`
	CC_Config_AUX_Divert_Rel_V    uint16              `sunspec:"offset=46,sf=V_SF"`
	CC_Config_AUX_Divert_Hyst_V   uint16              `sunspec:"offset=47,sf=V_SF"`
	CC_Config_MajorFWRev          uint16              `sunspec:"offset=48"`
	CC_Config_MidFWRev            uint16              `sunspec:"offset=49"`
	CC_Config_MinorFWRev          uint16              `sunspec:"offset=50"`
	CC_Config_DataLog_Day_offset  uint16              `sunspec:"offset=51"`
	CC_Config_DataLog_Cur_Day_off uint16              `sunspec:"offset=52"`
	CC_Config_DataLog_Daily_AH    uint16              `sunspec:"offset=53"`
	CC_Config_DataLog_Daily_KWH   uint16              `sunspec:"offset=54,sf=KWH_SF"`
	CC_Config_DataLog_Max_Out_A   uint16              `sunspec:"offset=55,sf=V_SF"`
	CC_Config_DataLog_Max_Out_W   uint16              `sunspec:"offset=56,sf=V_SF"`
	CC_Config_DataLog_Absorb_T    uint16              `sunspec:"offset=57"`
	CC_Config_DataLog_Float_T     uint16              `sunspec:"offset=58"`
	CC_Config_DataLog_Min_Batt_V  uint16              `sunspec:"offset=59,sf=V_SF"`
	CC_Config_DataLog_Max_Batt_V  uint16              `sunspec:"offset=60,sf=V_SF"`
	CC_Config_DataLog_Max_Input_V uint16              `sunspec:"offset=61,sf=V_SF"`
	CC_Config_DataLog_Clear       uint16              `sunspec:"offset=62"`
	CC_Config_DataLog_Clr_Comp    uint16              `sunspec:"offset=63"`
}

func (block *Block64112) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 64,
		Blocks: []smdx.BlockElement{
			{
				Length: 64,
				Points: []smdx.PointElement{
					{Id: Port, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Port Number", Description: ""},
					{Id: V_SF, Offset: 1, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: C_SF, Offset: 2, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: H_SF, Offset: 3, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: P_SF, Offset: 4, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: AH_SF, Offset: 5, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: KWH_SF, Offset: 6, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: CC_Config_fault, Offset: 7, Type: typelabel.Bitfield16, Mandatory: true, Label: "Faults", Description: ""},
					{Id: CC_Config_absorb_V, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Absorb", Description: ""},
					{Id: CC_Config_absorb_Hr, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "H_SF", Units: "Tmh", Mandatory: true, Label: "Absorb Time", Description: ""},
					{Id: CC_Config_absorb_End_A, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "A", Mandatory: true, Label: "Absorb End", Description: ""},
					{Id: CC_Config_rebulk_V, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Rebulk", Description: ""},
					{Id: CC_Config_float_V, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Float", Description: ""},
					{Id: CC_Config_max_Chg_A, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "A", Mandatory: true, Label: "Maximum Charge", Description: ""},
					{Id: CC_Config_equalize_V, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Equalize", Description: ""},
					{Id: CC_Config_equalize_Hr, Offset: 15, Type: typelabel.Uint16, Units: "Tmh", Mandatory: true, Label: "Equalize Time", Description: ""},
					{Id: CC_Config_auto_equalize, Offset: 16, Type: typelabel.Uint16, Units: "Tmd", Mandatory: true, Label: "Auto Equalize Interval", Description: ""},
					{Id: CC_Config_MPPT_mode, Offset: 17, Type: typelabel.Enum16, Mandatory: true, Label: "MPPT mode", Description: ""},
					{Id: CC_Config_sweep_width, Offset: 18, Type: typelabel.Enum16, Mandatory: true, Label: "Sweep Width", Description: ""},
					{Id: CC_Config_sweep_max, Offset: 19, Type: typelabel.Enum16, Mandatory: true, Label: "Sweep Maximum", Description: ""},
					{Id: CC_Config_U_Pick_Duty_cyc, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "Pct", Mandatory: true, Label: "U-Pick PWM Duty Cycle", Description: ""},
					{Id: CC_Config_grid_tie, Offset: 21, Type: typelabel.Enum16, Mandatory: true, Label: "Grid Tie Mode", Description: ""},
					{Id: CC_Config_temp_comp, Offset: 22, Type: typelabel.Enum16, Mandatory: true, Label: "Temp Comp Mode", Description: ""},
					{Id: CC_Config_temp_comp_llimt, Offset: 23, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Temp Comp Lower Limit", Description: ""},
					{Id: CC_Config_temp_comp_hlimt, Offset: 24, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Temp Comp Upper Limit", Description: ""},
					{Id: CC_Config_auto_restart, Offset: 25, Type: typelabel.Enum16, Mandatory: true, Label: "Auto Restart Mode", Description: ""},
					{Id: CC_Config_wakeup_VOC, Offset: 26, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Wakeup VOC Change", Description: ""},
					{Id: CC_Config_snooze_mode_A, Offset: 27, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "A", Mandatory: true, Label: "Snooze Mode", Description: ""},
					{Id: CC_Config_wakeup_interval, Offset: 28, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "Wakeup Interval", Description: ""},
					{Id: CC_Config_AUX_mode, Offset: 29, Type: typelabel.Enum16, Mandatory: true, Label: "AUX Output Mode", Description: ""},
					{Id: CC_Config_AUX_control, Offset: 30, Type: typelabel.Enum16, Mandatory: true, Label: "AUX Output Control", Description: ""},
					{Id: CC_Config_AUX_state, Offset: 31, Type: typelabel.Enum16, Mandatory: true, Label: "AUX Output State", Description: ""},
					{Id: CC_Config_AUX_polarity, Offset: 32, Type: typelabel.Enum16, Mandatory: true, Label: "AUX Output Polarity", Description: ""},
					{Id: CC_Config_AUX_L_Batt_disc, Offset: 33, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Low Battery Disconnect", Description: ""},
					{Id: CC_Config_AUX_L_Batt_rcon, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Low Battery Reconnect", Description: ""},
					{Id: CC_Config_AUX_L_Batt_dly, Offset: 35, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "AUX Low Battery Disconnect Delay", Description: ""},
					{Id: CC_Config_AUX_Vent_fan_V, Offset: 36, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Vent Fan", Description: ""},
					{Id: CC_Config_AUX_PV_triggerV, Offset: 37, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX PV Trigger", Description: ""},
					{Id: CC_Config_AUX_PV_trg_h_tm, Offset: 38, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "AUX PV Trigger Hold Time", Description: ""},
					{Id: CC_Config_AUX_Nlite_ThrsV, Offset: 39, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Night Light Threshold", Description: ""},
					{Id: CC_Config_AUX_Nlite_On_tm, Offset: 40, Type: typelabel.Uint16, ScaleFactor: "H_SF", Units: "Tmh", Mandatory: true, Label: "AUX Night Light On Time", Description: ""},
					{Id: CC_Config_AUX_Nlite_On_hist, Offset: 41, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "AUX Night Light On Hysteresis", Description: ""},
					{Id: CC_Config_AUX_Nlite_Off_hist, Offset: 42, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "AUX Night Light Off Hysteresis", Description: ""},
					{Id: CC_Config_AUX_Error_batt_V, Offset: 43, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Error Output Low Battery", Description: ""},
					{Id: CC_Config_AUX_Divert_h_time, Offset: 44, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "Tms", Mandatory: true, Label: "AUX Divert Hold Time", Description: ""},
					{Id: CC_Config_AUX_Divert_dly_time, Offset: 45, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "AUX Divert Delay Time", Description: ""},
					{Id: CC_Config_AUX_Divert_Rel_V, Offset: 46, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Divert Relative", Description: ""},
					{Id: CC_Config_AUX_Divert_Hyst_V, Offset: 47, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "AUX Divert Hysteresis", Description: ""},
					{Id: CC_Config_MajorFWRev, Offset: 48, Type: typelabel.Uint16, Mandatory: true, Label: "FM CC Major Firmware Number", Description: ""},
					{Id: CC_Config_MidFWRev, Offset: 49, Type: typelabel.Uint16, Mandatory: true, Label: "FM CC Mid Firmware Number", Description: ""},
					{Id: CC_Config_MinorFWRev, Offset: 50, Type: typelabel.Uint16, Mandatory: true, Label: "FM CC Minor Firmware Number", Description: ""},
					{Id: CC_Config_DataLog_Day_offset, Offset: 51, Type: typelabel.Uint16, Units: "Tmd", Mandatory: true, Label: "Set Data Log Day Offset", Description: ""},
					{Id: CC_Config_DataLog_Cur_Day_off, Offset: 52, Type: typelabel.Uint16, Units: "Tmd", Mandatory: true, Label: "Current Data Log Day Offset", Description: ""},
					{Id: CC_Config_DataLog_Daily_AH, Offset: 53, Type: typelabel.Uint16, Units: "Ah", Mandatory: true, Label: "Data Log Daily (Ah)", Description: ""},
					{Id: CC_Config_DataLog_Daily_KWH, Offset: 54, Type: typelabel.Uint16, ScaleFactor: "KWH_SF", Units: "kWh", Mandatory: true, Label: "Data Log Daily (kWh)", Description: ""},
					{Id: CC_Config_DataLog_Max_Out_A, Offset: 55, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "A", Mandatory: true, Label: "Data Log Daily Maximum Output (A)", Description: ""},
					{Id: CC_Config_DataLog_Max_Out_W, Offset: 56, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "W", Mandatory: true, Label: "Data Log Daily Maximum Output (W)", Description: ""},
					{Id: CC_Config_DataLog_Absorb_T, Offset: 57, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "Data Log Daily Absorb Time", Description: ""},
					{Id: CC_Config_DataLog_Float_T, Offset: 58, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "Data Log Daily Float Time", Description: ""},
					{Id: CC_Config_DataLog_Min_Batt_V, Offset: 59, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Data Log Daily Minimum Battery", Description: ""},
					{Id: CC_Config_DataLog_Max_Batt_V, Offset: 60, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Data Log Daily Maximum Battery", Description: ""},
					{Id: CC_Config_DataLog_Max_Input_V, Offset: 61, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Data Log Daily Maximum Input", Description: ""},
					{Id: CC_Config_DataLog_Clear, Offset: 62, Type: typelabel.Uint16, Mandatory: true, Label: "Data Log Clear", Description: ""},
					{Id: CC_Config_DataLog_Clr_Comp, Offset: 63, Type: typelabel.Uint16, Mandatory: true, Label: "Data Log Clear Complement", Description: ""},
				},
			},
		}})
}
