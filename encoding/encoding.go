package encoding

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/skosovsky/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML.
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON.
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData.
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML.
func (j *JSONData) Encoding() error { //nolint:dupl // for example
	fileInput, err := os.ReadFile(j.FileInput)
	if err != nil {
		err = fmt.Errorf("filed to read file: %w", err)
		log.Println(err)
		return err
	}

	var jsonData models.DockerCompose

	err = json.Unmarshal(fileInput, &jsonData)
	if err != nil {
		err = fmt.Errorf("filed to unmarshal: %w", err)
		log.Println(err)
		return err
	}

	yamlData, err := yaml.Marshal(jsonData)
	if err != nil {
		err = fmt.Errorf("filed to marshal: %w", err)
		log.Println(err)
		return err
	}

	fileOutput, err := os.Create(j.FileOutput)
	if err != nil {
		err = fmt.Errorf("filed to create file: %w", err)
		log.Println(err)
		return err
	}
	defer fileOutput.Close()

	_, err = fileOutput.Write(yamlData)
	if err != nil {
		err = fmt.Errorf("filed to write file: %w", err)
		log.Println(err)
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON.
func (y *YAMLData) Encoding() error { //nolint:dupl // for example
	fileInput, err := os.ReadFile(y.FileInput)
	if err != nil {
		err = fmt.Errorf("filed to read file: %w", err)
		log.Println(err)
		return err
	}

	var yamlData models.DockerCompose

	err = yaml.Unmarshal(fileInput, &yamlData)
	if err != nil {
		err = fmt.Errorf("filed to unmarshal: %w", err)
		log.Println(err)
		return err
	}

	jsonData, err := json.Marshal(yamlData)
	if err != nil {
		err = fmt.Errorf("filed to marshal: %w", err)
		log.Println(err)
		return err
	}

	fileOutput, err := os.Create(y.FileOutput)
	if err != nil {
		err = fmt.Errorf("filed to create file: %w", err)
		log.Println(err)
		return err
	}
	defer fileOutput.Close()

	_, err = fileOutput.Write(jsonData)
	if err != nil {
		err = fmt.Errorf("filed to write file: %w", err)
		log.Println(err)
		return err
	}

	return nil
}
