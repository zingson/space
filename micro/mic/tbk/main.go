package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"runtime"

	"tbk/graph"
)

// 服务默认配置信息
var appConfig = `
{
	"name":"tbk",
	"version":"0.0.1",
	"server": {
		"port":5803
	}
}
`

func main() {
	fmt.Println("******************************************************************************************")
	defer PrintStack()
	args()
	ListenAndServe()
}

func args() {
	var (
		// 服务名与端口号
		name         string
		port         int
		templatePath string
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))
	flag.StringVar(&templatePath, "t", app.Name, "页面渲染模板路径")
	flag.Parse()
	app.Name = name
	app.Port = port

	if templatePath == "" {
		// 测试时，需要编译输出到当前路径下才可以访问到
		app.TemplatePath = "./template/tmall-shop-site-v1"
	}
}

type Application struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	TemplatePath string `json:"templatePath"`
}

var app = &Application{Name: "tbk", Host: "0.0.0.0", Port: 5800, TemplatePath: ""}

func ListenAndServe() {
	Handles()
	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))

	if err := recover(); err != nil {
		log.Fatalf("** Main Fatalf-> %s", err)
	}
}

// HttpHandle
func Handles() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		w.Write([]byte("MicroService:" + app.Name))
	})

	// Graphql服务
	http.Handle("/graphql", graph.GraphqlHttpHandler())

	// 静态文件服务
	http.Handle("/assets/page.css", http.FileServer(http.Dir("public")))

	// 网站首页，访问URL：http://127.0.0.1:5803/xxx/index.html
	http.HandleFunc("xxx/index.html", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./")
		if err != nil {
			log.Print(err)
			io.WriteString(w, err.Error())
			return
		}
		t.Execute(w, nil)
	})

}
