package model6

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block6 - Secure Write Sequential Request - Include a digital signature along with the control data

const (
	ModelID          = 6
	ModelLabel       = "Secure Write Sequential Request"
	ModelDescription = "Include a digital signature along with the control data"
)

const (
	Alg   = "Alg"
	DS    = "DS"
	Ms    = "Ms"
	N     = "N"
	Off   = "Off"
	Role  = "Role"
	Rsrvd = "Rsrvd"
	Seq   = "Seq"
	Ts    = "Ts"
	Val1  = "Val1"
	Val10 = "Val10"
	Val11 = "Val11"
	Val12 = "Val12"
	Val13 = "Val13"
	Val14 = "Val14"
	Val15 = "Val15"
	Val16 = "Val16"
	Val17 = "Val17"
	Val18 = "Val18"
	Val19 = "Val19"
	Val2  = "Val2"
	Val20 = "Val20"
	Val21 = "Val21"
	Val22 = "Val22"
	Val23 = "Val23"
	Val24 = "Val24"
	Val25 = "Val25"
	Val26 = "Val26"
	Val27 = "Val27"
	Val28 = "Val28"
	Val29 = "Val29"
	Val3  = "Val3"
	Val30 = "Val30"
	Val31 = "Val31"
	Val32 = "Val32"
	Val33 = "Val33"
	Val34 = "Val34"
	Val35 = "Val35"
	Val36 = "Val36"
	Val37 = "Val37"
	Val38 = "Val38"
	Val39 = "Val39"
	Val4  = "Val4"
	Val40 = "Val40"
	Val41 = "Val41"
	Val42 = "Val42"
	Val43 = "Val43"
	Val44 = "Val44"
	Val45 = "Val45"
	Val46 = "Val46"
	Val47 = "Val47"
	Val48 = "Val48"
	Val49 = "Val49"
	Val5  = "Val5"
	Val50 = "Val50"
	Val51 = "Val51"
	Val52 = "Val52"
	Val53 = "Val53"
	Val54 = "Val54"
	Val55 = "Val55"
	Val56 = "Val56"
	Val57 = "Val57"
	Val58 = "Val58"
	Val59 = "Val59"
	Val6  = "Val6"
	Val60 = "Val60"
	Val61 = "Val61"
	Val62 = "Val62"
	Val63 = "Val63"
	Val64 = "Val64"
	Val65 = "Val65"
	Val66 = "Val66"
	Val67 = "Val67"
	Val68 = "Val68"
	Val69 = "Val69"
	Val7  = "Val7"
	Val70 = "Val70"
	Val71 = "Val71"
	Val72 = "Val72"
	Val73 = "Val73"
	Val74 = "Val74"
	Val75 = "Val75"
	Val76 = "Val76"
	Val77 = "Val77"
	Val78 = "Val78"
	Val79 = "Val79"
	Val8  = "Val8"
	Val80 = "Val80"
	Val9  = "Val9"
	X     = "X"
)

type Block6Repeat struct {
	DS uint16 `sunspec:"offset=0,access=rw"`
}

type Block6 struct {
	X     uint16         `sunspec:"offset=0,access=rw"`
	Off   uint16         `sunspec:"offset=1,access=rw"`
	Val1  uint16         `sunspec:"offset=2,access=rw"`
	Val2  uint16         `sunspec:"offset=3,access=rw"`
	Val3  uint16         `sunspec:"offset=4,access=rw"`
	Val4  uint16         `sunspec:"offset=5,access=rw"`
	Val5  uint16         `sunspec:"offset=6,access=rw"`
	Val6  uint16         `sunspec:"offset=7,access=rw"`
	Val7  uint16         `sunspec:"offset=8,access=rw"`
	Val8  uint16         `sunspec:"offset=9,access=rw"`
	Val9  uint16         `sunspec:"offset=10,access=rw"`
	Val10 uint16         `sunspec:"offset=11,access=rw"`
	Val11 uint16         `sunspec:"offset=12,access=rw"`
	Val12 uint16         `sunspec:"offset=13,access=rw"`
	Val13 uint16         `sunspec:"offset=14,access=rw"`
	Val14 uint16         `sunspec:"offset=15,access=rw"`
	Val15 uint16         `sunspec:"offset=16,access=rw"`
	Val16 uint16         `sunspec:"offset=17,access=rw"`
	Val17 uint16         `sunspec:"offset=18,access=rw"`
	Val18 uint16         `sunspec:"offset=19,access=rw"`
	Val19 uint16         `sunspec:"offset=20,access=rw"`
	Val20 uint16         `sunspec:"offset=21,access=rw"`
	Val21 uint16         `sunspec:"offset=22,access=rw"`
	Val22 uint16         `sunspec:"offset=23,access=rw"`
	Val23 uint16         `sunspec:"offset=24,access=rw"`
	Val24 uint16         `sunspec:"offset=25,access=rw"`
	Val25 uint16         `sunspec:"offset=26,access=rw"`
	Val26 uint16         `sunspec:"offset=27,access=rw"`
	Val27 uint16         `sunspec:"offset=28,access=rw"`
	Val28 uint16         `sunspec:"offset=29,access=rw"`
	Val29 uint16         `sunspec:"offset=30,access=rw"`
	Val30 uint16         `sunspec:"offset=31,access=rw"`
	Val31 uint16         `sunspec:"offset=32,access=rw"`
	Val32 uint16         `sunspec:"offset=33,access=rw"`
	Val33 uint16         `sunspec:"offset=34,access=rw"`
	Val34 uint16         `sunspec:"offset=35,access=rw"`
	Val35 uint16         `sunspec:"offset=36,access=rw"`
	Val36 uint16         `sunspec:"offset=37,access=rw"`
	Val37 uint16         `sunspec:"offset=38,access=rw"`
	Val38 uint16         `sunspec:"offset=39,access=rw"`
	Val39 uint16         `sunspec:"offset=40,access=rw"`
	Val40 uint16         `sunspec:"offset=41,access=rw"`
	Val41 uint16         `sunspec:"offset=42,access=rw"`
	Val42 uint16         `sunspec:"offset=43,access=rw"`
	Val43 uint16         `sunspec:"offset=44,access=rw"`
	Val44 uint16         `sunspec:"offset=45,access=rw"`
	Val45 uint16         `sunspec:"offset=46,access=rw"`
	Val46 uint16         `sunspec:"offset=47,access=rw"`
	Val47 uint16         `sunspec:"offset=48,access=rw"`
	Val48 uint16         `sunspec:"offset=49,access=rw"`
	Val49 uint16         `sunspec:"offset=50,access=rw"`
	Val50 uint16         `sunspec:"offset=51,access=rw"`
	Val51 uint16         `sunspec:"offset=52,access=rw"`
	Val52 uint16         `sunspec:"offset=53,access=rw"`
	Val53 uint16         `sunspec:"offset=54,access=rw"`
	Val54 uint16         `sunspec:"offset=55,access=rw"`
	Val55 uint16         `sunspec:"offset=56,access=rw"`
	Val56 uint16         `sunspec:"offset=57,access=rw"`
	Val57 uint16         `sunspec:"offset=58,access=rw"`
	Val58 uint16         `sunspec:"offset=59,access=rw"`
	Val59 uint16         `sunspec:"offset=60,access=rw"`
	Val60 uint16         `sunspec:"offset=61,access=rw"`
	Val61 uint16         `sunspec:"offset=62,access=rw"`
	Val62 uint16         `sunspec:"offset=63,access=rw"`
	Val63 uint16         `sunspec:"offset=64,access=rw"`
	Val64 uint16         `sunspec:"offset=65,access=rw"`
	Val65 uint16         `sunspec:"offset=66,access=rw"`
	Val66 uint16         `sunspec:"offset=67,access=rw"`
	Val67 uint16         `sunspec:"offset=68,access=rw"`
	Val68 uint16         `sunspec:"offset=69,access=rw"`
	Val69 uint16         `sunspec:"offset=70,access=rw"`
	Val70 uint16         `sunspec:"offset=71,access=rw"`
	Val71 uint16         `sunspec:"offset=72,access=rw"`
	Val72 uint16         `sunspec:"offset=73,access=rw"`
	Val73 uint16         `sunspec:"offset=74,access=rw"`
	Val74 uint16         `sunspec:"offset=75,access=rw"`
	Val75 uint16         `sunspec:"offset=76,access=rw"`
	Val76 uint16         `sunspec:"offset=77,access=rw"`
	Val77 uint16         `sunspec:"offset=78,access=rw"`
	Val78 uint16         `sunspec:"offset=79,access=rw"`
	Val79 uint16         `sunspec:"offset=80,access=rw"`
	Val80 uint16         `sunspec:"offset=81,access=rw"`
	Ts    uint32         `sunspec:"offset=82,access=rw"`
	Ms    uint16         `sunspec:"offset=84,access=rw"`
	Seq   uint16         `sunspec:"offset=85,access=rw"`
	Role  uint16         `sunspec:"offset=86,access=rw"`
	Rsrvd sunspec.Pad    `sunspec:"offset=87,access=rw"`
	Alg   sunspec.Enum16 `sunspec:"offset=88,access=rw"`
	N     uint16         `sunspec:"offset=89,access=rw"`

	Repeats []Block6Repeat
}

func (block *Block6) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 91,
		Blocks: []smdx.BlockElement{
			{
				Length: 90,
				Points: []smdx.PointElement{
					{Id: X, Offset: 0, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "X", Description: "Number of (offset, value) pairs being written"},
					{Id: Off, Offset: 1, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Offset", Description: "Starting offset for write operation"},
					{Id: Val1, Offset: 2, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Value1", Description: "Value to write to control register at offset"},
					{Id: Val2, Offset: 3, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val3, Offset: 4, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val4, Offset: 5, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val5, Offset: 6, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val6, Offset: 7, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val7, Offset: 8, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val8, Offset: 9, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val9, Offset: 10, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val10, Offset: 11, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val11, Offset: 12, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val12, Offset: 13, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val13, Offset: 14, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val14, Offset: 15, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val15, Offset: 16, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val16, Offset: 17, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val17, Offset: 18, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val18, Offset: 19, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val19, Offset: 20, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val20, Offset: 21, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val21, Offset: 22, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val22, Offset: 23, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val23, Offset: 24, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val24, Offset: 25, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val25, Offset: 26, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val26, Offset: 27, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val27, Offset: 28, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val28, Offset: 29, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val29, Offset: 30, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val30, Offset: 31, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val31, Offset: 32, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val32, Offset: 33, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val33, Offset: 34, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val34, Offset: 35, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val35, Offset: 36, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val36, Offset: 37, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val37, Offset: 38, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val38, Offset: 39, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val39, Offset: 40, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val40, Offset: 41, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val41, Offset: 42, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val42, Offset: 43, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val43, Offset: 44, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val44, Offset: 45, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val45, Offset: 46, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val46, Offset: 47, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val47, Offset: 48, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val48, Offset: 49, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val49, Offset: 50, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val50, Offset: 51, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val51, Offset: 52, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val52, Offset: 53, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val53, Offset: 54, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val54, Offset: 55, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val55, Offset: 56, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val56, Offset: 57, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val57, Offset: 58, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val58, Offset: 59, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val59, Offset: 60, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val60, Offset: 61, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val61, Offset: 62, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val62, Offset: 63, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val63, Offset: 64, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val64, Offset: 65, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val65, Offset: 66, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val66, Offset: 67, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val67, Offset: 68, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val68, Offset: 69, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val69, Offset: 70, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val70, Offset: 71, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val71, Offset: 72, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val72, Offset: 73, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val73, Offset: 74, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val74, Offset: 75, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val75, Offset: 76, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val76, Offset: 77, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val77, Offset: 78, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val78, Offset: 79, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val79, Offset: 80, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Val80, Offset: 81, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					{Id: Ts, Offset: 82, Type: typelabel.Uint32, Access: "rw", Mandatory: true, Label: "Timestamp", Description: "Timestamp value is the number of seconds since January 1, 2000"},
					{Id: Ms, Offset: 84, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Milliseconds", Description: "Millisecond counter 0-999"},
					{Id: Seq, Offset: 85, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Sequence", Description: "Sequence number of request"},
					{Id: Role, Offset: 86, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Role", Description: "Signing key used 0-5"},
					{Id: Rsrvd, Offset: 87, Type: typelabel.Pad, Access: "rw", Mandatory: true},
					{Id: Alg, Offset: 88, Type: typelabel.Enum16, Access: "rw", Mandatory: true, Label: "Algorithm", Description: "Algorithm used to compute the digital signature"},
					{Id: N, Offset: 89, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "N", Description: "Number of registers comprising the digital signature."},
				},
			},
			{
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: DS, Offset: 0, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "DS", Description: "Digital Signature"},
				},
			},
		}})
}
