package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
	"github.com/spf13/cobra"
)

const dirPerm os.FileMode = 0755
const filePerm os.FileMode = 0644

func NewApp(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Usage()
		log.Fatalln("Path for new site needs to be provided")
	}

	appPath, err := filepath.Abs(filepath.Clean(args[0]))
	if err != nil {
		cmd.Usage()
		log.Fatalln(err)
	}
	appName := filepath.Base(appPath)

	mkdirAll(appPath, "docs")
	createMainGo(appPath, appName)
	createRouterGo(appPath)
	createRoutesGo(appPath)
	createAPIBaseGo(appPath)
}

func createMainGo(appPath, appName string) {
	data := []byte(`package main

import (
	"flag"
	"fmt"

	"github.com/codegangsta/negroni"
)

func main() {
	var (
		portFlag = flag.Int("port", 9090, "port to serve ` + appName + `.")
	)
	flag.Parse()
	port := *portFlag
	addr := fmt.Sprintf(":%d", port)
	handler := InitializeRouter()
	n := negroni.Classic()
	n.UseHandler(handler)
	n.Run(addr)
}`)
	mainGoFile := filepath.Join(appPath, "main.go")
	ioutil.WriteFile(mainGoFile, data, filePerm)
}

func createRouterGo(appPath string) {
	data := []byte(`package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitializeRouter sets up all the handlers and returns a *mux.Router.
func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router
}`)
	routerGoFile := filepath.Join(appPath, "router.go")
	ioutil.WriteFile(routerGoFile, data, filePerm)
}

func createRoutesGo(appPath string) {
	data := []byte(`package main

import "net/http"

// Route represents a REST API route.
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes is a slice of Route.
type Routes []Route

var routes = Routes{
	Route{"ReadAPI", "GET", "/api", ReadAPI},
	Route{"DescribeAPI", "OPTIONS", "/api", DescribeAPI},
}`)
	routesGoFile := filepath.Join(appPath, "routes.go")
	ioutil.WriteFile(routesGoFile, data, filePerm)
}

func createAPIBaseGo(appPath string) {
	data := []byte(`package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	readLimit = 1048576 // 1 Mbyte
)

// ReadAPI is the handler for the GET /api endpoint.
func ReadAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// DescribeAPI is the handler for the OPTIONS /api endpoint.
func DescribeAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET,HEAD,OPTIONS")
	w.WriteHeader(http.StatusOK)
}`)
	apiBaseGoFile := filepath.Join(appPath, "api_base.go")
	ioutil.WriteFile(apiBaseGoFile, data, filePerm)
}

func mkdirAll(x ...string) {
	err := helpers.MkdirAll(dirPerm, x...)
	if err != nil {
		log.Fatalf("Error making the directory %s", filepath.Join(x...))
	}
}

func main() {

	var newCmd = &cobra.Command{
		Use:   "new <app-name>",
		Short: "Create directory <app-name> and initialize as REST API",
		Run:   NewApp,
	}
	var rootCmd = &cobra.Command{Use: "grab"}
	rootCmd.AddCommand(newCmd)
	rootCmd.Execute()

}
