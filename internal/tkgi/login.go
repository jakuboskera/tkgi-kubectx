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
func getTkgiApi(newContext string) string {
	var config Config
	config.get()
	for _, v := range config.Clusters {
		if newContext == v.Name {
			return v.TkgiAPI
		}
	}
	return ""
}

// getCredentials return credentials username and password from file according newContext name
func getCredentials(newContext string) (string, string) {
	var username string
	var config Config
	config.get()

	for _, v := range config.Clusters {
		if newContext == v.Name {
			username = v.Creds
		}
	}

	var creds Credentials
	creds.get()
	for _, v := range creds.Credentials {
		if username == v.Username {
			return v.Username, v.Password
		}
	}
	return "", ""
}

func isClusterAdmin(username string) bool {
	var creds Credentials
	creds.get()
	for _, v := range creds.Credentials {
		if username == v.Username && v.ClusterAdmin {
			return true
		}
	}
	return false
}

func clusterAdminLogin(tkgiPath, newTkgiApi, newContext, username, password string) (string, error) {
	fmt.Printf("Login first before switching context to \"%v\"...\n", newContext)
	cmd1 := exec.Command(tkgiPath, "login", "-a", newTkgiApi, "-u", username, "-p", password, "-k")

	b, err := cmd1.CombinedOutput()
	output := string(b)
	if err != nil {
		return "", errors.New(strings.Replace(output, "\nError: ", "", 1))
	}

	cmd2 := exec.Command(tkgiPath, "get-credentials", newContext)
	cmd2.Env = os.Environ()
	cmd2.Env = append(cmd2.Env, fmt.Sprintf("PKS_USER_PASSWORD=%v", password))

	b, err = cmd2.CombinedOutput()
	output = fmt.Sprintf("%v%v\n", output, string(b))
	if err != nil {
		return "", errors.New(output)
	}
	return output, nil
}

func nonClusterAdminLogin(tkgiPath, newTkgiApi, newContext, username, password string) (string, error) {
	fmt.Printf("Login first before switching context to \"%v\"...\n", newContext)
	cmd := exec.Command(tkgiPath, "get-kubeconfig", newContext, "-a", newTkgiApi, "-u", username, "-p", password, "-k")

	b, err := cmd.CombinedOutput()
	output := string(b)
	if err != nil {
		return "", errors.New(strings.Replace(output, "\nError: ", "", 1))
	} else {
		return output, nil
	}
}

func Login(currentContext, newContext string) (string, error) {
	currTkgiApi := getTkgiApi(currentContext)
	newTkgiApi := getTkgiApi(newContext)

	// if empty then for this newContext no login needed
	if newTkgiApi == "" {
		return "", nil
	}

	// if TKGI API is same for current and new context then no login needed
	if currTkgiApi == newTkgiApi {
		return "", nil
	}

	tkgi, err := tkgiPath()
	if err != nil {
		return "", err
	}

	username, password := getCredentials(newContext)

	if isClusterAdmin(username) {
		return clusterAdminLogin(tkgi, newTkgiApi, newContext, username, password)
	}
	return nonClusterAdminLogin(tkgi, newTkgiApi, newContext, username, password)
}
