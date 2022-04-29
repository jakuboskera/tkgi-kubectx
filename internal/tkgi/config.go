package tkgi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ahmetb/kubectx/internal/cmdutil"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Clusters []struct {
		Name    string `yaml:"name"`
		Creds   string `yaml:"creds"`
		TkgiAPI string `yaml:"tkgiApi"`
	} `yaml:"clusters"`
}

func (c *Config) get() *Config {
	path, err := tkgiKubectxConfigFile()
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

func tkgiKubectxConfigFile() (string, error) {
	home := cmdutil.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".kube", "tkgi-kubectx", "config.yaml"), nil
}
