package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const tpl = `{{ $Name := .Name }}Name is {{$Name}}.
The age is {{.Age}}.
{{ range .Emails }}An email is {{. | emailExpand}}. 
{{ end }}

{{ with .Jobs}}
	{{ range . }}Employer is {{ .Employer }} and the role is {{ .Role }}.
	{{ end }}
{{ end }}
`

func emailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substr := strings.Split(s, "@")
	if len(substr) != 2 {
		return s
	}

	return substr[0] + " at " + substr[1]
}

func main() {
	job1 := Job{Employer: "Employer 1", Role: "No one"}
	job2 := Job{Employer: "Employer 2", Role: "Boss"}

	p := Person{
		Name:   "Someone",
		Age:    30,
		Emails: []string{"first@email.com", "second@email.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	t := template.New("Person template")
	t = t.Funcs(template.FuncMap{"emailExpand": emailExpander})

	t, err := t.Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse template:", err)
		os.Exit(1)
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot execute template:", err)
		os.Exit(1)
	}
}
