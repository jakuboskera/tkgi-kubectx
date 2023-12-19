package tkgi

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tkgi []struct {
		URL      string   `yaml:"url"`
		Creds    string   `yaml:"creds"`
		Clusters []string `yaml:"clusters"`
	} `yaml:"tkgi"`
}

// get method returns content of config.yaml file
func (c *Config) get() *Config {
	path, err := getTkgiKubectxFile("config.yaml")
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
