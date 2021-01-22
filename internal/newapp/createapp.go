package newapp

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cumulusware/grab/internal/helpers"
)

const dirPerm os.FileMode = 0755
const filePerm os.FileMode = 0644

// CreateApp needs a better description.
func CreateApp(appPath string) {
	appName := filepath.Base(appPath)
	mkdirAll(appPath, "docs")
	mkdirAll(appPath, "public")
	createMain(appPath, appName)
	createRouter(appPath)
	createRoutes(appPath)
	createAPIBase(appPath)

}

func mkdirAll(x ...string) {
	err := helpers.MkdirAll(dirPerm, x...)
	if err != nil {
		log.Fatalf("Error making the directory %s", filepath.Join(x...))
	}
}
