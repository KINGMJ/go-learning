package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

func main() {
	demo10()
}

// 基础用法，解析模板
func demo1() {
	tmpl := template.New("myTemplate")
	tmpl, err := tmpl.Parse("Hello, {{.}}!")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, "Alice")
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

type UserData struct {
	Name  string
	Age   int
	Email string
}

// 从结构体中解析模板
func demo2() {
	tmpl := `
	{
		"name": "{{.Name}}",	
		"age": "{{.Age}}"
	}
	`
	data := UserData{Name: "Alice", Age: 30}
	t, err := template.New("user").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 解析文件模版
func demo3() {
	tmpl := template.Must(template.ParseFiles("user.json"))
	data := UserData{Name: "Alice", Age: 30}
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func demo4() {
	tmpl := `
Name: {{.Name}}
Age: {{.Age1}}
IsAdmin: {{.IsAdmin}}
`
	data := map[string]interface{}{
		"Name": "Alice",
		"Age":  30,
	}
	t, err := template.New("user").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func demo5() {
	tmpl := `
<ul>
{{range .Todos}} 
	{{if .Done1}}
	<li class="done">{{.Title}}</li>
	{{else}}
	<li>{{.Title}}</li>
	{{end}} 
{{end}}
<ul>
`
	// data := map[string]interface{}{
	// 	"Name": "Alice",
	// 	"Age":  30,
	// }
	data := TodoPageData{
		PageTitle: "My Todo List",
		Todos: []Todo{
			{Title: "Task1", Done: false},
			{Title: "Task2", Done: true},
		},
	}

	t, err := template.New("user").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 内置的函数使用
func demo6() {
	tmpl := `
Hello, {{.Name}}!
Your email is: {{ .Email}}
{{call .X 1 2}}
{{html .Html}}
`
	// data := []string{"apple", "banana", "cherry"}
	data := map[string]interface{}{
		"Name":     "Alice",
		"LastName": "Chen",
		"Age":      30,
		"IsAdmin":  false,
		"Arr":      [][]int{{1, 2, 3}, {4, 5, 6}},
		"Slice":    []int{1, 2, 3, 4},
		"Email":    `<script>alert("123")</script>`,
		"Html":     `<a href="mailto:alice@example.com">alice@example.com</a>`,
		"X": func(a, b int) int {
			return a + b
		},
	}

	t, err := template.New("builtin").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = t.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 自定义函数
func demo7() {
	tmpl := `
Hello, {{.Name}}!
Your email is: {{.Email}}
The sum of 1 and 2 is : {{add 1 2}}
The pow of 2 and 3 is : {{pow 2 3}}
`
	data := map[string]interface{}{
		"Name":  "Alice",
		"Email": "alice@example.com",
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"pow": func(a, b int) float64 {
			return math.Pow(float64(a), float64(b))
		},
	}

	t, err := template.New("customFuncs").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 管道操作
func demo8() {
	tmpl := `Hello, {{with $upperCase := . | toUpperCase}}
{{repeat $upperCase 3}}
{{end}}
`
	funcMap := template.FuncMap{
		"toUpperCase": strings.ToUpper,
		"repeat": func(s string, count int) string {
			return strings.Repeat(s, count)
		},
	}
	t, err := template.New("greeting").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = t.Execute(w, "Alice")
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8082", nil)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) -----------
// 模板嵌套
func demo9() {
	userTmpl, err := template.ParseFiles("user.tmpl", "avatar.tmpl")
	if err != nil {
		panic(err)
	}
	user := UserData{Name: "Alice", Email: "alice@example.com"}
	err = userTmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) -----------

func demo10() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Todo List",
			Todos: []Todo{
				{Title: "Task1", Done: false},
				{Title: "Task2", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
