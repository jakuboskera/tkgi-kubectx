package tkgi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ahmetb/kubectx/internal/cmdutil"
	"gopkg.in/yaml.v3"
)

type Credentials struct {
	Credentials []struct {
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		ClusterAdmin bool   `yaml:"clusterAdmin"`
	} `yaml:"credentials"`
}

func (c *Credentials) get() *Credentials {
	path, err := tkgiKubectxCredentialsFile()
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

func tkgiKubectxCredentialsFile() (string, error) {
	home := cmdutil.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".kube", "tkgi-kubectx", "credentials.yaml"), nil
}
