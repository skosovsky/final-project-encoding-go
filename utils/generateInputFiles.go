package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/skosovsky/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

func CreateJSONFile() { //nolint:dupl // for example
	jsonFile, err := os.Create("jsonInput.json")
	if err != nil {
		err = fmt.Errorf("filed to create file: %w", err)
		log.Println(err)
	}

	defer jsonFile.Close()

	ports := []string{"5000:5000"}
	volumes := []string{"/usercode/:/code"}
	links := []string{"database:backenddb"}

	web := models.Web{
		Build:   ".",
		Ports:   ports,
		Volumes: volumes,
		Links:   links,
	}

	environment := []string{
		"MYSQL_ROOT_PASSWORD=root",
		"MYSQL_USER=testuser",
		"MYSQL_PASSWORD=admin123",
		"MYSQL_DATABASE=backend",
	}

	volumes = []string{"/usercode/db/init.sql:/docker-entrypoint-initdb.d/init.sql"}

	database := models.Database{
		Image:       "mysql/mysql-server:5.7",
		Environment: environment,
		Volumes:     volumes,
	}

	services := models.Services{
		Web:      web,
		Database: database,
	}

	dockerCompose := models.DockerCompose{
		Version:  "3",
		Services: services,
	}

	out, err := json.Marshal(&dockerCompose)
	if err != nil {
		err = fmt.Errorf("filed to json encoding: %w", err)
		log.Println(err)
	}

	_, err = jsonFile.Write(out)
	if err != nil {
		err = fmt.Errorf("filed to write file: %w", err)
		log.Println(err)
	}
}

func CreateYAMLFile() { //nolint:dupl // for example
	yamlFile, err := os.Create("yamlInput.yml")
	if err != nil {
		err = fmt.Errorf("filed to create file: %w", err)
		log.Println(err)
	}

	defer yamlFile.Close()

	ports := []string{"5000:5000"}
	volumes := []string{"/usercode/:/code"}
	links := []string{"database:backenddb"}

	web := models.Web{
		Build:   ".",
		Ports:   ports,
		Volumes: volumes,
		Links:   links,
	}

	environment := []string{
		"MYSQL_ROOT_PASSWORD=root",
		"MYSQL_USER=testuser",
		"MYSQL_PASSWORD=admin123",
		"MYSQL_DATABASE=backend",
	}

	volumes = []string{"/usercode/db/init.sql:/docker-entrypoint-initdb.d/init.sql"}

	database := models.Database{
		Image:       "mysql/mysql-server:5.7",
		Environment: environment,
		Volumes:     volumes,
	}

	services := models.Services{
		Web:      web,
		Database: database,
	}

	dockerCompose := models.DockerCompose{
		Version:  "3",
		Services: services,
	}

	out, err := yaml.Marshal(&dockerCompose)
	if err != nil {
		err = fmt.Errorf("filed to yaml encoding: %w", err)
		log.Println(err)
	}

	_, err = yamlFile.Write(out)
	if err != nil {
		err = fmt.Errorf("filed to write file: %w", err)
		log.Println(err)
	}
}
