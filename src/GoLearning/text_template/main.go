package main

import (
	"os"
	"text/template"
)

type Todo struct {
	Name string
	Desc string
}

func main() {
	td := Todo{"test task", "demo of the text/template library"}

	tpl := template.Must(template.New("todos").Parse("You have a task named \"{{ .Name }}\" with description: \"{{ .Desc }}\"\n"))
	/*if err != nil {
		panic(err)
	}*/

	if err := tpl.Execute(os.Stdout, td); err != nil {
		panic(err)
	}

	td2 := Todo{"Go text todo", "template for a go text/template package"}
	if err := tpl.Execute(os.Stdout, td2); err != nil {
		panic(err)
	}

}
