package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	prose "gopkg.in/jdkato/prose.v2"
)

type trainResult struct {
	Label string
}

type labeledEntity struct {
	Text   string
	Entity string
	Label  string
}

func readFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return data
}

func mapLabeled(jsonl []byte) []labeledEntity {
	decoder := json.NewDecoder(bytes.NewReader(jsonl))
	entries := []labeledEntity{}

	for {
		entity := labeledEntity{}
		err := decoder.Decode(&entity)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		entries = append(entries, entity)
	}

	return entries
}

func mapSpans(entity labeledEntity) []prose.LabeledEntity {
	name := entity.Entity
	start := strings.Index(entity.Text, name)
	end := start + len(name)

	spans := []prose.LabeledEntity{{
		Start: start,
		End:   end,
		Label: entity.Label}}

	return spans
}

func mapContext(labeled []labeledEntity) (string, []prose.EntityContext) {
	label := ""
	if len(labeled) > 0 {
		label = labeled[0].Label
	}

	entities := []prose.EntityContext{}
	for _, entity := range labeled {
		entities = append(entities, prose.EntityContext{
			Text:   entity.Text,
			Spans:  mapSpans(entity),
			Accept: true})
	}

	return label, entities
}

func saveModel(label string, model *prose.Model) error {
	modelDir := filepath.Join(".", "models")
	modelPath := filepath.Join(modelDir, label)

	os.MkdirAll(modelPath, os.ModePerm)
	os.RemoveAll(filepath.Join(modelPath, "/Maxent"))

	return model.Write(modelPath)
}

func train(trainFileName string) (trainResult, error) {
	label, train := mapContext(mapLabeled(readFile(trainFileName)))
	model := prose.ModelFromData(label, prose.UsingEntities(train))

	err := saveModel(label, model)
	if err != nil {
		return trainResult{}, err
	}

	return trainResult{Label: label}, nil
}
