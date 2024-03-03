package main

import (
	"fmt"
	"log"

	"github.com/skosovsky/final-project-encoding-go/encoding"
	"github.com/skosovsky/final-project-encoding-go/utils"
)

func Encode(data encoding.MyEncoder) error {
	err := data.Encoding()
	if err != nil {
		err = fmt.Errorf("filed to encoding: %w", data.Encoding())
		log.Println(err)
	}

	return err
}

func main() {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	jsonData := encoding.JSONData{
		DockerCompose: nil,
		FileInput:     "jsonInput.json",
		FileOutput:    "yamlOutput.yml",
	}
	err := Encode(&jsonData)
	if err != nil {
		err = fmt.Errorf("filed to encoding JSON to YAML: %w", err)
		log.Println(err)
	}

	yamlData := encoding.YAMLData{
		DockerCompose: nil,
		FileInput:     "yamlInput.yml",
		FileOutput:    "jsonOutput.json",
	}
	err = Encode(&yamlData)
	if err != nil {
		err = fmt.Errorf("filed to encoding YAML to JSON: %w", err)
		log.Println(err)
	}
}
