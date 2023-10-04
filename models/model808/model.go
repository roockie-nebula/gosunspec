package model808

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block808 - Flow Battery Module Model -

const (
	ModelID          = 808
	ModelLabel       = "Flow Battery Module Model"
	ModelDescription = ""
)

const (
	ModuleTBD = "ModuleTBD"
	StackTBD  = "StackTBD"
)

type Block808Repeat struct {
	StackTBD uint16 `sunspec:"offset=0"`
}

type Block808 struct {
	ModuleTBD uint16 `sunspec:"offset=0"`

	Repeats []Block808Repeat
}

func (block *Block808) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "flow_battery_module",
		Length: 2,
		Blocks: []smdx.BlockElement{
			{
				Length: 1,
				Points: []smdx.PointElement{
					{Id: ModuleTBD, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Module Points To Be Determined", Description: ""},
				},
			},
			{Name: "stack",
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: StackTBD, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Stack Points To Be Determined", Description: ""},
				},
			},
		}})
}
