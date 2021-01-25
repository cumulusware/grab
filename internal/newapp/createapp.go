package newapp

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cumulusware/grab/internal/helpers"
	"github.com/manifoldco/promptui"
)

const dirPerm os.FileMode = 0755
const filePerm os.FileMode = 0644

// CreateApp needs a better description.
func CreateApp() error {
	appName, err := getAppName()
	if err != nil {
		return err
	}
	log.Printf("app name = %s", appName)

	appPath, err := getAppPath(appName)
	if err != nil {
		return err
	}
	log.Printf("app path = %s", appPath)

	// Create the various directories
	mkdirAll(appPath, "cmd")
	mkdirAll(appPath, "internal")
	mkdirAll(appPath, "docs")
	mkdirAll(appPath, "public")

	// Select database.
	db, err := getDB()
	if err != nil {
		return err
	}
	log.Printf("db = %s", db)

	// Run go mod init
	goModule, err := getGoModPath()
	if err != nil {
		return err
	}
	log.Printf("go mod path = %s", goModule)
	cmd := exec.Command("go", "mod", "init", goModule)
	cmd.Dir = fmt.Sprintf("./%s", appName)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error with go mod init: %s", err)
	}

	createMain(appPath, appName)
	createRouter(appPath)
	createRoutes(appPath)
	createAPIBase(appPath)
	return nil
}

func mkdirAll(x ...string) {
	err := helpers.MkdirAll(dirPerm, x...)
	if err != nil {
		log.Fatalf("Error making the directory %s", filepath.Join(x...))
	}
}

func getAppName() (string, error) {
	validate := func(input string) error {
		if input != strings.ToLower(input) {
			return errors.New("app name should be lower case")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter app name",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	fmt.Printf("You chose %q\n", result)
	return result, nil
}

func getAppPath(appName string) (string, error) {
	// FIXME(mdr): Check to see if the path already exists.
	defaultPath, err := filepath.Rel(".", appName)
	if err != nil {
		return "", err
	}

	validate := func(input string) error {
		if input == "" {
			return nil
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Enter app path",
		Validate:  validate,
		Default:   defaultPath,
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	if result == "" {
		result = defaultPath
	}

	fmt.Printf("You chose %q\n", result)
	return result, nil
}

func getDB() (string, error) {
	prompt := promptui.Select{
		Label: "Select Database",
		Items: []string{"PostgreSQL", "MongoDB"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("db prompt failed %v", err)
	}

	fmt.Printf("You chose %q\n", result)
	return result, nil
}

func getGoModPath() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Enter the Go module path",
		Default:   "github.com/",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	fmt.Printf("You chose %q\n", result)
	return result, nil
}
