package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
)

const dirPerm os.FileMode = 0755
const filePerm os.FileMode = 0644

// CreateApp creates a grab app in the given directory.
func CreateApp(appPath string) {
	appName := filepath.Base(appPath)
	fmt.Printf("Creating %s app\n", appName)
	mkdirAll(appPath, "routers")
	mkdirAll(appPath, "controllers")
	mkdirAll(appPath, "models")
	mkdirAll(appPath, "services")
	mkdirAll(appPath, "core", "authentication")
	mkdirAll(appPath, "api", "parameters")
	err := createMain(appPath, appName)
	if err != nil {
		log.Fatalf("Error creating main.go: %s", err)
	}
	createAuthRoutes(appPath)
	createAuthController(appPath)
	createAuthMiddlewares(appPath)
	createAuthService(appPath)
	createAuthParameters(appPath)
	createHelloRoutes(appPath)
	createHelloController(appPath)
	createUserModel(appPath)
	createCoreAuthJWTBackend(appPath)
}

func mkdirAll(x ...string) {
	err := helpers.MkdirAll(dirPerm, x...)
	if err != nil {
		log.Fatalf("Error making the directory %s", filepath.Join(x...))
	}
}
