package model801

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/roockie-nebula/gosunspec"
	"github.com/roockie-nebula/gosunspec/smdx"
	"github.com/roockie-nebula/gosunspec/typelabel"
)

// Block801 - Energy Storage Base Model (DEPRECATED) - This model has been deprecated.

const (
	ModelID          = 801
	ModelLabel       = "Energy Storage Base Model (DEPRECATED)"
	ModelDescription = "This model has been deprecated."
)

const (
	DEPRECATED = "DEPRECATED"
)

type Block801 struct {
	DEPRECATED sunspec.Enum16 `sunspec:"offset=0"`
}

func (block *Block801) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "storage",
		Length: 1,
		Blocks: []smdx.BlockElement{
			{
				Length: 1,
				Points: []smdx.PointElement{
					{Id: DEPRECATED, Offset: 0, Type: typelabel.Enum16, Mandatory: true, Label: "Deprecated Model", Description: "This model has been deprecated."},
				},
			},
		}})
}
