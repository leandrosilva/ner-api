package main

import (
	"path/filepath"

	prose "gopkg.in/jdkato/prose.v2"
)

type recognizeResult struct {
	Entities []entityCount
}

type entityCount struct {
	Entity prose.Entity `json:entity`
	Count  int          `json:count`
}

func recognize(text string, modelName string) (recognizeResult, error) {
	var result recognizeResult

	doc, err := getDocument(text, modelName)
	if err != nil {
		return result, err
	}

	if len(doc.Entities()) == 0 {
		return result, nil
	}
	result = recognizeResult{Entities: distinctEntities(doc.Entities())}

	return result, nil
}

func getDocument(text string, modelName string) (*prose.Document, error) {
	modelPath := filepath.Join(".", "models", modelName)
	model := prose.ModelFromDisk(modelPath)

	doc, err := prose.NewDocument(text, prose.UsingModel(model))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func distinctEntities(entities []prose.Entity) []entityCount {
	counter := map[string]entityCount{}
	for _, entity := range entities {
		value, found := counter[entity.Text]
		if found {
			value.Count = value.Count + 1
			counter[entity.Text] = value
		} else {
			counter[entity.Text] = entityCount{Entity: entity, Count: 1}
		}
	}

	distinct := []entityCount{}
	for _, value := range counter {
		distinct = append(distinct, value)
	}

	return distinct
}
