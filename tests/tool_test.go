package test

import (
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	_ = tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	_ = tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	_ = tIfElse.Execute(os.Stdout, nil)

}

type app struct{
	Name string
}

func (a *app)NameGetter() string {
	return "name from getter"
}

func TestTemplate2(t *testing.T) {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("${{.Name}} | {{.NameGetter}}$"))
	_ = tEmpty.Execute(os.Stdout, &app{Name:"name from mumber"})
}