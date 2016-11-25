// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model305

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block305 - GPS - Include to support location measurements

const (
	ModelID = 305
)

const (
	Alt  = "Alt"
	Date = "Date"
	Lat  = "Lat"
	Loc  = "Loc"
	Long = "Long"
	Tm   = "Tm"
)

type Block305 struct {
	Tm   string `sunspec:"offset=0,len=6"`
	Date string `sunspec:"offset=6,len=4"`
	Loc  string `sunspec:"offset=10,len=20"`
	Lat  int32  `sunspec:"offset=30,sf=-7"`
	Long int32  `sunspec:"offset=32,sf=-7"`
	Alt  int32  `sunspec:"offset=34"`
}

func (self *Block305) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "location",
		Length: 36,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 36,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Tm, Offset: 0, Type: typelabel.String, Units: "hhmmss.sssZ", Length: 6},
					smdx.PointElement{Id: Date, Offset: 6, Type: typelabel.String, Units: "YYYYMMDD", Length: 4},
					smdx.PointElement{Id: Loc, Offset: 10, Type: typelabel.String, Units: "text", Length: 20},
					smdx.PointElement{Id: Lat, Offset: 30, Type: typelabel.Int32, ScaleFactor: "-7", Units: "Degrees"},
					smdx.PointElement{Id: Long, Offset: 32, Type: typelabel.Int32, ScaleFactor: "-7", Units: "Degrees"},
					smdx.PointElement{Id: Alt, Offset: 34, Type: typelabel.Int32, Units: "meters"},
				},
			},
		}})
}
