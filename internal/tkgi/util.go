package tkgi

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ahmetb/kubectx/internal/cmdutil"
)

// exists returns true if file or folder exists
func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// getTkgiKubectxFile returns full path for given file
func getTkgiKubectxFile(file string) (string, error) {
	home := cmdutil.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".kube", "tkgi-kubectx", file), nil
}
