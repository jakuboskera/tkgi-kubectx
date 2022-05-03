package tkgi

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Credentials struct {
	Credentials []struct {
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		ClusterAdmin bool   `yaml:"clusterAdmin"`
	} `yaml:"credentials"`
}

// get method returns content of credentials.yaml file
func (c *Credentials) get() *Credentials {
	path, err := getTkgiKubectxFile("credentials.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fileExists, _ := exists(path)

	if fileExists {
		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("here", err)
		}
		err = yaml.Unmarshal(yamlFile, c)
		if err != nil {
			fmt.Println(err)
		}
	}
	return c
}
