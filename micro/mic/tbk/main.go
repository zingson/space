package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"tbk/graph"
)

func main() {
	fmt.Println("******************************************************************************************")
	defer PrintStack()
	args()
	app.ListenAndServe()
}

func args() {
	var (
		// 服务名与端口号
		name string
		port int
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))
	flag.Parse()

	app.Name = name
	app.Port = port
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))

	if err := recover(); err != nil {
		log.Fatalf("** Main Fatalf-> %s", err)
	}
}

type Application struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

var app = &Application{Name: "config", Host: "0.0.0.0", Port: 5800}

func (app *Application) ListenAndServe() {
	app.Handles()
	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}

func (app *Application) Handles() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		w.Write([]byte("MicroService:config"))
	})

	// Graphql服务
	http.Handle("/graphql", graph.GraphqlHttpHandler())

}
