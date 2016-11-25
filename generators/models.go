package main

import (
	"fmt"
	"github.com/crabmusket/gosunspec/smdx"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
)

const file = `
{{define "point"}}{{.Id}} {{.Type|goType}} ` + "`" + `sunspec:"offset={{.Offset}}{{if gt .Length 0}},len={{.Length}}{{end}}{{if ne .ScaleFactor ""}},sf={{.ScaleFactor}}{{end}}{{if ne .Access ""}},access={{.Access}}{{end}}"` + "`" + `{{end}}
{{define "model"}}
{{$model := index .Models 0 }}
{{$strings := index .Strings 0 }}

// Block{{$model.Id}} - {{$strings.ModelStrings.Label}} - {{$strings.ModelStrings.Description}}

const (
	ModelID = {{$model.Id}}
)

const (
{{range (uniqueNames $model)}}   {{.|title}} = "{{.}}"
{{end}}
)

{{range (repeatingBlocks $model)}}
type Block{{$model.Id}}Repeat struct {
	{{range .Points}}{{template "point" .}}
{{end}}
}
{{end}}

type Block{{$model.Id}} struct {
{{range (fixedBlocks $model)}}
{{range .Points}}    {{template "point" .}}
{{end}}
	{{end}}
{{range (repeatingBlocks $model)}}
  Repeats []Block{{$model.Id}}Repeat
{{end}}
}

func (self *Block{{$model.Id}}) GetId() sunspec.ModelId {
  return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
	Id: ModelID,
	Name: "{{$model.Name}}",
	Length: {{$model.Length}},
	Blocks: []smdx.BlockElement{ {{range $model.Blocks}}
			smdx.BlockElement{ {{if gt (len .Name) 0 }}Name: "{{.Name}}",{{end}}
				Length: {{.Length}},
				{{if gt (len .Type) 0 }}Type: "{{.Type}}",{{end}}
				Points: []smdx.PointElement{ {{range .Points}}
					smdx.PointElement{Id: {{.Id|title}},Offset: {{.Offset}},Type: typelabel.{{.Type|title}}{{optF "ScaleFactor" .ScaleFactor}}{{optF "Units" .Units}}{{optF "Access" .Access}}{{if gt (.Length) 0}},Length: {{.Length}}{{end}}{{if .Mandatory}},Mandatory: {{.Mandatory}}{{end}},}, {{end}}
				},
			}, {{end}}
	}})
}
{{end}}
// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model{{(index .Model.Models 0).Id}}

import (
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/typelabel"
)
{{template "model" .Model}}
`

const loader = `
// NOTICE
// This file was automatically generated by ../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package models

import (
{{range .Names}}   _ "github.com/crabmusket/gosunspec/models/{{.}}"
{{end}}
)
`

func main() {

	smdxDir := "../spec/smdx/"
	files, err := ioutil.ReadDir(smdxDir)
	if err != nil {
		log.Fatal(err)
	}

	models := []smdx.ModelDefinitionElement{}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "smdx_") {
			smdxFile, err := os.OpenFile(smdxDir+file.Name(), os.O_RDONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			def, err := smdx.FromXML(smdxFile)
			if err != nil {
				smdxFile.Close()
				log.Fatal(err)
			}
			smdxFile.Close()

			if len(def.Models) < 1 {
				log.Printf("failed to find any models in %s", file.Name())
				continue
			} else {
				if len(def.Models) > 1 {
					log.Printf("found additional models in %s", file.Name())
				}
				models = append(models, def)
			}
		}
	}

	t := template.New("file")
	t.Funcs(map[string]interface{}{
		"goType": func(sstype string) string {
			switch sstype {
			case "uint16", "uint32", "uint64", "int16", "float32", "int32", "int64", "string":
				return sstype
			case "sunssf":
				return "sunspec.ScaleFactor"
			case "eui48":
				return "sunspec." + strings.ToUpper(sstype)
			default:
				return "sunspec." + strings.Title(sstype)
			}
		},
		"repeatingBlocks": func(m smdx.ModelElement) []smdx.BlockElement {
			if len(m.Blocks) == 1 && m.Blocks[0].Type == "repeating" {
				return m.Blocks
			} else if len(m.Blocks) > 1 {
				return m.Blocks[1:]
			} else {
				return []smdx.BlockElement{}
			}
		},
		"fixedBlocks": func(m smdx.ModelElement) []smdx.BlockElement {
			if len(m.Blocks) > 0 && m.Blocks[0].Type != "repeating" {
				return m.Blocks[0:1]
			} else {
				return []smdx.BlockElement{}
			}
		},
		"optF": func(n, v string) string {
			if len(v) > 0 {
				return "," + n + ": \"" + v + "\""
			} else {
				return ""
			}
		},
		"title": strings.Title,
		"uniqueNames": func(m smdx.ModelElement) []string {
			seen := map[string]bool{}
			for _, b := range m.Blocks {
				for _, p := range b.Points {
					seen[p.Id] = true
				}
			}
			result := []string{}
			for n, _ := range seen {
				result = append(result, n)
			}
			sort.Strings(result)
			return result
		},
	})
	modelTemplate := template.Must(t.Parse(file))

	// write out all the model files

	for _, m := range models {
		dir := fmt.Sprintf("../models/model%d", m.Models[0].Id)
		if err := os.MkdirAll(dir, 0777); err != nil {
			log.Fatalf("%v", err)
		}
		outputFilename := dir + "/model.go.tmp"
		outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if err := modelTemplate.Execute(outputFile, map[string]interface{}{
			"Model": m,
		}); err != nil {
			log.Fatalf("template execution failed: %v", err)
		}
		outputFile.Close()
		cmd := exec.Command("/bin/sh", "-c", "gofmt -w "+outputFilename)
		if err := cmd.Run(); err != nil {
			log.Fatalf("gofmt failed: %v", err)
		}
		if err := os.Rename(outputFilename, dir+"/model.go"); err != nil {
			log.Fatalf("replacing models.go failed: %v", err)
		}
	}

	// write out a loader module

	files, err = ioutil.ReadDir("../models/")
	if err != nil {
		log.Fatal(err)
	}
	names := []string{}
	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), "model") {
			names = append(names, f.Name())
		}
	}
	outputFile, err := os.OpenFile("../models/loader.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	loaderTemplate := template.Must(t.Parse(loader))
	loaderTemplate.Execute(outputFile, map[string]interface{}{
		"Names": names,
	})
	outputFile.Close()
	cmd := exec.Command("/bin/sh", "-c", "gofmt -w ../models/loader.go")
	if err := cmd.Run(); err != nil {
		log.Fatalf("gofmt failed: %v", err)
	}

}
