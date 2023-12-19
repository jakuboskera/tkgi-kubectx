package tkgi

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// tkgiAvailable returns path to tkgi binary or error if not found in $PATH
func tkgiPath() (string, error) {
	path, err := exec.LookPath("tkgi")
	if err != nil {
		return "", err
	}
	return path, nil
}

// getTkgiApi returns tkgi API URI from a file
func getTkgiApi(context string) string {
	var config Config
	for _, t := range config.get().Tkgi {
		for _, c := range t.Clusters {
			if context == c {
				return t.URL
			}
		}
	}
	return ""
}

// getCredentials return credentials username and password from file according context name
func getCredentials(context string) (string, string) {
	var (
		username string
		config   Config
		creds    Credentials
	)

	for _, t := range config.get().Tkgi {
		for _, c := range t.Clusters {
			if context == c {
				username = t.Creds
			}
		}
	}

	for _, v := range creds.get().Credentials {
		if username == v.Username {
			return v.Username, v.Password
		}
	}
	return "", ""
}

// isClusterAdmin returns true if username is cluster admin, otherwise false
func isClusterAdmin(username string) bool {
	var creds Credentials
	for _, v := range creds.get().Credentials {
		if username == v.Username && v.ClusterAdmin {
			return true
		}
	}
	return false
}

// clusterAdminLogin perform login commands for cluster admin
func clusterAdminLogin(tkgiPath, tkgiApi, context, username, password string) (string, error) {
	fmt.Printf("Login first before switching context to \"%v\"...\n", context)
	cmd := exec.Command(tkgiPath, "login", "-a", tkgiApi, "-u", username, "-p", password, "-k")

	b, err := cmd.CombinedOutput()
	output := string(b)
	if err != nil {
		return "", errors.New(strings.Replace(output, "\nError: ", "", 1))
	}

	cmd = exec.Command(tkgiPath, "get-credentials", context)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("PKS_USER_PASSWORD=%v", password))

	b, err = cmd.CombinedOutput()
	output = fmt.Sprintf("%v%v\n", output, string(b))
	if err != nil {
		return "", errors.New(output)
	}
	return output, nil
}

// nonClusterAdminLogin perform login commands for non cluster admin
func nonClusterAdminLogin(tkgiPath, tkgiApi, context, username, password string) (string, error) {
	fmt.Printf("Login first before switching context to \"%v\"...\n", context)
	cmd := exec.Command(tkgiPath, "get-kubeconfig", context, "-a", tkgiApi, "-u", username, "-p", password, "-k")

	b, err := cmd.CombinedOutput()
	output := string(b)
	if err != nil {
		return "", errors.New(strings.Replace(output, "\nError: ", "", 1))
	} else {
		return output, nil
	}
}

// Login perform login to given context
func Login(context string) (string, error) {
	tkgi, err := tkgiPath()
	if err != nil {
		return "", err
	}

	tkgiApi := getTkgiApi(context)

	// if tkgiApi is empty it means that this context is not specified
	// in ~/.kube/tkgi-kubectx/config.yaml for tkgi-kubectx, so no login is needed
	if tkgiApi == "" {
		return "", nil
	}

	username, password := getCredentials(context)

	if isClusterAdmin(username) {
		return clusterAdminLogin(tkgi, tkgiApi, context, username, password)
	}
	return nonClusterAdminLogin(tkgi, tkgiApi, context, username, password)
}
